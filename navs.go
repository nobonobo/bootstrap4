package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// NavList ...
type NavList struct {
	vecty.Core
	ID        string                `vecty:"prop"`
	Role      string                `vecty:"prop"`
	Tabs      bool                  `vecty:"prop"`
	Pills     bool                  `vecty:"prop"`
	Fill      bool                  `vecty:"prop"`
	Justified bool                  `vecty:"prop"`
	Markup    vecty.MarkupList      `vecty:"prop"`
	Children  vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavList) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"nav":           true,
				"nav-tabs":      c.Tabs,
				"nav-pills":     c.Pills,
				"nav-fill":      c.Fill,
				"nav-justified": c.Justified,
			},
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
		),
		c.Markup,
		c.Children,
	)
}

// NavItem ...
type NavItem struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Dropdown bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"nav-item": true,
				"active":   c.Active,
				"dropdown": c.Dropdown,
			},
		),
		c.Markup,
		c.Children,
	)
}

// Nav ...
type Nav struct {
	vecty.Core
	ID        string                `vecty:"prop"`
	Tabs      bool                  `vecty:"prop"`
	Pills     bool                  `vecty:"prop"`
	Fill      bool                  `vecty:"prop"`
	Justified bool                  `vecty:"prop"`
	Markup    vecty.MarkupList      `vecty:"prop"`
	Children  vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Nav) Render() vecty.ComponentOrHTML {
	return elem.Navigation(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"nav":           true,
				"nav-tabs":      c.Tabs,
				"nav-pills":     c.Pills,
				"nav-fill":      c.Fill,
				"nav-justified": c.Justified,
			},
		),
		c.Markup,
		c.Children,
	)
}

// NavLink ...
type NavLink struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Toggle   string                `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *NavLink) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"nav-link": true,
				"active":   c.Active,
				"disabled": c.Disabled,
			},
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.MarkupIf(len(c.Toggle) > 0, vecty.Data("toggle", c.Toggle)),
		),
		c.Markup,
		c.Children,
	)
}
