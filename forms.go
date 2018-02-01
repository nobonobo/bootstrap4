package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FormGroup ...
type FormGroup struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *FormGroup) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("form-group")),
		c.Markup,
		c.Children,
	)
}

// FormCheck ...
type FormCheck struct {
	vecty.Core
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *FormCheck) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("form-check")),
		c.Markup,
		c.Children,
	)
}

// Label ...
type Label struct {
	vecty.Core
	For      string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Label) Render() vecty.ComponentOrHTML {
	return elem.Label(
		vecty.Markup(
			vecty.MarkupIf(len(c.For) > 0, vecty.Attribute("for", c.For)),
		),
		c.Markup,
		c.Children,
	)
}

// Input ...
type Input struct {
	vecty.Core
	Type        prop.InputType        `vecty:"prop"`
	ID          string                `vecty:"prop"`
	Name        string                `vecty:"prop"`
	Value       string                `vecty:"prop"`
	PlaceHolder string                `vecty:"prop"`
	Title       string                `vecty:"prop"`
	TabIndex    int                   `vecty:"prop"`
	Large       bool                  `vecty:"prop"`
	Small       bool                  `vecty:"prop"`
	Readonly    bool                  `vecty:"prop"`
	PlainText   bool                  `vecty:"prop"`
	Markup      vecty.MarkupList      `vecty:"prop"`
	Children    vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Input) Render() vecty.ComponentOrHTML {
	return elem.Input(
		vecty.Markup(
			vecty.MarkupIf(len(c.Type) > 0, prop.Type(c.Type)),
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Name) > 0, vecty.Attribute("name", c.Name)),
			vecty.MarkupIf(len(c.Value) > 0, prop.Value(c.Value)),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"form-control":           c.Type != prop.TypeFile,
				"form-control-file":      c.Type == prop.TypeFile,
				"form-control-plaintext": c.PlainText,
				"form-control-lg":        c.Large,
				"form-control-sm":        c.Small,
			},
			vecty.MarkupIf(len(c.PlaceHolder) > 0,
				vecty.Attribute("placeholder", c.PlaceHolder),
			),
			vecty.MarkupIf(c.Readonly || c.PlainText,
				vecty.Attribute("readonly", ""),
			),
			vecty.MarkupIf(len(c.Value) > 0,
				vecty.Attribute("value", c.Value),
			),
		),
		c.Markup,
		c.Children,
	)
}

// CheckInput ...
type CheckInput struct {
	vecty.Core
	Type     prop.InputType        `vecty:"prop"`
	ID       string                `vecty:"prop"`
	Name     string                `vecty:"prop"`
	Checked  bool                  `vecty:"prop"`
	Title    string                `vecty:"prop"`
	TabIndex int                   `vecty:"prop"`
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	NoLabel  bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CheckInput) Render() vecty.ComponentOrHTML {
	return elem.Input(
		vecty.Markup(
			vecty.MarkupIf(len(c.Type) > 0,
				prop.Type(c.Type),
			),
			vecty.MarkupIf(len(c.ID) > 0,
				vecty.Attribute("id", c.ID),
			),
			vecty.MarkupIf(len(c.Name) > 0,
				vecty.Attribute("name", c.Name),
			),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"form-check-input": true,
				"position-static":  c.NoLabel,
				"form-control-lg":  c.Large,
				"form-control-sm":  c.Small,
			},
			vecty.MarkupIf(c.Disabled,
				vecty.Attribute("disabled", ""),
			),
			vecty.MarkupIf(c.Checked,
				vecty.Attribute("checked", ""),
			),
		),
		c.Markup,
		c.Children,
	)
}

// TextArea ...
type TextArea struct {
	vecty.Core
	ID          string                `vecty:"prop"`
	Name        string                `vecty:"prop"`
	PlaceHolder string                `vecty:"prop"`
	Title       string                `vecty:"prop"`
	TabIndex    int                   `vecty:"prop"`
	Large       bool                  `vecty:"prop"`
	Small       bool                  `vecty:"prop"`
	Readonly    bool                  `vecty:"prop"`
	Rows        int                   `vecty:"prop"`
	Markup      vecty.MarkupList      `vecty:"prop"`
	Children    vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *TextArea) Render() vecty.ComponentOrHTML {
	return elem.TextArea(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0,
				vecty.Attribute("id", c.ID),
			),
			vecty.MarkupIf(len(c.Name) > 0,
				vecty.Attribute("name", c.Name),
			),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"form-control":    true,
				"form-control-lg": c.Large,
				"form-control-sm": c.Small,
			},
			vecty.MarkupIf(len(c.PlaceHolder) > 0,
				vecty.Attribute("placeholder", c.PlaceHolder),
			),
			vecty.MarkupIf(c.Readonly,
				vecty.Attribute("readonly", ""),
			),
			vecty.MarkupIf(c.Rows > 0,
				vecty.Attribute("rows", c.Rows),
			),
		),
		c.Markup,
		c.Children,
	)
}

// Select ...
type Select struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Name     string                `vecty:"prop"`
	Title    string                `vecty:"prop"`
	TabIndex int                   `vecty:"prop"`
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Readonly bool                  `vecty:"prop"`
	Multiple bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Select) Render() vecty.ComponentOrHTML {
	return elem.Select(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0,
				vecty.Attribute("id", c.ID),
			),
			vecty.MarkupIf(len(c.Name) > 0,
				vecty.Attribute("name", c.Name),
			),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"form-control":    true,
				"form-control-lg": c.Large,
				"form-control-sm": c.Small,
			},
			vecty.MarkupIf(c.Multiple,
				vecty.Attribute("multiple", ""),
			),
			vecty.MarkupIf(c.Readonly,
				vecty.Attribute("readonly", ""),
			),
		),
		c.Markup,
		c.Children,
	)
}

// Custom forms not supported.
