package controllers

import (
	"net/http"
	"reflect"

	"github.com/revel/revel"
)

var requireAuth = map[string][]string{
	"UserController": []string{"PUT"},
}

func authorize(c *revel.Controller) revel.Result {
	if methods, ok := requireAuth[c.Name]; ok {
		if HasElem(methods, c.Request.Method) {
			if c.Args[claimKey] == nil {
				if c.Args[currentUserKey] == nil {
					c.Response.Status = http.StatusUnauthorized
					return c.RenderText(http.StatusText(http.StatusUnauthorized))
				}
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
	revel.InterceptFunc(authorize, revel.BEFORE, &UserController{})
}
