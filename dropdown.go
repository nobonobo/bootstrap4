package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// DropdownMenu ...
type DropdownMenu struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	For      string                `vecty:"prop"`
	Right    bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownMenu) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"dropdown-menu":       true,
				"dropdown-menu-right": c.Right,
			},
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
	ID       string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownLinkItem) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
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
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *DropdownHeader) Render() vecty.ComponentOrHTML {
	return elem.Heading6(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("dropdown-header"),
		),
		c.Markup,
		c.Children,
	)
}
