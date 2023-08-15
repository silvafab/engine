package main

import "errors"

type Map struct {
	Name    string
	SizeX   int
	SizeY   int
	Spaces  [][]Spawnable
	Players []Player
}

type Space struct {
}

func NewMap(name string, sizeX, sizeY int, players []Player) Map {

	s := make([][]Spawnable, sizeX)

	for i := range s {
		s[i] = make([]Spawnable, sizeY)
	}

	return Map{
		Name:    name,
		SizeX:   sizeX,
		SizeY:   sizeY,
		Spaces:  s,
		Players: players,
	}
}

func (m *Map) IsSpaceOccupied(x, y int) bool {
	return m.Spaces[x][y] != nil
}

func (m *Map) SpawnUnit(unit Spawnable, x, y int) error {
	if m.IsSpaceOccupied(x, y) {
		return errors.New("Space occupied by another unit")
	}
	unit.SetLocation(x, y)
	m.Spaces[x][y] = unit
	return nil
}

func (m *Map) GetUnitInSpace(x, y int) Spawnable {
	return m.Spaces[x][y]
}

func (m *Map) MoveUnit(unit Spawnable, destX, destY int) error {

	attack := m.Spaces[destX][destY] != nil
	move := unit.CanMove()

	if !move {
		return errors.New("Unit cannot move")
	}

	if attack {
		// unitInSpace := m.Spaces[destX][destY]
		// if unitInSpace != nil {

		//TODO: IMPLEMENT ATTACK
		return errors.New("Space occupied by another unit")

		// }
		// move, _ = m.UnitAttack(unit, unitInSpace)
	}

	x, y := unit.GetLocation()
	m.Spaces[x][y] = nil
	m.Spaces[destX][destY] = unit
	unit.SetLocation(destX, destY)

	return nil
}

func (m *Map) UnitAttack(attacker, defender *Unit) (bool, error) {

	if attacker.Attack > defender.Defense {
		return true, nil
	}

	return false, nil
}
