package main

import (
	phys "github.com/chunqian/go-raylib/physac"
	rl "github.com/chunqian/go-raylib/raylib"

	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "Physac [raylib] - Physics demo")
	defer rl.CloseWindow()

	logoX := screenWidth - rl.MeasureText("Physac", 30) - 10
	logoY := int32(15)

	phys.InitPhysics()

	floor := phys.CreatePhysicsBodyRectangle(
		phys.NewVector2(float32(screenWidth/2), float32(screenHeight)),
		500,
		100,
		10,
	)
	floor.Enabled = false

	circle := phys.CreatePhysicsBodyCircle(
		phys.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		45,
		10,
	)
	circle.Enabled = false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		phys.UpdatePhysics()

		if rl.IsKeyPressed(int32(rl.KEY_R)) {

			phys.ResetPhysics()

			floor = phys.CreatePhysicsBodyRectangle(
				phys.NewVector2(float32(screenWidth/2), float32(screenHeight)),
				500,
				100,
				10,
			)
			floor.Enabled = false

			circle := phys.CreatePhysicsBodyCircle(
				phys.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
				45,
				10,
			)
			circle.Enabled = false
		}

		var pos rl.Vector2
		if rl.IsMouseButtonPressed(int32(rl.MOUSE_BUTTON_LEFT)) {
			pos = rl.GetMousePosition()

			phys.CreatePhysicsBodyPolygon(
				*(*phys.Vector2)(unsafe.Pointer(&pos)),
				float32(rl.GetRandomValue(20, 80)),
				rl.GetRandomValue(3, 8),
				10,
			)
		} else if rl.IsMouseButtonPressed(int32(rl.MOUSE_BUTTON_RIGHT)) {
			pos = rl.GetMousePosition()

			phys.CreatePhysicsBodyCircle(
				*(*phys.Vector2)(unsafe.Pointer(&pos)),
				float32(rl.GetRandomValue(10, 45)),
				10,
			)
		}

		bodiesCount := phys.GetPhysicsBodiesCount()
		for i := bodiesCount - 1; i >= 0; i-- {
			body := phys.GetPhysicsBody(i)
			if body != nil && int32(body.Position.Y) > screenHeight*2 {
				phys.DestroyPhysicsBody(body)
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawFPS(screenWidth-90, screenHeight-30)

		bodiesCount = phys.GetPhysicsBodiesCount()
		for i := int32(0); i < bodiesCount; i++ {
			body := phys.GetPhysicsBody(i)

			if body != nil {
				vertexCount := phys.GetPhysicsShapeVerticesCount(i)
				for j := int32(0); j < vertexCount; j++ {
					vertexA := phys.GetPhysicsShapeVertex(body, j)

					var jj int32
					if j+1 < vertexCount {
						jj = j + 1
					} else {
						jj = 0
					}

					vertexB := phys.GetPhysicsShapeVertex(body, jj)

					rl.DrawLineV(
						*(*rl.Vector2)(unsafe.Pointer(&vertexA)),
						*(*rl.Vector2)(unsafe.Pointer(&vertexB)),
						rl.Green,
					)
				}
			}
		}

		rl.DrawText("Left mouse button to create a polygon", 10, 10, 10, rl.White)
		rl.DrawText("Right mouse button to create a circle", 10, 25, 10, rl.White)
		rl.DrawText("Press 'R' to reset example", 10, 40, 10, rl.White)

		rl.DrawText("Physac", logoX, logoY, 30, rl.White)
		rl.DrawText("Powered by", logoX+50, logoY-7, 10, rl.White)

		rl.EndDrawing()
	}

	phys.ClosePhysics()
}
