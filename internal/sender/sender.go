package sender

import (
	"context"
	"itube/users/internal/db"
	"itube/users/pkg/api"
	"itube/users/pkg/email"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
)

// Sender отвечает за отправку писем с токенами.
type Sender struct {
	db     *db.Adapter      // доступ к базе данных
	tmplts *email.Templates // шаблоны писем
}

// New возвращает инициализированный отправщик почтовых сообщений с токенами.
func New(db *db.Adapter, t *email.Templates) *Sender {
	return &Sender{db: db, tmplts: t}
}

// Send запрашивает список ключей для отправки и отправляет их по почте.
func (s Sender) Send(ctx context.Context) error {
	// запрашиваем список токенов для отправки
	tokens, err := s.db.TokensToSend(ctx)
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil // нечего отсылать
	}
	// устанавливаем соединение с SMTP-сервером
	sender, err := s.tmplts.Dialer.Dial()
	if err != nil {
		return err
	}
	defer sender.Close()
	var msg = mail.NewMessage() // инициализируем почтовое сообщение
	// перебираем все токены
	for _, token := range tokens {
		domain, err := s.tmplts.Domain(token.Domain)
		if err != nil {
			log.WithError(err).Warn("ignore token for domain")
			continue
		}
		email, err := domain.Email(api.TokenType(token.Type).String())
		if err != nil {
			log.WithError(err).Warn("ignore token for token type")
			continue
		}
		// формируем почтовое сообщение
		msg.Reset()
		msg.SetHeader("From", domain.From)
		msg.SetHeader("To", token.Email)
		err = email.WithToken(token.Token).Apply(msg)
		if err != nil {
			log.WithError(err).Warn("ignore email template")
			continue
		}
		// отсылаем письмо
		err = sender.Send(domain.From, []string{token.Email}, msg)
		if err != nil {
			return err
		}
		// ставим метку, что письмо отправлено
		err = s.db.TokenSended(ctx, token.Token)
		if err != nil {
			return err
		}
	}
	return nil
}
