package data

import (
	"github.com/devexps/go-micro/v2/log"
	"your_project/api/gen/go/common/conf"
)

// Data .
type Data struct {
	log *log.Helper

	// TODO wrapped database client
}

// NewData .
func NewData(cfg *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log: l,
	}, cleanup, nil
}
