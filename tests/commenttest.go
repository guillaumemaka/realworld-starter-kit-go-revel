package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/controllers"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
)

type CommentController struct {
	AppTest
}

func (t *CommentController) TestGetComments() {
	t.Get("/api/articles/" + articles[0].Slug + "/comments")
	t.AssertOk()

	var CommentsJSON = controllers.CommentsJSON{}
	json.Unmarshal(t.ResponseBody, &CommentsJSON)

	t.AssertEqual(3, len(CommentsJSON.Comments))
}

func (t *CommentController) TestCreateCommentSucceed() {
	jsonBody, _ := json.Marshal(M{
		"comment": M{
			"body": "New comment",
		},
	})

	t.MakePostRequest("/api/articles/"+articles[0].Slug+"/comments", bytes.NewBuffer(jsonBody), http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	})

	t.AssertOk()

	var CommentJSON = controllers.CommentJSON{}
	json.Unmarshal(t.ResponseBody, &CommentJSON)

	t.AssertEqual("New comment", CommentJSON.Comment.Body)
	t.AssertEqual(users[0].Username, CommentJSON.Comment.Author.Username)
}

func (t *CommentController) TestCreateCommentValidation() {
	jsonBody, _ := json.Marshal(M{
		"comment": M{
			"body": "",
		},
	})

	t.MakePostRequest("/api/articles/"+articles[0].Slug+"/comments", bytes.NewBuffer(jsonBody), http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	})

	t.AssertStatus(http.StatusUnprocessableEntity)

	var ErrorJSON = ErrorJSON{}
	json.Unmarshal(t.ResponseBody, &ErrorJSON)

	msgs, present := ErrorJSON.Errors["body"]
	t.Assert(present)
	t.AssertEqual(models.EMPTY_MSG, msgs[0])
}

func (t *CommentController) TestCreateCommentUnauthorized() {
	t.Post("/api/articles/"+articles[0].Slug+"/comments", "application/json", nil)
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *CommentController) TestDeleteCommentForbidden() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[0].Slug + "/comments/" + strconv.Itoa(articles[0].Comments[0].ID))
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	}
	request.Send()
	t.AssertStatus(http.StatusForbidden)
}

func (t *CommentController) TestDeleteCommentSucceed() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[0].Slug + "/comments/" + strconv.Itoa(articles[0].Comments[0].ID))
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(articles[0].Comments[0].User.Username))},
	}
	request.Send()

	t.AssertOk()
}

func (t *CommentController) TestDeleteCommentUnauthorized() {
	t.Delete("/api/articles/" + articles[0].Slug + "/comments/" + strconv.Itoa(articles[0].Comments[0].ID))
	t.AssertStatus(http.StatusUnauthorized)
}
