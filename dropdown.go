package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// DropdownMenu ...
type DropdownMenu struct {
	vecty.Core
	For      string                `vecty:"prop"`
	Right    bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownMenu) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"dropdown-menu":       true,
				"dropdown-menu-right": c.Right,
			},
			vecty.Attribute("aria-labelledby", c.For),
		),
		c.Markup,
		c.Children,
	)
}

// DropdownDivider ...
type DropdownDivider struct {
	vecty.Core
}

// Render ...
func (c *DropdownDivider) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Class("dropdown-divider")))
}

// DropdownLinkItem ...
type DropdownLinkItem struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownLinkItem) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.Class("dropdown-item"),
		),
		c.Markup,
		c.Children,
	)
}

// DropdownHeader ...
type DropdownHeader struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownHeader) Render() vecty.ComponentOrHTML {
	return elem.Heading6(
		vecty.Markup(
			vecty.Class("dropdown-header"),
		),
		c.Markup,
		c.Children,
	)
}
