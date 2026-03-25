package main

type InstrumentType int

const (
	GUITAR InstrumentType = iota + 1
	BANJO
	DOBRO
	FIDDLE
	BASS
	MANDOLIN
)

func (it InstrumentType) String() string {
	switch it {
	case GUITAR:
		return "Guitar"
	case BANJO:
		return "Banjo"
	case DOBRO:
		return "Dobro"
	case FIDDLE:
		return "Fiddle"
	case BASS:
		return "Bass"
	case MANDOLIN:
		return "Mandolin"
	default:
		return "unspecified"
	}
}
