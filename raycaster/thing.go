package raycaster

// "Thing" is anything that can have a SpriteStruct.
type Thing struct {
	Sprite *SpriteStruct
	X, Y   float64
}
