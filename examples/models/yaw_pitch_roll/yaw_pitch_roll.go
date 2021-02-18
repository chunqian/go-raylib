package main

import (
	rl "go-raylib/raylib"

	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - plane rotations (yaw, pitch, roll)")
	defer rl.CloseWindow()

	texAngleGauge := rl.LoadTexture("../models/resources/angle_gauge.png")
	texBackground := rl.LoadTexture("../models/resources/background.png")
	texPitch := rl.LoadTexture("../models/resources/pitch.png")
	texPlane := rl.LoadTexture("../models/resources/plane.png")

	framebuffer := rl.LoadRenderTexture(192, 192)

	model := rl.LoadModel("../models/resources/plane.obj")
	defer func() {
		rl.UnloadTexture(rl.Texture2D(model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture))
		rl.UnloadModel(model)

		rl.UnloadRenderTexture(framebuffer)

		rl.UnloadTexture(texAngleGauge)
		rl.UnloadTexture(texBackground)
		rl.UnloadTexture(texPitch)
		rl.UnloadTexture(texPlane)
	}()
	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = rl.Texture(rl.LoadTexture("../models/resources/plane_diffuse.png"))

	rl.GenTextureMipmaps((*rl.Texture2D)(&model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture))

	camera := rl.NewCamera(
		rl.NewVector3(0, 60.0, -120.0),
		rl.NewVector3(0, 12.0, 0),
		rl.NewVector3(0, 1.0, 0),
		30.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	pitch := float32(0)
	roll := float32(0)
	yaw := float32(0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			roll += 1.0
		} else if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			roll -= 1.0
		} else {
			if roll > 0 {
				roll -= 0.5
			} else if roll < 0 {
				roll += 0.5
			}
		}

		if rl.IsKeyDown(int32(rl.KEY_S)) {
			yaw += 1.0
		} else if rl.IsKeyDown(int32(rl.KEY_A)) {
			yaw -= 1.0
		} else {
			if yaw > 0 {
				yaw -= 0.5
			} else if yaw < 0 {
				yaw += 0.5
			}
		}

		if rl.IsKeyDown(int32(rl.KEY_DOWN)) {
			pitch += 0.6
		} else if rl.IsKeyDown(int32(rl.KEY_UP)) {
			pitch -= 0.6
		} else {
			if pitch > 0.3 {
				pitch -= 0.3
			} else if pitch < -0.3 {
				pitch += 0.3
			}
		}

		pitchOffset := int32(pitch)
		for pitchOffset > 180 {
			pitchOffset -= 360
		}
		for pitchOffset < -180 {
			pitchOffset += 360
		}
		pitchOffset *= 10

		model.Transform = rl.MatrixRotateXYZ(
			rl.NewVector3(
				rl.DEG2RAD*pitch,
				rl.DEG2RAD*yaw,
				rl.DEG2RAD*roll,
			),
		)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		centerX := framebuffer.Texture.Width / 2
		centerY := framebuffer.Texture.Height / 2
		scaleFactor := float32(0.5)

		rl.BeginTextureMode(framebuffer)

		rl.BeginBlendMode(int32(rl.BLEND_ALPHA))

		rl.DrawTexturePro(
			texBackground,
			rl.NewRectangle(0, 0, float32(texBackground.Width), float32(texBackground.Height)),
			rl.NewRectangle(float32(centerX), float32(centerY), float32(texBackground.Width)*scaleFactor, float32(texBackground.Height)*scaleFactor),
			rl.NewVector2(float32(texBackground.Width)/2*scaleFactor, float32(texBackground.Height)/2*scaleFactor+float32(pitchOffset)*scaleFactor),
			roll,
			rl.White,
		)

		rl.DrawTexturePro(
			texPitch,
			rl.NewRectangle(0, 0, float32(texPitch.Width), float32(texPitch.Height)),
			rl.NewRectangle(float32(centerX), float32(centerY), float32(texPitch.Width)*scaleFactor, float32(texPitch.Height)*scaleFactor),
			rl.NewVector2(float32(texPitch.Width)/2*scaleFactor, float32(texBackground.Height)/2*scaleFactor+float32(pitchOffset)*scaleFactor),
			roll,
			rl.White,
		)

		rl.DrawTexturePro(
			texPlane,
			rl.NewRectangle(0, 0, float32(texPlane.Width), float32(texPlane.Height)),
			rl.NewRectangle(float32(centerX), float32(centerY), float32(texPlane.Width)*scaleFactor, float32(texPlane.Height)*scaleFactor),
			rl.NewVector2(float32(texPlane.Width)/2*scaleFactor, float32(texPlane.Height)/2*scaleFactor+float32(pitchOffset)*scaleFactor),
			0,
			rl.White,
		)

		rl.EndBlendMode()

		rl.EndTextureMode()

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.NewVector3(0, 6.0, 0), 1.0, rl.White)

		rl.DrawGrid(10, 10.0)

		rl.EndMode3D()

		DrawAngleGauge(texAngleGauge, 80, 70, roll, "roll", rl.Red)
		DrawAngleGauge(texAngleGauge, 190, 70, pitch, "pitch", rl.Green)
		DrawAngleGauge(texAngleGauge, 300, 70, yaw, "yaw", rl.SkyBlue)

		rl.DrawRectangle(30, 360, 260, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(30, 360, 260, 70, rl.Fade(rl.DarkBlue, 0.5))
		rl.DrawText("Pitch controlled with: KEY_UP / KEY_DOWN", 40, 370, 10, rl.DarkGray)
		rl.DrawText("Roll controlled with: KEY_LEFT / KEY_RIGHT", 40, 390, 10, rl.DarkGray)
		rl.DrawText("Yaw controlled with: KEY_A / KEY_S", 40, 410, 10, rl.DarkGray)

		rl.DrawTextureRec(
			rl.Texture2D(framebuffer.Texture),
			rl.NewRectangle(0, 0, float32(framebuffer.Texture.Width), -float32(framebuffer.Texture.Height)),
			rl.NewVector2(float32(screenWidth-framebuffer.Texture.Width-20), 20),
			rl.Fade(rl.White, 0.8),
		)
		rl.DrawRectangleLines(
			screenWidth-framebuffer.Texture.Width-20,
			20,
			framebuffer.Texture.Width,
			framebuffer.Texture.Height,
			rl.DarkGray,
		)

		rl.EndDrawing()
	}
}

func DrawAngleGauge(angleGauge rl.Texture2D, x float32, y float32, angle float32, title string, color rl.Color) {

	srcRec := rl.NewRectangle(0, 0, float32(angleGauge.Width), float32(angleGauge.Height))
	dstRec := rl.NewRectangle(x, y, float32(angleGauge.Width), float32(angleGauge.Height))
	origin := rl.NewVector2(float32(angleGauge.Width)/2, float32(angleGauge.Height)/2)
	textSize := int32(20)

	rl.DrawTexturePro(angleGauge, srcRec, dstRec, origin, angle, color)

	rl.DrawText(
		fmt.Sprintf("%5.1f", angle),
		int32(x)-rl.MeasureText(fmt.Sprintf("%5.1f", angle), textSize)/2,
		int32(y)+10,
		textSize,
		rl.DarkGray,
	)
	rl.DrawText(title, int32(x)-rl.MeasureText(title, textSize)/2, int32(y)+60, textSize, rl.DarkGray)
}
