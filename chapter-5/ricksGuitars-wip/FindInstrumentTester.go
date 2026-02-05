package main

import (
	"fmt"
)

func main() {
	inventory := NewInventory()
	initializeInventory(inventory)

	properties := make(map[string]interface{})
	properties["builder"] = GIBSON
	properties["backWood"] = MAPLE
	whatLikes := NewInstrumentSpec(properties)
	matchingInstruments := inventory.Search(whatLikes)

	if len(matchingInstruments) > 0 {
		fmt.Println("You might like these instruments:")
		for _, instrument := range matchingInstruments {
			spec := instrument.getSpec()

			fmt.Printf("We have a %v with the following properties:\n", spec.GetProperty("instrumentType"))

			for propertyName, value := range spec.GetProperties() {
				if propertyName == "instrumentType" {
					continue
				}
				fmt.Printf("    %s: %v\n", propertyName, value)
			}

			fmt.Printf("  You can have this %v for $%.2f\n  ----\n", spec.GetProperty("instrumentType"), instrument.getPrice())
		}
	} else {
		fmt.Println("Sorry, we have nothing for you.")
	}
}

func initializeInventory(inventory *Inventory) {
	properties := make(map[string]interface{})
	properties["instrumentType"] = GUITAR
	properties["builder"] = COLLINGS
	properties["model"] = "CJ"
	properties["type"] = ACOUSTIC
	properties["numStrings"] = 6
	properties["topWood"] = INDIAN_ROSEWOOD
	properties["backWood"] = SITKA
	inventory.addInstrument("11277", 3999.95, NewInstrumentSpec(properties))

	properties["builder"] = MARTIN
	properties["model"] = "D-18"
	properties["topWood"] = MAHOGANY
	properties["backWood"] = ADIRONDACK
	inventory.addInstrument("122784", 5495.95, NewInstrumentSpec(properties))

	properties["builder"] = FENDER
	properties["model"] = "Stratocastor"
	properties["type"] = ELECTRIC
	properties["topWood"] = ALDER
	properties["backWood"] = ALDER
	inventory.addInstrument("V95693", 1499.95, NewInstrumentSpec(properties))
	inventory.addInstrument("V9512", 1549.95, NewInstrumentSpec(properties))

	properties["builder"] = GIBSON
	properties["model"] = "Les Paul"
	properties["topWood"] = MAPLE
	properties["backWood"] = MAPLE
	inventory.addInstrument("70108276", 2295.95, NewInstrumentSpec(properties))

	properties["model"] = "SG '61 Reissue"
	properties["topWood"] = MAHOGANY
	properties["backWood"] = MAHOGANY
	inventory.addInstrument("82765501", 1890.95, NewInstrumentSpec(properties))

	properties["instrumentType"] = MANDOLIN
	properties["type"] = ACOUSTIC
	properties["model"] = "F-5G"
	properties["backWood"] = MAPLE
	properties["topWood"] = MAPLE
	delete(properties, "numStrings")
	inventory.addInstrument("9019920", 5495.99, NewInstrumentSpec(properties))

	properties["instrumentType"] = BANJO
	properties["model"] = "RB-3 Wreath"
	delete(properties, "topWood")
	properties["numStrings"] = 5
	inventory.addInstrument("8900231", 2945.95, NewInstrumentSpec(properties))
}
