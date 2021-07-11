package server_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/najork/apollo-server/config"
	"github.com/najork/apollo-server/server"
	"github.com/palantir/conjure-go-runtime/v2/conjure-go-client/httpclient"
	"github.com/palantir/pkg/httpserver"
	"github.com/palantir/witchcraft-go-logging/wlog/wapp"
	wconfig "github.com/palantir/witchcraft-go-server/v2/config"
	"github.com/palantir/witchcraft-go-server/v2/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	port, err := httpserver.AvailablePort()
	require.NoError(t, err)
	server := server.New().
		WithInstallConfig(config.Install{
			Install: wconfig.Install{
				ProductName: "apollo-server",
				Server: wconfig.Server{
					Address:        "localhost",
					ContextPath:    "/apollo",
					Port:           port,
					ManagementPort: port,
				},
				UseConsoleLog: true,
			},
		}).
		WithRuntimeConfig(config.Runtime{}).
		WithSelfSignedCertificate()
	go func() {
		if err := wapp.RunWithFatalLogging(context.Background(), func(ctx context.Context) error {
			return server.Start()
		}); err != nil {
			fmt.Println("server failed:", err)
		}
	}()
	defer func() {
		if err := server.Close(); err != nil {
			fmt.Println("failed to close server:", err)
		}
	}()

	client, err := httpclient.NewHTTPClient(
		httpclient.WithHTTPTimeout(5*time.Second),
		httpclient.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	require.NoError(t, err)

	assert.True(t, <-httpserver.Ready(func() (*http.Response, error) {
		return client.Get(fmt.Sprintf("https://localhost:%d/%s/%s", port, "apollo", status.LivenessEndpoint))
	}))
}
