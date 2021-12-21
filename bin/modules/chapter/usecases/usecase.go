package usecases

import "komikku-api/bin/pkg/utils"

type ChapterUsecase interface {
	GetChapterDetail(endpoint string) utils.Result
}
