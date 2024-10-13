package route

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/middleware"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, redisClient redis.Client, cm *bootstrap.ClientManager) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewFogetPWRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewProductRouter(env, timeout, db, publicRouter, redisClient)
	NewOrderRouter(env, timeout, db, publicRouter)
	NewImageUploadRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewLogoutRouter(env, timeout, db, protectedRouter)
	NewPromoteRouter(env, timeout, db, protectedRouter)

	//websocket router
	wsRouter := gin.Group("")
	NewChatRouter(env, timeout, db, cm, wsRouter)

	// static file server
	gin.Static("/uploads", "./uploads")
}
