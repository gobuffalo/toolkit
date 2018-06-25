package vcs

import (
	"context"
)

const (
	PLUGIN     = "plugin"
	GENERATOR  = "generator"
	MIDDLEWARE = "middleware"
	EXAMPLE    = "example"
	OTHER      = "other"
)

type VCS interface {
	Name() string
	Search(context.Context, ...string) ([]Repository, error)
}
