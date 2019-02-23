package actions

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/toolkit/models/discovery"
)

var r *render.Engine
var assetsBox = packr.New("../public", "../public")

var Helpers = render.Helpers{
	"knownTagsMD":   knownTagsMD,
	"topicLinks":    topicLinks,
	"knownTagsHTML": knownTagsHTML,
}

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("../templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: Helpers,
	})
}

func knownTagsMD() string {
	bb := &bytes.Buffer{}
	for _, tag := range discovery.KnownTags {
		bb.WriteString("* `" + tag.Tag + "` - " + tag.Description + "\n")
	}

	return bb.String()
}

func knownTagsHTML() (template.HTML, error) {
	t := tags.New("span", nil)
	var as []string
	for _, tag := range discovery.KnownTags {
		as = append(as, tags.New("a", tags.Options{
			"href":  "/tools?topic=" + tag.Tag,
			"body":  tag.Tag,
			"class": "category-link",
		}).String())
	}

	t.Append(strings.Join(as, " "))
	return t.HTML(), nil
}

func topicLinks(topics []string) (template.HTML, error) {
	t := tags.New("span", nil)
	var as []string
	for _, topic := range topics {
		as = append(as, tags.New("a", tags.Options{
			"href": "/tools?topic=" + topic,
			"body": topic,
		}).String())
	}
	t.Append(strings.Join(as, " "))
	return t.HTML(), nil
}
