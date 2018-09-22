package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/openware/barong/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
