package main

import (
	rl "goray/raylib"

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
	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - fog")
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
	defer func() {
		rl.UnloadTexture(texture)
		rl.UnloadModel(modelA)
		rl.UnloadModel(modelB)
		rl.UnloadModel(modelC)
	}()

	modelA.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	modelB.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	modelC.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture

	shader := rl.LoadShader(
		fmt.Sprintf("../shaders/resources/shaders/glsl%d/base_lighting.vs", GLSL_VERSION),
		fmt.Sprintf("../shaders/resources/shaders/glsl%d/fog.fs", GLSL_VERSION),
	)
	defer rl.UnloadShader(shader)

	*shader.Locser(int32(rl.LOC_MATRIX_MODEL)) = rl.GetShaderLocation(shader, "matModel")
	*shader.Locser(int32(rl.LOC_VECTOR_VIEW)) = rl.GetShaderLocation(shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(shader, "ambient")
	rl.SetShaderValue(shader, ambientLoc, unsafe.Pointer(&[4]float32{0.2, 0.2, 0.2, 1.0}), int32(rl.UNIFORM_VEC4))

	fogDensity := float32(0.15) // must be float32
	fogDensityLoc := rl.GetShaderLocation(shader, "fogDensity")
	rl.SetShaderValue(shader, fogDensityLoc, unsafe.Pointer(&fogDensity), int32(rl.UNIFORM_FLOAT))

	modelA.Materialser(0).Shader = shader
	modelB.Materialser(0).Shader = shader
	modelC.Materialser(0).Shader = shader

	CreateLight(LIGHT_POINT, rl.NewVector3(0, 2, 6), rl.Vector3Zero(), rl.White, shader)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsKeyDown(int32(rl.KEY_UP)) {
			fogDensity += 0.001
			if fogDensity > 1.0 {
				fogDensity = 1.0
			}
		}

		if rl.IsKeyDown(int32(rl.KEY_DOWN)) {
			fogDensity -= 0.001
			if fogDensity < 0.0 {
				fogDensity = 0.0
			}
		}

		rl.SetShaderValue(shader, fogDensityLoc, unsafe.Pointer(&fogDensity), int32(rl.UNIFORM_FLOAT))

		modelA.Transform = rl.MatrixMultiply(modelA.Transform, rl.MatrixRotateX(-0.025))
		modelA.Transform = rl.MatrixMultiply(modelA.Transform, rl.MatrixRotateZ(0.012))

		rl.SetShaderValue(shader, *shader.Locser(int32(rl.LOC_VECTOR_VIEW)), unsafe.Pointer(&camera.Position.X), int32(rl.UNIFORM_VEC3))

		rl.BeginDrawing()

		rl.ClearBackground(rl.Gray)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(modelA, rl.Vector3Zero(), 1.0, rl.White)
		rl.DrawModel(modelB, rl.NewVector3(-2.6, 0, 0), 1.0, rl.White)
		rl.DrawModel(modelC, rl.NewVector3(2.6, 0, 0), 1.0, rl.White)

		for i := -20; i < 20; i += 2 {

			rl.DrawModel(modelA, rl.NewVector3(float32(i), 0, 2), 1.0, rl.White)
		}

		rl.EndMode3D()

		rl.DrawText(fmt.Sprintf("Use KEY_UP/KEY_DOWN to change fog density [%.2f]", fogDensity), 10, 10, 20, rl.RayWhite)

		rl.EndDrawing()
	}
}
