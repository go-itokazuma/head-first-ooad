package main

import "fmt"

type Inventory struct {
	inventory []*Instrument
}

func NewInventory() *Inventory {
	return &Inventory{
		inventory: make([]*Instrument, 0),
	}
}

func (inv *Inventory) addInstrument(serialNumber string, price float64, spec interface{}) error {
	instrument, err := NewInstrument(serialNumber, price, spec)
	if err != nil {
		return fmt.Errorf("failed to add instrument: %w", err)
	}

	switch s := spec.(type) {
	case *GuitarSpec:
		guitar, err := NewGuitar(serialNumber, price, s)
		if err != nil {
			return fmt.Errorf("failed to add guitar: %w", err)
		}
		instrument = guitar.Instrument
	case *MandolinSpec:
		mandolin, err := NewMandolin(serialNumber, price, s)
		if err != nil {
			return fmt.Errorf("failed to add mandolin: %w", err)
		}
		instrument = mandolin.Instrument
	}
	inv.inventory = append(inv.inventory, instrument)
	return nil
}

func (inv *Inventory) getInstrument(serialNumber string) *Instrument {
	for _, instrument := range inv.inventory {
		if instrument.getSerialNumber() == serialNumber {
			return instrument
		}
	}
	return nil
}

func (inv *Inventory) searchGuitar(searchSpec *GuitarSpec) []*Guitar {
	matchingGuitars := make([]*Guitar, 0)
	for _, instrument := range inv.inventory {
		if spec, ok := instrument.getSpec().(*GuitarSpec); ok {
			if spec.MatchesGuitar(searchSpec) {
				matchingGuitars = append(matchingGuitars, &Guitar{Instrument: instrument})
			}
		}
	}
	return matchingGuitars
}

func (inv *Inventory) searchMandolin(searchSpec *MandolinSpec) []*Mandolin {
	matchingMandolins := make([]*Mandolin, 0)
	for _, instrument := range inv.inventory {
		if spec, ok := instrument.getSpec().(*MandolinSpec); ok {
			if spec.MatchesMandolin(searchSpec) {
				matchingMandolins = append(matchingMandolins, &Mandolin{Instrument: instrument})
			}
		}
	}
	return matchingMandolins
}
