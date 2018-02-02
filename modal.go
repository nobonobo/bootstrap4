package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Modal ...
type Modal struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *Modal) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Attribute("tabindex", -1),
			vecty.ClassMap{
				"modal": true,
			},
			vecty.Attribute("role", "dialog"),
		),
		c.Markup,
		c.Children,
	)
}

// ModalDialog ...
type ModalDialog struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Centered bool                  `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ModalDialog) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.ClassMap{
				"modal-dialog":          true,
				"modal-dialog-centered": c.Centered,
			},
			vecty.Attribute("role", "document"),
		),
		c.Markup,
		c.Children,
	)
}

// ModalContent ...
type ModalContent struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ModalContent) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("modal-content"),
		),
		c.Markup,
		c.Children,
	)
}

// ModalHeader ...
type ModalHeader struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ModalHeader) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("modal-header"),
		),
		c.Markup,
		c.Children,
	)
}

// ModalBody ...
type ModalBody struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ModalBody) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("modal-body"),
		),
		c.Markup,
		c.Children,
	)
}

// ModalFooter ...
type ModalFooter struct {
	vecty.Core
	ID       string                `vecty:"prop"`
	Markup   vecty.MarkupList      `vecty:"prop"`
	Children vecty.ComponentOrHTML `vecty:"prop"`
}

// Render ...
func (c *ModalFooter) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
			vecty.Class("modal-footer"),
		),
		c.Markup,
		c.Children,
	)
}
