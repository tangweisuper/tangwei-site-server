package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jieqiserver/consts"
)

const (
	// CLIName is the name of the CLI
	CLIName = "jieqiserver"
)

func InitApp() *cobra.Command {
	var version = consts.GetVersion()
	rootCmd := &cobra.Command{
		Use:     CLIName,
		Short:   "jieqi dynamic route server",
		Version: fmt.Sprintf("%s", version.Version),
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				zap.S().Fatal(zap.Error(err))
			}
		},
	}

	rootCmd.PersistentFlags().StringP("config-file", "c", "config.yaml", "配置文件")
	_ = viper.BindPFlag("config-file", rootCmd.PersistentFlags().Lookup("config-file"))

	rootCmd.AddCommand(NewServerCommand())
	rootCmd.AddCommand(NewVersionCommand())
	return rootCmd
}
