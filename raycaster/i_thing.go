package raycaster

// "Thing" is anything that can have a SpriteStruct.
type Thing interface {
	GetSprite() *SpriteStruct
	GetCoords() (float64, float64)
}
