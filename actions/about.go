package actions

import (
	"github.com/gobuffalo/buffalo"
)

// AboutHandler returns a template for information about this demo app
func AboutHandler(c buffalo.Context) error {

	return c.Render(200, r.HTML("about.html"))
}
