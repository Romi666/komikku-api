package usecases

import "komiku-srapper/bin/pkg/utils"

type ChapterUsecase interface {
	GetChapterDetail(endpoint string) utils.Result
}
