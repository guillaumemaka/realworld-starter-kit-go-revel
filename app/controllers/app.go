package controllers

import (
	"net/http"
	"reflect"
	"time"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/lib/auth"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type ApplicationController struct {
	GormController
	JWT auth.Tokener
}

type errorJSON struct {
	Errors models.ValidationErrors `json:"errors"`
}

const (
	currentUserKey    = "current_user"
	fetchedArticleKey = "article"
	claimKey          = "claim"
)

func (c *ApplicationController) Init() revel.Result {
	c.JWT = auth.NewJWT()
	return nil
}

func (c *ApplicationController) AddUser() revel.Result {
	user, err := c.currentUser()
	if err != nil {
		return c.NotFound(err.Error())
	}
	c.Args[currentUserKey] = user
	return nil
}

func (c *ApplicationController) ExtractArticle() revel.Result {
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

func (c *ApplicationController) currentUser() (*models.User, error) {
	var user = &models.User{}

	claims, _ := c.JWT.CheckRequest(c.Request)

	if claims != nil {
		c.Args[claimKey] = claims

		user, err := c.DB.FindUserByUsername(claims.Username)
		if err != nil {
			revel.INFO.Println("currentUser", err)
			return user, err
		}

		return user, nil
	}

	return user, nil
}

func (err *errorJSON) Build(errMap map[string]*revel.ValidationError) *errorJSON {
	err.Errors = models.ValidationErrors{}
	for _, validationError := range errMap {
		err.Errors[validationError.Key] = []string{validationError.Message}
	}
	return err
}

func (c *ApplicationController) buildArticleJSON(a *models.Article, u *models.User) Article {
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
