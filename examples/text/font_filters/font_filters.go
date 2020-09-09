package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - font filters")
	defer rl.CloseWindow()

	msg := "Loaded Font"

	font := rl.LoadFontEx("../text/resources/KAISG.ttf", 96, nil, 0)
	defer rl.UnloadFont(font)

	rl.GenTextureMipmaps(&font.Texture)

	fontSize := font.BaseSize
	fontPosition := rl.NewVector2(40, float32(screenHeight)/2-80)
	textSize := rl.NewVector2(0, 0)

	rl.SetTextureFilter(font.Texture, int32(rl.FILTER_POINT))
	currentFontFilter := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		fontSize += rl.GetMouseWheelMove() * 4.0

		if rl.IsKeyPressed(int32(rl.KEY_ONE)) {

			rl.SetTextureFilter(font.Texture, int32(rl.FILTER_POINT))
			currentFontFilter = 0
		} else if rl.IsKeyPressed(int32(rl.KEY_TWO)) {

			rl.SetTextureFilter(font.Texture, int32(rl.FILTER_BILINEAR))
			currentFontFilter = 1
		} else if rl.IsKeyPressed(int32(rl.KEY_THREE)) {

			rl.SetTextureFilter(font.Texture, int32(rl.FILTER_TRILINEAR))
			currentFontFilter = 2
		}

		textSize = rl.MeasureTextEx(font, msg, float32(fontSize), 0)

		if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			fontPosition.X -= 10
		} else if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			fontPosition.X += 10
		}

		if rl.IsFileDropped() {

			count := int32(0)
			droppedFiles := rl.GetDroppedFiles(&count)
			droppedFilePath := rl.ToString(droppedFiles, 0)

			if rl.IsFileExtension(droppedFilePath, ".ttf") {
				rl.UnloadFont(font)
				font = rl.LoadFontEx(droppedFilePath, fontSize, nil, 0)
				rl.ClearDroppedFiles()
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Use mouse wheel to change font size", 20, 20, 10, rl.Gray)
		rl.DrawText("Use KEY_RIGHT and KEY_LEFT to move text", 20, 40, 10, rl.Gray)
		rl.DrawText("Use 1, 2, 3 to change texture filter", 20, 60, 10, rl.Gray)
		rl.DrawText("Drop a new TTF font for dynamic loading", 20, 80, 10, rl.DarkGray)

		rl.DrawTextEx(font, msg, fontPosition, float32(fontSize), 0, rl.Black)

		rl.DrawRectangle(0, screenHeight-80, screenWidth, 80, rl.LightGray)
		rl.DrawText(fmt.Sprintf("Font size: %2.2f", float32(fontSize)), 20, screenHeight-50, 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("Text size: [%2.2f, %2.2f]", textSize.X, textSize.Y), 20, screenHeight-30, 10, rl.DarkGray)
		rl.DrawText("CURRENT TEXTURE FILTER:", 250, 400, 20, rl.Gray)

		if currentFontFilter == 0 {
			rl.DrawText("POINT", 570, 400, 20, rl.Black)
		} else if currentFontFilter == 1 {
			rl.DrawText("BILINEAR", 570, 400, 20, rl.Black)
		} else if currentFontFilter == 2 {
			rl.DrawText("TRILINEAR", 570, 400, 20, rl.Black)
		}

		rl.EndDrawing()
	}

	rl.ClearDroppedFiles()
}
