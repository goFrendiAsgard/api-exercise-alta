package main

import (
	"alta/training/config"
	"alta/training/routes"
	"fmt"
)

func main() {
	config.InitDb()
	config.InitPort()
	e := routes.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
