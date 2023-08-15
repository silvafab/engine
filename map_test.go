package main

import (
	"testing"
)

func TestInitializeMap(t *testing.T) {

	var testCases = []struct {
		name  string
		sizeX int
		sizeY int
	}{
		{"EmptyMap", 5, 6},
	}

	p1 := Player{Name: "Player 1"}
	players := []Player{p1}

	for _, m := range testCases {
		newMap := NewMap(m.name, m.sizeX, m.sizeY, players)

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
}

func TestSpawnUnit(t *testing.T) {

	p1 := Player{Name: "Player 1"}
	players := []Player{p1}

	unit1X := 3
	unit1Y := 3
	unit1 := &Unit{Name: "Soldier", Player: p1}

	unit2X := 3
	unit2Y := 3
	unit2 := &Unit{Name: "Rebel", Player: p1}
	newMap := NewMap("Map1", 10, 10, players)

	err := newMap.SpawnUnit(unit1, unit1X, unit1Y)

	if err != nil {
		t.Error("Unit should spawn successfully")
	}
	unitInMap := newMap.Spaces[unit1X][unit1Y]

	if unitInMap == nil {
		t.Error("Map should contain the unit")
	}

	unitInMapX, unitInMapY := unitInMap.GetLocation()
	if unitInMapX != unit1X {
		t.Error("Unit should have right X space")
	}

	if unitInMapY != unit1Y {
		t.Error("Unit should have right Y space")
	}

	err = newMap.SpawnUnit(unit2, unit2X, unit2Y)

	if err == nil {
		t.Error("A unit cannot spawn on an occupied space")
	}

}

func TestMoveSpawnable(t *testing.T) {
	p1 := Player{Name: "Player 1"}
	players := []Player{p1}
	newMap := NewMap("Map1", 10, 10, players)

	unit1X := 3
	unit1Y := 3
	unit1DestX := 4
	unit1DestY := 3
	unit1 := &Unit{Name: "Soldier"}

	unit2X := 4
	unit2Y := 4
	unit2 := &Unit{Name: "Rebel"}

	newMap.SpawnUnit(unit2, unit2X, unit2Y)

	buildingX := 2
	buildingY := 2
	buildingDestX := 8
	buildingDestY := 8
	building := &Building{Name: "Barracks"}
	newMap.SpawnUnit(building, buildingX, buildingY)

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

	err = newMap.MoveUnit(unit1, building.X, building.Y)
	if err == nil {
		t.Error("Error should be returned signaling that the destination space is occupied by another unit/building")
	}

	err = newMap.MoveUnit(building, buildingDestX, buildingDestY)
	if err == nil {
		t.Error("Building can't move")
	}

}

func TestSpawnUnitFromBuilding(t *testing.T) {

	p1 := Player{Name: "Player 1"}
	players := []Player{p1}

	buildingX := 2
	buildingY := 2
	building := &Building{Name: "Barracks"}

	unit1X := 3
	unit1Y := 3
	unit1 := &Unit{Name: "Soldier", Player: p1}

	newMap := NewMap("Map1", 10, 10, players)
	newMap.SpawnUnit(building, buildingX, buildingY)

	err := newMap.SpawnUnitFromBuilding(building, unit1, unit1X, unit1Y)

	if err != nil {
		t.Error("Building should be able to spawn unit")
	}

	err = newMap.SpawnUnitFromBuilding(building, unit1, buildingX, buildingY)

	if err == nil {
		t.Error("Unit should not be able to spawn in occupied space")
	}
}
