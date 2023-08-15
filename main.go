package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// a := make([]Unit, 8)

	// a[2] = Unit{Name: "adssad"}

	// fmt.Print(a[2])

	// a[2].Name = "jjjjaaa"
	// processBase(&a[2])
	// fmt.Print(a[2])

	unit1X := 3
	unit1Y := 3
	unit1 := &Unit{Name: "Soldier"}
	p1 := Player{Name: "Player 1"}
	players := []Player{p1}

	newMap := NewMap("Map1", 10, 10, players)

	newMap.SpawnUnit(unit1, unit1X, unit1Y)

	x, y := unit1.GetLocation()
	fmt.Println(x, y)
	// fmt.Println(unit1)
	// fmt.Println(&unit1)
	fmt.Println(newMap.GetUnitInSpace(unit1X, unit1Y))

	unit := newMap.GetUnitInSpace(unit1X, unit1Y)
	// unit.SetLocation(4, 4)
	// fmt.Println(newMap.GetUnitInSpace(unit1X, unit1Y))

	newMap.MoveUnit(unit, 5, 8)
	fmt.Println(newMap.GetUnitInSpace(unit1X, unit1Y))
	fmt.Println(newMap.GetUnitInSpace(5, 8))

}

type BaseI interface{}

type Concrete struct {
	Name string
}

func processBase(u Spawnable) {
	u.SetLocation(1, 12)
}

/*
Initial objects
	Map
	[Representation of game board, has a fixed size, units will spawn inside it]
		size
		name
	Unit (Interface)
	[Smallest entity controllable within the map. Example: ship, soldier]
		name
	Movable (Interface)
	[It allows for a Unit to move across the map]
		movement_units
	Building (Interface)
	[Similar to Unit, it's a separate object to handle static units better. Example: planet, base, watchtower]
		name
*/
