package main

import (
	"fmt"
	rg "goray/raygui"
	rl "goray/raylib"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func DrawStyleEditControls() {

	rg.GuiGroupBox(rg.Rectangle{X: 550, Y: 170, Width: 220, Height: 205}, "SCROLLBAR STYLE")

	style := rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.BORDER_WIDTH))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 195, Width: 110, Height: 10}, "BORDER_WIDTH")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 190, Width: 90, Height: 20}, "", &style, 0, 6, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.BORDER_WIDTH), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_SIZE))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 220, Width: 110, Height: 10}, "ARROWS_SIZE")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 215, Width: 90, Height: 20}, "", &style, 4, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_SIZE), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 245, Width: 110, Height: 10}, "SLIDER_PADDING")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 240, Width: 90, Height: 20}, "", &style, 0, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING), style)

	style = int32(
		rg.Bool2int(
			rg.GuiCheckBox(
				rg.Rectangle{X: 565, Y: 280, Width: 20, Height: 20},
				"ARROWS_VISIBLE",
				rg.Int2bool(
					int(
						rg.GuiGetStyle(
							int32(rg.SCROLLBAR),
							int32(rg.ARROWS_VISIBLE),
						),
					),
				),
			),
		),
	)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_VISIBLE), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 325, Width: 110, Height: 10}, "SLIDER_PADDING")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 320, Width: 90, Height: 20}, "", &style, 0, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_WIDTH))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 350, Width: 110, Height: 10}, "SLIDER_WIDTH")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 345, Width: 90, Height: 20}, "", &style, 2, 100, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_WIDTH), style)

	var text string
	if rg.GuiGetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_SIDE)) == int32(rg.SCROLLBAR_LEFT_SIDE) {
		text = "SCROLLBAR: LEFT"
	} else {
		text = "SCROLLBAR: RIGHT"
	}
	style = int32(
		rg.Bool2int(
			rg.GuiToggle(
				rg.Rectangle{X: 560, Y: 110, Width: 200, Height: 35},
				text,
				rg.Int2bool(
					int(
						rg.GuiGetStyle(
							int32(rg.LISTVIEW),
							int32(rg.SCROLLBAR_SIDE),
						),
					),
				),
			),
		),
	)
	rg.GuiSetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_SIDE), style)

	rg.GuiGroupBox(rg.Rectangle{X: 550, Y: 20, Width: 220, Height: 135}, "SCROLLPANEL STYLE")

	style = rg.GuiGetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_WIDTH))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 35, Width: 110, Height: 10}, "SCROLLBAR_WIDTH")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 30, Width: 90, Height: 20}, "", &style, 6, 30, false)
	rg.GuiSetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_WIDTH), style)

	style = rg.GuiGetStyle(int32(rg.DEFAULT), int32(rg.BORDER_WIDTH))
	rg.GuiLabel(rg.Rectangle{X: 555, Y: 60, Width: 110, Height: 10}, "BORDER_WIDTH")
	rg.GuiSpinner(rg.Rectangle{X: 670, Y: 55, Width: 90, Height: 20}, "", &style, 0, 20, false)
	rg.GuiSetStyle(int32(rg.DEFAULT), int32(rg.BORDER_WIDTH), style)
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raygui - GuiScrollPanel()")

	panelRec := rg.Rectangle{X: 20, Y: 40, Width: 200, Height: 150}
	panelContentRec := rg.Rectangle{X: 0, Y: 0, Width: 340, Height: 340}
	panelScroll := rg.Vector2{X: 99, Y: -20}

	showContentArea := true

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.DrawText(fmt.Sprintf("[%f, %f]", panelScroll.Self().X, panelScroll.Self().Y), 4, 4, 20, *rl.Red)

		view := rg.GuiScrollPanel(panelRec, panelContentRec, &panelScroll)

		rl.BeginScissorMode(int32(view.Self().X), int32(view.Self().Y), int32(view.Self().Width), int32(view.Self().Height))
		rg.GuiGrid(
			rg.Rectangle{
				X: panelRec.Self().X + panelScroll.Self().X,
				Y: panelRec.Self().Y + panelScroll.Self().Y,
				Width: panelContentRec.Self().Width,
				Height: panelContentRec.Self().Height,
			}, 16, 3)
		rl.EndScissorMode()

		if showContentArea {
			rl.DrawRectangle(
				int32(panelRec.Self().X+panelScroll.Self().X),
				int32(panelRec.Self().Y+panelScroll.Self().Y),
				int32(panelContentRec.Self().Width),
				int32(panelContentRec.Self().Height),
				rl.Fade(*rl.Red, 0.1),
			)
		}

		DrawStyleEditControls()

		showContentArea = rg.GuiCheckBox(
			rg.Rectangle{X: 565, Y: 80, Width: 20, Height: 20},
			"SHOW CONTENT AREA",
			showContentArea,
		)

		panelContentRec.Self().Width = rg.GuiSliderBar(
			rg.Rectangle{X: 590, Y: 385, Width: 145, Height: 15},
			"WIDTH",
			fmt.Sprintf("%f", panelContentRec.Self().Width),
			1, 600, 1,
		)

		panelContentRec.Self().Height = rg.GuiSliderBar(
			rg.Rectangle{X: 590, Y: 410, Width: 145, Height: 15},
			"HEIGHT",
			fmt.Sprintf("%f", panelContentRec.Self().Height),
			1, 400, 1,
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
