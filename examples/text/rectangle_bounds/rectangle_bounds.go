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

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - draw text inside a rectangle")
	defer rl.CloseWindow()

	text := "Text cannot escape\tthis container\t...word wrap also works when active so here's" +
		"a long text for testing.\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
		"tempor incididunt ut labore et dolore magna aliqua. Nec ullamcorper sit amet risus nullam eget felis eget."

	resizing := false
	wordWrap := true

	container := rl.NewRectangle(25, 25, float32(screenWidth-50), float32(screenHeight-250))
	resizer := rl.NewRectangle(container.X+container.Width-17, container.Y+container.Height-17, 14, 14)

	minWidth := int32(60)
	minHeight := int32(60)
	maxWidth := screenWidth - 50
	maxHeight := screenHeight - 160

	lastMouse := rl.NewVector2(0, 0)
	borderColor := rl.Maroon

	font := rl.GetFontDefault()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			wordWrap = !wordWrap
		}
		mouse := rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mouse, container) {
			borderColor = rl.Fade(rl.Maroon, 0.4)
		} else if !resizing {
			borderColor = rl.Maroon
		}

		if resizing {
			if rl.IsMouseButtonReleased(int32(rl.MOUSE_LEFT_BUTTON)) {
				resizing = false
			}

			width := int32(container.Width + (mouse.X - lastMouse.X))
			if width > maxWidth {
				width = maxWidth
			}
			if width < minWidth {
				width = minWidth
			}
			container.Width = float32(width)

			height := int32(container.Height + (mouse.Y - lastMouse.Y))
			if height > maxHeight {
				height = maxHeight
			}
			if height < minHeight {
				height = minHeight
			}
			container.Height = float32(height)
		} else {
			if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) && rl.CheckCollisionPointRec(mouse, resizer) {
				resizing = true
			}
		}

		resizer.X = container.X + container.Width - 17
		resizer.Y = container.Y + container.Height - 17

		lastMouse = mouse

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleLinesEx(container, 3, borderColor)

		rl.DrawTextRec(
			font,
			text,
			rl.NewRectangle(
				container.X+4,
				container.Y+4,
				container.Width-4,
				container.Height-4,
			),
			20,
			2,
			wordWrap,
			rl.Gray,
		)
		rl.DrawRectangleRec(resizer, borderColor)

		rl.DrawRectangle(0, screenHeight-54, screenWidth, 54, rl.Gray)
		rl.DrawRectangleRec(rl.NewRectangle(382, float32(screenHeight-34), 12, 12), rl.Maroon)

		rl.DrawText("Word Wrap: ", 313, screenHeight-115, 20, rl.Black)
		if wordWrap {
			rl.DrawText("ON", 447, screenHeight-115, 20, rl.Red)
		} else {
			rl.DrawText("OFF", 447, screenHeight-115, 20, rl.Black)
		}

		rl.DrawText("Press [SPACE] to toggle word wrap", 218, screenHeight-86, 20, rl.Gray)

		rl.DrawText("Click hold & drag the    to resize the container", 155, screenHeight-38, 20, rl.RayWhite)

		rl.EndDrawing()
	}
}
