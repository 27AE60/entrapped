package entrapped

import (
	"testing"
)

func TestCreateEmptyMineField(t *testing.T) {
	field := CreateEmptyMineField(3, 2)

	if field.width != 3 || field.height != 2 {
		t.Error("wrong mine field size")
		return
	}
	if len(field.field) != (3 * 2) {
		t.Error("wrong mine field size")
		return
	}
	for i := 0; i < len(field.field); i++ {
		if field.field[i].opened != false || field.field[i].contains != TypeEmpty {
			t.Error("wrong initialization")
			return
		}
	}
}

func TestChangeBlockType(t *testing.T) {
	validTypeMap := map[int]int{3: 4, 4: 5, 1: 2}
	invalidTypeMap := map[int]int{3: 4, 6: 5}

	field := CreateEmptyMineField(3, 2).ChangeBlockType(validTypeMap)
	for i := 0; i < len(field.field); i++ {
		if field.field[i].opened != false {
			t.Error("wrong initialization")
			return
		}

		if val, ok := validTypeMap[i]; ok {
			if field.field[i].contains != val {
				t.Error("wrong initialization")
				return
			}
		} else {
			if field.field[i].contains != TypeEmpty {
				t.Error("wrong initialization")
				return
			}
		}
	}

	oField := CreateEmptyMineField(3, 2).ChangeBlockType(invalidTypeMap)
	for i := 0; i < len(oField.field); i++ {
		if oField.field[i].opened != false {
			t.Error("wrong initialization")
			return
		}

		if val, ok := invalidTypeMap[i]; ok {
			if oField.field[i].contains != val {
				t.Error("wrong initialization")
				return
			}
		} else {
			if oField.field[i].contains != TypeEmpty {
				t.Error("wrong initialization")
				return
			}
		}
	}
}

func TestOpenAndCheckBlock(t *testing.T) {
	field := CreateEmptyMineField(3, 2)

	openOne, typeOne, typeErr := field.CheckBlock(8)
	if typeErr == nil || typeErr != ErrIndexOutOfRange {
		t.Error("wrong index access")
		return
	}
	if openOne != false || typeOne != TypeEmpty {
		t.Error("open not set to true")
		return
	}

	openTwo, typeTwo, typeErr := field.OpenBlock(3)
	if typeErr != nil || openTwo != true || typeTwo != TypeEmpty {
		t.Error("open not set to true")
		return
	}
}
