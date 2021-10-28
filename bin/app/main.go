package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"komiku-srapper/bin/config"
	"komiku-srapper/bin/modules/manga/handlers"
)

func main() {
	//Echo instance
	e := echo.New()

	e.Use(middleware.CORS())

	comicHttp := handlers.New()
	comicHttp.Mount(e)

	listenerPort := fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}