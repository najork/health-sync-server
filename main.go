package main

import (
	"os"

	"github.com/najork/health-sync-server/cmd"
	"github.com/palantir/pkg/signals"
)

func main() {
	signals.RegisterStackTraceWriter(os.Stderr, nil)
	os.Exit(cmd.Execute())
}
