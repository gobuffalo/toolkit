package discovery

import (
	"context"
)

const (
	PLUGIN     = "plugin"
	GENERATOR  = "generator"
	MIDDLEWARE = "middleware"
	POP        = "pop"
	TEMPLATING = "templating"
	GRIFTS     = "grifts"
	DEPLOYMENT = "deployment"
	WEBPACK    = "webpack"
	EXAMPLE    = "example"
	OTHER      = "other"
)

var KnownTags = []struct {
	Tag         string
	Description string
}{
	{PLUGIN, "Plugins"},
	{GENERATOR, "Generators"},
	{MIDDLEWARE, "Middleware"},
	{POP, "Pop/Soda"},
	{TEMPLATING, "Templating"},
	{GRIFTS, "Grift Tasks"},
	{DEPLOYMENT, "Deployment"},
	{WEBPACK, "Webpack/Front-End"},
}

type Engine interface {
	Name() string
	Search(context.Context, ...string) ([]Project, error)
}
