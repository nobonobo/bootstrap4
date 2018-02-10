package bootstrap4

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/prop"
)

// Kind ...
type Kind string

const (
	KindPrimary     Kind = "primary"
	KindSecondary   Kind = "secondary"
	KindSuccess     Kind = "success"
	KindDanger      Kind = "danger"
	KindWarning     Kind = "warning"
	KindInfo        Kind = "info"
	KindLight       Kind = "light"
	KindDark        Kind = "dark"
	KindTransparent Kind = "transparent" // background only
)

func (k Kind) String() string { return string(k) }

// DropDir ...
type DropDir string

const (
	DropDown  DropDir = "dropdown"
	DropUp    DropDir = "dropup"
	DropLeft  DropDir = "dropleft"
	DropRight DropDir = "dropright"
)

func (dd DropDir) String() string { return string(dd) }

// Size ...
type Size string

const (
	SizeXLarge Size = "xl"
	SizeLarge  Size = "lg"
	SizeMidium Size = "md"
	SizeSmall  Size = "sm"
)

func (sz Size) String() string { return string(sz) }

// BasicProp ...
type BasicProp struct {
	ID          string
	For         string
	Placeholder string
	Checked     bool
	Autoforcus  bool
	Name        string
	Title       string
	Value       string
	TabIndex    int
	Readonly    bool
	Disabled    bool
}

// Markup ...
func (c BasicProp) Markup() vecty.MarkupList {
	return vecty.Markup(
		vecty.MarkupIf(len(c.ID) > 0, prop.ID(c.ID)),
		vecty.MarkupIf(len(c.For) > 0, prop.For(c.For)),
		vecty.MarkupIf(len(c.Placeholder) > 0, prop.Placeholder(c.Placeholder)),
		vecty.MarkupIf(len(c.Value) > 0, prop.Value(c.Value)),
		prop.Checked(c.Checked),
		prop.Autofocus(c.Autoforcus),
		vecty.MarkupIf(len(c.Name) > 0, vecty.Attribute("name", c.Name)),
		vecty.MarkupIf(len(c.Title) > 0, vecty.Attribute("title", c.Title)),
		vecty.MarkupIf(c.TabIndex != 0, vecty.Attribute("tabindex", c.TabIndex)),
		vecty.MarkupIf(c.Readonly, vecty.Attribute("readonly", "")),
		vecty.MarkupIf(c.Disabled, vecty.Attribute("disabled", "")),
	)
}

// RefProp ...
type RefProp struct {
	Href string
	Src  string
}

// Markup ...
func (c RefProp) Markup() vecty.MarkupList {
	return vecty.Markup(
		vecty.MarkupIf(len(c.Href) > 0, prop.Href(c.Href)),
		vecty.MarkupIf(len(c.Src) > 0, prop.Src(c.Src)),
	)
}
