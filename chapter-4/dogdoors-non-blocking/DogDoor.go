package main

import (
	"fmt"
	"time"
)

type DogDoor struct {
	open         bool
	allowedBarks []*Bark
}

func NewDogDoor() *DogDoor {
	return &DogDoor{
		open:         false,
		allowedBarks: make([]*Bark, 0),
	}
}

func (d *DogDoor) AddAllowedBark(bark *Bark) {
	d.allowedBarks = append(d.allowedBarks, bark)
}

func (d *DogDoor) GetAllowedBarks() []*Bark {
	return d.allowedBarks
}

func (d *DogDoor) Open(done chan<- struct{}) {
	fmt.Println("犬用ドアが開く。")
	d.open = true

	time.AfterFunc(5*time.Second, func() {
		d.Close(done)
	})
}

func (d *DogDoor) Close(done chan<- struct{}) {
	fmt.Println("犬用ドアが閉まる。")
	d.open = false
	done <- struct{}{}
}

func (d *DogDoor) IsOpen() bool {
	return d.open
}
