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
	router.GET("/comic/:endpoint", m.GetComicInfo)
	router.GET("/comic/chapter/:endpoint", m.GetChapterDetail)
	router.GET("/comic/search", m.SearchManga)
	router.GET("/comic/genre", m.GetAllGenre)
}

func(m *MangaHandler) GetAllComic(c echo.Context) error {
	result := m.mangaCommandUsecase.GetAllComic()
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get All Comic", http.StatusOK, c)
}

func(m *MangaHandler) GetComicInfo(c echo.Context) error {
	endpoint := c.Param("endpoint")
	result := m.mangaCommandUsecase.GetComicInfo(endpoint)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get Comic Info", http.StatusOK, c)
}

func(m *MangaHandler) GetChapterDetail(c echo.Context) error {
	endpoint := c.Param("endpoint")
	result := m.mangaCommandUsecase.GetChapterDetail(endpoint)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get Chapter Detail", http.StatusOK, c)
}

func(m *MangaHandler) SearchManga(c echo.Context) error {
	query := c.QueryParam("query")
	result := m.mangaCommandUsecase.SearchManga(query)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Search manga", http.StatusOK, c)
}

func (m *MangaHandler) GetAllGenre(c echo.Context) error {
	result := m.mangaCommandUsecase.GetAllGenre()

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get all genre", http.StatusOK, c)
}