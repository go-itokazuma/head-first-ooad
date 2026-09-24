package main

import (
	"fmt"
	"os"
)

func testStationEqualityAndHashCode() bool {
	station1 := NewStation("Ajax急流")
	station2 := NewStation("ajax急流")

	return station1.Equals(station2) &&
		station1.HashCode() == station2.HashCode()
}

func testNetwork() bool {
	subway := NewSubway()

	subway.AddStation("Ajax急流")
	subway.AddStation("HTML高地")

	subway.AddConnection("Ajax急流", "HTML高地", "テスト線")

	station1 := subway.findStation("Ajax急流")
	station2 := subway.findStation("HTML高地")

	if station1 == nil || station2 == nil {
		return false
	}

	has1to2 := false
	for _, station := range subway.network[station1.HashCode()] {
		if station.Equals(station2) {
			has1to2 = true
			break
		}
	}

	has2to1 := false
	for _, station := range subway.network[station2.HashCode()] {
		if station.Equals(station1) {
			has2to1 = true
			break
		}
	}

	return has1to2 && has2to1
}

func main() {
	loader := NewSubwayLoader()
	objectville, err := loader.LoadFromFile("ObjectvilleSubway.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	subway := NewSubway()

	fmt.Println("\n駅の存在確認のテスト...")
	fmt.Println(subway.HasStation("Ajax急流")) // false
	subway.AddStation("Ajax急流")
	fmt.Println(subway.HasStation("Ajax急流")) // true
	fmt.Println(subway.HasStation("ajax急流")) // true

	fmt.Println("\nStationのEqualsとHashCodeの等値性のテスト...")
	if testStationEqualityAndHashCode() {
		fmt.Println("...等値性のテストに成功")
	} else {
		fmt.Println("...等値性のテストに失敗")
		os.Exit(-1)
	}

	fmt.Println("\n駅のテスト...")
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

	fmt.Println("\nネットワークのテスト...")
	if testNetwork() {
		fmt.Println("...ネットワークのテストに成功")
	} else {
		fmt.Println("...ネットワークのテストに失敗")
		os.Exit(-1)
	}
}
