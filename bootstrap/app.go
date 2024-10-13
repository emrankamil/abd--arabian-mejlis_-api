package bootstrap

import (
	"abduselam-arabianmejlis/mongo"

	"abduselam-arabianmejlis/redis"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Redis redis.Client
} 

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	// app.Redis = redis.NewClient(app.Env.REDISADDR)
	app.Redis = nil
	return *app
}

func (app *Application) Close() {
	CloseMongoDBConnection(app.Mongo)
	// app.Redis.Close()
}
