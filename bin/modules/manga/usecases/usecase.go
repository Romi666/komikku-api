package usecases

import "komikku-api/bin/pkg/utils"

type MangaUsecase interface {
	GetAllComic(filter string) utils.Result
	GetComicInfo(endpoint string) utils.Result
	SearchManga(query string) utils.Result
	GetAllGenre() utils.Result
	GetPopularManga(page int) utils.Result
	GetRecommendedManga(page int) utils.Result
	GetNewestManga(page int) utils.Result
	GetByGenre(endpoint string, page int) utils.Result
}