package queries

import "komikku-api/bin/pkg/utils"

type ChapterQuery interface {
	GetListChapter(endpoint string) utils.Result
	DetailChapter(endpoint string) utils.Result
}