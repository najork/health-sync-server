module github.com/najork/health-sync-server

go 1.16

require (
	github.com/abrander/garmin-connect v0.0.0-00010101000000-000000000000
	github.com/palantir/conjure-go-runtime/v2 v2.17.0
	github.com/palantir/pkg/cobracli v1.0.1
	github.com/palantir/pkg/httpserver v1.0.1
	github.com/palantir/pkg/safejson v1.0.1
	github.com/palantir/pkg/safeyaml v1.0.1
	github.com/palantir/pkg/signals v1.0.1
	github.com/palantir/witchcraft-go-error v1.5.0
	github.com/palantir/witchcraft-go-health v1.8.0
	github.com/palantir/witchcraft-go-logging v1.14.0
	github.com/palantir/witchcraft-go-server/v2 v2.19.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.30.0
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	go.uber.org/atomic v1.7.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/abrander/garmin-connect => github.com/najork/garmin-connect v0.0.0-20210822201646-f07485200955
