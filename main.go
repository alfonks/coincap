package main

import (
	"coincap/internal/initmodule"
	"coincap/internal/rest_api"
	"coincap/pkg/cfg"
	"context"
)

func main() {
	ctx := context.Background()
	config := cfg.Init()

	dep := initmodule.InitAppServer(ctx, config)

	ginRouter := rest_api.Router(dep, config)

	ginRouter.Run(config.Server.Address)
}
