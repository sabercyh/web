package cache_service

import (
	"strconv"
	"strings"
	"web/pkg/e"
)

type Article struct {
	ID    int
	TagID int
	// State int
	Theyear  string
	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	// if a.ID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.ID))
	// }
	// if a.TagID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.TagID))
	// }
	// if a.State >= 0 {
	// 	keys = append(keys, strconv.Itoa(a.State))
	// }
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
func (a *Article) GetAllArticlesKeyByYear() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
		"Year",
	}

	// if a.ID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.ID))
	// }
	// if a.TagID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.TagID))
	// }
	// if a.State >= 0 {
	// 	keys = append(keys, strconv.Itoa(a.State))
	// }
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
func (a *Article) GetArticlesKeyByYear() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
		"Year",
		a.Theyear,
	}

	// if a.ID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.ID))
	// }
	// if a.TagID > 0 {
	// 	keys = append(keys, strconv.Itoa(a.TagID))
	// }
	// if a.State >= 0 {
	// 	keys = append(keys, strconv.Itoa(a.State))
	// }
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
