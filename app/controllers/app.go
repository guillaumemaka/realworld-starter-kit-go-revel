package controllers

import (
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
