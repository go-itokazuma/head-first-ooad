package main

import (
	"fmt"
	"time"
)

func main() {
	door := NewDogDoor()
	remote := NewRemote(door)
	done := make(chan bool)

	fmt.Println("Fidoが外に出たいと吠える。。。")
	remote.PressButton(done)

	fmt.Println("\nFidoが外に出る。。。")
	fmt.Println("\nFidoが用を済ます。。。")

	time.Sleep(10 * time.Second)
	fmt.Println("\n...しかし、まだ外にいる！")
	fmt.Println("\nFidoが吠え始める...")
	fmt.Println("\nそこでGinaがリモコンをつかむ。")
	remote.PressButton(done)
	fmt.Println("\nFidoが家の中に戻る。。。")
	<-done
	<-done
}
