package grifts

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/toolkit/actions"
	. "github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = Namespace("generate", func() {

	Desc("readme", "generates the README.md from templates/README.md")
	Add("readme", func(c *Context) error {
		plush.Helpers.AddMany(actions.Helpers)
		ctx := plush.NewContextWithContext(c)
		tf, err := ioutil.ReadFile(filepath.Join("templates", "README.md"))
		if err != nil {
			return errors.WithStack(err)
		}

		ctx.Set("markdown", func(s string) string {
			return s
		})
		s, err := plush.Render(string(tf), ctx)
		if err != nil {
			return errors.WithStack(err)
		}

		mf, err := os.Create(filepath.Join("README.md"))
		if err != nil {
			return errors.WithStack(err)
		}
		defer mf.Close()
		_, err = mf.WriteString(s)
		return err
	})

})
