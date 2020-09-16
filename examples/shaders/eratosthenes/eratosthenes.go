package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

const GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - Sieve of Eratosthenes")
	defer rl.CloseWindow()

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	shader := rl.LoadShader("", fmt.Sprintf("../shaders/resources/shaders/glsl%d/eratosthenes.fs", GLSL_VERSION))
	defer rl.UnloadShader(shader)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginTextureMode(target)

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Black)

		rl.EndTextureMode()

		rl.BeginShaderMode(shader)

		rl.DrawTextureRec(
			target.Texture,
			rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)),
			rl.NewVector2(0, 0),
			rl.White,
		)

		rl.EndShaderMode()

		rl.EndDrawing()
	}
}
