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

	var testCases = []struct {
		name  string
		sizeX int
		sizeY int
	}{
		{"EmptyMap", 5, 6},
	}

	for _, m := range testCases {
		newMap := NewMap(m.name, m.sizeX, m.sizeY)

		if len(newMap.Spaces) != m.sizeX {
			t.Error("Map does not have expected X side")
		}

		for i := range newMap.Spaces {
			if len(newMap.Spaces[i]) != m.sizeY {
				t.Error("Map does not have expected Y side")
			}

			for ii := range newMap.Spaces[i] {
				if newMap.Spaces[i][ii] != nil {
					t.Error("Map is expected to be empty")
				}
			}
		}
	}

	// err := newMap.MoveUnit(unit1, unit1X, unit1Y, 3, 3)

	// if err != nil {
	// 	t.Error(err)
	// }

	// if newMap.GetUnitInSpace(unit1X, unit1Y) != nil {
	// 	t.Error("Unit should have been moved")
	// }

}

func TestSpawnUnit(t *testing.T) {
	unit1X := 3
	unit1Y := 3
	unit1 := &Unit{Name: "Soldier"}

	unit2X := 3
	unit2Y := 3
	unit2 := &Unit{Name: "Rebel"}

	newMap := NewMap("Map1", 10, 10)

	err := newMap.SpawnUnit(unit1, unit1X, unit1Y)

	if err != nil {
		t.Error("Unit should spawn successfully")
	}
	unitInMap := newMap.Spaces[unit1X][unit1Y]

	if unitInMap == nil {
		t.Error("Map should contain the unit")
	}

	if unitInMap.X != unit1X {
		t.Error("Unit should have right X space")
	}

	if unitInMap.Y != unit1Y {
		t.Error("Unit should have right Y space")
	}

	err = newMap.SpawnUnit(unit2, unit2X, unit2Y)

	if err == nil {
		t.Error("A unit cannot spawn on an occupied space")
	}

}
func TestMoveUnit(t *testing.T) {
	newMap := NewMap("Map1", 10, 10)

	unit1X := 3
	unit1Y := 3
	unit1DestX := 4
	unit1DestY := 3
	unit1 := &Unit{Name: "Soldier"}

	unit2X := 4
	unit2Y := 4
	unit2 := &Unit{Name: "Rebel"}

	newMap.SpawnUnit(unit1, unit1X, unit1Y)
	newMap.SpawnUnit(unit2, unit2X, unit2Y)
	err := newMap.MoveUnit(unit1, unit1DestX, unit1DestY)

	if err != nil {
		t.Error(err)
	}

	if newMap.GetUnitInSpace(unit1X, unit1Y) == unit1 {
		t.Error("Unit should have been removed from original space")
	}

	if newMap.GetUnitInSpace(unit1DestX, unit1DestY) != unit1 {
		t.Error("Unit should have been moved to destination space")
	}

	err = newMap.MoveUnit(unit1, unit2.X, unit2.Y)
	if err == nil {
		t.Error("Error should be returned signaling that the destination space is occupied by another unit")
	}

}
