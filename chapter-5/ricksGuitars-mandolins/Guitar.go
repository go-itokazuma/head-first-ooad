package main

import (
	"errors"
)

type Guitar struct {
	*Instrument
}

func NewGuitar(serialNumber string, price float64, spec *GuitarSpec) (*Guitar, error) {
	if spec == nil {
		return nil, errors.New("spec must not be nil")
	}

	return &Guitar{
		Instrument: &Instrument{
			serialNumber: serialNumber,
			price:        price,
			spec:         spec,
		},
	}, nil
}
