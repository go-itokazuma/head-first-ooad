package main

type Style int

const (
	A Style = iota + 1
	F
)

func (s Style) String() string {
	switch s {
	case A:
		return "A style"
	case F:
		return "F style"
	default:
		return "unspecified"
	}
}
