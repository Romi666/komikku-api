package queries

import (
	"github.com/gocolly/colly"
	"komiku-srapper/bin/config"
	"komiku-srapper/bin/modules/manga/models/domain"
	"komiku-srapper/bin/pkg/utils"
	"strings"
)

type MangaQueryImpl struct {
	URL			string
	Collector	*colly.Collector
}

func NewMangaQuery(url string, collector *colly.Collector) MangaQuery {
	return &MangaQueryImpl{
		URL: url,
		Collector: collector,
	}
}

func (g MangaQueryImpl) GetAllComic() <- chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var result []domain.Comic
		g.URL = config.GlobalEnv.BaseURL + "/daftar-komik/"
		g.Collector.AllowURLRevisit = true
		g.Collector.OnHTML("div.ls4", func(e *colly.HTMLElement) {
			var comic domain.Comic
			e.ForEach("div.ls4v", func(i int, element *colly.HTMLElement) {
				comic.Endpoint = element.ChildAttr("a", "href")
				comic.Image = strings.TrimSuffix(element.ChildAttr("img", "data-src"), "?resize=240,170")
			})
			e.ForEach("div.ls4j", func(i int, element *colly.HTMLElement) {
				comic.Title = element.ChildText("h4")
			})
			result = append(result, comic)
		})
		err := g.Collector.Visit(g.URL)
		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}
		output <- utils.Result{
			Data: result,
		}
	}()
	return output
}

func (g MangaQueryImpl) GetComicInfo(endpoint string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		var comicInfo domain.ComicInfo
		g.URL = config.GlobalEnv.BaseURL + endpoint
		g.Collector.AllowURLRevisit = true
		g.Collector.OnHTML("table.inftable", func(e *colly.HTMLElement) {
			e.ForEach("tbody", func(i int, element *colly.HTMLElement) {
				comicInfo.Title = element.ChildText("tr:nth-child(1) > td:nth-child(2)")
				comicInfo.Type = element.ChildText("tr:nth-child(2) > td:nth-child(2)")
				comicInfo.Author = element.ChildText("tr:nth-child(4) > td:nth-child(2)")
				comicInfo.Status = element.ChildText("tr:nth-child(5) > td:nth-child(2)")
				comicInfo.Rating = element.ChildText("tr:nth-child(6) > td:nth-child(2)")

			})
		})
		g.Collector.OnHTML("ul.genre", func(e *colly.HTMLElement) {
			e.ForEach("li.genre", func(i int, element *colly.HTMLElement) {
				comicInfo.Genre = append(comicInfo.Genre, element.Text)
			})
		})

		g.Collector.OnHTML("div.ims", func(e *colly.HTMLElement) {
			comicInfo.Thumbnail = strings.TrimSuffix(e.ChildAttr("img", "src"), "?w=225&quality=60")
		})
		err := g.Collector.Visit(g.URL)
		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: comicInfo,
		}
	}()

	return output
}
