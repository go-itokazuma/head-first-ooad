package main

import (
	"fmt"
)

func main() {
	door := NewDogDoor()
	remote := NewRemote(door)
	done := make(chan bool)

	fmt.Println("Fidoが外に出たいと吠える。。。")

	remote.PressButton(done)
	fmt.Println("\nFidoが外に出る。。。")

	fmt.Println("\nFidoが用を済ます。。。")

	fmt.Println("\nFidoが家の中に戻る。。。")
	<-done
}
