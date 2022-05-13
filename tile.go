package main

type tile struct {
	tileCode        string
	tileSlideAmount float64
}

func (t *tile) getStaticData() *tileStaticData {
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
}

var tileStaticTable = map[string]*tileStaticData{
	"WALL": {
		passable: false,
		opaque:   true,
	},
	"FLOOR": {
		passable: true,
		opaque:   false,
	},
	"DOOR": {
		passable: false,
		opaque:   true,
		thin:     true,
		openable: true,
	},
}
