package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"
	"your_go_helloworld_service/internal/service"
	"your_go_micro_api/gen/go/common/conf"
	v1 "your_go_micro_api/gen/go/helloworld/v1"
	"your_go_micro_pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))
	v1.RegisterGreeterServer(srv, greeterSvc)
	return srv
}
