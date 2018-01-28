package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Input ...
type Input struct {
	vecty.Core
	Type     prop.InputType
	ID       string
	Name     string
	Label    string
	Disabled bool
	Checked  bool
}

// Render ...
func (c *Input) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Label(),
		elem.Input(),
	)
}
