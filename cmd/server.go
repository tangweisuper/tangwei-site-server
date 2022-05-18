package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jieqiserver/config"
	"jieqiserver/server"
)

func NewServerCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "server",
		Short: "start server",
		Run: func(cmd *cobra.Command, args []string) {
			zap.S().Info("initializing registry server ...")
			conf, err := config.MakeConfig(viper.GetString("config-file"))
			if err != nil {
				zap.S().Fatal("Init service failed.", zap.Error(err))
			}

			ctx := context.Background()
			_, cancel := context.WithCancel(ctx)

			s := server.NewServer(conf)

			if s == nil {
				zap.S().Fatal("Init service failed.", zap.Error(err))
				cancel()
				return
			}

			err = s.Start()
			if err != nil {
				cancel()
				return
			}
			cancel()
		},
	}

	return &cmd
}
