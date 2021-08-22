package cmd

import (
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"
)

var (
	// Version of the program, set via ldflags by godel during build
	Version = "unspecified"

	rootCmd = &cobra.Command{
		Use: "health-sync-server",
	}
)

func Execute() int {
	return cobracli.ExecuteWithDefaultParams(rootCmd, cobracli.VersionFlagParam(Version))
}
