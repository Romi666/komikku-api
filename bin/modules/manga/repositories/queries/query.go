package queries

import "komiku-srapper/bin/pkg/utils"

type MangaQuery interface {
	GetAllComic() utils.Result
	GetComicInfo(endpoint string) utils.Result
	GetListChapter(endpoint string) utils.Result
	DetailChapter(endpoint string) utils.Result
}