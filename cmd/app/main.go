package main

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/api"
	"github.com/kerim-dauren/smart-terminal/config"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()

	application := &api.App{}

	if err := application.Start(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}
