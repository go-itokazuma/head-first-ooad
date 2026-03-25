package main

import (
	"errors"
)

type Instrument struct {
	serialNumber string
	price        float64
	spec         *InstrumentSpec
}

func NewInstrument(serialNumber string, price float64, spec *InstrumentSpec) (*Instrument, error) {
	if spec == nil {
		return nil, errors.New("spec must not be nil")
	}

	return &Instrument{
		serialNumber: serialNumber,
		price:        price,
		spec:         spec,
	}, nil
}

func (i *Instrument) getSerialNumber() string {
	return i.serialNumber
}

func (i *Instrument) getPrice() float64 {
	return i.price
}

func (i *Instrument) setPrice(newPrice float64) {
	i.price = newPrice
}

func (i *Instrument) getSpec() *InstrumentSpec {
	return i.spec
}
