package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jieqiserver/cmd"
	"math/rand"
	"os"
	"time"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	logger, err := config.Build()
	if err != nil {
		return
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	defer func() {
		if err := recover(); err != nil {
			zap.S().Error(err)
			os.Exit(1)
		}
	}()

	c := cmd.InitApp()
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
