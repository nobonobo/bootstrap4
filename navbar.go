package bootstrap4

//<nav class="navbar navbar-expand-lg navbar-light bg-light">

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Navbar ...
func Navbar(size Size, expand, light, dark, stickytop, fixedtop, fixedbottom bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &navbar{
		Size:        size,
		Expand:      expand,
		Light:       light,
		Dark:        dark,
		StickyTop:   stickytop,
		FixedTop:    fixedtop,
		FixedBottom: fixedbottom,
		Children:    children,
	}
}

type navbar struct {
	vecty.Core
	Size        Size                  `vecty:"prop"`
	Expand      bool                  `vecty:"prop"`
	Light       bool                  `vecty:"prop"`
	Dark        bool                  `vecty:"prop"`
	StickyTop   bool                  `vecty:"prop"`
	FixedTop    bool                  `vecty:"prop"`
	FixedBottom bool                  `vecty:"prop"`
	Children    []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navbar) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
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
	}
	return elem.Navigation(append(markup, c.Children...)...)
}

// NavbarToggler ...
func NavbarToggler(tp prop.InputType, light, dark bool) vecty.Component {
	return &navbarToggler{
		Type:  tp,
		Light: light,
		Dark:  dark,
	}
}

type navbarToggler struct {
	vecty.Core
	Type   prop.InputType `vecty:"prop"`
	Target string         `vecty:"prop"`
	Light  bool           `vecty:"prop"`
	Dark   bool           `vecty:"prop"`
}

func (c *navbarToggler) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
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
		elem.Span(
			vecty.Markup(vecty.Class("navbar-toggler-icon")),
		),
	}
	return elem.Button(markup...)
}

// NavbarCollapse ...
func NavbarCollapse(light, dark bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &navbarCollapse{
		Light:    light,
		Dark:     dark,
		Children: children,
	}
}

type navbarCollapse struct {
	vecty.Core
	Light    bool                  `vecty:"prop"`
	Dark     bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navbarCollapse) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"navbar-collapse": true,
				"collapse":        true,
				"bg-light":        c.Light,
				"navbar-light":    c.Light,
				"bg-dark":         c.Dark,
				"navbar-dark":     c.Dark,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// NavbarNav ...
func NavbarNav(children ...vecty.MarkupOrChild) vecty.Component {
	return &navbarNav{
		Children: children,
	}
}

type navbarNav struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navbarNav) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"navbar-nav": true,
			},
		),
	}
	return elem.UnorderedList(append(markup, c.Children...)...)
}

// NavbarBrand ...
func NavbarBrand(href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &navbarBrand{
		Href:     href,
		Children: children,
	}
}

type navbarBrand struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navbarBrand) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"navbar-brand": true,
			},
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}
