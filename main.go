package main

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/route"
	"abduselam-arabianmejlis/domain"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	// Create text index for search
	bootstrap.CreateTextIndex(db, domain.ProductsCollection)

	// if app.Redis == nil {
	// 	panic("Redis is not connected")
	// }
	// redisClient := app.Redis
	// defer app.Close()

	// start the client manager routine
	clientManager := bootstrap.NewClientManager()
	go clientManager.Start()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	// gin.SetMode(gin.ReleaseMode)
	
	router := gin.Default()

	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour

    router.Use(cors.New(config))

	route.Setup(env, timeout, db, router, nil, clientManager)

	// Setup WebSocket routes using Gin
	// setupRoutes(router)

	router.Run(env.ServerAddress) 
}
