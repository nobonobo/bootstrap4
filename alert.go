package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Alert ...
func Alert(kind Kind, dismiss bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &alert{
		Kind:     kind,
		Dismiss:  dismiss,
		Children: children,
	}
}

type alert struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Dismiss  bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *alert) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"alert":                    true,
				"alert-" + c.Kind.String(): true,
			},
			vecty.Attribute("role", "alert"),
		),
	}
	markup = append(markup, c.Children...)
	markup = append(markup,
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
	return elem.Div(markup...)
}

// AlertLink ...
func AlertLink(href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &alertLink{
		Href:     href,
		Children: children,
	}
}

type alertLink struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *alertLink) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			prop.Href(c.Href),
			vecty.Class("alert-link"),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}

// AlertHeading ...
func AlertHeading(children ...vecty.MarkupOrChild) vecty.Component {
	return &alertHeading{Children: children}
}

type alertHeading struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *alertHeading) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("alert-heading"),
		),
	}
	return elem.Heading4(append(markup, c.Children...)...)
}
