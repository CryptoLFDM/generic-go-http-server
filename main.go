package main

import (
	"generic-http-server/cli"
	"generic-http-server/config"
	"generic-http-server/db"
	"generic-http-server/engine"
	"generic-http-server/server"
	"github.com/go-co-op/gocron"
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
