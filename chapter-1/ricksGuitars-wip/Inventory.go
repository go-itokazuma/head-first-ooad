package main

type Inventory struct {
	guitars []*Guitar
}

func (i *Inventory) addGuitar(serialNumber string, price float64, spec *GuitarSpec) {
	newguitar := NewGuitar(serialNumber, price, spec)
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
		if guitar.getSpec().Matches(searchSpec) { // GuitarSpecに比較を委譲
			matchingGuitars = append(matchingGuitars, guitar)
		}
	}
	return matchingGuitars
}
