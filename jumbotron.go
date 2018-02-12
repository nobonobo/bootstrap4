package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Jumbotron ...
type Jumbotron struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Fluid    bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Jumbotron) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"jumbotron":       true,
				"jumbotron-fluid": c.Fluid,
			},
		),
		c.Markup,
		c.Children,
	)
}
