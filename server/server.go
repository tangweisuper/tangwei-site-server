package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"jieqiserver/config"
	"jieqiserver/core"
	"time"
)

type JieqiServer struct {
	Conf *config.Config
	App  *fiber.App
}

func (s *JieqiServer) Init() {
	s.App.Use("/", func(ctx *fiber.Ctx) error {
		prefix := s.getPrefix()
		if prefix != "" {
			url := string(ctx.Request().RequestURI())
			newUrl := "/" + prefix + url
			zap.S().Info(fmt.Sprintf("%s   =>   %s", url, newUrl))
			ctx.Request().SetRequestURI(newUrl)
		} else {
			zap.S().Warn("get no prefix")
		}
		return ctx.Next()
	})

	s.App.Static("/", s.Conf.DocumentRoot, fiber.Static{
		Compress:      false,
		ByteRange:     true,
		Browse:        false,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
	})
}

func (s *JieqiServer) getPrefix() string {
	jieqi := core.GetJieqi()
	return s.Conf.Route[jieqi]
}

func (s *JieqiServer) Start() error {
	err := s.App.Listen(fmt.Sprintf(":%d", s.Conf.Port))
	return err
}

func NewServer(conf *config.Config) *JieqiServer {
	server := &JieqiServer{
		Conf: conf,
		App:  fiber.New(),
	}

	server.Init()

	return server
}
