package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"math"
	"runtime"
)

var (
	G               = 400
	PLAYER_JUMP_SPD = float32(350.0)
	PLAYER_HOR_SPD  = float32(200.0)
)

type Player struct {
	Position rl.Vector2
	Speed    float32
	CanJump  bool
}

type EnvItem struct {
	Rect     rl.Rectangle
	Blocking bool
	Color    rl.Color
}

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 2d camera")
	defer rl.CloseWindow()

	player := Player{
		Position: rl.NewVector2(400, 280),
		Speed:    0,
		CanJump:  false,
	}

	envItems := []*EnvItem{
		{
			Rect:     rl.NewRectangle(0, 0, 1000, 400),
			Blocking: false,
			Color:    rl.LightGray,
		},
		{
			Rect:     rl.NewRectangle(0, 400, 1000, 200),
			Blocking: true,
			Color:    rl.Gray,
		},
		{
			Rect:     rl.NewRectangle(300, 200, 400, 10),
			Blocking: true,
			Color:    rl.Gray,
		},
		{
			Rect:     rl.NewRectangle(250, 300, 100, 10),
			Blocking: true,
			Color:    rl.Gray,
		},
		{
			Rect:     rl.NewRectangle(650, 300, 100, 10),
			Blocking: true,
			Color:    rl.Gray,
		},
	}

	camera := rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		player.Position,
		0,
		1.0,
	)

	cameraOption := 0
	cameraDescriptions := []string{
		"Follow player center",
		"Follow player center, but clamp to map edges",
		"Follow player center; smoothed",
		"Follow player center horizontally; updateplayer center vertically after landing",
		"Player push camera on getting too close to screen edge",
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		deltaTime := rl.GetFrameTime()

		UpdatePlayer(&player, envItems, deltaTime)

		camera.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if camera.Zoom > 3.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom < 0.25 {
			camera.Zoom = 0.25
		}

		if rl.IsKeyPressed(int32(rl.KEY_R)) {
			camera.Zoom = 1.0
			player.Position = rl.NewVector2(400, 280)
		}

		if rl.IsKeyPressed(int32(rl.KEY_C)) {
			cameraOption = (cameraOption + 1) % 5
		}
		switch {
		case cameraOption == 0:
			UpdateCameraCenter(&camera, &player, screenWidth, screenHeight)
		case cameraOption == 1:
			UpdateCameraCenterInsideMap(&camera, &player, envItems, screenWidth, screenHeight)
		case cameraOption == 2:
			UpdateCameraCenterSmoothFollow(&camera, &player, deltaTime, screenWidth, screenHeight)
		case cameraOption == 3:
			UpdateCameraEvenOutOnLanding(&camera, &player, deltaTime, screenWidth, screenHeight)
		case cameraOption == 4:
			UpdateCameraPlayerBoundsPush(&camera, &player, screenWidth, screenHeight)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.LightGray)

		rl.BeginMode2D(camera)

		for _, item := range envItems {
			rl.DrawRectangleRec(item.Rect, item.Color)
		}

		playerRect := rl.NewRectangle(player.Position.X-20, player.Position.Y-40, 40, 40)
		rl.DrawRectangleRec(playerRect, rl.Red)

		rl.EndMode2D()

		rl.DrawText("Controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Right/Left to move", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Space to jump", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel to Zoom in-out, R to reset zoom", 40, 80, 10, rl.DarkGray)
		rl.DrawText("- C to change camera mode", 40, 100, 10, rl.DarkGray)
		rl.DrawText("Current camera mode:", 20, 120, 10, rl.Black)
		rl.DrawText(cameraDescriptions[cameraOption], 40, 140, 10, rl.DarkGray)

		rl.EndDrawing()
	}
}

func UpdatePlayer(player *Player, envItems []*EnvItem, delta float32) {

	if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
		player.Position.X -= PLAYER_HOR_SPD * delta
	}
	if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
		player.Position.X += PLAYER_HOR_SPD * delta
	}
	if rl.IsKeyDown(int32(rl.KEY_SPACE)) && player.CanJump {
		player.Speed = -PLAYER_JUMP_SPD
		player.CanJump = false
	}

	hitObstacle := false
	for _, ei := range envItems {

		if ei.Blocking &&
			ei.Rect.X <= player.Position.X &&
			ei.Rect.X+ei.Rect.Width >= player.Position.X &&
			ei.Rect.Y >= player.Position.Y &&
			ei.Rect.Y < player.Position.Y+player.Speed*delta {

			hitObstacle = true
			player.Speed = 0.0
			player.Position.Y = ei.Rect.Y
		}
	}

	if !hitObstacle {
		player.Position.Y += player.Speed * delta
		player.Speed += float32(G) * delta
		player.CanJump = false
	} else {
		player.CanJump = true
	}
}

func UpdateCameraCenter(camera *rl.Camera2D, player *Player, width int32, height int32) {

	camera.Offset = rl.NewVector2(float32(width/2), float32(height/2))
	camera.Target = player.Position
}

func UpdateCameraCenterInsideMap(camera *rl.Camera2D, player *Player, envItems []*EnvItem, width int32, height int32) {

	camera.Offset = rl.NewVector2(float32(width/2), float32(height/2))

	minX, minY, maxX, maxY := 1000.0, 1000.0, -1000.0, -1000.0

	for _, ei := range envItems {

		minX = math.Min(float64(ei.Rect.X), minX)
		maxX = math.Max(float64(ei.Rect.X+ei.Rect.Width), maxX)
		minY = math.Min(float64(ei.Rect.Y), minY)
		maxY = math.Max(float64(ei.Rect.Y+ei.Rect.Height), maxY)
	}

	max := rl.GetWorldToScreen2D(rl.NewVector2(float32(maxX), float32(maxY)), *camera)

	min := rl.GetWorldToScreen2D(rl.NewVector2(float32(minX), float32(minY)), *camera)

	if max.X < float32(width) {
		camera.Offset.X = float32(width - (int32(max.X) - width/2))
	}
	if max.Y < float32(height) {
		camera.Offset.Y = float32(height - (int32(max.Y) - height/2))
	}
	if min.X > 0 {
		camera.Offset.X = float32(width/2) - min.X
	}
	if min.Y > 0 {
		camera.Offset.Y = float32(height/2) - min.Y
	}
}

func UpdateCameraCenterSmoothFollow(camera *rl.Camera2D, player *Player, delta float32, width int32, height int32) {

	minSpeed := 30.0
	minEffectLength := 10.0
	fractionSpeed := 0.8

	camera.Offset = rl.NewVector2(float32(width/2), float32(height/2))
	diff := rl.Vector2Subtract(player.Position, camera.Target)
	length := rl.Vector2Length(diff)

	if length > float32(minEffectLength) {
		speed := math.Max(fractionSpeed*float64(length), minSpeed)
		camera.Target = rl.Vector2Add(camera.Target, rl.Vector2Scale(diff, float32(speed)*delta/length))
	}
}

func UpdateCameraEvenOutOnLanding(camera *rl.Camera2D, player *Player, delta float32, width int32, height int32) {

	evenOutSpeed := float32(700)
	eveningOut := false
	evenOutTarget := float32(0)

	camera.Offset = rl.NewVector2(float32(width/2), float32(height/2))
	camera.Target.X = player.Position.X

	if eveningOut {
		if evenOutTarget > camera.Target.Y {
			camera.Target.Y += evenOutSpeed * delta

			if camera.Target.Y > evenOutTarget {
				camera.Target.Y = evenOutTarget
				eveningOut = false
			}
		} else {
			camera.Target.Y -= evenOutSpeed * delta

			if camera.Target.Y < evenOutTarget {
				camera.Target.Y = evenOutTarget
				eveningOut = false
			}
		}
	} else {

		if (player.CanJump && player.Speed == 0) && (player.Position.Y != camera.Target.Y) {
			eveningOut = true
			evenOutTarget = player.Position.Y
		}
	}
}

func UpdateCameraPlayerBoundsPush(camera *rl.Camera2D, player *Player, width int32, height int32) {
	bbox := rl.NewVector2(0.2, 0.2)

	bboxWorldMin := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1.0-bbox.X)*0.5*float32(width),
			(1-bbox.Y)*0.5*float32(height),
		),
		*camera,
	)

	bboxWorldMax := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1.0-bbox.X)*0.5*float32(width),
			(1-bbox.Y)*0.5*float32(height),
		),
		*camera,
	)

	camera.Offset = rl.NewVector2(
		(1.0-bbox.X)*0.5*float32(width),
		(1-bbox.Y)*0.5*float32(height),
	)

	if player.Position.X < bboxWorldMin.X {
		camera.Target.X = player.Position.X
	}
	if player.Position.Y < bboxWorldMin.Y {
		camera.Target.Y = player.Position.Y
	}
	if player.Position.X > bboxWorldMax.Y {
		camera.Target.X = bboxWorldMin.X + (player.Position.X - bboxWorldMax.X)
	}
	if player.Position.Y > bboxWorldMax.Y {
		camera.Target.Y = bboxWorldMin.Y + (player.Position.Y - bboxWorldMax.Y)
	}
}
