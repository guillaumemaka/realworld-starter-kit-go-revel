package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type UserController struct {
	ApplicationController
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// UserJSON is the wrapper around User to give it a key "user"
type UserJSON struct {
	User *User `json:"user"`
}

func (c UserController) GetUser() revel.Result {
	return c.RenderJSON(c.Args[currentUserKey])
}

func (c UserController) Register() revel.Result {
	body := struct {
		User struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	bodyUser := &body.User

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	defer c.Request.Body.Close()
	revel.TRACE.Println("User Binding:", bodyUser)
	u, errs := models.NewUser(bodyUser.Email, bodyUser.Username, bodyUser.Password)
	if errs != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(&errorJSON{errs})
	}

	err = c.DB.CreateUser(u)
	if err != nil {
		// TODO: Error JSON
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	res := &UserJSON{
		&User{
			Username: u.Username,
			Email:    u.Email,
			Token:    c.JWT.NewToken(u.Username),
		},
	}

	return c.RenderJSON(res)
}

func (c UserController) Login() revel.Result {
	body := struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}{}
	bodyUser := &body.User

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	revel.TRACE.Println("Body:", bodyUser)

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}
	defer c.Request.Body.Close()

	c.Validation.Required(bodyUser.Email).Key("email").Message(models.EMPTY_MSG)
	c.Validation.Email(bodyUser.Email).Key("email")
	c.Validation.Required(bodyUser.Password).Key("password").Message(models.EMPTY_MSG)

	if c.Validation.HasErrors() {
		errs := &errorJSON{}
		errs = errs.Build(c.Validation.ErrorMap())
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON(errs)
	}

	u, err := c.DB.FindUserByEmail(bodyUser.Email)
	if err != nil {
		// TODO: Error JSON
		c.Response.Status = http.StatusNotFound
		return c.RenderText(http.StatusText(http.StatusNotFound))
	}
	revel.TRACE.Println(u)
	match := u.MatchPassword(bodyUser.Password)
	if !match {
		// TODO: Error JSON
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	res := &UserJSON{
		&User{
			Username: u.Username,
			Email:    u.Email,
			Token:    c.JWT.NewToken(u.Username),
			Bio:      u.Bio,
			Image:    u.Image,
		},
	}

	return c.RenderJSON(res)
}

func (c UserController) UpdateUser() revel.Result {
	return c.Todo()
}
