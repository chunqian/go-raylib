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
	cameraT.Target = *player.Position.Ref()
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

func UpdateCameraCenterSmoothFollow(camera *rl.Camera2D, player *Player, width int32, height int32) {

	// minSpeed := 30.0
	// minEffectLength := 10.0
	// fractionSpeed := 0.8
}

func UpdateCameraEvenOutOnLanding(camera *rl.Camera2D, player *Player, delta float32, width int32, height int32) {

	// evenOutSpeed := float32(700)
	// eveningOut := false
	// evenOutTarget := float32(0)
}
