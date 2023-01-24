package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hopers-backend/config"
	"hopers-backend/docs"
	"hopers-backend/routes"
	"log"
	"time"
)

func engine() *gin.Engine {
	gin.ForceConsoleColor()
	server := gin.New()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		docs.SwaggerInfo.BasePath = "/api/v1"

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	server.Use(gin.Recovery())
	api := server.Group("/api")
	{
		api.GET("/health", routes.Health)
	}

	DefiExperience := api.Group("/DeFi")
	{
		DefiExperience.GET("/coins/:token", routes.GetCoinsPrice)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return server
}

func GoGinServer() {
	server := engine()
	server.Use(gin.Logger())
	if err := engine().Run(":" + fmt.Sprint(config.Cfg.APIPort)); err != nil {
		log.Fatal("Unable to start:", err)
	}
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}
