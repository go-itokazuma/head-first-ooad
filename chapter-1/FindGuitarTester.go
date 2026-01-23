package main

import "fmt"

func main() {
	//ギターの在庫を設定する
	inventory := Inventory{}
	initializeInventory(&inventory)

	whatErinLikes := Guitar{
		serialNumber: "",
		price:        0,
		builder:      "fender",
		model:        "stratocastor",
		guitarType:   "electric",
		backWood:     "Alder",
		topWood:      "Alder",
	}
	guitar := inventory.search(whatErinLikes)

	if guitar != nil {
		fmt.Printf("Erin, you might like this %v %v %v guitar:\n%v back and sides,\n%v top.\nYou can have it for only $%.2f!\n",
			guitar.getBuilder(),
			guitar.getModel(),
			guitar.getGuitarType(),
			guitar.getBackWood(),
			guitar.getTopWood(),
			guitar.getPrice())
	} else {
		fmt.Println("Sorry, Erin, we have nothing for you.")
	}

}

func initializeInventory(inventory *Inventory) {
	inventory.addGuitar("V95693", 1499.95, "Fender", "Stratocastor", "electric", "Alder", "Alder")
	//inventory.addGuitar("V9512", 1549.95, "fender", "stratocastor", "electric", "Alder", "Alder") //検索通るか確認用
}
