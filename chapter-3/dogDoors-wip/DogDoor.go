package main

import (
	"fmt"
	"time"
)

type DogDoor struct {
	open bool
}

func NewDogDoor() *DogDoor {
	return &DogDoor{
		open: false,
	}
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
