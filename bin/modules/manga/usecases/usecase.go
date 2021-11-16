package usecases

import "komiku-srapper/bin/pkg/utils"

type MangaUsecase interface {
	GetAllComic() utils.Result
	GetComicInfo(endpoint string) utils.Result
}