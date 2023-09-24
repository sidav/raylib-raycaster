package main

type tileState uint8

const (
	tileStateIdle tileState = iota
	tileStateOpening
	tileStateWaitsToClose
	tileStateClosing
)

type tile struct {
	tileCode          string
	tileSlideAmount   float64
	tileStateCooldown int
	state             tileState
}

func (t *tile) getStaticData() *tileStaticData {
	if tileStaticTable[t.tileCode] == nil {
		panic("No tile data for " + t.tileCode)
	}
	return tileStaticTable[t.tileCode]
}

func (t *tile) isPassable() bool {
	if t.getStaticData().openable {
		return t.tileSlideAmount > 0.5
	}
	return t.getStaticData().passable
}

func (t *tile) isOpened() bool {
	return t.tileSlideAmount >= 0.9
}

func (t *tile) isClosed() bool {
	return t.tileSlideAmount <= 0.0001
}

type tileStaticData struct {
	passable, opaque, thin, openable bool
	opensVertically                  bool // false implies "horizontally"
}

var tileStaticTable = map[string]*tileStaticData{
	"WALL": {
		passable: false,
		opaque:   true,
	},
	"WALLELEC": {
		passable: false,
		opaque:   true,
	},
	"FLOOR": {
		passable: true,
		opaque:   false,
	},
	"DOORVERT": {
		passable:        false,
		opaque:          true,
		thin:            true,
		openable:        true,
		opensVertically: true,
	},
	"DOORHORIZ": {
		passable: false,
		opaque:   true,
		thin:     true,
		openable: true,
	},
}
