package handlers

import (
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"komikku-api/bin/config"
	chapterQ "komikku-api/bin/modules/chapter/repositories/queries"
	chapterU "komikku-api/bin/modules/chapter/usecases"
	"komikku-api/bin/pkg/utils"
	"net/http"
)

type ChapterHandler struct {
	chapterCommandUsecase chapterU.ChapterUsecase
}


func New() *ChapterHandler {
	collector := colly.NewCollector()
	queryChapter := chapterQ.NewChapterQuery(config.GlobalEnv.BaseURL, collector)
	queryUsecase := chapterU.CreateNewChapterUsecase(queryChapter)
	return &ChapterHandler{
		chapterCommandUsecase: queryUsecase,
	}
}

func (c *ChapterHandler) Mount(router *echo.Echo)  {
	api := router.Group("/api")
	api.GET("/comic/chapter/:endpoint", c.GetChapterDetail)
}

func(m *ChapterHandler) GetChapterDetail(c echo.Context) error {
	endpoint := c.Param("endpoint")
	result := m.chapterCommandUsecase.GetChapterDetail(endpoint)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Get Chapter Detail", http.StatusOK, c)
}
