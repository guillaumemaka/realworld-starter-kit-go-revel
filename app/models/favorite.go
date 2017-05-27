package models

import "github.com/jinzhu/gorm"

type Favorite struct {
	ID        int
	User      User
	UserID    int `gorm:"index:index_favorites_on_user_id"`
	Article   Article
	ArticleID int `gorm:"index:index_favorites_on_article_id"`
}

func (f *Favorite) AfterCreate(db *gorm.DB) (err error) {
	var a = &Article{ID: f.ArticleID}
	err = db.First(&a).Update("favorites_count", gorm.Expr("favorites_count + ?", 1)).Error
	return
}

func (f *Favorite) AfterDelete(db *gorm.DB) (err error) {
	var a = &Article{ID: f.ArticleID}
	err = db.First(&a).Update("favorites_count", gorm.Expr("favorites_count - ?", 1)).Error
	return
}
