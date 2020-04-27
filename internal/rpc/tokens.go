package rpc

import (
	"context"
	"itube/users/internal/db"
	"itube/users/pkg/api"
)

// проверка, что сервис поддерживает все методы сервиса
var _ api.TokensServer = new(Tokens)

// Tokens реализует grpc-сервис для генерации и проверки токенов для сброса
// пароля пользователя или подтверждения почтового адреса.
type Tokens struct {
	db *db.Adapter
}

// Generate создает запрос для проверки адреса email пользователя или
// сброса пароля. При вызове сервер отправляет соответствующее письмо
// на email адрес пользователя с токеном для верификации.
// Повторный вызов с теми же значениями параметров заменяет токен на новый,
// а действие старого отменяет.
//
// Возвращает ошибки:
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Tokens) Generate(ctx context.Context, req *api.VerifyRequest) (*api.TokenInfo, error) {
	token, err := s.db.TokenGenerate(ctx, req.Domain, req.Email, int32(req.Type))
	if err != nil {
		return nil, statusError(err)
	}
	return &api.TokenInfo{
		Domain: req.Domain,
		Token:  token,
		Type:   req.Type,
	}, nil
}

// Verify проверяет токен и возвращает зарегистрированного пользователя.
// Если токен неверен, то возвращается ошибка NotFound. После проверки
// токен автоматически удаляется и повторное его использование невозможно.
//
// Так же автоматически подтверждает почтовый адрес, через который был
// отправлен данный токен.
//
// Возвращает ошибки:
//  - NotFound - пользователь не зарегистрирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Tokens) Verify(ctx context.Context, req *api.TokenInfo) (*api.User, error) {
	user, err := s.db.TokenVerify(ctx, req.Token)
	if err != nil {
		return nil, statusError(err)
	}
	return apiUser(req.Domain, user)
}
