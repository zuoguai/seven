package entity

import "gorm.io/gorm"

type Article struct {
}
type ArticleEntity struct {
}
type IArticleEntity interface {
}

func (a Article) TableName() string {
	return "t_article"
}

var ArticleEntityInstance = &ArticleEntity{}

type ArticleOptionFn func(*gorm.DB)
type SiteFindOptionInstance struct{}

var SiteFindOptionIns SiteFindOptionInstance
