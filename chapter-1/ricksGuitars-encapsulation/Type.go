package main

type Type int

const (
	ACOUSTIC Type = iota
	ELECTRIC
)

func (t Type) String() string {
	switch t {
	case ACOUSTIC:
		return "acoustic"
	case ELECTRIC:
		return "electric"
	default:
		return "unspecified"
	}
}
