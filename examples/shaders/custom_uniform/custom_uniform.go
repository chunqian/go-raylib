package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
	"unsafe"
)

const GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - custom uniform variable")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(8.0, 8.0, 8.0),
		rl.NewVector3(0, 1.5, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	model := rl.LoadModel("../shaders/resources/models/barracks.obj")
	texture := rl.LoadTexture("../shaders/resources/models/barracks_diffuse.png")
	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	defer func() {
		rl.UnloadTexture(texture)
		rl.UnloadModel(model)
	}()

	position := rl.NewVector3(0, 0, 0)

	shader := rl.LoadShader("", fmt.Sprintf("../shaders/resources/shaders/glsl%d/swirl.fs", GLSL_VERSION))
	defer rl.UnloadShader(shader)

	swirlCenterLoc := rl.GetShaderLocation(shader, "center")

	swirlCenter := [2]float32{float32(screenWidth / 2), float32(screenHeight / 2)}

	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		mousePosition := rl.GetMousePosition()

		swirlCenter[0] = mousePosition.X
		swirlCenter[1] = mousePosition.Y

		rl.SetShaderValue(shader, swirlCenterLoc, unsafe.Pointer(&swirlCenter), int32(rl.UNIFORM_VEC2))

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginTextureMode(target)

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, position, 0.5, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText("TEXT DRAWN IN RENDER TEXTURE", 200, 10, 30, rl.Red)

		rl.EndTextureMode()

		rl.BeginShaderMode(shader)

		rl.DrawTextureRec(
			target.Texture,
			rl.NewRectangle(
				0,
				0,
				float32(target.Texture.Width),
				-float32(target.Texture.Height),
			),
			rl.NewVector2(0, 0),
			rl.White,
		)

		rl.EndShaderMode()

		rl.DrawText("(c) Barracks 3D model by Alberto Cano", screenWidth-220, screenHeight-20, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
