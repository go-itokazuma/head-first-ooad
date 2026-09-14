package main

import "fmt"

type UnitTester struct {
}

func (ut *UnitTester) TestType(unit *Unit, unitType string, expectedOutputType string) {
	fmt.Println("\nTesting setting/getting the type property.")
	unit.SetType(unitType)
	outputType := unit.GetType()
	if expectedOutputType == outputType {
		fmt.Println("Test passed")
	} else {
		fmt.Printf("Test failed: %s didn't match %s\n", outputType, expectedOutputType)
	}
}

func (ut *UnitTester) TestUnitSpecificProperty(unit *Unit, propertyName string, inputValue interface{}, expectedOutputType interface{}) {
	fmt.Println("\nTesting setting/getting a unit-specific property.")
	unit.SetProperty(propertyName, inputValue)
	outputValue := unit.GetProperty(propertyName)
	if expectedOutputType == outputValue {
		fmt.Println("Test passed")
	} else {
		fmt.Printf("Test failed: %v didn't match %v\n", outputValue, expectedOutputType)
	}
}

func (ut *UnitTester) TestChangeProperty(unit *Unit, propertyName string, inputValue interface{}, expectedOutputType interface{}) {
	fmt.Println("\nTesting changing a unit-specific property.")
	unit.SetProperty(propertyName, inputValue)
	outputValue := unit.GetProperty(propertyName)
	if inputValue == outputValue {
		fmt.Println("Test passed")
	} else {
		fmt.Printf("Test failed: %v didn't match %v\n", outputValue, inputValue)
	}
}

func (ut *UnitTester) TestNonExistentProperty(unit *Unit, propertyName string) {
	fmt.Println("\nTesting getting a non-existent property's value.")
	outputValue := unit.GetProperty(propertyName)
	if outputValue == nil {
		fmt.Println("Test passed")
	} else {
		fmt.Printf("Test failed with value of %v\n", outputValue)
	}
}

// ID10
func (ut *UnitTester) TestCreateUnitGroup(unitList []*Unit) {
	fmt.Println("\nTesting creation of a UnitGroup.")

	group := NewUnitGroup(unitList)
	allMatched := true

	for _, unit := range unitList {
		outputUnit := group.GetUnit(unit.GetId())
		if outputUnit != unit {
			allMatched = false
			break
		}
	}

	if allMatched {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

// ID11
func (ut *UnitTester) TestAddUnitToUnitGroup(unit *Unit) {
	fmt.Println("\nTesting adding a unit to a UnitGroup.")
	group := NewUnitGroup([]*Unit{})
	group.AddUnit(unit)
	outputUnit := group.GetUnit(unit.GetId())
	if outputUnit == unit {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

// ID12
func (ut *UnitTester) TestGetUnitById(unitId int) {
	fmt.Println("\nTesting getting a unit by ID.")
	unit := NewUnit(unitId)
	group := NewUnitGroup([]*Unit{unit})
	outputUnit := group.GetUnit(unit.GetId())
	if outputUnit == unit {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func (ut *UnitTester) TestGetUnits(unitList []*Unit) {
	fmt.Println("\nTesting getting all units from a UnitGroup.")

	group := NewUnitGroup(unitList)
	outputUnits := group.GetUnits()

	outputByID := make(map[int]*Unit)

	allMatched := true
	for _, unit := range outputUnits {
		outputByID[unit.GetId()] = unit
	}

	if len(outputUnits) != len(unitList) {
		allMatched = false
	} else {
		for _, unit := range unitList {
			if outputByID[unit.GetId()] != unit {
				allMatched = false
				break
			}
		}
	}

	if allMatched {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func (ut *UnitTester) TestRemoveUnitByIdFromUnitGroup(unitId int) {
	fmt.Println("\nTesting removing a unit by ID from a UnitGroup.")
	unit := NewUnit(unitId)
	group := NewUnitGroup([]*Unit{unit})
	group.RemoveUnitById(unitId)
	outputUnit := group.GetUnit(unitId)
	if outputUnit == nil {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func (ut *UnitTester) TestRemoveUnitFromUnitGroup(unit *Unit) {
	fmt.Println("\nTesting removing a unit from a UnitGroup.")
	group := NewUnitGroup([]*Unit{unit})
	group.RemoveUnit(unit)
	outputUnit := group.GetUnit(unit.GetId())
	if outputUnit == nil {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func main() {
	tester := &UnitTester{}
	unit := NewUnit(1000)
	tester.TestType(unit, "infantry", "infantry")
	tester.TestUnitSpecificProperty(unit, "hitPoints", 25, 25)
	tester.TestChangeProperty(unit, "hitPoints", 15, 15)
	tester.TestNonExistentProperty(unit, "strength")

	// Testing UnitGroup
	// ID10
	unitA := NewUnit(100)
	unitB := NewUnit(200)
	unitList := []*Unit{
		unitA,
		unitB,
	}
	tester.TestCreateUnitGroup(unitList)

	//ID11
	unitC := NewUnit(100)
	tester.TestAddUnitToUnitGroup(unitC)

	// ID12
	tester.TestGetUnitById(100)

	//ID13
	unitD := NewUnit(300)
	unitE := NewUnit(400)

	unitList = []*Unit{
		unitD,
		unitE,
	}
	tester.TestGetUnits(unitList)

	//ID14
	tester.TestRemoveUnitByIdFromUnitGroup(100)

	//ID15
	unitF := NewUnit(100)
	tester.TestRemoveUnitFromUnitGroup(unitF)
}
