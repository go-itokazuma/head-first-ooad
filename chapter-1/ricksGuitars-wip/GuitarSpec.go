package main

// GuitarSpec構造体
type GuitarSpec struct {
	builder    Builder // 製造者
	model      string  // モデル
	guitarType Type    // 種別(アコースティックまたはエレキ)　※「type」は予約語のため命名変更
	backWood   Wood    // ギターに使用される木材
	topWood    Wood    // ギターに使用される木材
}

func NewGuitarSpec(builder Builder, model string, guitarType Type, backWood Wood, topWood Wood) *GuitarSpec {
	return &GuitarSpec{
		builder:    builder,
		model:      model,
		guitarType: guitarType,
		backWood:   backWood,
		topWood:    topWood,
	}
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
