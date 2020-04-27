package db

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// bcryptCost задает "стоимость" хеширования пароля пользователя.
// Т.к. это внутренний сервис то, в принципе, стоимость можно и снизить, ведь
// нет необходимости исскуственно замедлять этот процесс.
const bcryptCost = bcrypt.MinCost

// Adapter отвечает за работу с базой данных.
//
// Проверки пустых значений и правильности форматов входящих данных при вызове
// методов не происходи, т.к. лучше это перенести на проверку самого запроса.
//
// В общем случае ошибки базы данных возвращаются как есть - pgconn.PgError.
// В частности, "22P02" отвечает за неверный формат данных. Для получения
// сообщения о таких ошибках лучше использовать поле Message, т.к. стандартное
// строковое представление слишком многословно.
type Adapter struct {
	*pgxpool.Pool // пулл подключений к postgres
}

// Register регистрирует нового пользователя с логином и паролем. Если
// пользователь с таким email уже зарегистрирован, но у него не задан пароль,
// то ему автоматически устанавливается новый пароль. Если данные не изменяются,
// то ошибка не генерится.
func (db *Adapter) Register(ctx context.Context,
	email, password string) (*UserInfo, error) {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return nil, ErrEmptyEmail
	}
	// шифруем пароль пользователя перед сохранением
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}
	// регистрируем пользователя с паролем и разбираем данные о нем
	user, err := scanUser(db.QueryRow(ctx, sqlInsertUser, email, hashed))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrAlreadyRegisterd
		}
		return nil, err
	}
	return user, nil
}

// Authorize возвращает информацию о пользователе по логину и паролю.
// Проверяются только те пользователи, которые зарегистрированы с паролем.
// Если пользователь зарегистрирован через внешнего провайдера авторизации и
// у него не задан пароль, то авторизация через этот метод не пройдет и будет
// возвращена ошибка ErrNotFound.
func (db *Adapter) Authorize(ctx context.Context,
	email, password string) (*UserInfo, error) {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return nil, ErrEmptyEmail
	}
	// получаем и разбираем данные о пользователе, а так же пароль
	var hashed []byte
	user, err := scanUser(db.QueryRow(ctx, sqlSelectPassword, email), &hashed)
	if err != nil {
		return nil, err
	}
	// проверяем, что пароль совпадает
	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			err = ErrInvalidPassword
		}
		return nil, err
	}
	// обновляем дату последней успешной авторизации (ошибку игнорируем)
	_ = oneRow(db.Exec(ctx, sqlLogged, user.UID))
	return user, nil
}

// SetPassword изменяет или задает пароль пользователя, если он до этого был
// не задан.
func (db *Adapter) SetPassword(ctx context.Context,
	uid, password string) error {
	// шифруем пароль пользователя перед сохранением
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	// сохраняем новый пароль пользователя
	return oneRow(db.Exec(ctx, sqlUpdatePassword, hashed, uid))
}

// Update обновляет информацию о пользователе. При изменении email может
// возникнуть ошибка, если другой пользователь с таким email уже зарегистрирован.
//
// Расширенные свойства могут быть любой строкой, которую postgres посчитает
// json. А это достаточно широкие пределы. Для "сброса" расширенных свойств
// можно передать пустую строку.
//
// Возвращает ErrAlreadyRegisterd при попытке сменить email адрес на другой,
// который уже зарегистрирован за другим пользователем.
func (db *Adapter) Update(ctx context.Context,
	uid, email, properties string) error {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return ErrEmptyEmail
	}
	// обновляем информацию о пользователе
	err := oneRow(db.Exec(ctx, sqlUpdateUser, email, null(properties), uid))
	var dbErr = new(pgconn.PgError)
	// проверяем, что это ошибка смены email на уже существующий
	if err != nil && errors.As(err, &dbErr) && dbErr.Code == "23505" {
		return ErrAlreadyRegisterd
	}
	return err
}

// GetUser возвращает информацию о пользователе по его email адресу или
// уникальному идентификатору. Нужно, чтобы хотя бы одно из значений совпало
// с тем, что зарегистрировано в базе данных.
func (db *Adapter) GetUser(ctx context.Context,
	uid, email string) (*UserInfo, error) {
	return scanUser(db.QueryRow(ctx, sqlSelectUserByEmailOrUID,
		null(uid), null(email)))
}

// BlockUser блокирует/разблокирует пользователя. Заблокированный пользователь
// остается зарегистрированным, но не может авторизоваться.
func (db *Adapter) BlockUser(ctx context.Context,
	uid string, blocked bool) error {
	return oneRow(db.Exec(ctx, sqlBlockUser, blocked, uid))
}

// Logged обновляет дату последней авторизации пользователя.
func (db *Adapter) Logged(ctx context.Context,
	uid string) error {
	return oneRow(db.Exec(ctx, sqlLogged, uid))
}

// OpenIDAuthorize авторизует пользователя по информации о внешней авторизации.
func (db *Adapter) OpenIDAuthorize(ctx context.Context,
	provider, subject string) (*UserInfo, error) {
	user, err := scanUser(db.QueryRow(ctx, sqlSelectUserOpenID, provider, subject))
	if err != nil {
		return nil, err
	}
	// обновляем дату последней успешной авторизации (ошибку игнорируем)
	_ = oneRow(db.Exec(ctx, sqlLogged, user.UID))
	return user, nil
}

// OpenIDRegister регистрирует пользователя по информации о внешней авторизации.
// Множественные регистрации одного и того же пользователя не приведут к ошибке,
// а только изменят расширенные свойства пользователя.
func (db *Adapter) OpenIDRegister(ctx context.Context,
	provider, subject, email string, verified bool, properties string) (*UserInfo, error) {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return nil, ErrEmptyEmail
	}
	// стартуем транзакцию
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	// добавляем email в список подтвержденных
	if verified {
		_, err = tx.Exec(ctx, sqlInsertVerifiedEmail, email)
		if err != nil {
			return nil, err
		}
	}
	// регистрируем нового пользователя
	user, err := scanUser(tx.QueryRow(ctx, sqlInsertUserOpenID, email,
		null(properties)))
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}
	// добавляем информацию об openid регистрации
	_, err = tx.Exec(ctx, sqlInsertOpenID, provider, subject, user.UID)
	if err != nil {
		return nil, err
	}
	// принимаем транзакцию
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// RegInfo добавляет в журнал информацию о регистрации пользователя.
// Каждый вызов этой функции добавляет в журнал новую запись о регистрации.
//
// Возвращает ошибку ErrNotFound, если пользователь не зарегистрирован.
func (db *Adapter) RegInfo(ctx context.Context,
	domain, uid, email, provider, referer, utm string) error {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return ErrEmptyEmail
	}
	// заносим запись в журнал регистрации
	_, err := db.Exec(ctx, sqlInsertRegInfo,
		domain, uid, email, null(provider), null(referer), null(utm))
	// проверяем, что пользователь с таким идентификатором действительно
	// зарегистрирован.
	var dbErr = new(pgconn.PgError)
	if err != nil && errors.As(err, &dbErr) && dbErr.Code == "23503" {
		return ErrNotFound
	}
	return err
}

// EmailVerified добавляет email-адрес в список подтвержденных.
// При повторном вызове с тем же адресом ошибки не возникает: обновляется
// только дата последнего подтверждения адреса.
func (db *Adapter) EmailVerified(ctx context.Context,
	email string) error {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return ErrEmptyEmail
	}
	_, err := db.Exec(ctx, sqlInsertVerifiedEmail, email)
	return err
}

// tokenCoder используется для кодирования/декодирования токенов в строковый
// формат
var tokenCoder = base64.RawURLEncoding

// TokenGenerate генерирует новый токен для проверки почты или сброса пароля.
// Если для одного и того же домена, почтового адреса и типа уже был
// сгенерирован токен, то он заменяется на новый, что отменяет действие
// предыдущего.
func (db *Adapter) TokenGenerate(ctx context.Context,
	domain, email string, tokenType int32) (string, error) {
	// т.к. email является ключевым идентификационным полем, то на всякий
	// случай проверяем, что оно задано
	if email == "" {
		return "", ErrEmptyEmail
	}
	var token []byte
	err := db.QueryRow(ctx, sqlInsertToken, domain, email, tokenType).Scan(&token)
	if err != nil {
		return "", err
	}
	// приводим токен в формату строки, безопасной для передачи в качестве
	// значения в url
	return tokenCoder.EncodeToString(token), nil
}

// TokenVerify проверяет токен. Если токен найден, то почтовый адрес, на который
// он был  отправлен, автоматически помечается подтвержденным.
// Cам токен автоматически удаляется и повторное его использование невозможно.
// Может возвращаеть ошибку ErrNotFound, если адрес пользователя с тех пор
// изменился. Так же может быть ошибка ErrBlocked, если пользователь
// заблокирован.
func (db *Adapter) TokenVerify(ctx context.Context,
	token string) (*UserInfo, error) {
	// декодируем токен в бинарный формат
	tokenUUID, err := tokenCoder.DecodeString(token)
	if err != nil {
		return nil, ErrBadToken
	}
	// удаляем токен из базы и получаем почтовый адрес и тип, с которыми он
	// был зарегистрирован
	var (
		email     string
		tokenType int32     // на текущий момент тип токена не проверяется
		created   time.Time // время его создания
	)
	// удаляем токен в любом случае, раз уж он проверяется, чтобы нельзя было
	// его повторно использовать, поэтому делаем это вне транзакции
	err = db.QueryRow(ctx, sqlDeleteToken, tokenUUID).Scan(
		&email, &tokenType, &created)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrBadToken
		}
		return nil, err
	}
	// стартуем транзакцию, чтобы добавление проверенного адреса было
	// гарантировано сохранено в базе до того, как будет осуществлена выборка
	// данных о пользователе и в том же потоке, что и сама выборка
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	// добавляем email в список подтвержденных, какой бы тип токена не
	// использовался, т.к. все равно передано по почте, что однозначно ее
	// подтверждает
	_, err = tx.Exec(ctx, sqlInsertVerifiedEmail, email)
	if err != nil {
		return nil, err
	}
	// запрашиваем и разбираем информацию о пользователе
	user, err := scanUser(tx.QueryRow(ctx, sqlSelectUser, email))
	if err != nil {
		return nil, err
	}
	// принимаем транзакцию
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	// обновляем дату последней успешной авторизации (ошибку игнорируем)
	_ = oneRow(db.Exec(ctx, sqlLogged, user.UID))
	return user, nil
}

// TokenSended помечает токен как отправленный.
func (db *Adapter) TokenSended(ctx context.Context,
	token string) error {
	// декодируем токен в бинарный формат
	tokenUUID, err := tokenCoder.DecodeString(token)
	if err != nil {
		return ErrBadToken
	}
	// помечаем токен как отправленный
	_, err = db.Exec(ctx, sqlUpdateToken, tokenUUID)
	return err
}

// TokensToSend возвращает канал с токенами, предназначенными для отправки.
// func (db *Adapter) TokensToSend(ctx context.Context) error {
// 	return nil
// }

// TokenInfo описывает информацию о токене.
type TokenInfo struct {
	Token  string // представлени токена в виде base64 строки
	Domain string // название домена
	Email  string // email адрес пользователя
	Type   int32  // тип токена
}

// TokensToSend возвращает список токенов для отсылки.
func (db *Adapter) TokensToSend(ctx context.Context) ([]TokenInfo, error) {
	rows, err := db.Query(ctx, sqlSelectTokens)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil // нет токенов для отправки
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tokens = make([]TokenInfo, 0)
	for rows.Next() {
		var (
			token TokenInfo
			id    = make([]byte, 0, 16)
		)
		err = rows.Scan(&id, &token.Domain, &token.Email, &token.Type)
		if err != nil {
			return nil, err
		}
		token.Token = tokenCoder.EncodeToString(id)
		tokens = append(tokens, token)
	}
	return tokens, nil
}
