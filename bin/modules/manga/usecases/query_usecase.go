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

func (m mangaCommandUsecase) GetChapterDetail(endpoint string) utils.Result {
	var result utils.Result

	resultGetChapterDetail := m.mangaQuery.DetailChapter(endpoint)
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

func (m mangaCommandUsecase) GetComicInfo(endpoint string) utils.Result {
	var result utils.Result

	resultGetComicInfo := m.mangaQuery.GetComicInfo(endpoint)
	if resultGetComicInfo.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", resultGetComicInfo.Error)
		result.Error = errObj
		return result
	}

	comicInfo := resultGetComicInfo.Data.(domain.ComicInfo)

	resultGetChapter := m.mangaQuery.GetListChapter(endpoint)
	if resultGetChapter.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", resultGetChapter.Error)
		result.Error = errObj
		return result
	}

	chapterList := resultGetChapter.Data.([]domain.Chapter)
	comicInfo.ChapterList = chapterList

	result.Data = comicInfo


	return result
}

func (m mangaCommandUsecase) GetAllComic() utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetAllComic()
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

func (m mangaCommandUsecase) SearchManga(query string) utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.SearchManga(query)

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetComic := queryRes.Data.([]domain.Comic)
	if len(resultGetComic) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetComic
	return result
}

func (m mangaCommandUsecase) GetAllGenre() utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetAllGenre()

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetGenre := queryRes.Data.([]domain.Genre)
	if len(resultGetGenre) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetGenre
	return result
}



func CreateNewMangaUsecase(mangaQuery queries.MangaQuery) MangaUsecase {
	return &mangaCommandUsecase{
		mangaQuery: mangaQuery,
	}
}
