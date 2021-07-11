package server

import (
	"context"

	"github.com/najork/apollo-server/config"
	"github.com/najork/apollo-server/conjure/api"
	"github.com/najork/apollo-server/service"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"github.com/palantir/witchcraft-go-logging/wlog/wapp"
	"github.com/palantir/witchcraft-go-server/v2/witchcraft"
)

// New creates and returns a witchcraft Server.
func New() *witchcraft.Server {
	return witchcraft.NewServer().
		WithInstallConfigType(config.Install{}).
		WithRuntimeConfigType(config.Runtime{}).
		WithECVKeyProvider(witchcraft.ECVKeyNoOp()).
		WithInitFunc(initFn).
		WithOrigin(svc1log.CallerPkg(0, 1))
}

func initFn(ctx context.Context, info witchcraft.InitInfo) (cleanup func(), rErr error) {
	svc := service.New()
	go func() {
		_ = wapp.RunWithFatalLogging(ctx, func(ctx context.Context) error {
			return api.RegisterRoutesApolloService(info.Router, svc)
		})
	}()

	info.Router.WithHealth(svc)
	return nil, nil
}
