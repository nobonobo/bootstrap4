package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// ButtonProp ...
type ButtonProp struct {
	Kind    Kind
	Size    Size
	Outline bool
	Block   bool
	Toggle  string
}

// Markup ...
func (c ButtonProp) Markup() vecty.MarkupList {
	if len(c.Kind) == 0 {
		c.Kind = KindPrimary
	}
	return vecty.Markup(
		vecty.ClassMap{
			"btn":                            true,
			"btn-block":                      c.Block,
			"btn-" + c.Kind.String():         !c.Outline,
			"btn-outline-" + c.Kind.String(): c.Outline,
			"btn-" + c.Size.String():         len(c.Size) > 0,
			"dropdown-toggle":                len(c.Toggle) > 0,
		},
		vecty.MarkupIf(len(c.Toggle) > 0, vecty.Data("toggle", c.Toggle)),
	)
}

// Button ...
func Button(prop ButtonProp, children ...vecty.MarkupOrChild) vecty.Component {
	return &button{
		Prop:     prop,
		Children: children,
	}
}

type button struct {
	vecty.Core
	Prop     ButtonProp            `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *button) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		c.Prop.Markup(),
	}
	return elem.Button(append(markup, c.Children...)...)
}

// ButtonLink ...
func ButtonLink(prop ButtonProp, href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &buttonLink{
		Prop:     prop,
		Children: children,
	}
}

type buttonLink struct {
	vecty.Core
	Prop     ButtonProp            `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *buttonLink) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		c.Prop.Markup(),
		vecty.Markup(
			prop.Href(c.Href),
			vecty.Attribute("role", "button"),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}
