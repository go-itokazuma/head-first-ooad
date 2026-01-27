package main

import "fmt"

type DogDoor struct {
	open bool
}

func NewDogDoor() *DogDoor {
	return &DogDoor{
		open: false,
	}
}

func (d *DogDoor) Open() {
	fmt.Println("犬用ドアが開く。")
	d.open = true
}

func (d *DogDoor) Close() {
	fmt.Println("犬用ドアが閉まる。")
	d.open = false
}

func (d *DogDoor) IsOpen() bool {
	return d.open
}
