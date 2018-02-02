package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Card ...
type Card struct {
	vecty.Core
	ID             string                `vecty:"prop"`
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	Markup         vecty.MarkupList      `vecty:"prop"`
	Children       vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Card) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"card": true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
			},
		),
		c.Markup,
		c.Children,
	)
}

// CardHeader ...
type CardHeader struct {
	vecty.Core
	ID             string                `vecty:"prop"`
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Markup         vecty.MarkupList      `vecty:"prop"`
	Children       vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardHeader) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"card-header":                     true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
		c.Markup,
		c.Children,
	)
}

// CardBody ...
type CardBody struct {
	vecty.Core
	ID             string                `vecty:"prop"`
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Markup         vecty.MarkupList      `vecty:"prop"`
	Children       vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardBody) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"card-body":                       true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
		c.Markup,
		c.Children,
	)
}

// CardFooter ...
type CardFooter struct {
	vecty.Core
	ID             string                `vecty:"prop"`
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Markup         vecty.MarkupList      `vecty:"prop"`
	Children       vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardFooter) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"card-footer":                     true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
		c.Markup,
		c.Children,
	)
}

// CardTitle ...
type CardTitle struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardTitle) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("card-title"),
		),
		c.Markup,
		c.Children,
	)
}

// CardText ...
type CardText struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardText) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("card-text"),
		),
		c.Markup,
		c.Children,
	)
}

// CardImage ...
type CardImage struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardImage) Render() vecty.ComponentOrHTML {
	return elem.Image(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("card-img-top"),
		),
		c.Markup,
		c.Children,
	)
}

// CardLink ...
type CardLink struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardLink) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("card-link"),
			prop.Href(c.Href),
		),
		c.Markup,
		c.Children,
	)
}

// CardGroup ...
type CardGroup struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *CardGroup) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("card-group"),
		),
		c.Markup,
		c.Children,
	)
}
