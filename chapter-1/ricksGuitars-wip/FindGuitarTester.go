package main

import (
	"fmt"
)

func main() {
	inventory, err := initializeInventory()
	if err != nil {
		fmt.Println("error initializing inventory:", err)
		return
	}

	whatErinLikes := NewGuitarSpec(FENDER, "Stratocastor", ELECTRIC, ALDER, ALDER, 12)
	matchingGuitars := inventory.search(whatErinLikes)

	if len(matchingGuitars) == 0 {
		fmt.Println("Sorry, Erin, we have nothing for you.")
		return
	}
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
}

func initializeInventory() (*Inventory, error) {
	inventory := &Inventory{}
	guitarsToAdd := []Guitar{
		{serialNumber: "11277", price: 3999.95, spec: nil},
		{serialNumber: "11278", price: 3509.85, spec: nil},
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

	for _, guitarToAdd := range guitarsToAdd {
		if err := inventory.addGuitar(guitarToAdd.serialNumber, guitarToAdd.price, guitarToAdd.spec); err != nil {
			fmt.Println("warning: could not add guitar", guitarToAdd.serialNumber, ":", err)
		}
	}

	if len(inventory.guitars) == 0 {
		return nil, fmt.Errorf("inventory is nil")
	}
	return inventory, nil
}
