package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// DropdownMenu ...
func DropdownMenu(right bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &dropdownMenu{
		Right:    right,
		Children: children,
	}
}

type dropdownMenu struct {
	vecty.Core
	Right    bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *dropdownMenu) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"dropdown-menu":       true,
				"dropdown-menu-right": c.Right,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// DropdownDivider ...
func DropdownDivider() vecty.Component {
	return &dropdownDivider{}
}

type dropdownDivider struct {
	vecty.Core
}

func (c *dropdownDivider) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("dropdown-divider")))
}

// DropdownLinkItem ...
func DropdownLinkItem(href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &dropdownLinkItem{
		Href:     href,
		Children: children,
	}
}

type dropdownLinkItem struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *dropdownLinkItem) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.Class("dropdown-item"),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}

// DropdownHeader ...
func DropdownHeader(children ...vecty.MarkupOrChild) vecty.Component {
	return &dropdownHeader{
		Children: children,
	}
}

type dropdownHeader struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *dropdownHeader) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("dropdown-header"),
		),
	}
	return elem.Heading6(append(markup, c.Children...)...)
}
