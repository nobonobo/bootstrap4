package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/prop"
)

// Button ...
type Button struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Type     prop.InputType        `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Toggle   string                `vecty:"prop"` // [button or modal]
	Outline  bool                  `vecty:"prop"`
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Block    bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Checked  bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Button) Render() vecty.ComponentOrHTML {
	return vecty.Tag("button", //elem.Button(
		vecty.Markup(
			vecty.MarkupIf(len(c.Type) == 0, prop.Type(prop.TypeButton)),
			vecty.MarkupIf(len(c.Type) > 0, prop.Type(c.Type)),
			vecty.ClassMap{
				"btn":       true,
				"btn-lg":    c.Large,
				"btn-sm":    c.Small,
				"btn-block": c.Block,
				"active":    c.Active,
			},
			vecty.MarkupIf(!c.Outline, vecty.Class("btn-"+c.Kind.String())),
			vecty.MarkupIf(c.Outline, vecty.Class("btn-outline-"+c.Kind.String())),
			vecty.MarkupIf(c.Disabled, vecty.Attribute("disabled", "true")),
			vecty.MarkupIf(len(c.Toggle) > 0, vecty.Attribute("data-toggle", c.Toggle)),
		),
		c.Markup,
		c.Children,
	)
}

// ButtonLinks ...
type ButtonLinks struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Outline  bool                  `vecty:"prop"`
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Block    bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Checked  bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ButtonLinks) Render() vecty.ComponentOrHTML {
	return vecty.Tag("a", //elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.ClassMap{
				"btn":       true,
				"btn-lg":    c.Large,
				"btn-sm":    c.Small,
				"btn-block": c.Block,
			},
			vecty.MarkupIf(!c.Outline, vecty.Class("btn-"+c.Kind.String())),
			vecty.MarkupIf(c.Outline, vecty.Class("btn-outline-"+c.Kind.String())),
			vecty.MarkupIf(c.Disabled, vecty.Attribute("disabled", "true")),
			vecty.Attribute("role", "button"),
		),
		c.Markup,
		c.Children,
	)
}
