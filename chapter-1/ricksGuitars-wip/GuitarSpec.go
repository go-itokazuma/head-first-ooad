package main

import "strings"

// GuitarSpec構造体
type GuitarSpec struct {
	builder    Builder // 製造者
	model      string  // モデル
	guitarType Type    // 種別(アコースティックまたはエレキ)　※「type」は予約語のため命名変更
	backWood   Wood    // ギターに使用される木材
	topWood    Wood    // ギターに使用される木材
	numStrings int     // 弦の本数(p42追加)
}

func NewGuitarSpec(builder Builder, model string, guitarType Type, backWood Wood, topWood Wood, numStrings int) *GuitarSpec {
	return &GuitarSpec{
		builder:    builder,
		model:      model,
		guitarType: guitarType,
		backWood:   backWood,
		topWood:    topWood,
		numStrings: numStrings,
	}
}

func (gs *GuitarSpec) Matches(otherSpec *GuitarSpec) bool {
	if gs.getBuilder() != otherSpec.getBuilder() {
		return false
	}
	model := strings.ToLower(gs.getModel())
	if model != "" && model != strings.ToLower(otherSpec.getModel()) {
		return false
	}
	if gs.getGuitarType() != otherSpec.getGuitarType() {
		return false
	}
	if gs.getBackWood() != otherSpec.getBackWood() {
		return false
	}
	if gs.getTopWood() != otherSpec.getTopWood() {
		return false
	}
	if gs.getTopWood() != otherSpec.getTopWood() {
		return false
	}
	if gs.getNumStrings() != otherSpec.getNumStrings() {
		return false
	}
	return true
}

func (gs *GuitarSpec) getBuilder() Builder {
	return gs.builder
}

func (gs *GuitarSpec) getModel() string {
	return gs.model
}

func (gs *GuitarSpec) getGuitarType() Type {
	return gs.guitarType
}

func (gs *GuitarSpec) getBackWood() Wood {
	return gs.backWood
}

func (gs *GuitarSpec) getTopWood() Wood {
	return gs.topWood
}

func (gs *GuitarSpec) getNumStrings() int {
	return gs.numStrings
}
