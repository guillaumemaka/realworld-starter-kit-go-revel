package tests

type TagControllerTest struct {
	AppTest
}

func (t *TagControllerTest) TestGetTags() {
	t.Get("/api/articles/tags")
	t.AssertOk()
}
