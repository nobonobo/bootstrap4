package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Pagination ...
type Pagination struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Size     Size                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Pagination) Render() vecty.ComponentOrHTML {
	return elem.UnorderedList(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"pagination":                    true,
				"pagination-" + c.Size.String(): len(c.Size) > 0,
			},
		),
		c.Markup,
		c.Children,
	)
}

// PageItem ...
type PageItem struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Disabled bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *PageItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"page-item": true,
				"disabled":  c.Disabled,
			},
		),
		c.Markup,
		c.Children,
	)
}

// PageLink ...
type PageLink struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Href     string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *PageLink) Render() vecty.ComponentOrHTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"page-link": true,
			},
			vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		),
		c.Markup,
		c.Children,
	)
}
