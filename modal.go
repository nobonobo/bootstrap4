package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Modal ...
func Modal(children ...vecty.MarkupOrChild) vecty.Component {
	return &modal{
		Children: children,
	}
}

type modal struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modal) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Attribute("tabindex", -1),
			vecty.ClassMap{
				"modal": true,
			},
			vecty.Attribute("role", "dialog"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ModalDialog ...
func ModalDialog(children ...vecty.MarkupOrChild) vecty.Component {
	return &modalDialog{
		Children: children,
	}
}

type modalDialog struct {
	vecty.Core
	Centered bool                  `vecty:"prop"`
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modalDialog) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.ClassMap{
				"modal-dialog":          true,
				"modal-dialog-centered": c.Centered,
			},
			vecty.Attribute("role", "document"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ModalContent ...
func ModalContent(children ...vecty.MarkupOrChild) vecty.Component {
	return &modalContent{
		Children: children,
	}
}

type modalContent struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modalContent) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("modal-content"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ModalHeader ...
func ModalHeader(children ...vecty.MarkupOrChild) vecty.Component {
	return &modalHeader{
		Children: children,
	}
}

type modalHeader struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modalHeader) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("modal-header"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ModalBody ...
func ModalBody(children ...vecty.MarkupOrChild) vecty.Component {
	return &modalBody{
		Children: children,
	}
}

type modalBody struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modalBody) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("modal-body"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}

// ModalFooter ...
func ModalFooter(children ...vecty.MarkupOrChild) vecty.Component {
	return &modalFooter{
		Children: children,
	}
}

type modalFooter struct {
	vecty.Core
	Children []vecty.MarkupOrChild `vecty:"prop"`
}

func (c *modalFooter) Render() vecty.ComponentOrHTML {
	markup := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("modal-footer"),
		),
	}
	return elem.Div(append(markup, c.Children...)...)
}
