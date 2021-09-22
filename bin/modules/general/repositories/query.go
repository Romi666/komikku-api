package repositories

import "komiku-srapper/bin/pkg/utils"

type GeneralQuery interface {
	GetAllComic() <- chan utils.Result
}