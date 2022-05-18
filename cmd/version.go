package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jieqiserver/consts"
)

func NewVersionCommand() *cobra.Command {
	var short bool
	cmd := cobra.Command{
		Use:   "version",
		Short: "print version information",
		Run: func(cmd *cobra.Command, args []string) {
			PrintVersion(CLIName, consts.GetVersion(), short)
		},
	}
	cmd.Flags().BoolVar(&short, "short", false, "print just the version number")
	return &cmd
}

func PrintVersion(cliName string, version consts.Version, short bool) {
	fmt.Printf("%s: %s\n", cliName, version.Version)
	if short {
		return
	}
	fmt.Printf("  BuildDate: %s\n", version.BuildDate)
	fmt.Printf("  GitCommit: %s\n", version.GitCommit)
	fmt.Printf("  GitTreeState: %s\n", version.GitTreeState)
	if version.GitTag != "" {
		fmt.Printf("  GitTag: %s\n", version.GitTag)
	}
	fmt.Printf("  GoVersion: %s\n", version.GoVersion)
	fmt.Printf("  Compiler: %s\n", version.Compiler)
	fmt.Printf("  Platform: %s\n", version.Platform)
}
