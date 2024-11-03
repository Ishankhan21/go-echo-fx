package main

import (
	"github.com/Ishankhan21/go-fiber-fx/controllers"
	"github.com/Ishankhan21/go-fiber-fx/db"
	"github.com/Ishankhan21/go-fiber-fx/services"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			db.GetRedisconnection,
			db.GetMongoConnection,
			services.NewUserService,
			services.NewOrderService,
			controllers.NewUserController),
		fx.Invoke(StartEchoServer)).Run()
}
