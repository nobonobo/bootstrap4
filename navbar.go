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
	ID          string                `vecty:"prop"`
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
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"navbar":                           true,
				"navbar-" + c.Size.String():        !c.Expand && len(c.Size) > 0,
				"navbar-expand-" + c.Size.String(): c.Expand && len(c.Size) > 0,
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
	ID     string           `vecty:"prop"`
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
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Type) > 0, prop.Type(c.Type)),
			vecty.ClassMap{
				"navbar-toggler": true,
				"bg-light":       c.Light,
				"navbar-light":   c.Light,
				"bg-dark":        c.Dark,
				"navbar-dark":    c.Dark,
			},
			vecty.Data("toggle", "collapse"),
			vecty.MarkupIf(len(c.Target) > 0, vecty.Data("target", c.Target)),
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
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"navbar-collapse": true,
				"collapse":        true,
				"bg-light":        c.Light,
				"navbar-light":    c.Light,
				"bg-dark":         c.Dark,
				"navbar-dark":     c.Dark,
			},
		),
		c.Markup,
		c.Children,
	)
}

// NavbarNav ...
type NavbarNav struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavbarNav) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"navbar-nav": true,
			},
		),
		c.Markup,
		c.Children,
	)
}

// NavbarBrand ...
type NavbarBrand struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavbarBrand) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"navbar-brand": true,
			},
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
		c.Markup,
		c.Children,
	)
}
