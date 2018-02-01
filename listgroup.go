package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// ListGroup ...
type ListGroup struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ListGroup) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group": true,
			},
		),
		c.Markup,
		c.Children,
	)
}

// ListGroupItem ...
type ListGroupItem struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	ItemKind Kind                  `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ListGroupItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group-item":                        true,
				"list-group-item-" + c.ItemKind.String(): len(c.ItemKind) > 0,
				"active":   c.Active,
				"disabled": c.Disabled,
			},
		),
		c.Markup,
		c.Children,
	)
}

// ListGroupAction ...
type ListGroupAction struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ListGroupAction) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group": true,
			},
		),
		c.Markup,
		c.Children,
	)
}

// ListGroupItemAction ...
type ListGroupItemAction struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	ItemKind Kind                  `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ListGroupItemAction) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group-item":                        true,
				"list-group-item-" + c.ItemKind.String(): len(c.ItemKind) > 0,
				"active":   c.Active,
				"disabled": c.Disabled,
			},
		),
		c.Markup,
		c.Children,
	)
}
