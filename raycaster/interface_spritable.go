package raycaster

// "Spritable" is anything that can have a SpriteStruct.
type Spritable interface {
	GetSprite() *SpriteStruct
	GetCoords() (float64, float64, float64)       // x, y and height coordinate of Spritable's center
	GetWidthAndHeightFactors() (float64, float64) // 1 is "whole cell"
}
