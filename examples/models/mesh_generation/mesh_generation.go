package main

import (
	rl "goray/raylib"

	"runtime"
)

const (
	NUM_MODELS = 8
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - mesh generation")
	defer rl.CloseWindow()

	checked := rl.GenImageChecked(2, 2, 1, 1, rl.Red, rl.Green)
	texture := rl.LoadTextureFromImage(checked)
	defer rl.UnloadTexture(texture)

	rl.UnloadImage(checked)

	models := [NUM_MODELS]rl.Model{}

	models[0] = rl.LoadModelFromMesh(rl.GenMeshPlane(2, 2, 5, 5))
	models[1] = rl.LoadModelFromMesh(rl.GenMeshCube(2, 1, 2))
	models[2] = rl.LoadModelFromMesh(rl.GenMeshSphere(2, 32, 32))
	models[3] = rl.LoadModelFromMesh(rl.GenMeshHemiSphere(2, 16, 16))
	models[4] = rl.LoadModelFromMesh(rl.GenMeshCylinder(1, 2, 16))
	models[5] = rl.LoadModelFromMesh(rl.GenMeshTorus(0.25, 4, 16, 32))
	models[6] = rl.LoadModelFromMesh(rl.GenMeshKnot(1, 2, 16, 128))
	models[7] = rl.LoadModelFromMesh(rl.GenMeshPoly(5, 2))
	defer func() {
		for i := range models {
			rl.UnloadModel(models[i])
		}
	}()

	for i := range models {
		models[i].Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = rl.Texture(texture)
	}

	camera := rl.NewCamera(
		rl.NewVector3(5.0, 5.0, 5.0),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		0,
	)

	position := rl.NewVector3(0, 0, 0)

	currentModel := 0

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsMouseButtonPressed(int32(rl.MOUSE_LEFT_BUTTON)) {
			currentModel = (currentModel + 1) % NUM_MODELS
		}

		if rl.IsKeyPressed(int32(rl.KEY_RIGHT)) {
			currentModel++
			if currentModel >= NUM_MODELS {
				currentModel = 0
			}
		} else if rl.IsKeyPressed(int32(rl.KEY_LEFT)) {
			currentModel--
			if currentModel < 0 {
				currentModel = NUM_MODELS - 1
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(models[currentModel], position, 1.0, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(30, 400, 310, 30, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(30, 400, 310, 30, rl.Fade(rl.DarkBlue, 0.5))
		rl.DrawText("MOUSE LEFT BUTTON to CYCLE PROCEDURAL MODELS", 40, 410, 10, rl.Blue)

		switch currentModel {
		case 0:
			rl.DrawText("PLANE", 680, 10, 20, rl.DarkBlue)
			break
		case 1:
			rl.DrawText("CUBE", 680, 10, 20, rl.DarkBlue)
			break
		case 2:
			rl.DrawText("SPHERE", 680, 10, 20, rl.DarkBlue)
			break
		case 3:
			rl.DrawText("HEMISPHERE", 680, 10, 20, rl.DarkBlue)
			break
		case 4:
			rl.DrawText("CYLINDER", 680, 10, 20, rl.DarkBlue)
			break
		case 5:
			rl.DrawText("TORUS", 680, 10, 20, rl.DarkBlue)
			break
		case 6:
			rl.DrawText("KNOT", 680, 10, 20, rl.DarkBlue)
			break
		case 7:
			rl.DrawText("POLY", 680, 10, 20, rl.DarkBlue)
			break
		default:
			break
		}

		rl.EndDrawing()
	}
}
