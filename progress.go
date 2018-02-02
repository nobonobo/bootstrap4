package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Progress ...
type Progress struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Height   string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Progress) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("progress"),
			vecty.MarkupIf(len(c.Height) > 0, vecty.Style("height", c.Height)),
		),
		c.Markup,
		c.Children,
	)
}

// ProgressBar ...
type ProgressBar struct {
	vecty.Core
	ID             string                `vecty:"prop"`
	BackgroundKind Kind                  `vecty:"prop"`
	Width          string                `vecty:"prop"`
	Striped        bool                  `vecty:"prop"`
	Animated       bool                  `vecty:"prop"`
	Markup         vecty.MarkupList      `vecty:"prop"`
	Children       vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ProgressBar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.ClassMap{
				"progress-bar":                    true,
				"progress-bar-striped":            c.Striped,
				"progress-bar-animated":           c.Animated,
				"bg-" + c.BackgroundKind.String(): len(c.BackgroundKind) > 0,
			},
			vecty.Attribute("role", "progress-bar"),
			vecty.MarkupIf(len(c.Width) > 0, vecty.Style("width", c.Width)),
		),
		c.Markup,
		c.Children,
	)
}
