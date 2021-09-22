package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

const BASE_URL = "https://komiku.id/"

type Scrapper struct {
	URL			string
	Collector	*colly.Collector
}

type Comic struct {
	Title		string	`json:"title"`
	Genre 		string	`json:"genre"`
	Status		string	`json:"status"`
	LastUpdate	string	`json:"last_update"`
}

func (s *Scrapper) GetAllComic() []Comic {
	var listComic []Comic
	s.URL = BASE_URL + "daftar-komik/"
	s.Collector.OnHTML("h4", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	err := s.Collector.Visit(s.URL)
	if err != nil {
		log.Println(err)
	}
	return listComic
}

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Second)
	newScrapper := Scrapper{
		URL: BASE_URL,
		Collector: c,
	}

	comic := newScrapper.GetAllComic()
	fmt.Println(comic)
}