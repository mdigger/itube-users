package main

import (
	"context"
	"fmt"
	"itube/users/internal/db"
	"itube/users/internal/rpc"
	"itube/users/internal/sender"
	"itube/users/pkg/api"
	"itube/users/pkg/email"
	"itube/users/pkg/openid"
	"itube/users/pkg/tools"
	"net"
	"time"

	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
)

var (
	// DefaultPort задает порт по умолчанию для gRPC сервера.
	DefaultPort = 50051
	// SMTPSleep определяет время ожидания между проверками новых токенов для
	// отправки.
	SMTPSleep = time.Minute * 5
)

func init() {
	flag.EnvironmentPrefix = "ITUBE_USERS"
	flag.String(flag.DefaultConfigFlagname, "", "path to config file")
}

func main() {
	// разбираем параметры для запуска сервиса
	var (
		logLevel = flag.String("log_level", "info", "logging level")
		port     = flag.Int("port", DefaultPort, "grpc port")
		dsn      = flag.String("dsn", "postgres://postgres@localhost/"+
			tools.DefaultDBName, "postgres DSN")
		googleClientID = flag.String("google_client_id", "", "google client id")
		googleSecret   = flag.String("google_secret", "", "google secret")
		smtp           = flag.String("smtp", "", "smtp server url")
		tmpltsPath     = flag.String("templates", "../templates/emails.yaml",
			"file with email templates")
	)
	flag.Parse()
	// устанавливаем уровень логирования
	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	// tools.LogLevel = pgx.LogLevelTrace
	// подключаемся к базе данных
	tools.Logger = log.WithField("system", "pgx")
	pool, err := tools.Connect(*dsn)
	if err != nil {
		log.WithError(err).Fatal("error connecting to PostgreSQL database")
	}
	defer pool.Close()
	// прослойка для работы с базой данных
	var adapter = &db.Adapter{Pool: pool}
	// инициализируем провайдера авторизации Google
	googleProvider, err := openid.NewGoogle(*googleClientID, *googleSecret)
	if err != nil {
		log.WithError(err).Fatal("authorization provider initialization error")
	}
	log.WithFields(log.Fields{
		"provider": googleProvider.String(),
		"clientID": *googleClientID,
	}).Info("authorization provider initialized")
	// инициализируем порт для grpc
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.WithError(err).Fatal("init grpc listener port error")
	}
	defer listener.Close()
	// регистриуем grpc сервисы
	var grpcServer = tools.InitGRPCServer(log.WithField("module", "grpc"))
	api.RegisterIdentityServer(grpcServer, rpc.NewIdentity(adapter))
	api.RegisterOpenIDServer(grpcServer, rpc.NewOpenID(adapter, googleProvider))
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.WithError(err).Error("grpc server error")
		}
	}()
	log.WithField("port", *port).Infof("grpc server started")

	// инициализируем почтовые шаблоны
	mailTemplates, err := email.Init(*smtp, *tmpltsPath)
	if err != nil {
		log.WithError(err).Fatal("email templates initializing error")
	}
	log.WithField("domains", mailTemplates.Domains()).Info("email templates initialized")
	// запускаем обработчик для отправки почтовых сообщений с токенами
	var sender = sender.New(adapter, mailTemplates)
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	go func() {
		var timer = time.NewTimer(SMTPSleep)
		defer timer.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C: // вызываем отправку токенов интрервалами
				err := sender.Send(ctx)
				if err != nil {
					log.WithError(err).Error("error sending tokens")
				}
				timer.Reset(SMTPSleep) // повторить
			}
		}
	}()

	// завершение работы по сигналу прерывания
	var sig = tools.WaitSignal() // ожидание сигнала о прерывании
	log.WithField("signal", sig.String()).Infof("interrupt received")
	grpcServer.GracefulStop() // останавливаем gRPC сервер
	log.Info("service finished its work")
}
