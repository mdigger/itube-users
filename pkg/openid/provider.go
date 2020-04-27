package openid

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// Provider отвечает за обработку запросов авторизации пользователя через
// внешние системы по протоколу OpenID Connect.
type Provider struct {
	name         string         // название
	provider     *oidc.Provider // провайдер авторизации
	oauth2Config oauth2.Config  // конфигурация OAuth2
	states       *states        // хранение списка nonce для авторизации
}

// String возвращает идетификатор провайдера (название).
func (p Provider) String() string {
	return p.name
}

// New инициализирует и возвращает обработчик авторизации по протоколу
// OpenID Connect. Context используется для запроса информации с адресами
// поддерживаемых сервисов провайдера.
func New(cfg Config) (*Provider, error) {
	// проверяем, что заданы ключи для инициализации провайдера
	if cfg.СlientID == "" || cfg.Secret == "" {
		return nil, ErrMissingProviderKeys
	}
	// инициализируем провайдера авторизации
	// с контекстом тут жопа: он сохраняется и с ним потом запрашиваются ключи
	var provider, err = oidc.NewProvider(context.Background(), cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("authorization provider initialization error: %w", err)
	}
	// Задаем настройки для OAuth2.
	var oauth2Config = oauth2.Config{
		ClientID:     cfg.СlientID,
		ClientSecret: cfg.Secret,
		Endpoint:     provider.Endpoint(), // вычисляется после запроса сервисов провайдера
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	// инициализируем хранилище случайных последовательностей, используемых
	// при авторизации
	var states = new(states)
	// запускаем процесс автоматического удаления устаревших сессионных
	// ключей для проверки авторизации
	go func() {
		for {
			var before = time.Now().Add(-StateTTL) // время для проверки устаревания
			states.Range(func(key string, value stateObj) bool {
				// проверяем, что не "устарело"
				if value.Created.Before(before) {
					states.Delete(key) // удаляем устаревшую сессию
				}
				return true // переходим к проверке следующей сессии
			})
			// в принципе, нам не нужно слишком часто очищать этот список
			// но пусть будет в качестве интервала очистки тот же интервал,
			// что используется для задания времени жизни
			time.Sleep(StateTTL) // на время прекращаем обработку
		}
	}()
	// инициализируем наш обработчик авторизации
	var auth = &Provider{
		name:         cfg.Name,
		provider:     provider,
		oauth2Config: oauth2Config,
		states:       states,
	}
	return auth, nil
}

// LoginURL формирует и возвращает уникальный URL для авторизации пользователя
// через провайдера авторизации.
//
// redirectURL задает адрес для возврата авторизационной информации. Если не
// задан, то необходимо потом самостоятельно добавить к его полученному url
// с помощью параметра redirect_uri.
//
// params позволяют задать дополнительные именованны параметры, используем
// при авторизации. Например: login_hint, hd, display
// Подробнее о дополнительный параметрах можно прочитать:
// https://developers.google.com/identity/protocols/oauth2/openid-connect#authenticationuriparameters
func (p *Provider) LoginURL(redirectURI string, params map[string]string, data interface{}) string {
	// формируем список дополнительных параметров запроса
	var opts = make([]oauth2.AuthCodeOption, 0, len(params)+1)
	for k, v := range params {
		opts = append(opts, oauth2.SetAuthURLParam(k, v))
	}
	// добавляем идентификатор сервера
	if ServerNonce != "" {
		opts = append(opts, oidc.Nonce(ServerNonce))
	}
	// сохраняем объект в сессии и получаем сессионный ключ
	// этот ключ одновременно служит для защиты авторизационной сессии
	var state = randomToken(48)
	p.states.Store(state, stateObj{
		Created:     time.Now(),
		RedirectURI: redirectURI,
		Data:        data,
	})
	// копируем конфигурацию для авторизации и добавляем адрес для редиректа
	var cfg = p.oauth2Config
	cfg.RedirectURL = redirectURI
	// формируем адрес для авторизации пользователя и перенаправляем на него
	return cfg.AuthCodeURL(state, opts...)
}

var (
	// ErrMissingProviderKeys возвращается, если не заданы ключи для
	// инициализации провайдера авторизации
	ErrMissingProviderKeys = errors.New("missing provider clientid or secret")
	// ErrBadState возвращается, если указан неверный сессионный ключ или его
	// время жизни уже прошло. При http-ответе лучше заменять на код
	// 403 (bad request)
	ErrBadState = errors.New("state did not match or expired")
	// ErrMissingIDToken возвращается, когда в полученном авторизационном ответе
	// нет токена идентификации
	ErrMissingIDToken = errors.New("missing identification token")
	// ErrNonce возвращается, когда в идентификационном токене nonce не
	// совпадает с заданным на сервере
	ErrNonce = errors.New("invalid identification token nonce")
)

// UserInfo запрашивает идентификационный токен, проверяет его валидности и,
// на основании полученных в нем данных, формирует информацию о пользователе.
//
// Так же вторым значением возвращается сохраненное в сессии значение
// вспомогательных данных.
func (p *Provider) UserInfo(ctx context.Context, state, code string) (*UserInfo, interface{}, error) {
	// проверяем сессионный ключ и получаем данные о сессии
	stateObj, ok := p.states.Load(state)
	if !ok || stateObj.Created.Before(time.Now().Add(-StateTTL)) {
		return nil, nil, ErrBadState
	}
	// копируем конфигурацию для авторизации и добавляем адрес для редиректа
	var cfg = p.oauth2Config
	cfg.RedirectURL = stateObj.RedirectURI
	// получаем токен авторизации OAuth2
	oauth2Token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, nil, fmt.Errorf("error receiving authorization token: %w", err)
	}
	// получаем идентификационный токен из OAuth2 токена
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return nil, nil, ErrMissingIDToken
	}
	// проверяем полученный идентификационный токен
	var verifier = p.provider.Verifier(&oidc.Config{
		ClientID: cfg.ClientID,
	})
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, nil, fmt.Errorf("identification token verification error: %w", err)
	}
	// проверяем nonce, если он задан на сервере
	if ServerNonce != "" && idToken.Nonce != ServerNonce {
		return nil, nil, ErrNonce
	}
	// формируем информацию о пользователе
	var userInfo = &UserInfo{
		Issuer:      p.name,
		Subject:     idToken.Subject,
		provider:    p.provider,
		oauth2Token: oauth2Token,
		data:        stateObj.Data,
	}
	// заполняем поля профиля пользователя из информации идентификационного токена
	if err = idToken.Claims(userInfo); err != nil {
		return nil, nil, fmt.Errorf("identification token parsing error: %w", err)
	}
	return userInfo, stateObj.Data, nil
}
