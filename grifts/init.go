package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/toolkit/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
