package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

const (
	FLT_MAX = 340282346638528859811704183484516925440.0 // Maximum value of a float, from bit pattern 01111111011111111111111111111111
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - mesh picking")
	defer rl.CloseWindow()

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(20, 20, 20)
	camera.Target = rl.NewVector3(0, 8, 0)
	camera.Up = rl.NewVector3(0, 1.6, 0)
	camera.Fovy = 45.0
	camera.Type = int32(rl.CAMERA_PERSPECTIVE)

	ray := rl.Ray{}

	tower := rl.LoadModel("../models/resources/models/turret.obj")
	texture := rl.LoadTexture("../models/resources/models/turret_diffuse.png")
	tower.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = rl.Texture(texture)
	defer func() {
		rl.UnloadTexture(texture)
		rl.UnloadModel(tower)
	}()

	towerPos := rl.NewVector3(0, 0, 0)
	towerBBox := rl.MeshBoundingBox(*tower.Mesheser(0))
	hitMeshBBox := false
	hitTriangle := false

	// Test triangle
	ta := rl.NewVector3(-25, 0.5, 0)
	tb := rl.NewVector3(-4, 2.5, 1.0)
	tc := rl.NewVector3(-8, 6.5, 0)

	bary := rl.NewVector3(0, 0, 0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_FREE))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		nearestHit := rl.RayHitInfo{}
		nearestHit.Distance = FLT_MAX
		nearestHit.Hit = false

		hitObjectName := "None"
		cursorColor := rl.White

		ray = rl.GetMouseRay(rl.GetMousePosition(), camera)

		groundHitInfo := rl.GetCollisionRayGround(ray, 0)

		if groundHitInfo.Hit && groundHitInfo.Distance < nearestHit.Distance {
			nearestHit = groundHitInfo
			cursorColor = rl.Green
			hitObjectName = "Ground"
		}

		triHitInfo := rl.GetCollisionRayTriangle(ray, ta, tb, tc)

		if triHitInfo.Hit && triHitInfo.Distance < nearestHit.Distance {
			nearestHit = triHitInfo
			cursorColor = rl.Purple
			hitObjectName = "Triangle"

			bary = rl.Vector3Barycenter(nearestHit.Position, ta, tb, tc)
			hitTriangle = true
		} else {
			hitTriangle = false
		}

		meshHitInfo := rl.RayHitInfo{}

		if rl.CheckCollisionRayBox(ray, towerBBox) {
			hitMeshBBox = true

			meshHitInfo = rl.GetCollisionRayModel(ray, tower)
			if meshHitInfo.Hit && meshHitInfo.Distance < nearestHit.Distance {
				nearestHit = meshHitInfo
				cursorColor = rl.Orange
				hitObjectName = "Mesh"
			}
		}

		hitMeshBBox = false

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(tower, towerPos, 1.0, rl.White)

		rl.DrawLine3D(ta, tb, rl.Purple)
		rl.DrawLine3D(tb, tc, rl.Purple)
		rl.DrawLine3D(tc, ta, rl.Purple)

		if hitMeshBBox {
			rl.DrawBoundingBox(towerBBox, rl.Lime)
		}

		if nearestHit.Hit {
			rl.DrawCube(nearestHit.Position, 0.3, 0.3, 0.3, cursorColor)
			rl.DrawCubeWires(nearestHit.Position, 0.3, 0.3, 0.3, rl.Red)

			normalEnd := rl.NewVector3(
				nearestHit.Position.X+nearestHit.Normal.X,
				nearestHit.Position.Y+nearestHit.Normal.Y,
				nearestHit.Position.Z+nearestHit.Normal.Z,
			)

			rl.DrawLine3D(nearestHit.Position, normalEnd, rl.Red)
		}

		rl.DrawRay(ray, rl.Maroon)

		rl.DrawGrid(10, 10)

		rl.EndMode3D()

		rl.DrawText(fmt.Sprintf("Hit Object: %s", hitObjectName), 10, 50, 10, rl.Black)

		if nearestHit.Hit {

			ypos := int32(70)

			rl.DrawText(fmt.Sprintf("Distance: %3.2f", nearestHit.Distance), 10, ypos, 10, rl.Black)
			rl.DrawText(
				fmt.Sprintf(
					"Hit Pos: %3.2f %3.2f %3.2f",
					nearestHit.Position.X,
					nearestHit.Position.Y,
					nearestHit.Position.Z,
				),
				10,
				ypos+15,
				10,
				rl.Black,
			)
			rl.DrawText(
				fmt.Sprintf(
					"Hit Norm: %3.2f %3.2f %3.2f",
					nearestHit.Position.X,
					nearestHit.Position.Y,
					nearestHit.Position.Z,
				),
				10,
				ypos+30,
				10,
				rl.Black,
			)
			if hitTriangle {
				rl.DrawText(fmt.Sprintf("Barycenter: %3.2f %3.2f %3.2f", bary.X, bary.Y, bary.Z), 10, ypos+45, 10, rl.Black)
			}

		}

		rl.DrawText("Use Mouse to Move Camera", 10, 430, 10, rl.Gray)
		rl.DrawText("(c) Turret 3D model by Alberto Cano", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
