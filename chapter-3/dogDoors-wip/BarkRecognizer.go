package main

import "fmt"

type BarkRecognizer struct {
	door *DogDoor
}

func NewBarkRecognizer(door *DogDoor) *BarkRecognizer {
	return &BarkRecognizer{
		door: door,
	}
}

func (br *BarkRecognizer) Recognize(bark string, done chan<- struct{}) {
	fmt.Printf("BarkRecognizer: 検知-> %s\n", bark)
	br.door.Open(done)
}
