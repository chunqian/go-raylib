package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

const GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - vr simulator")
	defer rl.CloseWindow()

	rl.InitVrSimulator()
	defer rl.CloseVrSimulator()

	hmd := rl.VrDeviceInfo{}

	hmd.HResolution = 2160
	hmd.VResolution = 1200
	hmd.HScreenSize = 0.133793
	hmd.VScreenSize = 0.0669
	hmd.VScreenCenter = 0.04678
	hmd.EyeToScreenDistance = 0.041
	hmd.LensSeparationDistance = 0.07
	hmd.InterpupillaryDistance = 0.07

	hmd.LensDistortionValues = [4]float32{1.0, 0.22, 0.24, 0}
	hmd.ChromaAbCorrection = [4]float32{0.996, -0.004, 1.014, 0}

	distortion := rl.LoadShader("", fmt.Sprintf("../core/resources/distortion%d.fs", GLSL_VERSION))
	defer rl.UnloadShader(distortion)

	rl.SetVrConfiguration(hmd, distortion)

	camera := rl.NewCamera(
		rl.NewVector3(5.0, 2.0, 5.0),
		rl.NewVector3(0, 2.0, 0),
		rl.NewVector3(0, 1.0, 0),
		60.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	cubePosition := rl.NewVector3(0, 0, 0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_FIRST_PERSON))

	rl.SetTargetFPS(90)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			rl.ToggleVrMode()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginVrDrawing()

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(40, 1.0)

		rl.EndMode3D()

		rl.EndVrDrawing()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
