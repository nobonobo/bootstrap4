package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Card ...
func Card(back, border Kind, children ...vecty.MarkupOrChild) vecty.Component {
	return &card{
		BackgroundKind: back,
		BorderKind:     border,
		Children:       children,
	}
}

type card struct {
	vecty.Core
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	Children       []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *card) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card": true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardHeader ...
func CardHeader(back, border, text Kind, children ...vecty.MarkupOrChild) vecty.Component {
	return &cardHeader{
		BackgroundKind: back,
		BorderKind:     border,
		TextKind:       text,
		Children:       children,
	}
}

type cardHeader struct {
	vecty.Core
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Children       []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardHeader) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-header":                     true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardBody ...
func CardBody(back, border, text Kind, children ...vecty.MarkupOrChild) vecty.Component {
	return &cardBody{
		BackgroundKind: back,
		BorderKind:     border,
		TextKind:       text,
		Children:       children,
	}
}

type cardBody struct {
	vecty.Core
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Children       []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardBody) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-body":                       true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardFooter ...
func CardFooter(back, border, text Kind, children ...vecty.MarkupOrChild) vecty.Component {
	return &cardBody{
		BackgroundKind: back,
		BorderKind:     border,
		TextKind:       text,
		Children:       children,
	}
}

type cardFooter struct {
	vecty.Core
	BackgroundKind Kind                  `vecty:"prop"`
	BorderKind     Kind                  `vecty:"prop"`
	TextKind       Kind                  `vecty:"prop"`
	Children       []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *cardFooter) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-footer":                     true,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
				"border-" + c.BorderKind.String(): len(c.BorderKind) > 0,
				"text-" + c.TextKind.String():     len(c.TextKind) > 0,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardTitle ...
func CardTitle(children ...vecty.MarkupOrChild) vecty.Component {
	return &cardTitle{
		Children: children,
	}
}

type cardTitle struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardTitle) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-title": true,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardText ...
func CardText(children ...vecty.MarkupOrChild) vecty.Component {
	return &cardText{
		Children: children,
	}
}

type cardText struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardText) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-text": true,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// CardImage ...
func CardImage(children ...vecty.MarkupOrChild) vecty.Component {
	return &cardImage{
		Children: children,
	}
}

type cardImage struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardImage) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-img-top": true,
			},
		),
	}
	return elem.Image(append(markup, c.Children...)...)
}

// CardLink ...
func CardLink(href string, children ...vecty.MarkupOrChild) vecty.Component {
	return &cardLink{
		Href:     href,
		Children: children,
	}
}

type cardLink struct {
	vecty.Core
	Href     string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *cardLink) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("card-link"),
			prop.Href(c.Href),
		),
	}
	return elem.Anchor(append(markup, c.Children...)...)
}

// CardGroup ...
func CardGroup(children ...vecty.MarkupOrChild) vecty.Component {
	return &cardGroup{
		Children: children,
	}
}

type cardGroup struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

// Render ...
func (c *cardGroup) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"card-group": true,
			},
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}
