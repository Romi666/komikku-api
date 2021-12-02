package queries

import (
	"fmt"
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

func (g MangaQueryImpl) GetAllComic() utils.Result {
	var output utils.Result
	var result []domain.Comic
	g.URL = config.GlobalEnv.BaseURL + "/daftar-komik"
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
		output = utils.Result{
			Error: err,
		}
	}
	output = utils.Result{
		Data: result,
	}

return output
}

func (g MangaQueryImpl) GetComicInfo(endpoint string) utils.Result {
	var output utils.Result

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
	output = utils.Result{
			Error: err,
		}
	}

	output = utils.Result{
		Data: comicInfo,
	}
return output
}

func (g MangaQueryImpl) GetListChapter(endpoint string) utils.Result {
	var output utils.Result
	var chapterList []domain.Chapter
	g.URL = config.GlobalEnv.BaseURL + endpoint
	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("tbody._3Rsjq", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			var chapter domain.Chapter
			if element.ChildText("td.judulseries") != "" {
				chapter.Endpoint = element.ChildAttr("a", "href")
				chapter.Name = element.ChildText("td.judulseries")
				chapterList = append(chapterList, chapter)
			}
		})
	})
	err := g.Collector.Visit(g.URL)
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

func (g MangaQueryImpl) DetailChapter(endpoint string) utils.Result {
	var output utils.Result
	var chapter domain.ChapterDetail

	g.URL = config.GlobalEnv.BaseURL + endpoint
	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("section[id=Baca_Komik]", func(e *colly.HTMLElement) {
		imageList := e.ChildAttrs("img", "src")
		chapter.Image = imageList
	})
	g.Collector.OnHTML("header[id=Judul]", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			chapter.Title = e.ChildText("h1")
		}
	})
	err := g.Collector.Visit(g.URL)
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

func (g MangaQueryImpl) SearchManga(query string) utils.Result {
	var output utils.Result
	var result []domain.Comic
	g.URL = fmt.Sprintf("https://data.komiku.id/cari/?post_type=manga&s=%s", query)
	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = e2.ChildAttr("img", "data-src")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Endpoint = "/" + comic.Endpoint

		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		result = append(result, comic)
	})
	err := g.Collector.Visit(g.URL)
	if err != nil {
		output = utils.Result{
			Error: err,
		}

		return output
	}

	output = utils.Result{
		Data:  result,
	}

	return output
}

func (g MangaQueryImpl) GetAllGenre() utils.Result {
	var output utils.Result
	var result []domain.Genre

	g.URL = fmt.Sprint("https://data.komiku.id/pustaka/")
	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("ul.genre", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e2 *colly.HTMLElement) {
			var genre domain.Genre
			genre.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			genre.Endpoint = "/" + genre.Endpoint
			genre.Title  = e2.ChildText("a")

			result = append(result, genre)
		})
	})

	err := g.Collector.Visit(g.URL)
	if err != nil {
		output = utils.Result{
			Error: err,
		}

		return output
	}

	output = utils.Result{
		Data:  result,
	}

	return output
}

