package main

import (
	rl "go-raylib/raylib"

	"fmt"
	"runtime"
	"unsafe"
)

const (
	GLSL_VERSION       = 330
	MAX_PALETTES       = 3
	COLORS_PER_PALETTE = 8
	VALUES_PER_COLOR   = 3
)

var palettes = [MAX_PALETTES][COLORS_PER_PALETTE * VALUES_PER_COLOR]int32{
	{ // 3-BIT RGB
		0, 0, 0,
		255, 0, 0,
		0, 255, 0,
		0, 0, 255,
		0, 255, 255,
		255, 0, 255,
		255, 255, 0,
		255, 255, 255,
	},
	{ // AMMO-8 (GameBoy-like)
		4, 12, 6,
		17, 35, 24,
		30, 58, 41,
		48, 93, 66,
		77, 128, 97,
		137, 162, 87,
		190, 220, 127,
		238, 255, 204,
	},
	{ // RKBV (2-strip film)
		21, 25, 26,
		138, 76, 88,
		217, 98, 117,
		230, 184, 193,
		69, 107, 115,
		75, 151, 166,
		165, 189, 194,
		255, 245, 247,
	},
}

var paletteText = []string{
	"3-BIT RGB",
	"AMMO-8 (GameBoy-like)",
	"RKBV (2-strip film)",
}

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - color palette switch")
	defer rl.CloseWindow()

	shader := rl.LoadShader("", fmt.Sprintf("../shaders/resources/shaders/glsl%d/palette_switch.fs", GLSL_VERSION))
	defer rl.UnloadShader(shader)

	paletteLoc := rl.GetShaderLocation(shader, "palette")

	currentPalette := 0
	lineHeight := screenHeight / COLORS_PER_PALETTE

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_RIGHT)) {
			currentPalette++
		} else if rl.IsKeyPressed(int32(rl.KEY_LEFT)) {
			currentPalette--
		}

		if currentPalette >= MAX_PALETTES {
			currentPalette = 0
		} else if currentPalette < 0 {
			currentPalette = MAX_PALETTES - 1
		}

		rl.SetShaderValueV(shader, paletteLoc, unsafe.Pointer(&palettes[currentPalette]), int32(rl.UNIFORM_IVEC3), COLORS_PER_PALETTE)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginShaderMode(shader)

		for i := int32(0); i < COLORS_PER_PALETTE; i++ {

			b := byte(i)
			rl.DrawRectangle(0, lineHeight*i, rl.GetScreenWidth(), lineHeight, rl.NewColor(b, b, b, 255))
		}

		rl.EndShaderMode()

		rl.DrawText("< >", 10, 10, 30, rl.DarkBlue)
		rl.DrawText("CURRENT PALETTE:", 60, 15, 20, rl.RayWhite)
		rl.DrawText(paletteText[currentPalette], 300, 15, 20, rl.Red)

		rl.DrawFPS(700, 15)

		rl.EndDrawing()
	}
}
