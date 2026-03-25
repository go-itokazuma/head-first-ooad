package main

import (
	"fmt"
	"sync"
	"time"
)

type DogDoor struct {
	open         bool
	allowedBarks []*Bark
	wg           sync.WaitGroup
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

func (d *DogDoor) Open() {
	if d.open {
		return
	}
	fmt.Println("犬用ドアが開く。")
	d.open = true

	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		time.Sleep(5 * time.Second)
		if d.IsOpen() {
			d.Close()
		}
	}()
}

func (d *DogDoor) Close() {
	if !d.open {
		return
	}
	fmt.Println("犬用ドアが閉まる。")
	d.open = false
}

func (d *DogDoor) IsOpen() bool {
	return d.open
}

func (d *DogDoor) WaitUntilClosed() {
	d.wg.Wait()
}
