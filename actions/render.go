package actions

import (
	"bytes"
	"html/template"
	"reflect"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"loop": loop,
		},
	})
}

func loop(name string, slice interface{}, help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", errors.New("loop requires a block")
	}
	bb := &bytes.Buffer{}

	rv := reflect.Indirect(reflect.ValueOf(slice))
	if rv.Kind() != reflect.Slice {
		return "", errors.Errorf("slice must type of slice")
	}
	for i := 0; i < rv.Len(); i++ {
		v := rv.Index(i).Interface()
		ctx := help.Context.New()
		ctx.Set(name, v)
		s, err := help.BlockWith(ctx)
		if err != nil {
			return "", errors.WithStack(err)
		}
		bb.WriteString(s)
	}

	return template.HTML(bb.String()), nil
}
