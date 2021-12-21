package handlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"komiku-srapper/bin/config"
	chapterQ "komiku-srapper/bin/modules/chapter/repositories/queries"
	mangaQ "komiku-srapper/bin/modules/manga/repositories/queries"
	mangaU "komiku-srapper/bin/modules/manga/usecases"
	httpError "komiku-srapper/bin/pkg/http-error"
	"komiku-srapper/bin/pkg/utils"
	"net/http"
	"strconv"
)

type MangaHandler struct {
	mangaCommandUsecase	mangaU.MangaUsecase
}

func New() *MangaHandler {
	collector := colly.NewCollector()
	queryManga := mangaQ.NewMangaQuery(config.GlobalEnv.BaseURL, collector)
	queryChapter := chapterQ.NewChapterQuery(config.GlobalEnv.BaseURL, collector)
	queryUsecase := mangaU.CreateNewMangaUsecase(queryManga, queryChapter)
	return &MangaHandler{
		mangaCommandUsecase: queryUsecase,
	}
}

//Mount function
func(m *MangaHandler) Mount(router *echo.Echo) {
	router.GET("/comic/list", m.GetAllComic)
	router.GET("/comic/:endpoint", m.GetComicInfo)
	router.GET("/comic/search", m.SearchManga)
	router.GET("/comic/genre", m.GetAllGenre)
	router.GET("/comic/popular", m.GetPopularManga)
	router.GET("/comic/recommended", m.GetRecommendedManga)
	router.GET("/comic/newest", m.GetNewestManga)
}

func(m *MangaHandler) GetAllComic(c echo.Context) error {
	filter := c.QueryParam("filter")
	result := m.mangaCommandUsecase.GetAllComic(filter)
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

func (m *MangaHandler) GetPopularManga(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"));

	if err != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%v", err.Error())
		return utils.ResponseError(errObj, c)
	}

	result := m.mangaCommandUsecase.GetPopularManga(page)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Search manga", http.StatusOK, c)
}

func (m *MangaHandler) GetRecommendedManga(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%v", err.Error())
		return utils.ResponseError(errObj, c)
	}

	result := m.mangaCommandUsecase.GetRecommendedManga(page)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Search manga", http.StatusOK, c)
}

func (m *MangaHandler) GetNewestManga(c echo.Context) error  {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%v", err.Error())
		return utils.ResponseError(errObj, c)
	}

	result := m.mangaCommandUsecase.GetNewestManga(page)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Search manga", http.StatusOK, c)
}