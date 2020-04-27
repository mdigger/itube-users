package db

import (
	"github.com/jackc/pgconn"
)

// используется для сохранения пустых строк в базе данных как NULL.
func null(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

// oneRow проверяет, что sql-команда применена к одной записи в базе данных.
// Возвращает ErrNotFound, если запрос выполнен без ошибок, но не "зацепил" ни
// одной записи.
func oneRow(ct pgconn.CommandTag, err error) error {
	if err != nil {
		return err
	}
	// пользователь не найден - возвращаем ошибку
	if ct.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}
