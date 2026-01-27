package main

import (
	"errors"
	"fmt"
)

func main() {
	// ギターの在庫を設定する
	inventory := Inventory{}
	err := initializeInventory(&inventory)

	if len(inventory.guitars) == 0 {
		fmt.Println("Inventory is nil:", err)
		return
	}

	if err != nil {
		fmt.Println("Error initializing inventory:", err)
		// return
	}

	// 検索するギターの仕様を設定する
	whatErinLikes := NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12)
	matchingGuitars := inventory.search(whatErinLikes)

	if len(matchingGuitars) > 0 {
		fmt.Println("Erin, you might like these guitars:")
		for _, guitar := range matchingGuitars {
			spec := guitar.getSpec()
			fmt.Printf("  We have a %v %v %v guitar:\n     %v back and sides,\n     %v top.\n     %v strings.\n  You can have it for only $%.2f!\n  ----\n",
				spec.getBuilder(),
				spec.getModel(),
				spec.getGuitarType(),
				spec.getBackWood(),
				spec.getTopWood(),
				spec.getNumStrings(),
				guitar.getPrice())
		}
	} else {
		fmt.Println("Sorry, Erin, we have nothing for you.")
	}
}

func initializeInventory(inventory *Inventory) error { // 戻り値エラー追加

	guitarsToAdd := []Guitar{
		{serialNumber: "11277", price: 3999.95, spec: nil},
		{serialNumber: "V95693", price: 1499.95, spec: NewGuitarSpec(COLLINGS, "CJ", ACOUSTIC, INDIAN_ROSEWOOD, SITKA, 6)},
		{serialNumber: "V9512", price: 1549.95, spec: NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12)},
		{serialNumber: "V9513", price: 1570.85, spec: NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12)},
		{serialNumber: "122784", price: 5495.95, spec: NewGuitarSpec(MARTIN, "D-18", ACOUSTIC, MAHOGANY, ADIRONDACK, 6)},
		{serialNumber: "76531", price: 6295.95, spec: NewGuitarSpec(MARTIN, "OM-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK, 6)},
		{serialNumber: "70108276", price: 2295.95, spec: NewGuitarSpec(GIBSON, "Les Paul", ELECTRIC, MAHOGANY, MAHOGANY, 6)},
		{serialNumber: "82765501", price: 1890.95, spec: NewGuitarSpec(GIBSON, "SG '61 Reissue", ELECTRIC, MAHOGANY, MAHOGANY, 6)},
		{serialNumber: "77023", price: 6275.95, spec: NewGuitarSpec(MARTIN, "D-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK, 6)},
		{serialNumber: "1092", price: 12995.95, spec: NewGuitarSpec(OLSON, "SJ", ACOUSTIC, INDIAN_ROSEWOOD, CEDAR, 6)},
		{serialNumber: "566-62", price: 8999.95, spec: NewGuitarSpec(RYAN, "Cathedral", ACOUSTIC, COCOBOLO, CEDAR, 6)},
	}

	var inverrs error
	var errCount int

	for _, guitarToAdd := range guitarsToAdd {
		err := inventory.addGuitar(guitarToAdd.serialNumber, guitarToAdd.price, guitarToAdd.spec)
		if err != nil {
			// fmt.Println("Error:", err) // デバッグ用
			inverrs = errors.Join(inverrs, err)
			errCount++
			continue // 次のギター追加へ
		}

	}

	if inverrs != nil {
		return fmt.Errorf("%d件追加できませんでした: %w", errCount, inverrs)
	}
	return nil

	/*
		// 一括でエラーハンドリングできる書き方にしたい
		err := inventory.addGuitar("11277", 3999.95, nil)
		if err != nil {
			return err
		}
		err = inventory.addGuitar("11277", 3999.95, NewGuitarSpec(COLLINGS, "CJ", ACOUSTIC, INDIAN_ROSEWOOD, SITKA, 6))
		if err != nil {
			return err
		}
		// 全てaddGuitarできたらnil返す
		return nil
	*/

	/*
		inventory.addGuitar("11277", 3999.95, NewGuitarSpec(COLLINGS, "CJ", ACOUSTIC, INDIAN_ROSEWOOD, SITKA, 6))
		inventory.addGuitar("V95693", 1499.95, NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12))
		inventory.addGuitar("V9512", 1549.95, NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12))
		inventory.addGuitar("V95693", 1499.95, NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 6))
		inventory.addGuitar("V9512", 1668.95, NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12))
		inventory.addGuitar("122784", 5495.95, NewGuitarSpec(MARTIN, "D-18", ACOUSTIC, MAHOGANY, ADIRONDACK, 6))
		inventory.addGuitar("76531", 6295.95, NewGuitarSpec(MARTIN, "OM-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK, 6))
		inventory.addGuitar("70108276", 2295.95, NewGuitarSpec(GIBSON, "Les Paul", ELECTRIC, MAHOGANY, MAHOGANY, 6))
		inventory.addGuitar("82765501", 1890.95, NewGuitarSpec(GIBSON, "SG '61 Reissue", ELECTRIC, MAHOGANY, MAHOGANY, 6))
		inventory.addGuitar("77023", 6275.95, NewGuitarSpec(MARTIN, "D-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK, 6))
		inventory.addGuitar("1092", 12995.95, NewGuitarSpec(OLSON, "SJ", ACOUSTIC, INDIAN_ROSEWOOD, CEDAR, 6))
		inventory.addGuitar("566-62", 8999.95, NewGuitarSpec(RYAN, "Cathedral", ACOUSTIC, COCOBOLO, CEDAR, 6))
		inventory.addGuitar("6 29584", 2100.95, NewGuitarSpec(PRS, "Dave Navarro Signature", ELECTRIC, MAHOGANY, MAPLE, 6))
	*/
}
