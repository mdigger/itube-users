package rpc

import (
	"encoding/json"
	"errors"
	"itube/users/internal/db"
	"itube/users/pkg/api"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jackc/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// apiUser преобразует информацию о пользователе в grpc формат.
func apiUser(domain string, userInfo *db.UserInfo) (*api.User, error) {
	var user = &api.User{
		Domain:   domain,
		UID:      userInfo.UID,
		Email:    userInfo.Email,
		Verified: userInfo.Verified,
	}
	// время последнего обновления, если установлено
	if !userInfo.Updated.IsZero() {
		user.Updated = &userInfo.Updated
	}
	// добавляем расширенные свойства, если они определены
	if userInfo.Properties != nil {
		user.Properties = new(types.Struct)
		err := jsonpb.UnmarshalString(*userInfo.Properties, user.Properties)
		if err != nil {
			return nil, status.Errorf(codes.Internal,
				"user properties: %s", err)
		}
	}
	return user, nil
}

// statusError подменяет ошибку на grpc-ошибку. В зависимости от типа, может
// преобразовывать ее представление.
func statusError(err error) error {
	// подменяем стандартные ошибки
	switch err {
	case db.ErrAlreadyRegisterd:
		return status.Error(codes.AlreadyExists, err.Error())
	case db.ErrBadToken, db.ErrEmptyEmail, db.ErrInvalidPassword:
		return status.Error(codes.InvalidArgument, err.Error())
	case db.ErrBlocked, db.ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	}
	var dbErr = new(pgconn.PgError)
	if !errors.As(err, &dbErr) {
		// другой тип ошибки
		return status.Error(codes.Internal, err.Error())
	}
	// разбираем ошибки базы данных
	if len(dbErr.Code) > 2 {
		switch dbErr.Code[:2] {
		case "22", "23": // Data Exception & Integrity Constraint Violation
			return status.Error(codes.InvalidArgument, dbErr.Message)
		}
	}
	// все равно для базы "упрощаем" стандартное описание ошибки
	return status.Error(codes.Internal, dbErr.Message)
}

// utm преобразует маркетинговую информацию в строку.
func utm(utm map[string]string) string {
	if len(utm) == 0 {
		return ""
	}
	var result string
	data, err := json.Marshal(utm)
	// игнорируем ошибку преобразования в формат json
	if err == nil {
		result = string(data)
	}
	return result
}
