package main

import (
	"fmt"
	"rl"
	"runtime"
)

var (
	MAX_POSTPRO_SHADERS = 12
)

type PostproShader int32

const (
	FX_GRAYSCALE PostproShader = iota
	FX_POSTERIZATION
	FX_DREAM_VISION
	FX_PIXELIZER
	FX_CROSS_HATCHING
	FX_CROSS_STITCHING
	FX_PREDATOR_VIEW
	FX_SCANLINES
	FX_FISHEYE
	FX_SOBEL
	FX_BLOOM
	FX_BLUR
)

var postproShaderText = []string{
	"GRAYSCALE",
	"POSTERIZATION",
	"DREAM_VISION",
	"PIXELIZER",
	"CROSS_HATCHING",
	"CROSS_STITCHING",
	"PREDATOR_VIEW",
	"SCANLINES",
	"FISHEYE",
	"SOBEL",
	"BLOOM",
	"BLUR",
}

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - postprocessing shader")

	camera := rl.Camera{
		Position: rl.Vector3{X: 2.0, Y: 3.0, Z: 2.0},
		Target:   rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Up:       rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:     45.0,
		Type:     int32(rl.CAMERA_CUSTOM),
	}

	model := rl.LoadModel("../resources/models/church.obj")
	texture := rl.LoadTexture("../resources/models/church_diffuse.png")

	model.GetMaterials(*model.GetMaterialCount())[0].GetMaps(rl.MAP_DIFFUSE + 1)[rl.MAP_DIFFUSE].SetTexture(&texture)

	shaders := make([]rl.Shader, MAX_POSTPRO_SHADERS)

	shaders[FX_GRAYSCALE] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/grayscale.fs", 330))
	shaders[FX_POSTERIZATION] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/posterization.fs", 330))
	shaders[FX_DREAM_VISION] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/dream_vision.fs", 330))
	shaders[FX_PIXELIZER] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/pixelizer.fs", 330))
	shaders[FX_CROSS_HATCHING] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/cross_hatching.fs", 330))
	shaders[FX_CROSS_STITCHING] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/cross_stitching.fs", 330))
	shaders[FX_PREDATOR_VIEW] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/predator.fs", 330))
	shaders[FX_SCANLINES] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/scanlines.fs", 330))
	shaders[FX_FISHEYE] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/fisheye.fs", 330))
	shaders[FX_SOBEL] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/sobel.fs", 330))
	shaders[FX_BLOOM] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/bloom.fs", 330))
	shaders[FX_BLUR] = rl.LoadShader("", fmt.Sprintf("../resources/shaders/glsl%d/blur.fs", 330))

	currentShader := FX_GRAYSCALE

	target := rl.LoadRenderTexture(screenWidth, screenHeight)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera)

		if rl.IsKeyPressed(int32(rl.KEY_RIGHT)) {
			currentShader++
		} else if rl.IsKeyPressed(int32(rl.KEY_LEFT)) {
			currentShader--
		}

		if currentShader >= PostproShader(MAX_POSTPRO_SHADERS) {
			currentShader = 0
		} else if currentShader < 0 {
			currentShader = PostproShader(MAX_POSTPRO_SHADERS) - 1
		}

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.BeginTextureMode(target)

		rl.ClearBackground(*rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 0.1, *rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.EndTextureMode()

		rl.BeginShaderMode(shaders[currentShader])

		rl.DrawTextureRec(*target.GetTexture(), rl.Rectangle{X: 0, Y: 0, Width: float32(*target.GetTexture().GetWidth()), Height: -float32(*target.GetTexture().GetHeight())}, rl.Vector2{X: 0, Y: 0}, *rl.White)

		rl.EndShaderMode()

		rl.DrawRectangle(0, 9, 580, 30, rl.Fade(*rl.LightGray, 0.7))

		rl.DrawText("(c) Church 3D model by Alberto Cano", screenWidth-200, screenHeight-20, 10, *rl.Gray)

		rl.DrawText("CURRENT POSTPRO SHADER:", 10, 15, 20, *rl.Black)

		rl.DrawText(postproShaderText[currentShader], 330, 15, 20, *rl.Red)

		rl.DrawText("< >", 540, 10, 30, *rl.DarkBlue)

		rl.DrawFPS(700, 15)

		rl.EndDrawing()
	}

	for i := 0; i < MAX_POSTPRO_SHADERS; i++ {
		rl.UnloadShader(shaders[i])
	}

	rl.UnloadTexture(texture)
	rl.UnloadModel(model)
	rl.UnloadRenderTexture(target)

	rl.CloseWindow()
}
