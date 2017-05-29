package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

type TagController struct {
	ApplicationController
}

type TagJSON struct {
	Tags []string `json:"tags"`
}

func (c *TagController) Index() revel.Result {
	tags, err := c.DB.FindTags()

	if err != nil {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderText(http.StatusText(http.StatusUnprocessableEntity))
	}

	var tagJson = TagJSON{}
	for _, tag := range tags {
		tagJson.Tags = append(tagJson.Tags, tag.Name)
	}

	return c.RenderJSON(tagJson)
}
