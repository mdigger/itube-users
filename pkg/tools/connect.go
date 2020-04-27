package tools

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	// DefaultDBName содержит имя базы данных, используемой по умолчанию, если
	// не задано другого имени в конфигурации для соединения.
	DefaultDBName = "itube_users"
	// ConnectionTimeout задает максимальное время ожидания подключения к
	// базе данных для тестирования подключения.
	ConnectionTimeout = time.Second * 10
	// Logger задает логгер для базы данных.
	Logger logrus.FieldLogger
	// LogLevel определяет уровень по умолчанию для вывода в лог.
	LogLevel = pgx.LogLevelWarn
)

// Connect подключается к базе данных и возвращает пулл соединений с ней.
// Так же осуществляется проверка реального подключения к базе данных.
//
// Если задан Logger, то он будет автоматически подключен для ведения логов
// работы с базой данных.
func Connect(dsn string) (*pgxpool.Pool, error) {
	// разбираем параметры подключения
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	// подставляем базу данных по умолчанию, если она не определена
	var cfg = config.ConnConfig
	if cfg.Database == "" {
		cfg.Database = DefaultDBName
	}
	// подключаем логгирование действий с базой данных
	if Logger != nil {
		cfg.Logger = logrusadapter.NewLogger(Logger)
		cfg.LogLevel = (pgx.LogLevel)(LogLevel) // ограничиваем вывод записей в лог
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	// проверяем реальное подключение к базе данных
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout)
	defer cancel()
	pconn, err := pool.Acquire(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}
	defer pconn.Release()
	if err = pconn.Conn().Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	// выводим в лог информацию о подключении
	if Logger != nil {
		Logger.WithFields(logrus.Fields{
			"host": cfg.Host,
			"db":   cfg.Database,
			"user": cfg.User,
		}).Info("db connected")
	}
	return pool, nil
}
