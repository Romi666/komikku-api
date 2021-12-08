package queries

import "komiku-srapper/bin/pkg/utils"

type ChapterQuery interface {
	GetListChapter(endpoint string) utils.Result
	DetailChapter(endpoint string) utils.Result
}