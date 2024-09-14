package kernel

const (
	TypeLowercase Type = iota + 1
	TypeCamel
	TypeUnderscore
	TypeUppercase
	TypeStrike
)

type Type uint8
