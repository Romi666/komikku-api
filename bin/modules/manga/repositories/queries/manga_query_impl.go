package queries

import (
	"fmt"
	"github.com/gocolly/colly"
	"komikku-api/bin/config"
	"komikku-api/bin/modules/manga/models/domain"
	"komikku-api/bin/pkg/utils"
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

func (g MangaQueryImpl) GetAllComic(filter string) utils.Result {
	var output utils.Result
	var result []domain.Comic
	if filter != "" {
		g.URL = config.GlobalEnv.BaseURL + "daftar-komik/?tipe=" + filter
	}else {
		g.URL = config.GlobalEnv.BaseURL + "daftar-komik"
	}
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

func (g MangaQueryImpl) SearchManga(query string) utils.Result {
	var output utils.Result
	var result []domain.Comic
	g.URL = fmt.Sprintf("https://data.komiku.id/cari/?post_type=manga&s=%s", query)
	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.TrimSuffix(e2.ChildAttr("img", "data-src"), "?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
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
			genre.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL + "genre/", "", -1)
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

func (g MangaQueryImpl) GetPopularManga(page int) utils.Result {
	var (
		output utils.Result
		result []domain.Comic
		err		error
	)

	if page != 1 {
		g.URL = fmt.Sprintf("https://data.komiku.id/other/hot/page/%d/", page)
	} else {
		g.URL = fmt.Sprint("https://data.komiku.id/other/hot/")
	}

	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.TrimSuffix(e2.ChildAttr("img", "data-src"), "?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
			comic.Endpoint = "/" + comic.Endpoint

		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		result = append(result, comic)
	})
	err = g.Collector.Visit(g.URL)
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

func (g MangaQueryImpl) GetRecommendedManga(page int) utils.Result {
	var (
		output utils.Result
		result []domain.Comic
		err		error
	)

	if page != 1 {
		g.URL = fmt.Sprintf("https://data.komiku.id/other/rekomendasi/page/%d/", page)
	}else {
		g.URL = fmt.Sprint("https://data.komiku.id/other/rekomendasi/")
	}

	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.Trim(e2.ChildAttr("img", "data-src"), "?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
			comic.Endpoint = "/" + comic.Endpoint

		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		result = append(result, comic)
	})
	err = g.Collector.Visit(g.URL)
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

func (g MangaQueryImpl) GetNewestManga(page int) utils.Result {
	var (
		output utils.Result
		result []domain.Comic
		err		error
	)

	if page != 1 {
		g.URL = fmt.Sprintf("https://data.komiku.id/pustaka/page/%d/", page)
	} else {
		g.URL = fmt.Sprint("https://data.komiku.id/pustaka/")
	}

	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.TrimSuffix(e2.ChildAttr("img", "data-src"),"?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
			comic.Endpoint = "/" + comic.Endpoint

		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		result = append(result, comic)
	})
	err = g.Collector.Visit(g.URL)
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

func (g MangaQueryImpl) GetByGenre(endpoint string, page int) utils.Result {
	var (
		output utils.Result
		result []domain.Comic
		err		error
	)

	if page != 1 {
		g.URL = fmt.Sprintf("https://komiku.id/genre/%s/page/%d/", endpoint, page)
	} else {
		g.URL = fmt.Sprintf("https://komiku.id/genre/%s/", endpoint)
	}
	fmt.Println(g.URL)

	g.Collector.AllowURLRevisit = true
	g.Collector.OnHTML("div.bge", func(e *colly.HTMLElement) {
		var comic domain.Comic
		comic.Desc = e.ChildText("p")
		e.ForEach("div.bgei", func(i int, e2 *colly.HTMLElement) {
			comic.Image = strings.TrimSuffix(e2.ChildAttr("img", "data-src"),"?resize=450,235&quality=60")
			comic.Endpoint = strings.Replace(e2.ChildAttr("a", "href"), config.GlobalEnv.BaseURL, "", 1)
			comic.Type = e2.ChildText("b")
			comic.Endpoint = "/" + comic.Endpoint

		})
		e.ForEach("div.kan", func(i int, e2 *colly.HTMLElement) {
			comic.Title = e2.ChildText("h3")
		})
		result = append(result, comic)
	})
	err = g.Collector.Visit(g.URL)
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



