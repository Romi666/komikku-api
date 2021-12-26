package handlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"komikku-api/bin/config"
	chapterQ "komikku-api/bin/modules/chapter/repositories/queries"
	mangaQ "komikku-api/bin/modules/manga/repositories/queries"
	mangaU "komikku-api/bin/modules/manga/usecases"
	httpError "komikku-api/bin/pkg/http-error"
	"komikku-api/bin/pkg/utils"
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
	api := router.Group("/api")
	api.GET("/comic/list", m.GetAllComic)
	api.GET("/comic/info/:endpoint", m.GetComicInfo)
	api.GET("/comic/search/:query", m.SearchManga)
	api.GET("/comic/genre", m.GetAllGenre)
	api.GET("/comic/popular", m.GetPopularManga)
	api.GET("/comic/recommended/page/:page", m.GetRecommendedManga)
	api.GET("/comic/newest/page/:page", m.GetNewestManga)
	api.GET("/comic/genres/:endpoint/page/:page", m.GetByGenre)
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
	query := c.Param("query")
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
	page, err := strconv.Atoi(c.Param("page"))
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
	page, err := strconv.Atoi(c.Param("page"))
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

func (m *MangaHandler) GetByGenre(c echo.Context) error {
	endpoint := c.Param("endpoint")
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%v", err.Error())
		return utils.ResponseError(errObj, c)
	}

	result := m.mangaCommandUsecase.GetByGenre(endpoint, page)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get Comic By Genre", http.StatusOK, c)
}