package entity

import "gorm.io/gorm"

type Article struct {
}
type ArticleEntity interface {
}
type IArticleEntity struct {
}

func (a Article) TableName() string {
	return "t_article"
}

var ArticleEntityInstance = &IArticleEntity{}

type ArticleOptionFn func(*gorm.DB)
type SiteFindOptionInstance struct{}

var SiteFindOptionIns SiteFindOptionInstance
