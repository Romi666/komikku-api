package repositories

import (
	"github.com/gocolly/colly"
	"komiku-srapper/bin/config"
	"komiku-srapper/bin/modules/general/models/domain"
	"komiku-srapper/bin/pkg/utils"
)

type GeneralQueryImpl struct {
	URL			string
	Collector	*colly.Collector
}

func NewGeneralQueryImpl(url string, collector *colly.Collector) GeneralQuery {
	return &GeneralQueryImpl{
		URL: url,
		Collector: collector,
	}
}

func (g GeneralQueryImpl) GetAllComic() <- chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var result []domain.Comic
		g.URL = config.GlobalEnv.BaseURL + "daftar-komik/"
		g.Collector.OnHTML("div.ls4", func(e *colly.HTMLElement) {
			var comic domain.Comic
			e.ForEach("div.ls4v", func(i int, element *colly.HTMLElement) {
				comic.Endpoint = element.ChildAttr("a", "href")
				comic.Image = element.ChildAttr("img", "data-src")
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
