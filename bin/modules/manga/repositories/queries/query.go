package queries

import "komiku-srapper/bin/pkg/utils"

type MangaQuery interface {
	GetAllComic() <- chan utils.Result
	GetComicInfo(endpoint string) <- chan utils.Result
}