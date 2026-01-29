package main

import (
	"fmt"
	"time"
)

type Remote struct {
	door *DogDoor
}

func NewRemote(door *DogDoor) *Remote {
	return &Remote{
		door: door,
	}
}

func (r *Remote) PressButton(done chan bool) {
	fmt.Println("リモコンのボタンが押された。。。")
	if r.door.IsOpen() {
		r.door.Close()
	} else {
		r.door.Open()
		time.AfterFunc(5*time.Second, func() {
			r.door.Close()
			done <- true
		})
	}
}
