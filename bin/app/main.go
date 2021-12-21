package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"komikku-api/bin/config"
	chapterH "komikku-api/bin/modules/chapter/handlers"
	mangaH "komikku-api/bin/modules/manga/handlers"
	"komikku-api/bin/pkg/middlewares"
)

func main() {
	//Echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middlewares.Logger())

	comicHttp := mangaH.New()
	comicHttp.Mount(e)

	chapterHttp :=chapterH.New()
	chapterHttp.Mount(e)

	listenerPort := fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}