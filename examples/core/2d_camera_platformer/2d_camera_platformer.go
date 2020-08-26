package main

import (
	rl "goray/raylib"

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

		cameraT := camera.Convert()
		playerPositionT := player.Position.Convert()

		cameraT.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if cameraT.Zoom > 3.0 {
			cameraT.Zoom = 3.0
		} else if cameraT.Zoom < 0.25 {
			cameraT.Zoom = 0.25
		}

		if rl.IsKeyPressed(int32(rl.KEY_R)) {
			cameraT.Zoom = 1.0
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

		playerRect := rl.NewRectangle(playerPositionT.X-20, playerPositionT.Y-40, 40, 40)
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
	positionT := player.Position.Convert()

	if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
		positionT.X -= PLAYER_HOR_SPD * delta
	}
	if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
		positionT.X += PLAYER_HOR_SPD * delta
	}
	if rl.IsKeyDown(int32(rl.KEY_SPACE)) && player.CanJump {
		player.Speed = -PLAYER_JUMP_SPD
		player.CanJump = false
	}

	hitObstacle := false
	for _, ei := range envItems {
		rectT := ei.Rect.Convert()

		if ei.Blocking &&
			rectT.X <= positionT.X &&
			rectT.X+rectT.Width >= positionT.X &&
			rectT.Y >= positionT.Y &&
			rectT.Y < positionT.Y+player.Speed*delta {

			hitObstacle = true
			player.Speed = 0.0
			positionT.Y = rectT.Y
		}
	}

	if !hitObstacle {
		positionT.Y += player.Speed * delta
		player.Speed += float32(G) * delta
		player.CanJump = false
	} else {
		player.CanJump = true
	}
}

func UpdateCameraCenter(camera *rl.Camera2D, player *Player, width int32, height int32) {

	cameraT := camera.Convert()

	cameraT.Offset, _ = rl.NewVector2(float32(width/2), float32(height/2)).PassValue()
	cameraT.Target, _ = player.Position.PassValue()
}

func UpdateCameraCenterInsideMap(camera *rl.Camera2D, player *Player, envItems []*EnvItem, width int32, height int32) {

	cameraT := camera.Convert()
	cameraT.Offset, _ = rl.NewVector2(float32(width/2), float32(height/2)).PassValue()
	cameraOffsetT := camera.GetOffset().Convert()

	minX, minY, maxX, maxY := 1000.0, 1000.0, -1000.0, -1000.0

	for _, ei := range envItems {
		rectT := ei.Rect.Convert()
		minX = math.Min(float64(rectT.X), minX)
		maxX = math.Max(float64(rectT.X+rectT.Width), maxX)
		minY = math.Min(float64(rectT.Y), minY)
		maxY = math.Max(float64(rectT.Y+rectT.Height), maxY)
	}

	max := rl.GetWorldToScreen2D(rl.NewVector2(float32(maxX), float32(maxY)), *camera)
	maxT := max.Convert()

	min := rl.GetWorldToScreen2D(rl.NewVector2(float32(minX), float32(minY)), *camera)
	minT := min.Convert()

	if maxT.X < float32(width) {
		cameraOffsetT.X = float32(width - (int32(maxT.X) - width/2))
	}
	if maxT.Y < float32(height) {
		cameraOffsetT.Y = float32(height - (int32(maxT.Y) - height/2))
	}
	if minT.X > 0 {
		cameraOffsetT.X = float32(width/2) - minT.X
	}
	if minT.Y > 0 {
		cameraOffsetT.Y = float32(height/2) - minT.Y
	}
}

func UpdateCameraCenterSmoothFollow(camera *rl.Camera2D, player *Player, delta float32, width int32, height int32) {

	minSpeed := 30.0
	minEffectLength := 10.0
	fractionSpeed := 0.8

	camera.Convert().Offset, _ = rl.NewVector2(float32(width/2), float32(height/2)).PassValue()
	diff := rl.Vector2Subtract(player.Position, *camera.GetTarget())
	length := rl.Vector2Length(diff)

	if length > float32(minEffectLength) {
		speed := math.Max(fractionSpeed*float64(length), minSpeed)
		camera.Convert().Target, _ = rl.Vector2Add(*camera.GetTarget(), rl.Vector2Scale(diff, float32(speed)*delta/length)).PassValue()
	}
}

func UpdateCameraEvenOutOnLanding(camera *rl.Camera2D, player *Player, delta float32, width int32, height int32) {

	evenOutSpeed := float32(700)
	eveningOut := false
	evenOutTarget := float32(0)

	cameraT := camera.Convert()
	cameraTargetT := camera.GetTarget().Convert()
	playerPositionT := player.Position.Convert()

	cameraT.Offset, _ = rl.NewVector2(float32(width/2), float32(height/2)).PassValue()
	cameraTargetT.X = playerPositionT.X

	if eveningOut {
		if evenOutTarget > cameraTargetT.Y {
			cameraTargetT.Y += evenOutSpeed * delta

			if cameraTargetT.Y > evenOutTarget {
				cameraTargetT.Y = evenOutTarget
				eveningOut = false
			}
		} else {
			cameraTargetT.Y -= evenOutSpeed * delta

			if cameraTargetT.Y < evenOutTarget {
				cameraTargetT.Y = evenOutTarget
				eveningOut = false
			}
		}
	} else {

		if (player.CanJump && player.Speed == 0) && (playerPositionT.Y != cameraTargetT.Y) {
			eveningOut = true
			evenOutTarget = playerPositionT.Y
		}
	}
}

func UpdateCameraPlayerBoundsPush(camera *rl.Camera2D, player *Player, width int32, height int32) {
	bbox := rl.NewVector2(0.2, 0.2)
	bboxT := bbox.Convert()

	bboxWorldMin := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1.0-bboxT.X)*0.5*float32(width),
			(1-bboxT.Y)*0.5*float32(height),
		),
		*camera,
	)

	bboxWorldMax := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1.0-bboxT.X)*0.5*float32(width),
			(1-bboxT.Y)*0.5*float32(height),
		),
		*camera,
	)

	camera.Convert().Offset, _ = rl.NewVector2(
		(1.0-bboxT.X)*0.5*float32(width),
		(1-bboxT.Y)*0.5*float32(height),
	).PassValue()

	playerPositionT := player.Position.Convert()
	bboxWorldMinT := bboxWorldMin.Convert()
	bboxWorldMaxT := bboxWorldMax.Convert()
	cameraTargetT := camera.GetTarget().Convert()

	if playerPositionT.X < bboxWorldMinT.X {
		cameraTargetT.X = playerPositionT.X
	}
	if playerPositionT.Y < bboxWorldMinT.Y {
		cameraTargetT.Y = playerPositionT.Y
	}
	if playerPositionT.X > bboxWorldMaxT.Y {
		cameraTargetT.X = bboxWorldMinT.X + (playerPositionT.X - bboxWorldMaxT.X)
	}
	if playerPositionT.Y > bboxWorldMaxT.Y {
		cameraTargetT.Y = bboxWorldMinT.Y + (playerPositionT.Y - bboxWorldMaxT.Y)
	}
}
