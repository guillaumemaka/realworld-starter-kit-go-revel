package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type CommentController struct {
	ApplicationController
}

type Comment struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Author    Author `json:"author"`
}

type CommentJSON struct {
	Comment Comment `json:"comment"`
}

type CommentsJSON struct {
	Comments []Comment `json:"comments"`
}

type commentBody struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func (c *CommentController) Index() revel.Result {
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	var comments []models.Comment
	err := c.DB.GetComments(a, &comments)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	var commentsJSON = CommentsJSON{}
	for _, comment := range comments {
		commentsJSON.Comments = append(commentsJSON.Comments, c.buildCommentJSON(&comment, u))
	}

	return c.RenderJSON(commentsJSON)
}

func (c *CommentController) Create(id int) revel.Result {
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	var commentBody commentBody
	if err := json.NewDecoder(c.Request.Body).Decode(&commentBody); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	defer c.Request.Body.Close()

	comment, errs := models.NewComment(a, u, commentBody.Comment.Body)

	if errs != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(errorJSON{errs})
	}

	err := c.DB.CreateComment(comment)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	commentJSON := CommentJSON{
		Comment: c.buildCommentJSON(comment, u),
	}

	return c.RenderJSON(commentJSON)
}

func (c *CommentController) Delete(id int) revel.Result {
	u := c.Args[currentUserKey].(*models.User)

	var comment = models.Comment{}
	err := c.DB.GetComment(id, &comment)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	if canDelete := comment.CanBeDeletedBy(u); !canDelete {
		c.Response.Status = http.StatusForbidden
		return c.RenderText("You can't delete this comment. Permission denied.")
	}

	err = c.DB.DeleteComment(&comment)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	c.Response.Status = http.StatusOK
	return nil
}

func (c *CommentController) buildCommentJSON(comment *models.Comment, user *models.User) Comment {
	following := false

	if (user != &models.User{}) {
		following = c.DB.IsFollowing(user.ID, comment.User.ID)
	}

	return Comment{
		ID:        comment.ID,
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
		Author: Author{
			Username:  comment.User.Username,
			Bio:       comment.User.Bio,
			Image:     comment.User.Image,
			Following: following,
		},
	}
}
