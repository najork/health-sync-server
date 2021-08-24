package server

import (
	"context"

	connect "github.com/abrander/garmin-connect"
	"github.com/najork/health-sync-server/config"
	"github.com/najork/health-sync-server/conjure/api"
	"github.com/najork/health-sync-server/service"
	werror "github.com/palantir/witchcraft-go-error"
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
		WithSelfSignedCertificate().
		WithInitFunc(initFn).
		WithOrigin(svc1log.CallerPkg(0, 1))
}

func initFn(ctx context.Context, info witchcraft.InitInfo) (cleanup func(), rErr error) {
	install, ok := info.InstallConfig.(config.Install)
	if !ok {
		return nil, werror.ErrorWithContextParams(ctx, "install config type assertion failed")
	}

	client := connect.NewClient(connect.Credentials(install.Credentials.Email, install.Credentials.Password))
	svc := service.New(client)
	go func() {
		_ = wapp.RunWithFatalLogging(ctx, func(ctx context.Context) error {
			return svc.Start(ctx)
		})
	}()

	go func() {
		_ = wapp.RunWithFatalLogging(ctx, func(ctx context.Context) error {
			return api.RegisterRoutesHealthSyncService(info.Router, svc)
		})
	}()

	info.Router.WithReadiness(svc)
	info.Router.WithHealth(svc)
	return nil, nil
}
