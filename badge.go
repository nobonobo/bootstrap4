package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Badge ...
type Badge struct {
	vecty.Core
	Kind     Kind                `vecty:"prop"`
	Pill     bool                `vecty:"prop"`
	Markup   vecty.MarkupList    `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *Badge) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
		),
		c.Markup,
		c.Children,
	)
}

// BadgeLinks ...
type BadgeLinks struct {
	vecty.Core
	Kind     Kind                `vecty:"prop"`
	Href     string              `vecty:"prop"`
	Pill     bool                `vecty:"prop"`
	Markup   vecty.MarkupList    `vecty:"prop"`
	Children vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *BadgeLinks) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.Class("badge"),
			vecty.MarkupIf(len(c.Kind) == 0, vecty.Class("badge-"+KindPrimary.String())),
			vecty.MarkupIf(len(c.Kind) > 0, vecty.Class("badge-"+c.Kind.String())),
			vecty.MarkupIf(c.Pill, vecty.Class("badge-pill")),
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
		c.Markup,
		c.Children,
	)
}
