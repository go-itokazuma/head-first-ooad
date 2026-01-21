package main

import "fmt"

//Guitar構造体
type Guitar struct {
	serialNumber string　//シリアル番号
	price        float64　//価格
	builder      string　//製造者
	model        string　//モデル
	guitarType   string //種別(アコースティックまたはエレキ)　※「type」は予約語のため命名変更
	backWood     string //ギターに使用される木材
	topWood      string //ギターに使用される木材
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

func (g *Guitar) getBuilder() string {
	return g.builder
}

func (g *Guitar) getModel() string {
	return g.model
}

func (g *Guitar) getGuitarType() string {
	return g.guitarType
}

func (g *Guitar) getBackWood() string {
	return g.backWood
}

func (g *Guitar) getTopWood() string {
	return g.topWood
}
