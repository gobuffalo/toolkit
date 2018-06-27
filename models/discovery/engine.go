package discovery

import "context"

type Service interface {
	Name() string
	Search(context.Context, ...string) ([]Project, error)
}
