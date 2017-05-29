package models

import (
	"errors"
	"time"
)

// Comment is the representation of a comment
type Comment struct {
	ID        int
	Body      string `gorm:"type:text`
	Article   Article
	ArticleID int `gorm:"index:index_comments_on_article_id"`
	User      User
	UserID    int `gorm:"index:index_comments_on_user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	errorCommentBodyIsEmpty = errors.New(EMPTY_MSG)
)

type CommentStorer interface {
	CreateComment(*Comment) error
	DeleteComment(*Comment) error
	GetComments(*Article, *[]Comment) error
	GetComment(int, *Comment) error
}

// NewComment initialize a new comment struct
func NewComment(article *Article, user *User, body string) (*Comment, ValidationErrors) {
	if body == "" {
		return nil, ValidationErrors{
			"body": []string{errorCommentBodyIsEmpty.Error()},
		}
	}

	return &Comment{
		Body:    body,
		User:    *user,
		Article: *article,
	}, nil
}

// CanBeDeletedBy check if the comment can be deleted by the given user
func (comment *Comment) CanBeDeletedBy(user *User) bool {
	return (user.Username == comment.User.Username)
}

// CreateComment persist a new comment in the database
func (db *DB) CreateComment(comment *Comment) (err error) {
	err = db.Create(&comment).Error
	return
}

// DeleteComment delete a comment from the database
func (db *DB) DeleteComment(comment *Comment) (err error) {
	err = db.Delete(&comment).Error
	return
}

// GetComments get all comments for the givan article
func (db *DB) GetComments(article *Article, comments *[]Comment) error {
	err := db.Model(&article).Preload("User").Related(&comments).Error
	return err
}

// GetComment get a comment for the given commentID
func (db *DB) GetComment(commentID int, comment *Comment) error {
	err := db.Preload("User").First(&comment, commentID).Error
	return err
}
