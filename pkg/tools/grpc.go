package tools

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// InitGRPCServer возвращает инициализированный сервис grpc, в который добавлены
// всякие "прокладки" для логгирования, восстановления ошибок и прочее.
// Если логгер не задан, то используется логгер по умолчанию.
func InitGRPCServer(log *logrus.Entry, opts ...grpc_logrus.Option) *grpc.Server {
	// инициализируем логгер, если он не задан
	if log != nil {
		log = logrus.NewEntry(logrus.StandardLogger())
	}
	grpc_logrus.ReplaceGrpcLogger(log)
	// инициализируем новый grpc сервер, добавляя в него всякие "плюшки"
	var grpcServer = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.
				WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(log, opts...),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.
				WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(log, opts...),
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	// добавляем сервис проверки "живости"
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// добавляет отдачу информацию от поддерживаемых методах и параметрах
	reflection.Register(grpcServer)
	return grpcServer
}
