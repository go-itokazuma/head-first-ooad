package main

import (
	"fmt"
	"time"
)

func main() {
	door := NewDogDoor()
	remote := NewRemote(door)

	fmt.Println("Fidoが外に出たいと吠える。。。")
	remote.PressButton()

	fmt.Println("\nFidoが外に出る。。。")

	fmt.Println("\nFidoが用を済ます。。。")

	fmt.Println("\nFidoが家の中に戻る。。。")
	// ドアが閉まるのを待つ(5秒以上の待機)
	time.Sleep(6 * time.Second)
}
