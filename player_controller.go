package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *game) workPlayerInput() {
	if rl.IsKeyDown(rl.KeyUp) {
		g.scene.Camera.MoveForward(0.03)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.scene.Camera.Rotate(3 * -3.141592654 / 180)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.scene.Camera.Rotate(3 * 3.141592654 / 180)
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		g.scene.Camera.MoveForward(-1)
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.scene.things = append(g.scene.things, &thing{
			x:          g.scene.Camera.X,
			y:          g.scene.Camera.Y,
			spriteCode: "proj",
		})
	}
}

func (g *game) movePlayerForward() {

}
