package usecases

import (
	"fmt"
	chapterM "komiku-srapper/bin/modules/chapter/models/domain"
	chapterQ "komiku-srapper/bin/modules/chapter/repositories/queries"
	mangaM "komiku-srapper/bin/modules/manga/models/domain"
	mangaQ "komiku-srapper/bin/modules/manga/repositories/queries"
	httpError "komiku-srapper/bin/pkg/http-error"
	"komiku-srapper/bin/pkg/utils"
)

type mangaCommandUsecase struct {
	mangaQuery	mangaQ.MangaQuery
	chapterQuery chapterQ.ChapterQuery
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

	comicInfo := resultGetComicInfo.Data.(mangaM.ComicInfo)

	resultGetChapter := m.chapterQuery.GetListChapter(endpoint)
	if resultGetChapter.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", resultGetChapter.Error)
		result.Error = errObj
		return result
	}

	chapterList := resultGetChapter.Data.([]chapterM.Chapter)
	comicInfo.ChapterList = chapterList

	result.Data = comicInfo


	return result
}

func (m mangaCommandUsecase) GetAllComic(filter string) utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetAllComic(filter)
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetManga := queryRes.Data.([]mangaM.Comic)
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

	resultGetComic := queryRes.Data.([]mangaM.Comic)
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

	resultGetGenre := queryRes.Data.([]mangaM.Genre)
	if len(resultGetGenre) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetGenre
	return result
}

func (m mangaCommandUsecase) GetPopularManga(page int) utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetPopularManga(page)

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetComic := queryRes.Data.([]mangaM.Comic)
	if len(resultGetComic) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetComic
	return result
}

func (m mangaCommandUsecase) GetRecommendedManga(page int) utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetRecommendedManga(page)

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetComic := queryRes.Data.([]mangaM.Comic)
	if len(resultGetComic) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetComic
	return result
}

func (m mangaCommandUsecase) GetNewestManga(page int) utils.Result {
	var result utils.Result

	queryRes := m.mangaQuery.GetNewestManga(page)

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", queryRes.Error)
		result.Error = errObj
		return result
	}

	resultGetComic := queryRes.Data.([]mangaM.Comic)
	if len(resultGetComic) == 0 {
		errObj := httpError.NewNotFound()
		errObj.Message = fmt.Sprintf("%v", "Data not found")
		result.Error = errObj
		return result
	}

	result.Data = resultGetComic
	return result
}


func CreateNewMangaUsecase(mangaQuery mangaQ.MangaQuery, chapterQuery chapterQ.ChapterQuery) MangaUsecase {
	return &mangaCommandUsecase{
		mangaQuery: mangaQuery,
		chapterQuery: chapterQuery,
	}
}
