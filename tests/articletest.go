package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/controllers"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
)

type ArticleController struct {
	AppTest
}

type M map[string]interface{}

type testValidation struct {
	errorKey string
	message  string
	body     M
}

func (t *ArticleController) TestGetArticles() {
	t.Get("/api/articles")
	t.AssertOk()

	var ArticlesJSON = controllers.ArticlesJSON{}
	json.Unmarshal(t.ResponseBody, &ArticlesJSON)

	t.AssertEqual(5, len(ArticlesJSON.Articles))
	t.AssertEqual(articles[4].Title, ArticlesJSON.Articles[0].Title)
	t.AssertEqual(articles[4].Body, ArticlesJSON.Articles[0].Body)
	t.AssertEqual(articles[4].Description, ArticlesJSON.Articles[0].Description)
	t.AssertEqual(articles[4].User.Username, ArticlesJSON.Articles[0].Author.Username)
	t.AssertEqual(articles[4].User.Bio, ArticlesJSON.Articles[0].Author.Bio)
	t.AssertEqual(false, ArticlesJSON.Articles[0].Favorited)
}

func (t *ArticleController) TestGetArticlesByTag() {
	t.Get("/api/articles?tag=" + tags[0].Name)
	t.AssertOk()

	var ArticlesJSON = controllers.ArticlesJSON{}
	json.Unmarshal(t.ResponseBody, &ArticlesJSON)

	t.AssertEqual(5, len(ArticlesJSON.Articles))
	t.AssertEqual(articles[4].Title, ArticlesJSON.Articles[0].Title)
	t.AssertEqual(articles[4].Body, ArticlesJSON.Articles[0].Body)
	t.AssertEqual(articles[4].Description, ArticlesJSON.Articles[0].Description)
	t.AssertEqual(articles[4].User.Username, ArticlesJSON.Articles[0].Author.Username)
	t.AssertEqual(articles[4].User.Bio, ArticlesJSON.Articles[0].Author.Bio)
	t.AssertEqual(articles[4].Tags[0].Name, ArticlesJSON.Articles[0].TagList[0])
}

func (t *ArticleController) TestGetArticlesFavoritedBy() {
	t.Get("/api/articles?favorited=" + users[7].Username)
	t.AssertOk()

	var ArticlesJSON = controllers.ArticlesJSON{}
	json.Unmarshal(t.ResponseBody, &ArticlesJSON)

	t.AssertEqual(5, len(ArticlesJSON.Articles))
	t.AssertEqual(articles[4].Title, ArticlesJSON.Articles[0].Title)
	t.AssertEqual(articles[4].Body, ArticlesJSON.Articles[0].Body)
	t.AssertEqual(articles[4].Description, ArticlesJSON.Articles[0].Description)
	t.AssertEqual(articles[4].User.Username, ArticlesJSON.Articles[0].Author.Username)
	t.AssertEqual(articles[4].User.Bio, ArticlesJSON.Articles[0].Author.Bio)
	t.AssertEqual(articles[4].Tags[0].Name, ArticlesJSON.Articles[0].TagList[0])
	t.AssertEqual(false, ArticlesJSON.Articles[0].Favorited)

	t.AssertEqual(articles[3].Title, ArticlesJSON.Articles[1].Title)
	t.AssertEqual(articles[3].Body, ArticlesJSON.Articles[1].Body)
	t.AssertEqual(articles[3].Description, ArticlesJSON.Articles[1].Description)
	t.AssertEqual(articles[3].User.Username, ArticlesJSON.Articles[1].Author.Username)
	t.AssertEqual(articles[3].User.Bio, ArticlesJSON.Articles[1].Author.Bio)
	t.AssertEqual(articles[3].Tags[0].Name, ArticlesJSON.Articles[1].TagList[0])
	t.AssertEqual(false, ArticlesJSON.Articles[1].Favorited)

	t.AssertEqual(articles[2].Title, ArticlesJSON.Articles[2].Title)
	t.AssertEqual(articles[2].Body, ArticlesJSON.Articles[2].Body)
	t.AssertEqual(articles[2].Description, ArticlesJSON.Articles[2].Description)
	t.AssertEqual(articles[2].User.Username, ArticlesJSON.Articles[2].Author.Username)
	t.AssertEqual(articles[2].User.Bio, ArticlesJSON.Articles[2].Author.Bio)
	t.AssertEqual(articles[2].Tags[0].Name, ArticlesJSON.Articles[2].TagList[0])
	t.AssertEqual(false, ArticlesJSON.Articles[2].Favorited)
}

func (t *ArticleController) TestGetArticlesMarkedAsFavorite() {
	request := t.GetCustom(t.BaseUrl() + "/api/articles?favorited=" + users[7].Username)
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[7].Username))},
	}
	request.Send()

	t.AssertOk()

	var ArticlesJSON = controllers.ArticlesJSON{}
	json.Unmarshal(t.ResponseBody, &ArticlesJSON)

	t.AssertEqual(5, len(ArticlesJSON.Articles))
	t.AssertEqual(articles[4].Title, ArticlesJSON.Articles[0].Title)
	t.AssertEqual(articles[4].Body, ArticlesJSON.Articles[0].Body)
	t.AssertEqual(articles[4].Description, ArticlesJSON.Articles[0].Description)
	t.AssertEqual(articles[4].User.Username, ArticlesJSON.Articles[0].Author.Username)
	t.AssertEqual(articles[4].User.Bio, ArticlesJSON.Articles[0].Author.Bio)
	t.AssertEqual(articles[4].Tags[0].Name, ArticlesJSON.Articles[0].TagList[0])
	t.AssertEqual(true, ArticlesJSON.Articles[0].Favorited)

	t.AssertEqual(articles[3].Title, ArticlesJSON.Articles[1].Title)
	t.AssertEqual(articles[3].Body, ArticlesJSON.Articles[1].Body)
	t.AssertEqual(articles[3].Description, ArticlesJSON.Articles[1].Description)
	t.AssertEqual(articles[3].User.Username, ArticlesJSON.Articles[1].Author.Username)
	t.AssertEqual(articles[3].User.Bio, ArticlesJSON.Articles[1].Author.Bio)
	t.AssertEqual(articles[3].Tags[0].Name, ArticlesJSON.Articles[1].TagList[0])
	t.AssertEqual(true, ArticlesJSON.Articles[1].Favorited)

	t.AssertEqual(articles[2].Title, ArticlesJSON.Articles[2].Title)
	t.AssertEqual(articles[2].Body, ArticlesJSON.Articles[2].Body)
	t.AssertEqual(articles[2].Description, ArticlesJSON.Articles[2].Description)
	t.AssertEqual(articles[2].User.Username, ArticlesJSON.Articles[2].Author.Username)
	t.AssertEqual(articles[2].User.Bio, ArticlesJSON.Articles[2].Author.Bio)
	t.AssertEqual(articles[2].Tags[0].Name, ArticlesJSON.Articles[2].TagList[0])
	t.AssertEqual(true, ArticlesJSON.Articles[2].Favorited)
}

func (t *ArticleController) TestGetArticlesAuthoredBy() {
	t.Get("/api/articles?author=" + users[0].Username)
	t.AssertOk()

	var ArticlesJSON = controllers.ArticlesJSON{}
	json.Unmarshal(t.ResponseBody, &ArticlesJSON)

	t.AssertEqual(1, len(ArticlesJSON.Articles))
	t.AssertEqual(articles[0].Title, ArticlesJSON.Articles[0].Title)
	t.AssertEqual(articles[0].Description, ArticlesJSON.Articles[0].Description)
	t.AssertEqual(articles[0].Body, ArticlesJSON.Articles[0].Body)
	t.AssertEqual(articles[0].User.Username, ArticlesJSON.Articles[0].Author.Username)
	t.AssertEqual(articles[0].User.Bio, ArticlesJSON.Articles[0].Author.Bio)
	t.AssertEqual(articles[0].Tags[0].Name, ArticlesJSON.Articles[0].TagList[0])
}

func (t *ArticleController) TestGetArticle() {
	t.Get("/api/articles/" + articles[0].Slug)
	t.AssertOk()

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.AssertEqual(articles[0].Title, ArticleJSON.Article.Title)
	t.AssertEqual(articles[0].Body, ArticleJSON.Body)
	t.AssertEqual(articles[0].Description, ArticleJSON.Description)
	t.AssertEqual(articles[0].User.Username, ArticleJSON.Author.Username)
	t.AssertEqual(articles[0].User.Bio, ArticleJSON.Author.Bio)
	t.AssertEqual(articles[0].Tags[0].Name, ArticleJSON.TagList[0])
}

func (t *ArticleController) TestGetArticleNotFound() {
	t.Get("/api/articles/not-found-slug")
	t.AssertNotFound()
}

func (t *ArticleController) TestCreateArticleSucceed() {
	jsonBody, _ := json.Marshal(controllers.ArticleJSON{
		controllers.Article{
			Title:       "New Article Title",
			Description: "New Article Description",
			Body:        "New Article Body",
			TagList:     []string{"tag2"},
		},
	})
	t.MakePostRequest("/api/articles", bytes.NewBuffer(jsonBody), http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	})

	t.AssertStatus(http.StatusCreated)

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.AssertEqual("new-article-title", ArticleJSON.Article.Slug)
	t.AssertEqual("New Article Title", ArticleJSON.Article.Title)
	t.AssertEqual("New Article Description", ArticleJSON.Article.Description)
	t.AssertEqual("New Article Body", ArticleJSON.Article.Body)
	t.AssertEqual([]string{"tag2"}, ArticleJSON.Article.TagList)
}

func (t *ArticleController) TestCreateArticleValidations() {
	tests := []testValidation{
		testValidation{
			errorKey: "title",
			message:  models.EMPTY_MSG,
			body: M{
				"article": M{
					"title":       "",
					"description": "Description",
					"body":        "Body",
				},
			},
		},
		testValidation{
			errorKey: "description",
			message:  models.EMPTY_MSG,
			body: M{
				"article": M{
					"title":       "Title",
					"description": "",
					"body":        "Body",
				},
			},
		},
		testValidation{
			errorKey: "body",
			message:  models.EMPTY_MSG,
			body: M{
				"article": M{
					"title":       "Title",
					"description": "Description",
					"body":        "",
				},
			},
		},
	}

	for _, test := range tests {
		jsonBody, _ := json.Marshal(test.body)
		t.MakePostRequest("/api/articles", bytes.NewBuffer(jsonBody), http.Header{
			"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[6].Username))},
		})
		t.AssertStatus(http.StatusUnprocessableEntity)
		var ErrorJSON = ErrorJSON{}
		json.Unmarshal(t.ResponseBody, &ErrorJSON)
		msgs, present := ErrorJSON.Errors[test.errorKey]
		t.Assert(present)
		t.AssertEqual(test.message, msgs[0])
	}

	jsonBody, _ := json.Marshal(M{
		"article": M{
			"title":       "",
			"description": "",
			"body":        "",
		},
	})

	t.MakePostRequest("/api/articles", bytes.NewBuffer(jsonBody), http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[6].Username))},
	})

	t.AssertStatus(http.StatusUnprocessableEntity)

	var ErrorJSON = ErrorJSON{}
	json.Unmarshal(t.ResponseBody, &ErrorJSON)

	for _, errorKey := range []string{"title", "description", "body"} {
		msgs, present := ErrorJSON.Errors[errorKey]
		t.Assert(present)
		t.AssertEqual(models.EMPTY_MSG, msgs[0])
	}
}

func (t *ArticleController) TestCreateArticleUnauthorized() {
	jsonBody, _ := json.Marshal(controllers.ArticleJSON{})
	t.Post("/api/articles", "application/json", bytes.NewBuffer(jsonBody))
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *ArticleController) TestUpdateArticleSucceed() {
	for _, field := range []string{"title", "description", "body"} {
		jsonBody, _ := json.Marshal(M{
			"article": M{
				field: "Updated",
			},
		})
		request := t.PutCustom(t.BaseUrl()+"/api/articles/"+articles[0].Slug, "application/json", bytes.NewBuffer(jsonBody))
		request.Header = http.Header{
			"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(articles[0].User.Username))},
		}
		request.Send()

		t.AssertOk()
		var ArticleJSON = controllers.ArticleJSON{}
		json.Unmarshal(t.ResponseBody, &ArticleJSON)

		value := reflect.ValueOf(ArticleJSON.Article)
		t.AssertEqual("updated", ArticleJSON.Article.Slug)
		t.AssertEqual(value.FieldByName(strings.Title(field)).String(), "Updated")

		DB.First(&articles[0])
	}
}

func (t *ArticleController) TestUpdateArticleUnauthorized() {
	jsonBody, _ := json.Marshal(controllers.ArticleJSON{})
	t.Put("/api/articles/"+articles[0].Slug, "application/json", bytes.NewBuffer(jsonBody))
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *ArticleController) TestUpdateArticleForbidden() {
	jsonBody, _ := json.Marshal(controllers.ArticleJSON{})
	request := t.PutCustom(t.BaseUrl()+"/api/articles/"+articles[2].Slug, "application/json", bytes.NewBuffer(jsonBody))
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(articles[1].User.Username))},
	}
	request.Send()
	t.AssertStatus(http.StatusForbidden)
}

func (t *ArticleController) TestDeleteArticleUnauthorized() {
	t.Delete("/api/articles/" + articles[0].Slug)
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *ArticleController) TestDeleteArticleForbidden() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[0].Slug)
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(articles[1].User.Username))},
	}
	request.Send()
	t.AssertStatus(http.StatusForbidden)
}

func (t *ArticleController) TestDeleteArticleSucceed() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[3].Slug)
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(articles[3].User.Username))},
	}
	request.Send()
	t.AssertStatus(http.StatusNoContent)
	DB.Save(&articles[0])
}
