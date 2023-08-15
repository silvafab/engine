package main

type Coordinates struct {
	X int
	Y int
}

type Spawnable interface {
	SetLocation(x, y int)
	GetLocation() (int, int)
	CanMove() bool
}

type Spawner interface {
	CanSpawn() bool
}

type Unit struct {
	Name    string
	Attack  int
	Defense int
	Player  Player
	Coordinates
}

// TODO: Implement both functions as generic functions
func (u *Unit) SetLocation(x, y int) {
	u.X = x
	u.Y = y
}

func (u Unit) GetLocation() (int, int) {
	return u.X, u.Y
}

func (u Unit) CanMove() bool {
	return true
}

func (u Unit) CanSpawn() bool {
	return false
}

type Building struct {
	Name   string
	Player Player
	Coordinates
}

func (u *Building) SetLocation(x, y int) {
	u.X = x
	u.Y = y
}

func (u Building) GetLocation() (int, int) {
	return u.X, u.Y
}

func (b Building) CanMove() bool {
	return false
}

func (b Building) CanSpawn() bool {
	return true
}
