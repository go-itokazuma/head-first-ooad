package main

import "strings"

type Inventory struct {
	guitars []*Guitar
}

func (i *Inventory) addGuitar(serialNumber string, price float64, builder Builder, model string, guitarType Type, backWood Wood, topWood Wood) {
	newguitar := NewGuitar(serialNumber, price, builder, model, guitarType, backWood, topWood)
	i.guitars = append(i.guitars, newguitar)
}

func (i *Inventory) getGuitar(serialNumber string) *Guitar {
	for _, guitar := range i.guitars {
		if guitar.getSerialNumber() == serialNumber {
			return guitar
		}
	}
	return nil
}

func (i *Inventory) search(searchSpec *GuitarSpec) []*Guitar {
	matchingGuitars := make([]*Guitar, 0)
	for _, guitar := range i.guitars {
		// シリアル番号は全て異なるので無視する
		// 価格は全て異なるので無視する
		guitarSpec := guitar.getSpec() // GuitarSpecを取得
		if searchSpec.getBuilder() != guitarSpec.getBuilder() {
			continue
		}
		model := strings.ToLower(searchSpec.getModel())
		if model != "" && model != strings.ToLower(guitarSpec.getModel()) {
			continue
		}
		if searchSpec.getGuitarType() != guitarSpec.getGuitarType() {
			continue
		}
		if searchSpec.getBackWood() != guitarSpec.getBackWood() {
			continue
		}
		if searchSpec.getTopWood() != guitarSpec.getTopWood() {
			continue
		}
		matchingGuitars = append(matchingGuitars, guitar)
	}
	return matchingGuitars
}
