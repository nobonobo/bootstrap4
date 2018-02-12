package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Alert ...
type Alert struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Kind     Kind                  `vecty:"prop"`
	Dismiss  bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Alert) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"alert":                    true,
				"alert-" + c.Kind.String(): true,
			},
			vecty.Attribute("role", "alert"),
		),
		c.Markup,
		c.Children,
		vecty.If(c.Dismiss,
			elem.Button(
				vecty.Markup(
					prop.Type(prop.TypeButton),
					vecty.Class("close"),
					vecty.Attribute("data-dismiss", "alert"),
					vecty.Attribute("aria-label", "Close"),
				),
				elem.Span(
					vecty.Markup(vecty.Attribute("aria-hidden", "true")),
					vecty.Text(`×`),
				),
			),
		),
	)
}

// AlertLink ...
type AlertLink struct {
	vecty.Core
	ID       string              `vecty:"prop"`
	Href     string              `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *AlertLink) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			prop.Href(c.Href),
			vecty.Class("alert-link"),
		),
		c.Children,
	)
}

// AlertHeading ...
type AlertHeading struct {
	vecty.Core
	ID       string              `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *AlertHeading) Render() vecty.ComponentOrHTML {
	return elem.Heading4(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("alert-heading"),
		),
		c.Children,
	)
}
