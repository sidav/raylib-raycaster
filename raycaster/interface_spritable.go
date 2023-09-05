package raycaster

// "Spritable" is anything that can have a SpriteStruct.
type Spritable interface {
	GetSprite() *SpriteStruct
	GetCoords() (float64, float64)
}
