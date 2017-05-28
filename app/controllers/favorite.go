package controllers

import (
	"net/http"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type FavoriteController struct {
	ApplicationController
}

func (c *FavoriteController) Post() revel.Result {
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	err := c.DB.FavoriteArticle(u, a)

	articleJSON := ArticleJSON{
		Article: c.buildArticleJSON(a, u),
	}

	c.Response.ContentType = "application/json"

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
	} else {
		c.Response.Status = http.StatusOK
	}

	return c.RenderJSON(articleJSON)
}

func (c *FavoriteController) Delete() revel.Result {
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	err := c.DB.UnfavoriteArticle(u, a)

	articleJSON := ArticleJSON{
		Article: c.buildArticleJSON(a, u),
	}

	c.Response.ContentType = "application/json"

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
	} else {
		c.Response.Status = http.StatusOK
	}

	return c.RenderJSON(articleJSON)
}
