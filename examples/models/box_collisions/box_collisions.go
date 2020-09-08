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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - box collisions")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(0, 10, 10),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	playerPosition := rl.NewVector3(0, 1.0, 2.0)
	playerSize := rl.NewVector3(1.0, 2.0, 1.0)
	playerColor := rl.Green

	enemyBoxPos := rl.NewVector3(-4.0, 1.0, 0)
	enemyBoxSize := rl.NewVector3(2.0, 2.0, 2.0)

	enemySpherePos := rl.NewVector3(4.0, 0, 0)
	enemySphereSize := float32(1.5)

	collision := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			playerPosition.X += 0.2
		} else if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			playerPosition.X -= 0.2
		} else if rl.IsKeyDown(int32(rl.KEY_DOWN)) {
			playerPosition.Z += 0.2
		} else if rl.IsKeyDown(int32(rl.KEY_UP)) {
			playerPosition.Z -= 0.2
		}

		collision = false

		if rl.CheckCollisionBoxes(
			rl.NewBoundingBox(
				rl.NewVector3(
					playerPosition.X-playerSize.X/2,
					playerPosition.Y-playerSize.Y/2,
					playerPosition.Z-playerSize.Z/2,
				),
				rl.NewVector3(
					playerPosition.X+playerSize.X/2,
					playerPosition.Y+playerSize.Y/2,
					playerPosition.Z+playerSize.Z/2,
				),
			),
			rl.NewBoundingBox(
				rl.NewVector3(
					enemyBoxPos.X-enemyBoxSize.X/2,
					enemyBoxPos.Y-enemyBoxSize.Y/2,
					enemyBoxPos.Z-enemyBoxSize.Z/2,
				),
				rl.NewVector3(
					enemyBoxPos.X+enemyBoxSize.X/2,
					enemyBoxPos.Y+enemyBoxSize.Y/2,
					enemyBoxPos.Z+enemyBoxSize.Z/2,
				),
			),
		) {
			collision = true
		}

		if rl.CheckCollisionBoxSphere(
			rl.NewBoundingBox(
				rl.NewVector3(
					playerPosition.X-playerSize.X/2,
					playerPosition.Y-playerSize.Y/2,
					playerPosition.Z-playerSize.Z/2,
				),
				rl.NewVector3(
					playerPosition.X+playerSize.X/2,
					playerPosition.Y+playerSize.Y/2,
					playerPosition.Z+playerSize.Z/2,
				),
			),
			enemySpherePos,
			enemySphereSize,
		) {
			collision = true
		}

		if collision {
			playerColor = rl.Red
		} else {
			playerColor = rl.Green
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.ToCamera3D(camera))

		rl.DrawCube(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.Gray)
		rl.DrawCubeWires(enemyBoxPos, enemyBoxSize.X, enemyBoxSize.Y, enemyBoxSize.Z, rl.DarkGray)

		rl.DrawSphere(enemySpherePos, enemySphereSize, rl.Gray)
		rl.DrawSphereWires(enemySpherePos, enemySphereSize, 16, 16, rl.DarkGray)

		rl.DrawCubeV(playerPosition, playerSize, playerColor)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText("Move player with cursors to collide", 220, 40, 20, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
