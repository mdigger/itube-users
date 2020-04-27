package email

import (
	"errors"
	"strings"

	"gopkg.in/mail.v2"
)

// TokenPlaceholder описывает строку, которая подменяется на реальный токен.
var TokenPlaceholder = "_TOKEN_PLACEHOLDER_"

// Template описывает содержимое письма.
type Template struct {
	Subject string // тема письма
	Text    string // текстровый вариант письма
	HTML    string // html вариант письма
}

// WithReplace заменяет пары строк в шаблоных письма через string.Replacer и
// возвращает копию письма.
func (e Template) WithReplace(oldnew ...string) Template {
	var repl = strings.NewReplacer(oldnew...)
	return Template{
		Subject: e.Subject,
		Text:    repl.Replace(e.Text),
		HTML:    repl.Replace(e.HTML),
	}
}

// WithToken возвращает копию шаблона письма, подставив токен вместо
// TokenPlaceholder.
//
// Внимание! Т.к. токен может (а обычно так и есть) использоваться в url, то
// необходимо убедиться, что он содержит только допустимые в url символы.
// Лучше всего предварительно использовать функции url.PathEscape или
// url.QueryEscape, в зависимости от того, где заменяется значение. Если
// необходимо заменить в разных местах разными значениями, то лучше
// воспользователься методом WithReplace, указав разные метки и дав им
// соответствующие значения.
func (e Template) WithToken(token string) Template {
	return e.WithReplace(TokenPlaceholder, token)
}

// ErrEmpty в случае пустого шаблона письма.
var ErrEmpty = errors.New("empty email")

// Apply заполняет почтовое сообщение данными из шаблона: заголовок и тело
// письма в формате текст и html, если они определены. Возвращает ошибку, если
// текст писем не задан ни в одном формате.
func (e Template) Apply(m *mail.Message) error {
	m.SetHeader("Subject", e.Subject)
	if e.Text != "" {
		m.SetBody("text/plain", e.Text)
		if e.HTML != "" {
			m.AddAlternative("text/html", e.HTML)
		}
	} else if e.HTML != "" {
		m.SetBody("text/html", e.HTML)
	} else {
		return ErrEmpty
	}
	return nil
}
