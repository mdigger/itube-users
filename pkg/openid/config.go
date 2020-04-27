package openid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

// Настройки, используемые для формирования и проверки случайной строки сессии
// авторизации и nonce.
var (
	StateLength = 48               // длина сессионой строки (рекомендуется не меньше 30)
	StateTTL    = time.Minute * 30 // время жизни сессии авторизации
	// строка с идентификацией сервера (сервиса), используемая при создании
	// токена авторизации внешним сервером авторизации
	ServerNonce = randomToken(8)
)

// randomToken generates a random @length length token.
func randomToken(length int) string {
	var bytes = make([]byte, base64.RawURLEncoding.DecodedLen(length))
	_, err := rand.Read(bytes)
	if err != nil {
		panic(fmt.Errorf("random string generation error: %w", err))
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

// Config описывает информацию для конфигурации OpenID Connect авторизации.
//
// В качестве провайдера указывается его название и url для обращение к серверу
// авторизации. Например, для Google используется адрес
// "https://accounts.google.com".
//
// ClienID и Secret необходимо получить на сервер провайдера при регистрации
// сервиса. Для Google нужно зарегистрировать в административной консоли
// свой проект, заполнить информацию о нем и сформировать пары для
// для идентификации клиента. Адрес сервера консоли:
// https://console.developers.google.com/project/<your-project-id>/apiui/credential
//
// RedirectURLs задает список зарегистрированных URL для возврата
// авторизационной информации. Если задан, то перед генерацией url для
// авторизации проверяется, что он находится в списке.
type Config struct {
	Name     string // Название провайдера
	URL      string // URL сервиса авторизации
	СlientID string // иденитфикатор клиента
	Secret   string // секретный код для подтверждения клиента
}
