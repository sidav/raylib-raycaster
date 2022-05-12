package main

type tile struct {
	tileCode string
}

func (t *tile) getStaticData() *tileStaticData {
	return tileStaticTable[t.tileCode]
}

type tileStaticData struct {
	passable, opaque, thin bool
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
	},
}
