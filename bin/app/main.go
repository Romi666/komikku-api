package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"komiku-srapper/bin/config"
	chapterH "komiku-srapper/bin/modules/chapter/handlers"
	mangaH "komiku-srapper/bin/modules/manga/handlers"
)

func main() {
	//Echo instance
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	comicHttp := mangaH.New()
	comicHttp.Mount(e)

	chapterHttp :=chapterH.New()
	chapterHttp.Mount(e)

	listenerPort := fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}