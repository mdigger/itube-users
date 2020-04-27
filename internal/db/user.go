package db

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
)

// UserInfo описывает данные о пользователе.
//
// Расширенные аттрибуты, чтобы избежать множественных перекодировок из/в
// json, представлены в виде строки.
type UserInfo struct {
	UID        string    `json:"uid"`                      // уникальный идентификатор
	Email      string    `json:"email"`                    // email-адрес
	Verified   bool      `json:"email_verified,omitempty"` // флаг, что email-адрес подтвержден
	Properties *string   `json:"properties,omitempty"`     // дополнительные атрибуты
	Updated    time.Time `json:"updated"`                  // дата и время последнего обновления
}

// scanUser разбирает полученные данные о пользователе.
func scanUser(row pgx.Row, fields ...interface{}) (*UserInfo, error) {
	var (
		user    = new(UserInfo)
		blocked bool
		params  = []interface{}{
			&user.UID,
			&user.Email,
			&user.Properties,
			&user.Updated,
			&blocked,
			&user.Verified,
		}
	)
	if len(fields) > 0 {
		params = append(params, fields...)
	}
	err := row.Scan(params...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if blocked {
		return nil, ErrBlocked
	}
	return user, nil
}
