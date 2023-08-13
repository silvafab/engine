package main

import "fmt"

func main() {
	fmt.Println("Hello world")
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
