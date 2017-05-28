package controllers

import (
	"net/http"
	"reflect"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

var requireAuth = map[string][]string{
	"UserController":     []string{"GET", "PUT"},
	"ArticleController":  []string{"POST", "PUT", "DELETE"},
	"FavoriteController": []string{"POST", "DELETE"},
}

func authorize(c *revel.Controller) revel.Result {
	if methods, ok := requireAuth[c.Name]; ok {
		if HasElem(methods, c.Request.Method) {
			if (reflect.DeepEqual(c.Args[currentUserKey].(*models.User), &models.User{})) {
				c.Response.Status = http.StatusUnauthorized
				return c.RenderText(http.StatusText(http.StatusUnauthorized))
			}
		}
	}

	return nil
}

func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Init, revel.BEFORE)
	revel.InterceptMethod((*ApplicationController).Init, revel.BEFORE)
	revel.InterceptMethod((*ApplicationController).AddUser, revel.BEFORE)
	revel.InterceptMethod((*ApplicationController).ExtractArticle, revel.BEFORE)
	revel.InterceptFunc(authorize, revel.BEFORE, &UserController{})
	revel.InterceptFunc(authorize, revel.BEFORE, &ArticleController{})
	revel.InterceptFunc(authorize, revel.BEFORE, &FavoriteController{})
}
