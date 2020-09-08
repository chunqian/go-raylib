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

    rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - scissor test")
    defer rl.CloseWindow()

    scissorArea := rl.NewRectangle(0, 0, 300, 300)
    scissorMode := true

    rl.SetTargetFPS(60)

    for !rl.WindowShouldClose() {

        if rl.IsKeyPressed(int32(rl.KEY_S)) {
            scissorMode = !scissorMode
        }

        scissorArea.X = float32(rl.GetMouseX()) - scissorArea.Width/2
        scissorArea.Y = float32(rl.GetMouseY()) - scissorArea.Height/2

        rl.BeginDrawing()

        rl.ClearBackground(rl.RayWhite)

        if scissorMode {
            rl.BeginScissorMode(
                int32(scissorArea.X),
                int32(scissorArea.Y),
                int32(scissorArea.Width),
                int32(scissorArea.Height),
            )
        }

        rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Red)
        rl.DrawText("Move the mouse around to reveal this text!", 190, 200, 20, rl.LightGray)

        if scissorMode {
            rl.EndScissorMode()
        }

        rl.DrawRectangleLinesEx(scissorArea, 1, rl.Black)
        rl.DrawText("Press S to toggle scissor test", 10, 10, 20, rl.Black)

        rl.EndDrawing()
    }
}
