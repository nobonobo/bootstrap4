package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FormProp ...
type FormProp struct {
	Type      prop.InputType
	Size      Size
	Plaintext bool
	NoLabel   bool
	Multiple  bool
	Rows      int
}

// Markup ...
func (c FormProp) Markup() vecty.MarkupList {
	return vecty.Markup(
		vecty.MarkupIf(len(c.Type) > 0, prop.Type(c.Type)),
		vecty.ClassMap{
			"form-control":                    c.Type != prop.TypeFile,
			"form-control-file":               c.Type == prop.TypeFile,
			"form-control-plaintext":          c.Plaintext,
			"form-control-" + c.Size.String(): len(c.Size) > 0,
			"position-static":                 c.NoLabel,
		},
		vecty.MarkupIf(c.Multiple, vecty.Attribute("multiple", "")),
		vecty.MarkupIf(c.Rows > 0, vecty.Attribute("rows", c.Rows)),
	)
}

// FormGroup ...
func FormGroup(children ...vecty.MarkupOrChild) vecty.Component {
	return &formGroup{
		Children: children,
	}
}

type formGroup struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *formGroup) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("form-group")),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// FormCheck ...
func FormCheck(children ...vecty.MarkupOrChild) vecty.Component {
	return &formCheck{Children: children}
}

type formCheck struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *formCheck) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("form-check")),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// TextInput ...
func TextInput(prop FormProp, children ...vecty.MarkupOrChild) vecty.Component {
	return &textInput{Prop: prop, Children: children}
}

type textInput struct {
	vecty.Core
	Prop     FormProp              `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *textInput) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{c.Prop.Markup()}
	return elem.Input(append(markup, c.Children...)...)
}

// CheckInput ...
func CheckInput(prop FormProp, children ...vecty.MarkupOrChild) vecty.Component {
	return &checkInput{Prop: prop, Children: children}
}

type checkInput struct {
	vecty.Core
	Prop     FormProp
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *checkInput) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		c.Prop.Markup(),
		vecty.Markup(vecty.Class("form-check-input")),
	}
	return elem.Input(append(markup, c.Children...)...)
}

// TextArea ...
func TextArea(prop FormProp, children ...vecty.MarkupOrChild) vecty.Component {
	return &textArea{Prop: prop, Children: children}
}

type textArea struct {
	vecty.Core
	Prop     FormProp              `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *textArea) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		c.Prop.Markup(),
	}
	return elem.TextArea(append(markup, c.Children...)...)
}

// SelectInput ...
func SelectInput(prop FormProp, children ...vecty.MarkupOrChild) vecty.Component {
	return &selectInput{Prop: prop, Children: children}
}

type selectInput struct {
	vecty.Core
	Prop     FormProp              `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *selectInput) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		c.Prop.Markup(),
	}
	return elem.Select(append(markup, c.Children...)...)
}

// Custom forms not supported.
