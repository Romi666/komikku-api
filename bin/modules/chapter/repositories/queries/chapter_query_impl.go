package queries

import (
	"github.com/gocolly/colly"
	"komikku-api/bin/config"
	"komikku-api/bin/modules/chapter/models/domain"
	"komikku-api/bin/pkg/utils"
)

type ChapterQueryImpl struct {
	URL			string
	Collector	*colly.Collector
}


func NewChapterQuery(url string, collector *colly.Collector) ChapterQuery {
	return &ChapterQueryImpl{
		URL: url,
		Collector: collector,
	}
}

func (c ChapterQueryImpl) GetListChapter(endpoint string) utils.Result {
	var output utils.Result
	var chapterList []domain.Chapter
	c.URL = config.GlobalEnv.BaseURL + endpoint
	c.Collector.AllowURLRevisit = true
	c.Collector.OnHTML("tbody._3Rsjq", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			var chapter domain.Chapter
			if element.ChildText("td.judulseries") != "" {
				chapter.Endpoint = element.ChildAttr("a", "href")
				chapter.Name = element.ChildText("td.judulseries")
				chapterList = append(chapterList, chapter)
			}
		})
	})
	err := c.Collector.Visit(c.URL)
	if err != nil {
		output = utils.Result{
			Error: err,
		}
	}
	output = utils.Result{
		Data: chapterList,
	}

	return output
}

func (c ChapterQueryImpl) DetailChapter(endpoint string) utils.Result {
	var output utils.Result
	var chapter domain.ChapterDetail

	c.URL = config.GlobalEnv.BaseURL + endpoint
	c.Collector.AllowURLRevisit = true
	c.Collector.OnHTML("section[id=Baca_Komik]", func(e *colly.HTMLElement) {
		imageList := e.ChildAttrs("img", "src")
		chapter.Image = imageList
	})
	c.Collector.OnHTML("header[id=Judul]", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			chapter.Title = e.ChildText("h1")
		}
	})
	err := c.Collector.Visit(c.URL)
	if err != nil {
		output = utils.Result{
			Error: err,
		}
		return output
	}
	output = utils.Result{
		Data: chapter,
	}
	return output
}
