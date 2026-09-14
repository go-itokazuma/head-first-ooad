package main

type Unit struct {
	unitType   string
	id         int
	name       string
	weapons    []*Weapon
	properties map[string]interface{}
}

func NewUnit(id int) *Unit {
	return &Unit{
		id:         id,
		weapons:    make([]*Weapon, 0),
		properties: make(map[string]interface{}),
	}
}

func (u *Unit) GetId() int {
	return u.id
}

func (u *Unit) SetName(name string) {
	u.name = name
}

func (u *Unit) GetName() string {
	return u.name
}
func (u *Unit) SetType(unitType string) {
	u.unitType = unitType
}

func (u *Unit) GetType() string {
	return u.unitType
}

func (u *Unit) AddWeapon(weapon *Weapon) {
	if weapon != nil {
		u.weapons = append(u.weapons, weapon)
	}
}

func (u *Unit) GetWeapons() []*Weapon {
	return u.weapons
}
func (u *Unit) SetProperty(property string, value interface{}) {
	if u.properties == nil {
		u.properties = make(map[string]interface{})
	}
	u.properties[property] = value
}

func (u *Unit) GetProperty(property string) interface{} {
	if u.properties == nil {
		return nil
	}
	return u.properties[property]
}
