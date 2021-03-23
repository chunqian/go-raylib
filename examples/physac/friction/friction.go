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
	rl.InitWindow(screenWidth, screenHeight, "Physac [raylib] - Physics friction")
	defer rl.CloseWindow()

	logoX := screenWidth - rl.MeasureText("Physac", 30) - 10
	logoY := 15

	phys.InitPhysics()
	defer phys.ClosePhysics()

	floor := phys.CreatePhysicsBodyRectangle(phys.NewVector2(float32(screenWidth/2), float32(screenHeight)), float32(screenWidth), 100, 10)
	floor.Enabled = false
	wall := phys.CreatePhysicsBodyRectangle(phys.NewVector2(float32(screenWidth/2), float32(screenHeight)*0.8), 10, 80, 10)
	wall.Enabled = false

	rectLeft := phys.CreatePhysicsBodyRectangle(phys.NewVector2(25, float32(screenHeight)-5), 250, 250, 10)
	rectLeft.Enabled = false
	phys.SetPhysicsBodyRotation(rectLeft, 30*rl.DEG2RAD)

	rectRight := phys.CreatePhysicsBodyRectangle(phys.NewVector2(float32(screenWidth)-25, float32(screenHeight)-5), 250, 250, 10)
	rectRight.Enabled = false
	phys.SetPhysicsBodyRotation(rectRight, 330*rl.DEG2RAD)

	bodyA := phys.CreatePhysicsBodyRectangle(phys.NewVector2(35, float32(screenHeight)*0.6), 40, 40, 10)
	bodyA.StaticFriction = 0.1
	bodyA.DynamicFriction = 0.1
	phys.SetPhysicsBodyRotation(bodyA, 30*rl.DEG2RAD)

	bodyB := phys.CreatePhysicsBodyRectangle(phys.NewVector2(float32(screenWidth)-35, float32(screenHeight)*0.6), 40, 40, 10)
	bodyB.StaticFriction = 1.0
	bodyB.DynamicFriction = 1.0
	phys.SetPhysicsBodyRotation(bodyB, 330*rl.DEG2RAD)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		phys.RunPhysicsStep()

		if rl.IsKeyPressed('R') {
			bodyA.Position = phys.NewVector2(35, float32(screenHeight)*0.6)
			bodyA.Velocity = phys.NewVector2(0, 0)
			bodyA.AngularVelocity = 0
			phys.SetPhysicsBodyRotation(bodyA, 30*rl.DEG2RAD)

			bodyB.Position = phys.NewVector2(float32(screenWidth)-35, float32(screenHeight)*0.6)
			bodyB.Velocity = phys.NewVector2(0, 0)
			bodyB.AngularVelocity = 0
			phys.SetPhysicsBodyRotation(bodyB, 330*rl.DEG2RAD)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawFPS(screenWidth-90, screenHeight-30)

		bodiesCount := phys.GetPhysicsBodiesCount()
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

		rl.DrawRectangle(0, screenHeight-49, screenWidth, 49, rl.Black)

		rl.DrawText("Friction amount", (screenWidth-rl.MeasureText("Friction amount", 30))/2, 75, 30, rl.White)
		rl.DrawText("0.1", int32(bodyA.Position.X)-rl.MeasureText("0.1", 20)/2, int32(bodyA.Position.Y)-7, 20, rl.White)
		rl.DrawText("1", int32(bodyB.Position.X)-rl.MeasureText("1", 20)/2, int32(bodyB.Position.Y)-7, 20, rl.White)

		rl.DrawText("Press 'R' to reset example", 10, 10, 10, rl.White)
		rl.DrawText("Physac", int32(logoX), int32(logoY), 30, rl.White)
		rl.DrawText("Powered by", int32(logoX)+50, int32(logoY)-7, 10, rl.White)

		rl.EndDrawing()
	}
}
