package bootstrap4

// Kind ...
type Kind string

const (
	KindPrimary   Kind = "primary"
	KindSecondary Kind = "secondary"
	KindSuccess   Kind = "success"
	KindDanger    Kind = "danger"
	KindWarning   Kind = "warning"
	KindInfo      Kind = "info"
	KindLight     Kind = "light"
	KindDark      Kind = "dark"
)

func (k Kind) String() string { return string(k) }
