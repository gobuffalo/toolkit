package grifts

import (
	"os"
	"os/exec"

	. "github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = Add("release", func(c *Context) error {
	if err := Run("generate:readme", c); err != nil {
		return errors.WithStack(err)
	}
	cmd := exec.CommandContext(c, "buffalo", "heroku", "deploy")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
})
