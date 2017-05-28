package tests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/controllers"
)

type FavoriteControllerTest struct {
	AppTest
}

func (t *FavoriteControllerTest) TestFavoriteSucceed() {
	t.MakePostRequest("/api/articles/"+articles[2].Slug+"/favorite", nil, http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	})
	t.AssertOk()

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.Assert(ArticleJSON.Article.Favorited)
	t.AssertEqual(4, ArticleJSON.FavoritesCount)
}

func (t *FavoriteControllerTest) TestFavoriteUnauthorized() {
	t.Post("/api/articles/"+articles[3].Slug+"/favorite", "application/json", nil)
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *FavoriteControllerTest) TestFavoriteAlreadyFavorited() {
	t.MakePostRequest("/api/articles/"+articles[1].Slug+"/favorite", nil, http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[7].Username))},
	})
	t.AssertStatus(http.StatusUnprocessableEntity)

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.Assert(ArticleJSON.Article.Favorited)
	t.AssertEqual(3, ArticleJSON.FavoritesCount)
}

func (t *FavoriteControllerTest) TestUnfavoriteSucceed() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[3].Slug + "/favorite")
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[7].Username))},
	}
	request.Send()

	t.AssertOk()

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.AssertEqual(false, ArticleJSON.Article.Favorited)
	t.AssertEqual(2, ArticleJSON.FavoritesCount)
}

func (t *FavoriteControllerTest) TestUnfavoriteUnauthorized() {
	t.Delete("/api/articles/" + articles[3].Slug + "/favorite")
	t.AssertStatus(http.StatusUnauthorized)
}

func (t *FavoriteControllerTest) TestUnfavoriteNotYetFavorited() {
	request := t.DeleteCustom(t.BaseUrl() + "/api/articles/" + articles[3].Slug + "/favorite")
	request.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Token %v", JWT.NewToken(users[0].Username))},
	}
	request.Send()

	t.AssertStatus(http.StatusUnprocessableEntity)

	var ArticleJSON = controllers.ArticleJSON{}
	json.Unmarshal(t.ResponseBody, &ArticleJSON)

	t.AssertEqual(false, ArticleJSON.Article.Favorited)
	t.AssertEqual(3, ArticleJSON.FavoritesCount)
}
