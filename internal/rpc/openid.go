package rpc

import (
	"context"
	"errors"
	"itube/users/internal/db"
	"itube/users/pkg/api"
	"itube/users/pkg/openid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// проверка, что сервис поддерживает все методы сервиса
var _ api.OpenIDServer = new(OpenID)

// OpenID реализует grpc-сервис для авторизации и регистрации пользователей
// с помощью внешних провайдеров авторизации по протоколу OpenID Connect.
type OpenID struct {
	db        *db.Adapter
	providers map[string]*openid.Provider // провайдеры авторизации
}

// NewOpenID возвращает инициализированный сервис для авторизации.
func NewOpenID(db *db.Adapter, providers ...*openid.Provider) *OpenID {
	var list = make(map[string]*openid.Provider, len(providers))
	for _, provider := range providers {
		list[provider.String()] = provider
	}
	return &OpenID{
		db:        db,
		providers: list,
	}
}

// Login выдает URL для перехода на авторизацию к провайдеру.
//
// Возвращает InvalidArgument, если указан неподдерживаемый идентификатор
// провайдера авторизации.
func (s *OpenID) Login(ctx context.Context, req *api.Provider) (*api.LoginURL, error) {
	// получаем провайдера, ответственного за авторизацию
	provider, ok := s.providers[req.Provider]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument,
			"unsupported provider: %s", req.Provider)
	}
	// формируем url для перехода на авторизацию
	var loginURL = provider.LoginURL(req.RedirectURI, req.Params, req.RegInfo)
	return &api.LoginURL{Domain: req.Domain, URL: loginURL}, nil
}

// Authorize проверяет авторизацию и возвращает информацию об
// авторизованном пользователе. Если пользователь не зарегистрирован,
// то происходит его автоматическая регистрация.
//
// Возвращает ошибки:
//  - NotFound - пользователь заблокирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *OpenID) Authorize(ctx context.Context, req *api.AuthCode) (*api.User, error) {
	// получаем провайдера, ответственного за авторизацию
	provider, ok := s.providers[req.Provider]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument,
			"unsupported provider: %s", req.Provider)
	}
	// запрашиваем информацию о пользователе у системы авторизации
	userinfo, data, err := provider.UserInfo(ctx, req.State, req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"openid authorization error: %s", err)
	}
	// запрашиваем из базы данные о пользователе
	user, err := s.db.OpenIDAuthorize(ctx, provider.String(), userinfo.Subject)
	if err == nil {
		return apiUser(req.Domain, user) // возвращаем информацию о пользователе
	}
	// произошла ошибка
	if !errors.Is(err, db.ErrNotFound) {
		return nil, statusError(err)
	}
	// пользователь не зарегистрирован - регистрируем
	user, err = s.db.OpenIDRegister(ctx, provider.String(), userinfo.Subject,
		userinfo.Email, userinfo.Verified, string(userinfo.JSON()))
	if err != nil {
		return nil, statusError(err)
	}
	// добавляем в журнал запись о регистрации (возможную ошибку игнорируем)
	var reginfo, _ = data.(api.RegInfo)
	_ = s.db.RegInfo(ctx, req.Domain, user.UID, user.Email, provider.String(),
		reginfo.Referer, utm(reginfo.UTM))
	return apiUser(req.Domain, user) // возвращаем информацию о пользователе
}
