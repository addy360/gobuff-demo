package grifts

import (
  "github.com/gobuffalo/buffalo"
	"buffdemo/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
