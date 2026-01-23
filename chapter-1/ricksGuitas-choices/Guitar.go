package main

// Guitar構造体
type Guitar struct {
	serialNumber string  // シリアル番号
	price        float64 // 価格
	builder      Builder // 製造者
	model        string  // モデル
	guitarType   Type    // 種別(アコースティックまたはエレキ)　※「type」は予約語のため命名変更
	backWood     Wood    // ギターに使用される木材
	topWood      Wood    // ギターに使用される木材
}

func NewGuitar(serialNumber string, price float64, builder Builder, model string, guitarType Type, backWood Wood, topWood Wood) *Guitar {
	return &Guitar{
		serialNumber: serialNumber,
		price:        price,
		builder:      builder,
		model:        model,
		guitarType:   guitarType,
		backWood:     backWood,
		topWood:      topWood,
	}
}

func (g *Guitar) getSerialNumber() string {
	return g.serialNumber
}

func (g *Guitar) getPrice() float64 {
	return g.price
}

func (g *Guitar) setPrice(newPrice float64) {
	g.price = newPrice
}

func (g *Guitar) getBuilder() Builder {
	return g.builder
}

func (g *Guitar) getModel() string {
	return g.model
}

func (g *Guitar) getGuitarType() Type {
	return g.guitarType
}

func (g *Guitar) getBackWood() Wood {
	return g.backWood
}

func (g *Guitar) getTopWood() Wood {
	return g.topWood
}
