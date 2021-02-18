package main

import (
	rl "go-raylib/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - drop files")
	defer rl.CloseWindow()

	count := int32(0)
	var droppedFiles **byte

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsFileDropped() {
			droppedFiles = rl.GetDroppedFiles(&count)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if count == 0 {
			rl.DrawText("Drop your files to this window!", 100, 40, 20, rl.DarkGray)
		} else {
			rl.DrawText("Dropped files:", 100, 40, 20, rl.DarkGray)

			for i := int32(0); i < count; i++ {

				if i%2 == 0 {
					rl.DrawRectangle(0, 85+40*i, screenWidth, 40, rl.Fade(rl.LightGray, 0.5))
				} else {
					rl.DrawRectangle(0, 85+40*i, screenWidth, 40, rl.Fade(rl.LightGray, 0.3))
				}
				rl.DrawText(rl.ToString(droppedFiles, i), 120, 100+40*i, 10, rl.Gray)
			}

			rl.DrawText("Drop new files...", 100, 110+40*count, 20, rl.DarkGray)
		}
		rl.EndDrawing()
	}

	rl.ClearDroppedFiles()
}
