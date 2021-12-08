package usecases

import "komiku-srapper/bin/pkg/utils"

type MangaUsecase interface {
	GetAllComic(filter string) utils.Result
	GetComicInfo(endpoint string) utils.Result
	SearchManga(query string) utils.Result
	GetAllGenre() utils.Result
	GetPopularManga() utils.Result
	GetRecommendedManga() utils.Result
	GetNewestManga() utils.Result
}