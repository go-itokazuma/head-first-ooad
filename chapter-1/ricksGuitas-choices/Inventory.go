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

func (i *Inventory) search(searchGuitar *Guitar) []*Guitar {
	matchingGuitars := make([]*Guitar, 0)
	for _, guitar := range i.guitars {
		// シリアル番号は全て異なるので無視する
		// 価格は全て異なるので無視する
		if searchGuitar.getBuilder() != guitar.getBuilder() {
			continue
		}
		model := strings.ToLower(searchGuitar.getModel())
		if model != "" && model != strings.ToLower(guitar.getModel()) {
			continue
		}
		if searchGuitar.getGuitarType() != guitar.getGuitarType() {
			continue
		}
		if searchGuitar.getBackWood() != guitar.getBackWood() {
			continue
		}
		if searchGuitar.getTopWood() != guitar.getTopWood() {
			continue
		}
		matchingGuitars = append(matchingGuitars, guitar)
	}
	return matchingGuitars
}
