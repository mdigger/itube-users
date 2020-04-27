//go:generate ../../scripts/protogen.sh

package rpc

import (
	"context"
	"io"
	"itube/users/internal/db"
	"itube/users/pkg/api"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// проверка, что сервис поддерживает все методы сервиса
var _ api.IdentityServer = new(Identity)

// Identity реализует grpc-сервис для авторизации и регистрации пользователей.
type Identity struct {
	db *db.Adapter
}

// NewIdentity инициализирует и возвращает серверный обработчик grpc для авторизации
// пользователей.
func NewIdentity(db *db.Adapter) *Identity {
	return &Identity{db: db}
}

// Register регистрирует и возвращает информацию о пользователе.
//
// Если пользователь уже зарегистрирован, но пароль для него не установлен, то
// устанавливает новый пароль. Данный случай возникает, если до этого
// пользователь был зарегистрирован через внешнего провайдера.
//
// Возвращает ошибки:
//  - AlreadyExists - пользователь уже зарегистрирован и у него задан пароль
//  - NotFound - пользователь заблокирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Identity) Register(ctx context.Context, req *api.Login) (*api.User, error) {
	user, err := s.db.Register(ctx, req.Email, req.Password)
	if err != nil {
		return nil, statusError(err)
	}
	// добавляем в журнал запись о регистрации (возможную ошибку игнорируем)
	_ = s.db.RegInfo(ctx, req.Domain, user.UID, user.Email, "",
		req.RegInfo.Referer, utm(req.RegInfo.UTM))
	return apiUser(req.Domain, user) // возвращаем информацию о пользователе
}

// Authorize авторизует пользователя по логину (email) и паролю. Возвращает
// информацию о пользователе в случае успешной авторизации. В противном случае
// возвращает ошибку.
//
// Возвращает ошибки:
//  - NotFound - пользователь не зарегистрирован или блокирован
//  - InvalidArgument - неверный пароль пользователя
//  - Internal - внутренние ошибки
func (s *Identity) Authorize(ctx context.Context, req *api.Login) (*api.User, error) {
	user, err := s.db.Authorize(ctx, req.Email, req.Password)
	if err != nil {
		return nil, statusError(err)
	}
	return apiUser(req.Domain, user) // возвращаем информацию о пользователе
}

// SetPassword заменяет пароль пользователя. Возвращает ошибку, если
// пользователь не зарегистрирован.
//
// Возвращает ошибки:
//  - NotFound - пользователь не зарегистрирован
//  - Internal - внутренние ошибки
func (s *Identity) SetPassword(ctx context.Context, req *api.Password) (*types.Empty, error) {
	err := s.db.SetPassword(ctx, req.UID, req.Password)
	if err != nil {
		return nil, statusError(err)
	}
	return new(types.Empty), nil
}

// Update обновляет информацию о пользователе. Возвращает ошибку,
// если пользователь не зарегистрирован. Информация, что email проверен, а
// так же дата обновления игнорируется.
//
// Возвращает ошибки:
//  - AlreadyExists - пользователь с таким email уже зарегистрирован
//  - NotFound - пользователь не зарегистрирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Identity) Update(ctx context.Context, req *api.User) (*types.Empty, error) {
	var err error
	// преобразуем формат расширенных свойств
	var properties string
	if req.Properties != nil {
		properties, err = new(jsonpb.Marshaler).MarshalToString(req.Properties)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"properties error: %s", err)
		}
	}
	err = s.db.Update(ctx, req.UID, req.Email, properties)
	if err != nil {
		return nil, statusError(err)
	}
	return new(types.Empty), nil
}

// Block используется для блокировки/разблокировки пользователя.
// Заблокированный пользователь продолжает оставаться зарегистрированных,
// но не может авторизоваться.
//
// Возвращает ошибки:
//  - NotFound - пользователь не зарегистрирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Identity) Block(ctx context.Context, req *api.BlockID) (*types.Empty, error) {
	err := s.db.BlockUser(ctx, req.UID, req.Blocked)
	if err != nil {
		return nil, statusError(err)
	}
	return new(types.Empty), nil
}

// Get возвращает информацию о пользователе по идентификатору или email.
//
// Возвращает ошибки:
//  - NotFound - пользователь не зарегистрирован
//  - InvalidArgument - неверный формат данных входящего запроса
//  - Internal - внутренние ошибки
func (s *Identity) Get(ctx context.Context, req *api.UserID) (*api.User, error) {
	user, err := s.db.GetUser(ctx, req.GetUID(), req.GetEmail())
	if err != nil {
		return nil, statusError(err)
	}
	return apiUser(req.Domain, user)
}

// List возвращает информацию о пользователях по идентификатору или email.
// Используется для получения информации о других пользователях в потоке.
func (s *Identity) List(stream api.Identity_ListServer) error {
	var ctx = stream.Context()
	for {
		var in, err = stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// получаем информацию о пользователе
		user, err := s.db.GetUser(ctx, in.GetUID(), in.GetEmail())
		if err != nil {
			return statusError(err)
		}
		// преобразуем данные о пользователе, если они есть
		result, err := apiUser(in.Domain, user)
		if err != nil {
			return err
		}
		// отправляем ответ
		err = stream.Send(result)
		if err != nil {
			return err
		}
	}
}
