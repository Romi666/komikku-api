package usecases

import (
	"fmt"
	"komikku-api/bin/modules/chapter/models/domain"
	chapterQ "komikku-api/bin/modules/chapter/repositories/queries"
	httpError "komikku-api/bin/pkg/http-error"
	"komikku-api/bin/pkg/utils"
)

type chapterCommandUsecase struct {
	chapterQuery chapterQ.ChapterQuery
}

func (c chapterCommandUsecase) GetChapterDetail(endpoint string) utils.Result {
	var result utils.Result

	resultGetChapterDetail := c.chapterQuery.DetailChapter(endpoint)
	if resultGetChapterDetail.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", resultGetChapterDetail.Error)
		result.Error = errObj
		return result
	}

	chapterDetail := resultGetChapterDetail.Data.(domain.ChapterDetail)

	result.Data = chapterDetail

	return result
}

func CreateNewChapterUsecase( chapterQuery chapterQ.ChapterQuery) ChapterUsecase {
	return &chapterCommandUsecase{
		chapterQuery: chapterQuery,
	}
}
