package entrapped

import (
	"errors"
)

const (
	TypeEmpty int = 0
)

var (
	ErrIndexOutOfRange error = errors.New("invalid position")
)

type block struct {
	opened   bool
	contains int
}

type MineField struct {
	field         []block
	width, height int
}

func CreateEmptyMineField(width, height int) *MineField {
	return &MineField{
		make([]block, (width * height)),
		width,
		height,
	}
}

func (m *MineField) ChangeBlockType(typeMap map[int]int) *MineField {
	for key, val := range typeMap {
		if key >= (m.width * m.height) {
			continue
		}

		m.field[key].contains = val
	}

	return m
}

func (m *MineField) OpenBlock(idx int) (bool, int, error) {
	if idx >= (m.height * m.height) {
		return false, 0, ErrIndexOutOfRange
	}

	m.field[idx].opened = true

	return m.field[idx].opened, m.field[idx].contains, nil
}

func (m *MineField) CheckBlock(idx int) (bool, int, error) {
	if idx >= (m.height * m.height) {
		return false, 0, ErrIndexOutOfRange
	}

	return m.field[idx].opened, m.field[idx].contains, nil
}
