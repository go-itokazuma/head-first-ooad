package main

import "fmt"

func main() {
	// ギターの在庫を設定する
	inventory := Inventory{}
	initializeInventory(&inventory)

	whatErinLikes := NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER)
	matchingGuitars := inventory.search(whatErinLikes)

	if len(matchingGuitars) > 0 {
		fmt.Println("Erin, you might like these guitars:")
		for _, guitar := range matchingGuitars {
			spec := guitar.getSpec()
			fmt.Printf("  We have a %v %v %v guitar:\n     %v back and sides,\n     %v top.\n  You can have it for only $%.2f!\n  ----\n",
				spec.getBuilder(),
				spec.getModel(),
				spec.getGuitarType(),
				spec.getBackWood(),
				spec.getTopWood(),
				guitar.getPrice())
		}
	} else {
		fmt.Println("Sorry, Erin, we have nothing for you.")
	}
}

func initializeInventory(inventory *Inventory) {
	inventory.addGuitar("11277", 3999.95, COLLINGS, "CJ", ACOUSTIC, INDIAN_ROSEWOOD, SITKA)
	inventory.addGuitar("V95693", 1499.95, FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER)
	inventory.addGuitar("V9512", 1549.95, FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER)
	inventory.addGuitar("122784", 5495.95, MARTIN, "D-18", ACOUSTIC, MAHOGANY, ADIRONDACK)
	inventory.addGuitar("76531", 6295.95, MARTIN, "OM-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK)
	inventory.addGuitar("70108276", 2295.95, GIBSON, "Les Paul", ELECTRIC, MAHOGANY, MAHOGANY)
	inventory.addGuitar("82765501", 1890.95, GIBSON, "SG '61 Reissue", ELECTRIC, MAHOGANY, MAHOGANY)
	inventory.addGuitar("77023", 6275.95, MARTIN, "D-28", ACOUSTIC, BRAZILIAN_ROSEWOOD, ADIRONDACK)
	inventory.addGuitar("1092", 12995.95, OLSON, "SJ", ACOUSTIC, INDIAN_ROSEWOOD, CEDAR)
	inventory.addGuitar("566-62", 8999.95, RYAN, "Cathedral", ACOUSTIC, COCOBOLO, CEDAR)
	inventory.addGuitar("6 29584", 2100.95, PRS, "Dave Navarro Signature", ELECTRIC, MAHOGANY, MAPLE)
}
