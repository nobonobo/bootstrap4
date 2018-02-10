package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Progress ...
func Progress(height string, children ...vecty.MarkupOrChild) vecty.Component {
	return &progress{
		Height:   height,
		Children: children,
	}
}

type progress struct {
	vecty.Core
	Height   string                `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *progress) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("progress"),
			vecty.MarkupIf(len(c.Height) > 0, vecty.Style("height", c.Height)),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ProgressBar ...
func ProgressBar(kind Kind, width string, striped, animated bool, children ...vecty.MarkupOrChild) vecty.Component {
	return &progressBar{
		Kind:     kind,
		Width:    width,
		Striped:  striped,
		Animated: animated,
		Children: children,
	}
}

type progressBar struct {
	vecty.Core
	Kind     Kind                  `vecty:"prop"`
	Width    string                `vecty:"prop"`
	Striped  bool                  `vecty:"prop"`
	Animated bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *progressBar) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"progress-bar":          true,
				"progress-bar-striped":  c.Striped,
				"progress-bar-animated": c.Animated,
				"bg-" + c.Kind.String(): len(c.Kind) > 0,
			},
			vecty.Attribute("role", "progress-bar"),
			vecty.MarkupIf(len(c.Width) > 0, vecty.Style("width", c.Width)),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}
