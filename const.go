package bootstrap4

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
