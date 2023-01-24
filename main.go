package main

import (
	"github.com/go-co-op/gocron"
	"hopers-backend/cli"
	"hopers-backend/config"
	"hopers-backend/db"
	"hopers-backend/engine"
	"hopers-backend/server"
	"time"
)

func main() {
	cliFilled := cli.Cli()
	config.LoadYamlConfig(cliFilled.FilePathConfig)
	//redis.InitRedis()
	db.Migrate()
	go func() {
		s := gocron.NewScheduler(time.Local)
		s.Every(1).Minutes().Do(engine.HarvestCoinPrice)
		s.StartAsync()
	}()
	server.GoGinServer()
}
