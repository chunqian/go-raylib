package main

import (
	rl "goray/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - font loading")

	msg := "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHI\nJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmn\nopqrstuvwxyz{|}~¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓ\nÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷\nøùúûüýþÿ"

	fontBm := rl.LoadFont("../text/resources/pixantiqua.fnt")
	defer rl.UnloadFont(fontBm)

	fontTtf := rl.LoadFontEx("../text/resources/pixantiqua.ttf", 32, nil, 250)
	defer rl.UnloadFont(fontTtf)

	useTtf := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_SPACE)) {
			useTtf = true
		} else {
			useTtf = false
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Hold SPACE to use TTF generated font", 20, 20, 20, rl.LightGray)

		if !useTtf {
			rl.DrawTextEx(
				fontBm,
				msg,
				rl.NewVector2(20, 100),
				float32(fontBm.This.BaseSize),
				2,
				rl.Maroon,
			)
			rl.DrawText("Using BMFont (Angelcode) imported", 20, rl.GetScreenHeight()-30, 20, rl.Gray)
		} else {
			rl.DrawTextEx(
				fontTtf,
				msg,
				rl.NewVector2(20, 100),
				float32(fontTtf.This.BaseSize),
				2,
				rl.Lime,
			)
			rl.DrawText("Using TTF font generated", 20, rl.GetScreenHeight()-30, 20, rl.Gray)
		}

		rl.EndDrawing()
	}
}
