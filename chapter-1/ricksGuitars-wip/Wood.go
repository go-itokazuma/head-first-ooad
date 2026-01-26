package main

type Wood int

const (
	INDIAN_ROSEWOOD Wood = iota
	BRAZILIAN_ROSEWOOD
	MAHOGANY
	MAPLE
	COCOBOLO
	CEDAR
	ADIRONDACK
	ALDER
	SITKA
)

func (w Wood) String() string {
	switch w {
	case INDIAN_ROSEWOOD:
		return "Indian Rosewood"
	case BRAZILIAN_ROSEWOOD:
		return "Brazilian Rosewood"
	case MAHOGANY:
		return "Mahogany"
	case MAPLE:
		return "Maple"
	case COCOBOLO:
		return "Cocobolo"
	case CEDAR:
		return "Cedar"
	case ADIRONDACK:
		return "Adirondack"
	case ALDER:
		return "Alder"
	case SITKA:
		return "Sitka"
	default:
		return "unspecified"
	}
}
