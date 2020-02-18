package article_service

import (
	"database/sql"
	"fmt"
	"github.com/DevOpserzhao/ops_gin/first/models"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) StatsDB() sql.DBStats {
	return models.StatsDB()

}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

//func (a *Article) Get() (*models.Article, error) {
//	var cacheArticle *models.Article
//
//	cache := cache_service.Article{ID: a.ID}
//	key := cache.GetArticleKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheArticle)
//			return cacheArticle, nil
//		}
//	}
//
//	article, err := models.GetArticle(a.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	gredis.Set(key, article, 3600)
//	return article, nil
//}

type transformedArticle struct {
	TagID int    `json:"tag_id" gorm:"index"`
	Tag   string `json:"tag"`

	Title string `json:"title"`
}

func (a *Article) Get() (interface{}, error) {
	//var cacheArticle *models.Article
	//
	//cache := cache_service.Article{ID: a.ID}
	//key := cache.GetArticleKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticle)
	//		return cacheArticle, nil
	//	}
	//}

	article, err := models.GetArticle(a.ID)

	_article := transformedArticle{TagID: article.TagID, Tag: article.Tag.Name, Title: article.Title}

	fmt.Printf("%v", article)

	if err != nil {
		return nil, err
	}

	//gredis.Set(key, article, 3600)
	return _article, nil
}

//func (a *Article) GetAll() ([]*models.Article, error) {
//	var (
//		articles, cacheArticles []*models.Article
//	)
//
//	cache := cache_service.Article{
//		TagID: a.TagID,
//		State: a.State,
//
//		PageNum:  a.PageNum,
//		PageSize: a.PageSize,
//	}
//	key := cache.GetArticlesKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheArticles)
//			return cacheArticles, nil
//		}
//	}
//
//	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
//	if err != nil {
//		return nil, err
//	}
//
//	gredis.Set(key, articles, 3600)
//	return articles, nil
//}

//func (a *Article) GetAll() ([] transformedArticle, error) {
func (a *Article) GetAll() (map[string]interface{}, error) {

	data := make(map[string]interface{})
	var (
		articles []*models.Article
	)

	var _articles []transformedArticle

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())

	for _, item := range articles {

		_articles = append(_articles, transformedArticle{TagID: item.TagID, Tag: item.Tag.Name, Title: item.Title})
	}

	if err != nil {
		return nil, err
	}
	data["articles"] = _articles
	data["StatsDB"] = a.StatsDB()
	return data, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
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
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
