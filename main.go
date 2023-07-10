package main

import (
	"alta.id/go-skeleton/config"
	"alta.id/go-skeleton/middlewares"
	"alta.id/go-skeleton/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
