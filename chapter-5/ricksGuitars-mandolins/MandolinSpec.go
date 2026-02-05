package main

// MandolinSpec構造体
type MandolinSpec struct {
	instrumentSpec InstrumentSpec
	style          Style
}

func NewMandolinSpec(builder Builder, model string, mandolinType Type, backWood Wood, topWood Wood, style Style) *MandolinSpec {
	return &MandolinSpec{
		instrumentSpec: *NewInstrumentSpec(builder, model, mandolinType, backWood, topWood),
		style:          style,
	}
}

func (ms *MandolinSpec) getStyle() Style {
	return ms.style
}

func (ms *MandolinSpec) MatchesMandolin(otherSpec *MandolinSpec) bool {
	if !ms.instrumentSpec.Matches(&otherSpec.instrumentSpec) {
		return false
	}
	if ms.style != otherSpec.style {
		return false
	}
	return true
}
