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

func (m *Map) IsSpaceOccupiedByUnit(x, y int) bool {
	return m.Spaces[x][y] != nil
}

func (m *Map) SpawnUnit(unit *Unit, x, y int) error {
	if m.IsSpaceOccupiedByUnit(x, y) {
		return errors.New("Space occupied by another unit")
	}
	unit.X = x
	unit.Y = y
	m.Spaces[x][y] = unit
	return nil
}

func (m *Map) GetUnitInSpace(x, y int) *Unit {
	return m.Spaces[x][y]
}

func (m *Map) MoveUnit(unit *Unit, destX, destY int) error {

	attack := m.Spaces[destX][destY] != nil
	move := true

	if attack {
		unitInSpace := m.Spaces[destX][destY]
		if unitInSpace != nil {
			return errors.New("Space occupied by another unit")
		}
		// move, _ = m.UnitAttack(unit, unitInSpace)
	}

	if move {
		m.Spaces[unit.X][unit.Y] = nil
		m.Spaces[destX][destY] = unit
		unit.X = destX
		unit.Y = destY
	}

	return nil
}

func (m *Map) UnitAttack(attacker, defender *Unit) (bool, error) {

	if attacker.Attack > defender.Defense {
		return true, nil
	}

	return false, nil
}
