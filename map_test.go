package main

import (
	"testing"
)

/*
Initialize map
	Should return a Map
Create unit
	Should insert a unit in the map
		Error: unit tries to start outside boundaries
		Error: unit tries to start in occupied space (either by unit or building)
Create building
	Should insert a building in the map
		Error: building tries to start outside boundaries
		Error: building tries to start in occupied space (either by unit or building)
Move unit
	Should move a unit inside the map
		Error: unit tries to move outside boundaries
		Error: unit tries to move to occupied space (either by unit or building)
*/

func TestInitializeMap(t *testing.T) {

	// var testCases = []struct {
	// 	name  string
	// 	sizeX int
	// 	sizeY int
	// }{
	// 	{"Name1", 5, 6},
	// 	{"Name2", 222, 665},
	// 	{"Name3", 10, 2},
	// }

	newMap := NewMap("namae", 5, 5)
	// if newMap.Name != "name" {
	// 	t.Error("Wrong name")
	// }

	unit1X := 1
	unit1Y := 1

	unit1 := &Unit{Name: "Soldier"}

	newMap.SpawnUnit(unit1, unit1X, unit1Y)

	unitInMap := newMap.GetUnitInSpace(unit1X, unit1Y)

	if unitInMap == nil {
		t.Error("Map should contain the unit")
	}

	err := newMap.MoveUnit(unit1, unit1X, unit1Y, 3, 3)

	if err != nil {
		t.Error(err)
	}

	if newMap.GetUnitInSpace(unit1X, unit1Y) != nil {
		t.Error("Unit should have been moved")
	}

}
