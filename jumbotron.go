package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Jumbotron ...
func Jumbotron(fluid bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &jumbotron{Fluid: fluid, Children: children}
}

type jumbotron struct {
	vecty.Core
	Fluid    bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *jumbotron) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"jumbotron":       true,
				"jumbotron-fluid": c.Fluid,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}
