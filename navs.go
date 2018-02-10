package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// NavList ...
func NavList(role string, tabs, pills, fill, justified bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &navList{
		Role:      role,
		Tabs:      tabs,
		Pills:     pills,
		Fill:      fill,
		Justified: justified,
		Children:  children,
	}
}

type navList struct {
	vecty.Core
	Role      string                `vecty:"prop"`
	Tabs      bool                  `vecty:"prop"`
	Pills     bool                  `vecty:"prop"`
	Fill      bool                  `vecty:"prop"`
	Justified bool                  `vecty:"prop"`
	Children  []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navList) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"nav":           true,
				"nav-tabs":      c.Tabs,
				"nav-pills":     c.Pills,
				"nav-fill":      c.Fill,
				"nav-justified": c.Justified,
			},
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
		),
	}
	return elem.UnorderedList(append(markup, c.Children...)...)
}

// NavItem ...
func NavItem(active, dropdown bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &navItem{
		Active:   active,
		Dropdown: dropdown,
		Children: children,
	}
}

type navItem struct {
	vecty.Core
	Active   bool                  `vecty:"prop"`
	Dropdown bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navItem) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"nav-item": true,
				"active":   c.Active,
				"dropdown": c.Dropdown,
			},
		),
	}
	return elem.ListItem(append(markup, c.Children...)...)
}

// Nav ...
func Nav(tabs, pills, fill, justified bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &nav{
		Tabs:      tabs,
		Pills:     pills,
		Fill:      fill,
		Justified: justified,
		Children:  children,
	}
}

type nav struct {
	vecty.Core
	Tabs      bool                  `vecty:"prop"`
	Pills     bool                  `vecty:"prop"`
	Fill      bool                  `vecty:"prop"`
	Justified bool                  `vecty:"prop"`
	Children  []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *nav) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"nav":           true,
				"nav-tabs":      c.Tabs,
				"nav-pills":     c.Pills,
				"nav-fill":      c.Fill,
				"nav-justified": c.Justified,
			},
		),
	}
	return elem.Navigation(append(markup, c.Children...)...)
}

// NavLink ...
func NavLink(href, toggle, role string, active, disabled bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &navLink{
		Href:     href,
		Toggle:   toggle,
		Role:     role,
		Active:   active,
		Disabled: disabled,
		Children: children,
	}
}

type navLink struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Toggle   string                `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *navLink) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"nav-link": true,
				"active":   c.Active,
				"disabled": c.Disabled,
			},
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.MarkupIf(len(c.Toggle) > 0, vecty.Data("toggle", c.Toggle)),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}
