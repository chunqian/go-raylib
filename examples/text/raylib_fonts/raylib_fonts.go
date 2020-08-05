package main

import (
	rl "goray/raylib"

	"runtime"
)

var MAX_FONTS = 8

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - raylib fonts")

	fonts := make([]rl.Font, MAX_FONTS)

	fonts[0] = rl.LoadFont("../text/resources/fonts/alagard.png")
	fonts[1] = rl.LoadFont("../text/resources/fonts/pixelplay.png")
	fonts[2] = rl.LoadFont("../text/resources/fonts/mecha.png")
	fonts[3] = rl.LoadFont("../text/resources/fonts/setback.png")
	fonts[4] = rl.LoadFont("../text/resources/fonts/romulus.png")
	fonts[5] = rl.LoadFont("../text/resources/fonts/pixantiqua.png")
	fonts[6] = rl.LoadFont("../text/resources/fonts/alpha_beta.png")
	fonts[7] = rl.LoadFont("../text/resources/fonts/jupiter_crash.png")

	messages := []string{
		"ALAGARD FONT designed by Hewett Tsoi",
		"PIXELPLAY FONT designed by Aleksander Shevchuk",
		"MECHA FONT designed by Captain Falcon",
		"SETBACK FONT designed by Brian Kent (AEnigma)",
		"ROMULUS FONT designed by Hewett Tsoi",
		"PIXANTIQUA FONT designed by Gerhard Grossmann",
		"ALPHA_BETA FONT designed by Brian Kent (AEnigma)",
		"JUPITER_CRASH FONT designed by Brian Kent (AEnigma)",
	}

	spacings := []int32{2, 4, 8, 4, 3, 4, 4, 1}

	positions := make([]rl.Vector2, MAX_FONTS)

	for i := 0; i < MAX_FONTS; i++ {
		fontT := fonts[i].Convert()
		positionT := positions[i].Convert()

		measureText := rl.MeasureTextEx(fonts[i], messages[i], float32(fontT.BaseSize*2), float32(spacings[i]))
		measureTextT := measureText.Convert()

		positionT.X = float32(screenWidth/2) - measureTextT.X/2
		positionT.Y = float32(int32(60) + fontT.BaseSize + int32(45*i))
	}

	positions[3].Convert().Y += 8
	positions[4].Convert().Y += 2
	positions[7].Convert().Y += 8

	colors := []rl.Color{
		*rl.Maroon,
		*rl.Orange,
		*rl.DarkGreen,
		*rl.DarkBlue,
		*rl.DarkPurple,
		*rl.Lime,
		*rl.Gold,
		*rl.Red,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.DrawText("free fonts included with raylib", 250, 20, 20, *rl.DarkGray)
		rl.DrawLine(220, 50, 590, 50, *rl.DarkGray)

		for i := 0; i < MAX_FONTS; i++ {
			rl.DrawTextEx(
				fonts[i],
				messages[i],
				positions[i],
				float32(fonts[i].Convert().BaseSize*2),
				float32(spacings[i]),
				colors[i],
			)
		}

		rl.EndDrawing()
	}

	for i := 0; i < MAX_FONTS; i++ {
		rl.UnloadFont(fonts[i])
	}

	rl.CloseWindow()

}
