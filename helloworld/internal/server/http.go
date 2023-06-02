package server

import (
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/http"
	"your_go_helloworld_service/internal/service"
	"your_go_micro_api/gen/go/common/conf"
	v1 "your_go_micro_api/gen/go/helloworld/v1"
	"your_go_micro_pkg/bootstrap"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	greeterSvc service.GreeterService,
) *http.Server {
	srv := bootstrap.CreateHttpServer(cfg, logging.Server(logger))
	v1.RegisterGreeterHTTPServer(srv, greeterSvc)
	return srv
}
