package main

import "fmt"

func main() {
	door := NewDogDoor()
	remote := NewRemote(door)

	fmt.Println("Fidoが外に出たいと吠える。。。")
	remote.PressButton()

	fmt.Println("\nFidoが外に出る。。。")
	remote.PressButton()

	fmt.Println("\nFidoが用を済ます。。。")
	remote.PressButton()

	fmt.Println("\nFidoが家の中に戻る。。。")
	remote.PressButton()
}
