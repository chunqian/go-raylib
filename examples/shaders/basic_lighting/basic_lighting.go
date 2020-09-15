package main

import (
	rl "goray/raylib"

	"fmt"
	"math"
	"runtime"
	"unsafe"
	// "runtime/debug"
)

const GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(2.0, 2.0, 6.0),
		rl.NewVector3(0, 0.5, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	modelA := rl.LoadModelFromMesh(rl.GenMeshTorus(0.4, 1.0, 16, 32))
	modelB := rl.LoadModelFromMesh(rl.GenMeshCube(1.0, 1.0, 1.0))
	modelC := rl.LoadModelFromMesh(rl.GenMeshSphere(0.5, 32, 32))

	texture := rl.LoadTexture("../shaders/resources/texel_checker.png")

	modelA.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	modelB.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	modelC.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture

	shader := rl.LoadShader(
		fmt.Sprintf("../shaders/resources/shaders/glsl%d/base_lighting.vs", GLSL_VERSION),
		fmt.Sprintf("../shaders/resources/shaders/glsl%d/lighting.fs", GLSL_VERSION),
	)
	defer func() {
		rl.UnloadShader(shader)
		rl.UnloadTexture(texture)

		rl.UnloadModel(modelA)
		rl.UnloadModel(modelB)
		rl.UnloadModel(modelC)
	}()

	*shader.Locser(int32(rl.LOC_MATRIX_MODEL)) = rl.GetShaderLocation(shader, "matModel")
	*shader.Locser(int32(rl.LOC_VECTOR_VIEW)) = rl.GetShaderLocation(shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(shader, "ambient")
	rl.SetShaderValue(shader, ambientLoc, unsafe.Pointer(&[4]float32{0.2, 0.2, 0.2, 1.0}), int32(rl.UNIFORM_VEC4))

	angle := float32(6.282)

	modelA.Materialser(0).Shader = shader
	modelB.Materialser(0).Shader = shader
	modelC.Materialser(0).Shader = shader

	lights := [MAX_LIGHTS]Light{}
	lights[0] = CreateLight(LIGHT_POINT, rl.NewVector3(4, 2, 4), rl.Vector3Zero(), rl.White, shader)
	lights[1] = CreateLight(LIGHT_POINT, rl.NewVector3(4, 2, 4), rl.Vector3Zero(), rl.Red, shader)
	lights[2] = CreateLight(LIGHT_POINT, rl.NewVector3(0, 4, 2), rl.Vector3Zero(), rl.Green, shader)
	lights[3] = CreateLight(LIGHT_POINT, rl.NewVector3(0, 4, 2), rl.Vector3Zero(), rl.Blue, shader)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_W)) {
			lights[0].enabled = !lights[0].enabled
		}
		if rl.IsKeyPressed(int32(rl.KEY_R)) {
			lights[1].enabled = !lights[1].enabled
		}
		if rl.IsKeyPressed(int32(rl.KEY_G)) {
			lights[2].enabled = !lights[2].enabled
		}
		if rl.IsKeyPressed(int32(rl.KEY_B)) {
			lights[3].enabled = !lights[3].enabled
		}

		rl.UpdateCamera(&camera)

		angle -= 0.02
		lights[0].position.X = float32(math.Cos(float64(angle)) * 4.0)
		lights[0].position.Z = float32(math.Sin(float64(angle)) * 4.0)

		lights[1].position.X = float32(math.Cos(float64(-angle*0.6)) * 4.0)
		lights[1].position.Z = float32(math.Sin(float64(-angle*0.6)) * 4.0)

		lights[2].position.X = float32(math.Cos(float64(angle*0.2)) * 4.0)
		lights[2].position.Z = float32(math.Sin(float64(angle*0.2)) * 4.0)

		lights[3].position.X = float32(math.Cos(float64(-angle*0.35)) * 4.0)
		lights[3].position.Z = float32(math.Sin(float64(-angle*0.35)) * 4.0)

		UpdateLightValues(shader, lights[0])
		UpdateLightValues(shader, lights[1])
		UpdateLightValues(shader, lights[2])
		UpdateLightValues(shader, lights[3])

		modelA.Transform = rl.MatrixMultiply(modelA.Transform, rl.MatrixRotateX(-0.025))
		modelA.Transform = rl.MatrixMultiply(modelA.Transform, rl.MatrixRotateZ(0.012))

		cameraPos := [3]float32{camera.Position.X, camera.Position.Y, camera.Position.Z}
		rl.SetShaderValue(shader, *shader.Locser(int32(rl.LOC_VECTOR_VIEW)), unsafe.Pointer(&cameraPos), int32(rl.UNIFORM_VEC3))

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(modelA, rl.Vector3Zero(), 1.0, rl.White)
		rl.DrawModel(modelB, rl.NewVector3(-1.6, 0, 0), 1.0, rl.White)
		rl.DrawModel(modelC, rl.NewVector3(1.6, 0, 0), 1.0, rl.White)

		if lights[0].enabled {
			rl.DrawSphereEx(lights[0].position, 0.2, 8, 8, rl.White)
		}
		if lights[1].enabled {
			rl.DrawSphereEx(lights[1].position, 0.2, 8, 8, rl.Red)
		}
		if lights[2].enabled {
			rl.DrawSphereEx(lights[2].position, 0.2, 8, 8, rl.Green)
		}
		if lights[3].enabled {
			rl.DrawSphereEx(lights[3].position, 0.2, 8, 8, rl.Blue)
		}

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.DrawText("Use keys RGBW to toggle lights", 10, 30, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
