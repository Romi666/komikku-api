package handlers

import (
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"komiku-srapper/bin/config"
	"komiku-srapper/bin/modules/manga/repositories/queries"
	"komiku-srapper/bin/modules/manga/usecases"
	"komiku-srapper/bin/pkg/utils"
	"net/http"
)

type MangaHandler struct {
	mangaCommandUsecase	usecases.MangaUsecase
}

func New() *MangaHandler {
	collector := colly.NewCollector()
	queryManga := queries.NewMangaQuery(config.GlobalEnv.BaseURL, collector)
	queryUsecase := usecases.CreateNewMangaUsecase(queryManga)
	return &MangaHandler{
		mangaCommandUsecase: queryUsecase,
	}
}

//Mount function
func(m *MangaHandler) Mount(router *echo.Echo) {
	router.GET("/comic/list", m.GetAllComic)
}

func(m *MangaHandler) GetAllComic(c echo.Context) error {
	result := m.mangaCommandUsecase.GetAllComic()
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get All Comic", http.StatusOK, c)
}