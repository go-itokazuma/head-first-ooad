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

func (br *BarkRecognizer) Recognize(bark *Bark) {
	fmt.Printf("BarkRecognizer: 検知-> %s\n", bark.GetSound())
	allowedBarks := br.door.GetAllowedBarks()
	for _, allowedBark := range allowedBarks {
		if allowedBark.Equals(bark) {
			br.door.Open()
			return
		}
	}
	fmt.Println("この犬は許可されていません。")
}
