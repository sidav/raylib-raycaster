package raycaster

import "container/list"

type Scene interface {
	GetFloorTextureForCoords(int, int) *Texture
	GetCeilingTextureForCoords(int, int) *Texture
	GetCamera() *Camera
	GetListOfThings() *list.List
	AreGridCoordsValid(int, int) bool
	IsTileOpaque(int, int) bool
	IsTileThin(int, int) bool
	GetTileSlideAmount(int, int) float64
	GetTileElevation(int, int) float64
	GetTextureForTile(int, int) *Texture
}
