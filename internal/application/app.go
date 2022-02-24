package application

import (
	"context"

	"exemplo.com/desafioGo/assets/rpc"
	"exemplo.com/desafioGo/config"
)

type App struct {
	Config        *config.Config
	EmpresaServer *rpc.TwirpServer
}

func New(ctc context.Context, c *config.Config, s *rpc.TwirpServer) *App {
	return &App{
		Config:        c,
		EmpresaServer: s,
	}
}

func NewTest() *App {
	return &App{}
}
