package tests

import (
	"encoding/json"

	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/controllers"
)

type TagControllerTest struct {
	AppTest
}

func (t *TagControllerTest) TestGetTags() {
	t.Get("/api/tags")
	t.AssertOk()

	var TagJSON = controllers.TagJSON{}
	json.Unmarshal(t.ResponseBody, &TagJSON)

	t.AssertEqual(3, len(TagJSON.Tags))
	t.AssertEqual([]string{"tag1", "tag2", "tag3"}, TagJSON.Tags)
}
