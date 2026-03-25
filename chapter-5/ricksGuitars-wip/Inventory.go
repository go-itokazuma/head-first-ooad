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

func (inv *Inventory) addInstrument(serialNumber string, price float64, spec *InstrumentSpec) error {
	instrument, err := NewInstrument(serialNumber, price, spec)
	if err != nil {
		return fmt.Errorf("failed to add instrument: %w", err)
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

func (inv *Inventory) Search(searchSpec *InstrumentSpec) []*Instrument {
	matchingInstruments := make([]*Instrument, 0)
	for _, instrument := range inv.inventory {
		if instrument.getSpec().Matches(searchSpec) {
			matchingInstruments = append(matchingInstruments, instrument)
		}
	}
	return matchingInstruments
}
