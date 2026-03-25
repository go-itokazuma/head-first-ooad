package main

import "strings"

type InstrumentSpec struct {
	builder        Builder
	model          string
	instrumentType Type
	backWood       Wood
	topWood        Wood
}

func NewInstrumentSpec(builder Builder, model string, instrumentType Type, backWood Wood, topWood Wood) *InstrumentSpec {
	return &InstrumentSpec{
		builder:        builder,
		model:          model,
		instrumentType: instrumentType,
		backWood:       backWood,
		topWood:        topWood,
	}
}

func (is *InstrumentSpec) Matches(otherSpec *InstrumentSpec) bool {
	if is.getBuilder() != otherSpec.getBuilder() {
		return false
	}
	model := strings.ToLower(is.getModel())
	if model != "" && model != strings.ToLower(otherSpec.getModel()) {
		return false
	}
	if is.getInstrumentType() != otherSpec.getInstrumentType() {
		return false
	}
	if is.getBackWood() != otherSpec.getBackWood() {
		return false
	}
	if is.getTopWood() != otherSpec.getTopWood() {
		return false
	}
	return true
}

func (is *InstrumentSpec) getBuilder() Builder {
	return is.builder
}

func (is *InstrumentSpec) getModel() string {
	return is.model
}

func (is *InstrumentSpec) getInstrumentType() Type {
	return is.instrumentType
}

func (is *InstrumentSpec) getBackWood() Wood {
	return is.backWood
}

func (is *InstrumentSpec) getTopWood() Wood {
	return is.topWood
}
