package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - text formatting")
	defer rl.CloseWindow()

	score := 100020
	hiscore := 200450
	lives := 5

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("Score: %d", score), 200, 80, 20, rl.Red)

		rl.DrawText(fmt.Sprintf("HiScore: %d", hiscore), 200, 120, 20, rl.Green)

		rl.DrawText(fmt.Sprintf("Lives: %d", lives), 200, 160, 40, rl.Blue)

		rl.DrawText(fmt.Sprintf("Elapsed Time: %2.2f ms", rl.GetFrameTime()*1000), 200, 220, 20, rl.Black)

		rl.EndDrawing()
	}
}
