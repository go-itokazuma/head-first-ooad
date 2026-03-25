package main

import (
	"fmt"
	"time"
)

func main() {
	door := NewDogDoor()
	// remote := NewRemote(door)

	recognizer := NewBarkRecognizer(door)
	done := make(chan struct{})

	// ハードウェアによる鳴き声検知をシミュレーションする
	fmt.Println("Fidoが外に出たいと吠える。。。")
	recognizer.Recognize("ウー", done)

	// remote.PressButton(done)

	fmt.Println("\nFidoが外に出る。。。")
	fmt.Println("\nFidoが用を済ます。。。")

	time.Sleep(10 * time.Second)
	fmt.Println("\n...しかし、まだ外にいる！")

	// ハードウェアによる鳴き声検知を再びシミュレーションする
	fmt.Println("\nFidoが吠え始める...")
	// fmt.Println("\nそこでGinaがリモコンをつかむ。")
	recognizer.Recognize("ウー", done)

	// remote.PressButton(done)
	fmt.Println("\nFidoが家の中に戻る。。。")
	<-done
	<-done
}
