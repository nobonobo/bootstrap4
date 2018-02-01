package bootstrap4

//<nav class="navbar navbar-expand-lg navbar-light bg-light">

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Navbar ...
type Navbar struct {
	vecty.Core
	Size        Size                  `vecty:"prop"`
	Expand      bool                  `vecty:"prop"`
	Light       bool                  `vecty:"prop"`
	Dark        bool                  `vecty:"prop"`
	StickyTop   bool                  `vecty:"prop"`
	FixedTop    bool                  `vecty:"prop"`
	FixedBottom bool                  `vecty:"prop"`
	Markup      vecty.MarkupList      `vecty:"prop"`
	Children    vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Navbar) Render() vecty.ComponentOrHTML {
	return elem.Navigation(
		vecty.Markup(
			vecty.ClassMap{
				"navbar":                           true,
				"navbar-" + c.Size.String():        !c.Expand,
				"navbar-expand-" + c.Size.String(): c.Expand,
				"bg-light":                         c.Light,
				"navbar-light":                     c.Light,
				"bg-dark":                          c.Dark,
				"navbar-dark":                      c.Dark,
				"sticky-top":                       c.StickyTop,
				"fixed-top":                        c.FixedTop,
				"fixed-bottom":                     c.FixedBottom,
			},
		),
		c.Markup,
		c.Children,
	)
}

// NavbarToggler ...
type NavbarToggler struct {
	vecty.Core
	Type   prop.InputType   `vecty:"prop"`
	Target string           `vecty:"prop"`
	Light  bool             `vecty:"prop"`
	Dark   bool             `vecty:"prop"`
	Markup vecty.MarkupList `vecty:"prop"`
}

// Render ...
func (c *NavbarToggler) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.ClassMap{
				"navbar-toggler": true,
				"bg-light":       c.Light,
				"navbar-light":   c.Light,
				"bg-dark":        c.Dark,
				"navbar-dark":    c.Dark,
			},
			prop.Type(c.Type),
			vecty.Data("toggle", "collapse"),
			vecty.Data("target", c.Target),
		),
		c.Markup,
		elem.Span(
			vecty.Markup(vecty.Class("navbar-toggler-icon")),
		),
	)
}

// NavbarCollapse ...
type NavbarCollapse struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Light    bool                  `vecty:"prop"`
	Dark     bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavbarCollapse) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.ClassMap{
				"navbar-collapse": true,
				"collapse":        true,
				"bg-light":        c.Light,
				"navbar-light":    c.Light,
				"bg-dark":         c.Dark,
				"navbar-dark":     c.Dark,
			},
			prop.ID(c.ID),
		),
		c.Markup,
		c.Children,
	)
}
