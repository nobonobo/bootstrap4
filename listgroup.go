package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// ListGroup ...
func ListGroup(role string, children ...vecty.MarkupOrChild) vecty.Component {
	return &listGroup{
		Role:     role,
		Children: children,
	}
}

type listGroup struct {
	vecty.Core
	Role     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *listGroup) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group": true,
			},
		),
	}
	return elem.UnorderedList(append(markup, c.Children...)...)
}

// ListGroupItem ...
func ListGroupItem(kind Kind, role string, children ...vecty.MarkupOrChild) vecty.Component {
	return &listGroupItem{
		Kind:     kind,
		Role:     role,
		Children: children,
	}
}

type listGroupItem struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *listGroupItem) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group-item":                    true,
				"list-group-item-" + c.Kind.String(): len(c.Kind) > 0,
			},
		),
	}
	return elem.ListItem(append(markup, c.Children...)...)
}

// ListGroupAction ...
func ListGroupAction(role string, children ...vecty.MarkupOrChild) vecty.Component {
	return &listGroupAction{
		Role:     role,
		Children: children,
	}
}

type listGroupAction struct {
	vecty.Core
	Role     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *listGroupAction) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group": true,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ListGroupItemAction ...
func ListGroupItemAction(kind Kind, role, href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &listGroupItemAction{
		Kind:     kind,
		Role:     role,
		Href:     href,
		Children: children,
	}
}

type listGroupItemAction struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Role     string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *listGroupItemAction) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.MarkupIf(len(c.Role) > 0, vecty.Attribute("role", c.Role)),
			vecty.ClassMap{
				"list-group-item":                    true,
				"list-group-item-" + c.Kind.String(): len(c.Kind) > 0,
			},
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}
