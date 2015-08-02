package entrapped

const (
	bombsPerCellRatio float32 = 0.25
	lifesPerCellRatio float32 = 0.15
	empty             int     = 0
	life              int     = 1
	mine              int     = -1
)

type Trap struct {
	trapMap *MineField
	health  int
}

type trapConf struct {
	width    int
	height   int
	numBombs int
	numLifes int
	health   int
}

func makeTrap(params trapConf) *trap {
	return &trap{
		trapMap: nil,
		health:  params.health,
	}
}

func (t *trap) open(idx int) (int, int, string) {
	if idx > (t.trapMap.size*t.trapMap.size)-1 {
		return 0, 0, "error:invalid idx"
	}

	if t.lifes > 0 {
		ele := t.trapMap.checkIndex(idx)
		if ele == mine {
			t.lifes--
			return ele, t.lifes, ""
		} else if ele == bonusLife {
			t.lifes++
			return ele, t.lifes, ""
		} else {
			return ele, t.lifes, ""
		}
	}

	return 0, 0, ""
}
