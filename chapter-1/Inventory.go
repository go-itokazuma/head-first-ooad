package main

type Inventory struct {
	guitars []*Guitar
}

func (i *Inventory) addGuitar(serialNumber string, price float64, builder string, model string, guitarType string, backWood string, topWood string) {
	newguitar := &Guitar{
		serialNumber: serialNumber,
		price:        price,
		builder:      builder,
		model:        model,
		guitarType:   guitarType,
		backWood:     backWood,
		topWood:      topWood,
	}
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

func (i *Inventory) search(searchGuitar Guitar) *Guitar {
	for _, guitar := range i.guitars {
		//シリアル番号は全て異なるので無視する
		//価格は全て異なるので無視する
		if searchGuitar.getBuilder() != "" && searchGuitar.getBuilder() != guitar.getBuilder() {
			continue
		}
		if searchGuitar.getModel() != "" && searchGuitar.getModel() != guitar.getModel() {
			continue
		}
		if searchGuitar.getGuitarType() != "" && searchGuitar.getGuitarType() != guitar.getGuitarType() {
			continue
		}
		if searchGuitar.getBackWood() != "" && searchGuitar.getBackWood() != guitar.getBackWood() {
			continue
		}
		if searchGuitar.getTopWood() != "" && searchGuitar.getTopWood() != guitar.getTopWood() {
			continue
		}
		return guitar
	}
	return nil
}
