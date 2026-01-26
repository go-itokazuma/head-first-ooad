package main

// Guitar構造体
type Guitar struct {
	serialNumber string      // シリアル番号
	price        float64     // 価格
	spec         *GuitarSpec // 製造者
}

func NewGuitar(serialNumber string, price float64, spec *GuitarSpec) *Guitar {
	return &Guitar{
		serialNumber: serialNumber,
		price:        price,
		spec:         spec,
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

func (g *Guitar) getSpec() *GuitarSpec {
	return g.spec
}
