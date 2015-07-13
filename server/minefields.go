package entrapped

const (
	minSize int = 7
	maxSize int = 10
)

const (
	covered int = 0
	empty   int = 1
	mine    int = -2
)

type mineField struct {
	field []int
	size  int
}

func createEmptyMineField(size int) *mineField {
	if size < minSize {
		size = minSize
	}
	if size > maxSize {
		size = maxSize
	}

	field := make([]int, (size * size))

	return &mineField{field, size}
}

func (m *mineField) addRandomBombs(numBombs int) *mineField {
	size := len(m.field)

	for i := 0; i < numBombs; i++ {
		x := randomInt(size - 1)

		val := m.field[x]
		if val == 0 {
			m.field[x] = -1
		} else {
			i--
		}
	}

	return m
}
