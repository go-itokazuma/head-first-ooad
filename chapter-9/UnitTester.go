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

func main() {
	tester := &UnitTester{}
	unit := NewUnit(1000)
	tester.TestType(unit, "infantry", "infantry")
	tester.TestUnitSpecificProperty(unit, "hitPoints", 25, 25)
	tester.TestChangeProperty(unit, "hitPoints", 15, 15)
	tester.TestNonExistentProperty(unit, "strength")
}
