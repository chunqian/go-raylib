package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
	"unsafe"
)

const GLSL_VERSION = 330

var POINTS_OF_INTEREST = [6][2]float32{
	{-0.348827, 0.607167},
	{-0.786268, 0.169728},
	{-0.8, 0.156},
	{0.285, 0},
	{-0.835, -0.2321},
	{-0.70176, -0.3842},
}

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - julia sets")
	defer rl.CloseWindow()

	shader := rl.LoadShader("", fmt.Sprintf("../shaders/resources/shaders/glsl%d/julia_set.fs", GLSL_VERSION))
	defer rl.UnloadShader(shader)

	c := [2]float32{POINTS_OF_INTEREST[0][0], POINTS_OF_INTEREST[0][1]}

	offset := [2]float32{-float32(screenWidth / 2), -float32(screenHeight / 2)}
	zoom := float32(1.0)

	offsetSpeed := rl.NewVector2(0, 0)

	cLoc := rl.GetShaderLocation(shader, "c")
	zoomLoc := rl.GetShaderLocation(shader, "zoom")
	offsetLoc := rl.GetShaderLocation(shader, "offset")

	screenDims := [2]float32{float32(screenWidth), float32(screenHeight)}

	rl.SetShaderValue(shader, rl.GetShaderLocation(shader, "screenDims"), unsafe.Pointer(&screenDims), int32(rl.UNIFORM_VEC2))

	rl.SetShaderValue(shader, cLoc, unsafe.Pointer(&c), int32(rl.UNIFORM_VEC2))
	rl.SetShaderValue(shader, zoomLoc, unsafe.Pointer(&zoom), int32(rl.UNIFORM_FLOAT))
	rl.SetShaderValue(shader, offsetLoc, unsafe.Pointer(&offset), int32(rl.UNIFORM_VEC2))

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	incrementSpeed := float32(0)
	showControls := true
	pause := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_ONE)) ||
			rl.IsKeyPressed(int32(rl.KEY_TWO)) ||
			rl.IsKeyPressed(int32(rl.KEY_THREE)) ||
			rl.IsKeyPressed(int32(rl.KEY_FOUR)) ||
			rl.IsKeyPressed(int32(rl.KEY_FIVE)) ||
			rl.IsKeyPressed(int32(rl.KEY_SIX)) {

			if rl.IsKeyPressed(int32(rl.KEY_ONE)) {
				c[0] = POINTS_OF_INTEREST[0][0]
				c[1] = POINTS_OF_INTEREST[0][1]
			} else if rl.IsKeyPressed(int32(rl.KEY_TWO)) {
				c[0] = POINTS_OF_INTEREST[1][0]
				c[1] = POINTS_OF_INTEREST[1][1]
			} else if rl.IsKeyPressed(int32(rl.KEY_THREE)) {
				c[0] = POINTS_OF_INTEREST[2][0]
				c[1] = POINTS_OF_INTEREST[2][1]
			} else if rl.IsKeyPressed(int32(rl.KEY_FOUR)) {
				c[0] = POINTS_OF_INTEREST[3][0]
				c[1] = POINTS_OF_INTEREST[3][1]
			} else if rl.IsKeyPressed(int32(rl.KEY_FIVE)) {
				c[0] = POINTS_OF_INTEREST[4][0]
				c[1] = POINTS_OF_INTEREST[4][1]
			} else if rl.IsKeyPressed(int32(rl.KEY_SIX)) {
				c[0] = POINTS_OF_INTEREST[5][0]
				c[1] = POINTS_OF_INTEREST[5][1]
			}

			rl.SetShaderValue(shader, cLoc, unsafe.Pointer(&c), int32(rl.UNIFORM_VEC2))
		}

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			pause = !pause
		}
		if rl.IsKeyPressed(int32(rl.KEY_F1)) {
			showControls = !showControls
		}

		if !pause {

			if rl.IsKeyPressed(int32(rl.KEY_RIGHT)) {
				incrementSpeed++
			} else if rl.IsKeyPressed(int32(rl.KEY_LEFT)) {
				incrementSpeed--
			}

			if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) ||
				rl.IsMouseButtonDown(int32(rl.MOUSE_RIGHT_BUTTON)) {

				if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {
					zoom += zoom * 0.003
				}
				if rl.IsMouseButtonDown(int32(rl.MOUSE_RIGHT_BUTTON)) {
					zoom -= zoom * 0.003
				}

				mousePos := rl.GetMousePosition()

				offsetSpeed.X = mousePos.X - float32(screenWidth/2)
				offsetSpeed.Y = mousePos.Y - float32(screenHeight/2)

				offset[0] += rl.GetFrameTime() * offsetSpeed.X * 0.8
				offset[1] += rl.GetFrameTime() * offsetSpeed.Y * 0.8
			} else {
				offsetSpeed = rl.NewVector2(0, 0)
			}

			rl.SetShaderValue(shader, zoomLoc, unsafe.Pointer(&zoom), int32(rl.UNIFORM_FLOAT))
			rl.SetShaderValue(shader, offsetLoc, unsafe.Pointer(&offset), int32(rl.UNIFORM_VEC2))

			amount := rl.GetFrameTime() * incrementSpeed * 0.0005
			c[0] += amount
			c[1] += amount

			rl.SetShaderValue(shader, cLoc, unsafe.Pointer(&c), int32(rl.UNIFORM_VEC2))
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginTextureMode(target)

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Black)

		rl.EndTextureMode()

		rl.BeginShaderMode(shader)

		rl.DrawTexture(target.Texture, 0, 0, rl.White)

		rl.EndShaderMode()

		if showControls {
			rl.DrawText("Press Mouse buttons right/left to zoom in/out and move", 10, 15, 10, rl.RayWhite)
			rl.DrawText("Press KEY_F1 to toggle these controls", 10, 30, 10, rl.RayWhite)
			rl.DrawText("Press KEYS [1 - 6] to change point of interest", 10, 45, 10, rl.RayWhite)
			rl.DrawText("Press KEY_LEFT | KEY_RIGHT to change speed", 10, 60, 10, rl.RayWhite)
			rl.DrawText("Press KEY_SPACE to pause movement animation", 10, 75, 10, rl.RayWhite)
		}

		rl.EndDrawing()
	}
}
