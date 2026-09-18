package main

import (
	"fmt"
	"os"
)

func main() {
	loader := NewSubwayLoader()
	objectville, err := loader.LoadFromFile("ObjectvilleSubway.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("駅のテスト...")
	if objectville.HasStation("Ajax急流") &&
		objectville.HasStation("LSPレーン") &&
		objectville.HasStation("Head Firstシアター") {
		fmt.Println("...駅のテストに成功")
	} else {
		fmt.Println("...駅のテストに失敗")
		os.Exit(-1)
	}

	fmt.Println("\n接続のテスト...")
	if objectville.HasConnection("HTML高地", "JavaBeans 大通り", "ブーチ線") &&
		objectville.HasConnection("OOA&Dオーバル", "Head Firstラウンジ", "ガンマ線") &&
		objectville.HasConnection("Java牧場", "JSPジャンクション", "ヤコブソン線") {
		fmt.Println("...接続のテストに成功")
	} else {
		fmt.Println("...接続のテストに失敗")
		os.Exit(-1)
	}
}
