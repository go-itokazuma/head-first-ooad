package main

type UnitGroup struct {
	units map[int]*Unit
}

func NewUnitGroup(unitList []*Unit) *UnitGroup {
	ug := &UnitGroup{
		units: make(map[int]*Unit),
	}
	for _, unit := range unitList {
		ug.units[unit.GetId()] = unit
	}
	return ug
}

func (ug *UnitGroup) AddUnit(unit *Unit) {
	ug.units[unit.GetId()] = unit
}

func (ug *UnitGroup) RemoveUnitById(id int) {
	delete(ug.units, id)
}

func (ug *UnitGroup) RemoveUnit(unit *Unit) {
	ug.RemoveUnitById(unit.GetId())
}

func (ug *UnitGroup) GetUnit(id int) *Unit {
	return ug.units[id]
}

func (ug *UnitGroup) GetUnits() []*Unit {
	unitList := make([]*Unit, 0, len(ug.units))
	for _, unit := range ug.units {
		unitList = append(unitList, unit)
	}
	return unitList
}
