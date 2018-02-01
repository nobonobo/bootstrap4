package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// ButtonToolbar ...
type ButtonToolbar struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ButtonToolbar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("btn-toolbar"),
			vecty.Attribute("role", "toolbar"),
		),
		c.Markup,
		c.Children,
	)
}

// ButtonGroup ...
type ButtonGroup struct {
	vecty.Core
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Vertical bool                  `vecty:"prop"`
	DropDir  DropDir               `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ButtonGroup) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"btn-group":          !c.Vertical,
				"btn-group-virtical": c.Vertical,
				"btn-group-lg":       c.Large,
				"btn-group-sm":       c.Small,
				c.DropDir.String():   len(c.DropDir) > 0,
			},
			vecty.Attribute("role", "group"),
		),
		c.Markup,
		c.Children,
	)
}

// InputGroup ...
type InputGroup struct {
	vecty.Core
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *InputGroup) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"input-group":    true,
				"input-group-lg": c.Large,
				"input-group-sm": c.Small,
			},
			vecty.Attribute("role", "group"),
		),
		c.Markup,
		c.Children,
	)
}

// InputGroupPrepend ...
type InputGroupPrepend struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *InputGroupPrepend) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"input-group-prepend": true,
			},
		),
		c.Markup,
		c.Children,
	)
}

// InputGroupAppend ...
type InputGroupAppend struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *InputGroupAppend) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"input-group-append": true,
			},
		),
		c.Markup,
		c.Children,
	)
}

// InputGroupText ...
type InputGroupText struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *InputGroupText) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"input-group-text": true,
			},
		),
		c.Markup,
		c.Children,
	)
}
