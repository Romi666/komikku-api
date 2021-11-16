package usecases

import (
	"fmt"
	"komiku-srapper/bin/modules/manga/models/domain"
	"komiku-srapper/bin/modules/manga/repositories/queries"
	httpError "komiku-srapper/bin/pkg/http-error"
	"komiku-srapper/bin/pkg/utils"
)

type mangaCommandUsecase struct {
	mangaQuery	queries.MangaQuery
}

func (m mangaCommandUsecase) GetComicInfo(endpoint string) utils.Result {
	var result utils.Result

	queryRes := <- m.mangaQuery.GetComicInfo(endpoint)
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	comicInfo := queryRes.Data.(domain.ComicInfo)
	result.Data = comicInfo

	return result
}

func (m mangaCommandUsecase) GetAllComic() utils.Result {
	var result utils.Result

	queryRes := <- m.mangaQuery.GetAllComic()
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetManga := queryRes.Data.([]domain.Comic)
	if len(resultGetManga) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	result.Data = resultGetManga
	return result
}

func CreateNewMangaUsecase(mangaQuery queries.MangaQuery) MangaUsecase {
	return &mangaCommandUsecase{
		mangaQuery: mangaQuery,
	}
}
