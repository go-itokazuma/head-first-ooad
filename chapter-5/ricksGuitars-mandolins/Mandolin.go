package main

import (
	"errors"
)

type Mandolin struct {
	*Instrument
}

func NewMandolin(serialNumber string, price float64, spec *MandolinSpec) (*Mandolin, error) {
	if spec == nil {
		return nil, errors.New("spec must not be nil")
	}

	return &Mandolin{
		Instrument: &Instrument{
			serialNumber: serialNumber,
			price:        price,
			spec:         spec,
		},
	}, nil
}
