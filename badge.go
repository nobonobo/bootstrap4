package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Badge ...
func Badge(kind Kind, pill bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &badge{
		Kind:     kind,
		Pill:     pill,
		Children: children,
	}
}

type badge struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Pill     bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *badge) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
		),
	}
	return elem.Span(append(markup, c.Children...)...)
}

// BadgeLinks ...
func BadgeLinks(kind Kind, pill bool, href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &badgeLinks{
		Kind:     kind,
		Pill:     pill,
		Href:     href,
		Children: children,
	}
}

type badgeLinks struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Pill     bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *badgeLinks) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}
