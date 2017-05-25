package controllers

import (
	"time"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/lib/auth"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
)

type ApplicationController struct {
	GormController
	JWT auth.Tokener
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
	if user := c.currentUser(); user != nil {
		c.Args[currentUserKey] = user
	}
	return nil
}

func (c *ApplicationController) currentUser() *models.User {
	if c.Args[currentUserKey] != nil {
		return c.Args[currentUserKey].(*models.User)
	}

	claims, err := c.JWT.CheckRequest(c.Request)
	if err != nil {
		revel.INFO.Println("JWT CheckRequest", err)
	}

	if claims != nil {
		c.Args[claimKey] = claims
		var user *models.User

		if err := cache.Get(claims.Username, &user); err != nil {
			user, err := c.DB.FindUserByUsername(claims.Username)
			if err != nil {
				revel.INFO.Println("currentUser", err)
			} else {
				go cache.Set(claims.Username, user, 24*time.Hour)
			}
		}

		return user
	}
	return nil
}
