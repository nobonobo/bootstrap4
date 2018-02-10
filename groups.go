package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// ButtonToolbar ...
func ButtonToolbar(children ...vecty.MarkupOrChild) vecty.Component {
	return &buttonToolbar{Children: children}
}

type buttonToolbar struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *buttonToolbar) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("btn-toolbar"),
			vecty.Attribute("role", "toolbar"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ButtonGroup ...
func ButtonGroup(size Size, vertical bool, dir DropDir, children ...vecty.MarkupOrChild) vecty.Component {
	return &buttonGroup{
		Size:     size,
		Vertical: vertical,
		DropDir:  dir,
		Children: children,
	}
}

type buttonGroup struct {
	vecty.Core
	Size     Size                  `vecty:"prop"`
	Vertical bool                  `vecty:"prop"`
	DropDir  DropDir               `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *buttonGroup) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"btn-group":                    !c.Vertical,
				"btn-group-virtical":           c.Vertical,
				"btn-group-" + c.Size.String(): len(c.Size) > 0,
				c.DropDir.String():             len(c.DropDir) > 0,
			},
			vecty.Attribute("role", "group"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// InputGroup ...
func InputGroup(size Size, children ...vecty.MarkupOrChild) vecty.Component {
	return &inputGroup{Size: size, Children: children}
}

type inputGroup struct {
	vecty.Core
	Size     Size                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *inputGroup) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"input-group":                    true,
				"input-group-" + c.Size.String(): len(c.Size) > 0,
			},
			vecty.Attribute("role", "group"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// InputGroupPrepend ...
func InputGroupPrepend(children ...vecty.MarkupOrChild) vecty.Component {
	return &inputGroupPrepend{Children: children}
}

type inputGroupPrepend struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *inputGroupPrepend) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("input-group-prepend")),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// InputGroupAppend ...
func InputGroupAppend(children ...vecty.MarkupOrChild) vecty.Component {
	return &inputGroupAppend{Children: children}
}

type inputGroupAppend struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *inputGroupAppend) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("input-group-append")),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// InputGroupText ...
func InputGroupText(children ...vecty.MarkupOrChild) vecty.Component {
	return &inputGroupText{Children: children}
}

type inputGroupText struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *inputGroupText) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("input-group-text")),
	}
	return elem.Div(append(markup, c.Children...)...)
}
