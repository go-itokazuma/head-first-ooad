package main

import "fmt"

type Inventory struct {
	guitars []*Guitar
}

func (i *Inventory) addGuitar(serialNumber string, price float64, spec *GuitarSpec) error {
	newguitar, err := NewGuitar(serialNumber, price, spec)
	// fmt.Println(err) //デバッグ用
	if err != nil {
		// fmt.Println(err) //デバッグ用//
		return fmt.Errorf("failed to add guitar: %w", err)
	}
	i.guitars = append(i.guitars, newguitar)
	return nil
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
