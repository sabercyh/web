package article_service

import (
	"encoding/json"
	"web/models"
	"web/pkg/gredis"
	"web/pkg/logging"
	"web/pkg/setting"

	"web/service/cache_service"
)

type Article struct {
	ID         int
	TagID      int
	Title      string
	Journal    string
	Author     string
	Authors    string
	Date       string
	Link       string
	Papercode  string
	Abstract   string
	Theyear    string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":     a.TagID,
		"title":      a.Title,
		"journal":    a.Journal,
		"author":     a.Author,
		"authors":    a.Authors,
		"date":       a.Date,
		"link":       a.Link,
		"papercode":  a.Papercode,
		"abstract":   a.Abstract,
		"theyear":    a.Theyear,
		"created_by": a.CreatedBy,
		"state":      a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	ok, err := gredis.Delete((&cache_service.Article{PageNum: 0, PageSize: setting.AppSetting.PageSize}).GetArticlesKey())
	if !ok || err != nil {
		return err
	}
	err = gredis.LikeDeletes((&cache_service.Article{Theyear: a.Theyear, PageNum: 0, PageSize: setting.AppSetting.PageSize}).GetArticlesKeyByYear())
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	err := models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":      a.TagID,
		"title":       a.Title,
		"journal":     a.Journal,
		"author":      a.Author,
		"authors":     a.Authors,
		"date":        a.Date,
		"link":        a.Link,
		"papercode":   a.Papercode,
		"abstract":    a.Abstract,
		"theyear":     a.Theyear,
		"modified_by": a.ModifiedBy,
		"state":       a.State,
	})
	if err != nil {
		return err
	}

	cache := &cache_service.Article{ID: a.ID, PageNum: 0, PageSize: setting.AppSetting.PageSize}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		ok, err := gredis.Delete(cache.GetArticleKey())
		if !ok || err != nil {
			logging.Info(err)
		}
	}
	err = gredis.LikeDeletes((cache.GetAllArticlesKeyByYear()))
	if err != nil {
		logging.Info(err)
	}
	ok, err := gredis.Delete(cache.GetArticlesKey())
	if !ok || err != nil {
		logging.Info(err)
	}
	return nil
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 36000)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cache_service.Article{
		TagID: 1,
		// State: a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 36000)
	return articles, nil
}

func (a *Article) GetAllByYear() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cache_service.Article{
		TagID: 1,
		// State: a.State,
		Theyear:  a.Theyear,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKeyByYear()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticlesByYear(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 36000)
	return articles, nil
}

func (a *Article) Delete() error {
	err := models.DeleteArticle(a.ID)
	if err != nil {
		return err
	}
	cache := &cache_service.Article{ID: a.ID, PageNum: 0, PageSize: setting.AppSetting.PageSize}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		ok, err := gredis.Delete(cache.GetArticleKey())
		if !ok || err != nil {
			logging.Info(err)
		}
	}
	ok, err := gredis.Delete((cache.GetAllArticlesKeyByYear()))
	if !ok || err != nil {
		logging.Info(err)
	}
	ok, err = gredis.Delete(cache.GetArticlesKey())
	if !ok || err != nil {
		logging.Info(err)
	}
	return nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	// if a.State != -1 {
	// 	maps["state"] = a.State
	// }
	if a.Theyear != "" {
		maps["theyear"] = a.Theyear
	}
	maps["tag_id"] = 1

	return maps
}
