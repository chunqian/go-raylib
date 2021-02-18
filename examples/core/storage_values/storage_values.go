package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

type StorageData int32

const (
	STORAGE_POSITION_SCORE StorageData = iota
	STORAGE_POSITION_HISCORE
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - storage save/load values")
	defer rl.CloseWindow()

	score := int32(0)
	hiscore := int32(0)
	framesCounter := int32(0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_R)) {

			score = rl.GetRandomValue(1000, 2000)
			hiscore = rl.GetRandomValue(2000, 4000)
		}

		if rl.IsKeyPressed(int32(rl.KEY_ENTER)) {

			rl.SaveStorageValue(uint32(STORAGE_POSITION_SCORE), score)
			rl.SaveStorageValue(uint32(STORAGE_POSITION_HISCORE), hiscore)
		} else if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {

			score = rl.LoadStorageValue(uint32(STORAGE_POSITION_SCORE))
			hiscore = rl.LoadStorageValue(uint32(STORAGE_POSITION_HISCORE))
		}

		framesCounter++

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("SCORE: %d", score), 280, 130, 40, rl.Maroon)
		rl.DrawText(fmt.Sprintf("HI-SCORE: %d", hiscore), 210, 200, 50, rl.Black)

		rl.DrawText(fmt.Sprintf("frames: %d", framesCounter), 10, 10, 20, rl.Lime)

		rl.DrawText("Press R to generate random numbers", 220, 40, 20, rl.LightGray)
		rl.DrawText("Press ENTER to SAVE values", 250, 310, 20, rl.LightGray)
		rl.DrawText("Press SPACE to LOAD values", 252, 350, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
