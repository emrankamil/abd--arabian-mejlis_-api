package main

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/route"
	"abduselam-arabianmejlis/domain"
	"time"

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

	router := gin.Default()
	route.Setup(env, timeout, db, router, nil, clientManager)

	// Setup WebSocket routes using Gin
	// setupRoutes(router)

	router.Run(env.ServerAddress) 
}
