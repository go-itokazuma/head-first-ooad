package main

type Builder int

const (
	Unspecified Builder = iota
	FENDER
	MARTIN
	GIBSON
	COLLINGS
	OLSON
	RYAN
	PRS
	ANY
)

func (b Builder) String() string {
	switch b {
	case FENDER:
		return "Fender"
	case MARTIN:
		return "Martin"
	case GIBSON:
		return "Gibson"
	case COLLINGS:
		return "Collings"
	case OLSON:
		return "Olson"
	case RYAN:
		return "Ryan"
	case PRS:
		return "PRS"
	default:
		return "Unspecified"
	}
}
