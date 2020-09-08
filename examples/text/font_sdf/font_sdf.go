package main

import (
	"fmt"
	rl "goray/raylib"

	"runtime"
)

var GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - SDF fonts")
	defer rl.CloseWindow()

	msg := "Signed Distance Fields"

	fontDefault := rl.Font{}

	fontDefault.BaseSize = 16
	fontDefault.CharsCount = 95
	*fontDefault.Chars(0) = *rl.LoadFontData("../text/resources/anonymous_pro_bold.ttf", 16, nil, 95, int32(rl.FONT_DEFAULT))

	recs := fontDefault.Recs(0)
	stlas := rl.GenImageFontAtlas(fontDefault.Chars(0), &recs, 95, 16, 4, 0)
	fontDefault.Texture = rl.LoadTextureFromImage(stlas)
	rl.UnloadImage(stlas)
	defer rl.UnloadFont(fontDefault)

	fontSDF := rl.Font{}

	fontSDF.BaseSize = 16
	fontSDF.CharsCount = 95
	*fontSDF.Chars(0) = *rl.LoadFontData("../text/resources/anonymous_pro_bold.ttf", 16, nil, 0, int32(rl.FONT_SDF))

	recs = fontSDF.Recs(0)
	stlas = rl.GenImageFontAtlas(fontSDF.Chars(0), &recs, 95, 16, 0, 1)
	fontSDF.Texture = rl.LoadTextureFromImage(stlas)
	rl.UnloadImage(stlas)
	defer rl.UnloadFont(fontSDF)

	shader := rl.LoadShader("", fmt.Sprintf("../text/resources/shaders/glsl%d/sdf.fs", GLSL_VERSION))
	defer rl.UnloadShader(shader)

	rl.SetTextureFilter(fontSDF.Texture, int32(rl.FILTER_BILINEAR))

	fontPosition := rl.NewVector2(40, float32(screenHeight/2-50))
	textSize := rl.NewVector2(0, 0)

	fontSize := float32(16.0)
	currentFont := int32(0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		fontSize += float32(rl.GetMouseWheelMove() * 8)

		if fontSize < 6 {
			fontSize = 6
		}

		if rl.IsKeyDown(int32(rl.KEY_SPACE)) {
			currentFont = 1
		} else {
			currentFont = 0
		}

		if currentFont == 0 {
			textSize = rl.MeasureTextEx(fontDefault, msg, fontSize, 0)
		} else {
			textSize = rl.MeasureTextEx(fontSDF, msg, fontSize, 0)
		}

		fontPosition.X = float32(rl.GetScreenWidth())/2 - textSize.X/2
		fontPosition.Y = float32(rl.GetScreenHeight())/2 - textSize.Y/2 + 80

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if currentFont == 1 {
			rl.BeginShaderMode(shader)
			rl.DrawTextEx(fontSDF, msg, fontPosition, fontSize, 0, rl.Black)
			rl.EndShaderMode()

			rl.DrawTexture(fontSDF.Texture, 10, 10, rl.Black)
		} else {
			rl.DrawTextEx(fontDefault, msg, fontPosition, fontSize, 0, rl.Black)
			rl.DrawTexture(fontDefault.Texture, 10, 10, rl.Black)
		}

		if currentFont == 1 {
			rl.DrawText("SDF!", 320, 20, 80, rl.Red)
		} else {
			rl.DrawText("default font", 315, 40, 30, rl.Gray)
		}

		rl.DrawText("FONT SIZE: 16.0", rl.GetScreenWidth()-240, 20, 20, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("RENDER SIZE: %.2f", fontSize), rl.GetScreenWidth()-240, 50, 20, rl.DarkGray)
		rl.DrawText("Use MOUSE WHEEL to SCALE TEXT!", rl.GetScreenWidth()-240, 90, 10, rl.DarkGray)

		rl.DrawText("HOLD SPACE to USE SDF FONT VERSION!", 340, rl.GetScreenHeight()-30, 20, rl.Maroon)

		rl.EndDrawing()
	}
}
