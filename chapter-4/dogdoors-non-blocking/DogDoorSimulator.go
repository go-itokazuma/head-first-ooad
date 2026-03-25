package main

import (
	"fmt"
	"time"
)

func main() {
	door := NewDogDoor()
	door.AddAllowedBark(NewBark("ロー"))
	door.AddAllowedBark(NewBark("ローーー"))
	door.AddAllowedBark(NewBark("ラル"))
	door.AddAllowedBark(NewBark("ウー"))
	recognizer := NewBarkRecognizer(door)
	remote := NewRemote(door)
	_ = remote
	done := make(chan struct{})

	// ハードウェアによる鳴き声検知をシミュレーションする
	fmt.Println("Bruceが吠え出す")
	recognizer.Recognize(NewBark("ロー"), done)

	// remote.PressButton(done)

	fmt.Println("\nBruceが外に出る。。。")
	time.Sleep(10 * time.Second)
	fmt.Println("\nBruceが用を済ます。。。")

	fmt.Println("\n...しかし外に閉め出される！")

	// ハードウェアによる鳴き声検知を再びシミュレーションする(Bruce以外の犬)
	smallDogBark := NewBark("キャン")
	fmt.Println("\n小さな犬が吠え出す")
	recognizer.Recognize(smallDogBark, done)

	time.Sleep(5 * time.Second)
	fmt.Println("\nBruceが吠え出す")
	recognizer.Recognize(NewBark("ローーー"), done)
	fmt.Println("\nBruceが家の中に戻る。。。")
	<-done
	<-done
}
