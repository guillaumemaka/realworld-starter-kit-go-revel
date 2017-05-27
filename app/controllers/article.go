package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type ArticleController struct {
	ApplicationController
}

type Article struct {
	Slug           string   `json:"slug"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int      `json:"favoritesCount"`
	TagList        []string `json:"tagList"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	Author         Author   `json:"user"`
}

type Author struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type ArticleJSON struct {
	Article `json:"article"`
}

type ArticlesJSON struct {
	Articles      []Article `json:"articles"`
	ArticlesCount int       `json:"articlesCount"`
}

func (c *ArticleController) ExtractArticle() revel.Result {
	if slug := c.Params.Route.Get("slug"); slug != "" {
		a, err := c.DB.GetArticle(slug)
		if err != nil {
			c.Response.Status = http.StatusNotFound
			return c.RenderText(err.Error())
		}

		c.Args[fetchedArticleKey] = a
	}

	return nil
}

func (c *ArticleController) Index(tag, favorited, author string, offset, limit int) revel.Result {
	revel.TRACE.Println("tag: ", tag, "favorited: ", favorited, "authored: ", author, "offset: ", offset, "limit: ", limit)
	var articles []models.Article

	query := c.DB.GetAllArticles()
	query = c.DB.Offset(query, offset)
	query = c.DB.Limit(query, limit)
	query = c.DB.FilterByTag(query, tag)
	query = c.DB.FilterAuthoredBy(query, author)
	query = c.DB.FilterFavoritedBy(query, favorited)

	err := query.Debug().Find(&articles).Error

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	c.Response.ContentType = "application/json"

	var articlesJSON = ArticlesJSON{
		Articles:      make([]Article, 0),
		ArticlesCount: 0,
	}

	if len(articles) == 0 {
		return c.RenderJSON(articlesJSON)
	}

	var u = c.Args[currentUserKey].(*models.User)

	for i := range articles {
		a := &articles[i]
		articlesJSON.Articles = append(articlesJSON.Articles, c.buildArticleJSON(a, u))
	}

	articlesJSON.ArticlesCount = len(articles)

	return c.RenderJSON(articlesJSON)
}

func (c *ArticleController) Create() revel.Result {
	var body struct {
		Article struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			Body        string   `json:"body"`
			TagList     []string `json:"tagList"`
		} `json:"article"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	defer c.Request.Body.Close()

	u := c.Args[currentUserKey].(*models.User)

	a := models.NewArticle(body.Article.Title, body.Article.Description, body.Article.Body, u)

	if valid, errs := a.IsValid(); !valid {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(errorJSON{errs})
	}

	for _, tagName := range body.Article.TagList {
		tag, _ := c.DB.FindTagOrInit(tagName)
		a.Tags = append(a.Tags, tag)
	}

	if err := c.DB.CreateArticle(a); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	articleJSON := ArticleJSON{
		Article: c.buildArticleJSON(a, u),
	}

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(articleJSON)
}

func (c *ArticleController) Read() revel.Result {
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	articleJSON := ArticleJSON{
		Article: c.buildArticleJSON(a, u),
	}

	return c.RenderJSON(articleJSON)
}

func (c *ArticleController) Update() revel.Result {
	var err error
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	if !a.IsOwnedBy(u.Username) {
		c.Response.Status = http.StatusForbidden
		err = fmt.Errorf("You don't have the permission to edit this article")
		return c.RenderText(err.Error())
	}

	var body map[string]map[string]interface{}

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	defer c.Request.Body.Close()

	if _, present := body["article"]; !present {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	var article map[string]interface{}

	article = body["article"]

	if title, present := article["title"]; present {
		a.Title = title.(string)
	}

	if description, present := article["description"]; present {
		a.Description = description.(string)
	}

	if body, present := article["body"]; present {
		a.Body = body.(string)
	}

	if valid, errs := a.IsValid(); !valid {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(errorJSON{errs})
	}

	if err := c.DB.SaveArticle(a); err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(err.Error())
	}

	articleJSON := ArticleJSON{
		Article: c.buildArticleJSON(a, u),
	}

	return c.RenderJSON(articleJSON)
}

func (c *ArticleController) Delete() revel.Result {
	var err error
	a := c.Args[fetchedArticleKey].(*models.Article)
	u := c.Args[currentUserKey].(*models.User)

	if !a.IsOwnedBy(u.Username) {
		c.Response.Status = http.StatusForbidden
		err = fmt.Errorf("You don't have the permission to delete this article")
		return c.RenderText(err.Error())
	}

	err = c.DB.DeleteArticle(a)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		err = fmt.Errorf("You don't have the permission to delete this article")
		return c.RenderText(err.Error())
	}

	c.Response.Status = http.StatusNoContent
	return c.RenderText(http.StatusText(http.StatusNoContent))
}

func (c *ArticleController) buildArticleJSON(a *models.Article, u *models.User) Article {
	following := false
	favorited := false
	//TODO: Remove reflection
	if !reflect.DeepEqual(u, &models.User{}) {
		following = c.DB.IsFollowing(u.ID, a.User.ID)
		favorited = c.DB.IsFavorited(u.ID, a.ID)
	}

	article := Article{
		Slug:           a.Slug,
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		Favorited:      favorited,
		FavoritesCount: a.FavoritesCount,
		CreatedAt:      a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      a.UpdatedAt.Format(time.RFC3339),
		Author: Author{
			Username:  a.User.Username,
			Bio:       a.User.Bio,
			Image:     a.User.Image,
			Following: following,
		},
	}

	for _, t := range a.Tags {
		article.TagList = append(article.TagList, t.Name)
	}

	return article
}
