package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Button ...
type Button struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Type     prop.InputType        `vecty:"prop"`
	ID       string                `vecty:"prop"`
	Name     string                `vecty:"prop"`
	Value    string                `vecty:"prop"`
	Title    string                `vecty:"prop"`
	TabIndex int                   `vecty:"prop"`
	Active   bool                  `vecty:"prop"`
	Outline  bool                  `vecty:"prop"`
	Large    bool                  `vecty:"prop"`
	Small    bool                  `vecty:"prop"`
	Block    bool                  `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Checked  bool                  `vecty:"prop"`
	Toggle   string                `vecty:"prop"`
	Target   string                `vecty:"prop"`
	Dismiss  string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Button) Render() vecty.ComponentOrHTML {
	if len(c.Kind) == 0 {
		c.Kind = KindPrimary
	}
	return elem.Button(
		vecty.Markup(
			vecty.MarkupIf(len(c.Type) == 0, prop.Type(prop.TypeButton)),
			vecty.MarkupIf(len(c.Type) > 0, prop.Type(c.Type)),
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Name) > 0, vecty.Attribute("name", c.Name)),
			vecty.MarkupIf(len(c.Value) > 0, prop.Value(c.Value)),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"btn":             true,
				"btn-lg":          c.Large,
				"btn-sm":          c.Small,
				"btn-block":       c.Block,
				"active":          c.Active,
				"dropdown-toggle": c.Toggle == "dropdown",
			},
			vecty.MarkupIf(!c.Outline, vecty.Class("btn-"+c.Kind.String())),
			vecty.MarkupIf(c.Outline, vecty.Class("btn-outline-"+c.Kind.String())),
			vecty.MarkupIf(c.Disabled, vecty.Attribute("disabled", nil)),
			vecty.MarkupIf(len(c.Toggle) > 0, vecty.Data("toggle", c.Toggle)),
			vecty.MarkupIf(len(c.Target) > 0, vecty.Data("target", c.Target)),
			vecty.MarkupIf(len(c.Dismiss) > 0, vecty.Data("dismiss", c.Dismiss)),
		),
		c.Markup,
		c.Children,
	)
}

// ButtonLinks ...
type ButtonLinks struct {
	vecty.Core
	Kind         Kind                  `vecty:"prop"`
	Href         string                `vecty:"prop"`
	ID           string                `vecty:"prop"`
	Title        string                `vecty:"prop"`
	TabIndex     int                   `vecty:"prop"`
	Outline      bool                  `vecty:"prop"`
	Large        bool                  `vecty:"prop"`
	Small        bool                  `vecty:"prop"`
	Block        bool                  `vecty:"prop"`
	Disabled     bool                  `vecty:"prop"`
	Checked      bool                  `vecty:"prop"`
	WithDropdown bool                  `vecty:"prop"`
	Toggle       string                `vecty:"prop"`
	Target       string                `vecty:"prop"`
	Dismiss      string                `vecty:"prop"`
	Markup       vecty.MarkupList      `vecty:"prop"`
	Children     vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ButtonLinks) Render() vecty.ComponentOrHTML {
	if len(c.Kind) == 0 {
		c.Kind = KindPrimary
	}
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
			vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
			vecty.ClassMap{
				"btn":             true,
				"btn-lg":          c.Large,
				"btn-sm":          c.Small,
				"btn-block":       c.Block,
				"disabled":        c.Disabled,
				"dropdown-toggle": c.Toggle == "dropdown",
			},
			vecty.MarkupIf(!c.Outline, vecty.Class("btn-"+c.Kind.String())),
			vecty.MarkupIf(c.Outline, vecty.Class("btn-outline-"+c.Kind.String())),
			vecty.Attribute("role", "button"),
			vecty.MarkupIf(len(c.Toggle) > 0, vecty.Data("toggle", c.Toggle)),
			vecty.MarkupIf(len(c.Target) > 0, vecty.Data("target", c.Target)),
			vecty.MarkupIf(len(c.Dismiss) > 0, vecty.Data("dismiss", c.Dismiss)),
		),
		c.Markup,
		c.Children,
	)
}
