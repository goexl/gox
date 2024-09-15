package kernel

const (
	TypeLowercase Type = iota + 1
	TypeCamel
	TypeUnderscore
	TypeUppercase
	TypeStrike

	TypeNonalpha
)

type Type uint8
