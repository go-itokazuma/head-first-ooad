package main

import (
	"fmt"
)

type Remote struct {
	door *DogDoor
}

func NewRemote(door *DogDoor) *Remote {
	return &Remote{
		door: door,
	}
}

func (r *Remote) PressButton(done chan<- struct{}) {
	fmt.Println("リモコンのボタンが押された。。。")
	if r.door.IsOpen() {
		r.door.Close(done)
	} else {
		r.door.Open(done)
	}
}
