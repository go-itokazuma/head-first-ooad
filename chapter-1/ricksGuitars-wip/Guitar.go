package main

import (
	"errors"
)

// Guitar構造体
type Guitar struct {
	serialNumber string      // シリアル番号
	price        float64     // 価格
	spec         *GuitarSpec // 製造者
}

func NewGuitar(serialNumber string, price float64, spec *GuitarSpec) (*Guitar, error) {
	// fmt.Println(spec) //デバッグ用(遷移しているか確認)
	if spec == nil {
		return nil, errors.New("spec must not be nil")
	}

	return &Guitar{
		serialNumber: serialNumber,
		price:        price,
		spec:         spec,
	}, nil
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
