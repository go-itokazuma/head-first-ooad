package main

// GuitarSpec構造体
type GuitarSpec struct {
	instrumentSpec InstrumentSpec
	numStrings     int // 弦の本数(p42追加)
}

func NewGuitarSpec(builder Builder, model string, guitarType Type, backWood Wood, topWood Wood, numStrings int) *GuitarSpec {
	return &GuitarSpec{
		instrumentSpec: *NewInstrumentSpec(builder, model, guitarType, backWood, topWood),
		numStrings:     numStrings,
	}
}

func (gs *GuitarSpec) getNumStrings() int {
	return gs.numStrings
}

func (gs *GuitarSpec) MatchesGuitar(otherSpec *GuitarSpec) bool {
	if !gs.instrumentSpec.Matches(&otherSpec.instrumentSpec) {
		return false
	}
	if gs.numStrings != otherSpec.numStrings {
		return false
	}
	return true
}
