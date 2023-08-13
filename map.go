package main

import "errors"

type Map struct {
	Name   string
	SizeX  int
	SizeY  int
	Spaces [][]*Unit
}

type Space struct {
}

func NewMap(name string, sizeX, sizeY int) Map {

	s := make([][]*Unit, sizeX)

	for i := range s {
		s[i] = make([]*Unit, sizeY)
	}

	return Map{
		Name:   name,
		SizeX:  sizeX,
		SizeY:  sizeY,
		Spaces: s,
	}
}

func (m *Map) SpawnUnit(unit *Unit, x, y int) error {
	m.Spaces[x][y] = unit
	return nil
}

func (m *Map) GetUnitInSpace(x, y int) *Unit {
	return m.Spaces[x][y]
}

func (m *Map) MoveUnit(unit *Unit, x, y int, destX, destY int) error {
	unitInSpace := m.Spaces[x][y]
	if unitInSpace != unit {
		return errors.New("Wrong unit")
	}

	attack := m.Spaces[destX][destY] != nil
	move := true
	// if m.Spaces[destX][destY] != nil {
	// 	return errors.New("Space occupied")
	// }

	if attack {
		move, _ = m.UnitAttack(unit, unitInSpace)
	}

	if move {
		m.Spaces[destX][destY] = unit
		m.Spaces[x][y] = nil
	}

	return nil
}

func (m *Map) UnitAttack(attacker, defender *Unit) (bool, error) {

	if attacker.Attack > defender.Defense {
		return true, nil
	}

	return false, nil
}
