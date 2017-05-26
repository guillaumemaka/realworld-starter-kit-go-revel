package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/controllers"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
)

type UserControllerTest struct {
	AppTest
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistrationBody struct {
	User UserRegister `json:"user"`
}

type UserLoginBody struct {
	User UserLogin `json:"user"`
}

type testRegistration struct {
	errorKey string
	message  string
	body     UserRegistrationBody
}

type testLogin struct {
	errorKey string
	message  string
	body     UserLoginBody
}

type ErrorJSON struct {
	Errors map[string][]string `json:"errors"`
}

func (t *AppTest) TestLoginSuccesFully() {
	bodyUser := UserLoginBody{
		UserLogin{
			Email:    "user1@example.com",
			Password: "password",
		},
	}

	jsonBody, _ := json.Marshal(bodyUser)

	t.MakePostRequest("/api/users/login", bytes.NewBuffer(jsonBody), nil)
	t.AssertOk()

	var UserJSON = controllers.UserJSON{}
	json.Unmarshal(t.ResponseBody, &UserJSON)

	t.AssertEqual(JWT.NewToken("user1"), UserJSON.User.Token)
	t.AssertEqual("user1", UserJSON.User.Username)
	t.AssertEqual(bodyUser.User.Email, UserJSON.User.Email)
}

func (t *AppTest) TestLoginFail() {
	tests := []testLogin{
		testLogin{
			errorKey: "email",
			message:  models.EMPTY_MSG,
			body: UserLoginBody{
				UserLogin{
					Email:    "",
					Password: "password",
				},
			},
		},
		testLogin{
			errorKey: "password",
			message:  models.EMPTY_MSG,
			body: UserLoginBody{
				UserLogin{
					Email:    "user1@example.com",
					Password: "",
				},
			},
		},
	}

	for _, test := range tests {
		jsonBody, _ := json.Marshal(test.body)

		t.MakePostRequest("/api/users/login", bytes.NewBuffer(jsonBody), nil)
		t.AssertStatus(422)

		var ErrorJSON = ErrorJSON{}
		json.Unmarshal(t.ResponseBody, &ErrorJSON)

		msg, ok := ErrorJSON.Errors[test.errorKey]
		t.Assert(ok)
		t.AssertEqual(test.message, msg[0])
	}

	jsonBody, _ := json.Marshal(UserLoginBody{})

	t.MakePostRequest("/api/users/login", bytes.NewBuffer(jsonBody), nil)
	t.AssertStatus(422)

	var ErrorJSON = ErrorJSON{}
	json.Unmarshal(t.ResponseBody, &ErrorJSON)

	var errorKeys = []string{"email", "password"}
	for _, errorKey := range errorKeys {
		msg, ok := ErrorJSON.Errors[errorKey]
		t.Assert(ok)
		t.AssertEqual(models.EMPTY_MSG, msg[0])
	}
}

func (t *AppTest) TestRegistrationSuccesFully() {
	bodyUser := UserRegistrationBody{
		UserRegister{
			Username: "newuser",
			Email:    "newuser@example.com",
			Password: "password",
		},
	}

	jsonBody, _ := json.Marshal(bodyUser)

	t.MakePostRequest("/api/users", bytes.NewBuffer(jsonBody), nil)
	t.AssertOk()

	var UserJSON = controllers.UserJSON{}
	json.Unmarshal(t.ResponseBody, &UserJSON)

	t.AssertEqual(JWT.NewToken("newuser"), UserJSON.User.Token)
	t.AssertEqual(bodyUser.User.Username, UserJSON.User.Username)
	t.AssertEqual(bodyUser.User.Email, UserJSON.User.Email)
}

func (t *AppTest) TestRegistrationFail() {
	tests := []testRegistration{
		testRegistration{
			errorKey: "username",
			message:  models.EMPTY_MSG,
			body: UserRegistrationBody{
				UserRegister{
					Username: "",
					Email:    "newuser@example.com",
					Password: "password",
				},
			},
		},
		testRegistration{
			errorKey: "email",
			message:  models.EMPTY_MSG,
			body: UserRegistrationBody{
				UserRegister{
					Username: "newuser",
					Email:    "",
					Password: "password",
				},
			},
		},
		testRegistration{
			errorKey: "password",
			message:  models.EMPTY_MSG,
			body: UserRegistrationBody{
				UserRegister{
					Username: "newuser",
					Email:    "newuser@example.com",
					Password: "",
				},
			},
		},
	}

	for _, test := range tests {
		jsonBody, _ := json.Marshal(test.body)

		t.MakePostRequest("/api/users", bytes.NewBuffer(jsonBody), nil)
		t.AssertStatus(422)

		var ErrorJSON = ErrorJSON{}
		json.Unmarshal(t.ResponseBody, &ErrorJSON)

		msg, ok := ErrorJSON.Errors[test.errorKey]
		t.Assert(ok)
		t.AssertEqual(test.message, msg[0])
	}

	jsonBody, _ := json.Marshal(UserRegistrationBody{})

	t.MakePostRequest("/api/users", bytes.NewBuffer(jsonBody), nil)
	t.AssertStatus(422)

	var ErrorJSON = ErrorJSON{}
	json.Unmarshal(t.ResponseBody, &ErrorJSON)

	var errorKeys = []string{"username", "email", "password"}
	for _, errorKey := range errorKeys {
		msg, ok := ErrorJSON.Errors[errorKey]
		t.Assert(ok)
		t.AssertEqual(models.EMPTY_MSG, msg[0])
	}
}

func (t *UserControllerTest) TestGetCurrentUserUnauthorized() {
	t.Get("/api/user")
	t.AssertStatus(401)
}

func (t *UserControllerTest) TestGetCurrentSuccess() {
	request := t.GetCustom(t.BaseUrl() + "/api/user")

	request.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	}
	request.Send()
	t.AssertOk()
}
