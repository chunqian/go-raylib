package main

import (
	"fmt"
	rl "github.com/chunqian/go-raylib/raylib"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func DrawStyleEditControls() {

	rl.GuiGroupBox(rl.NewRectangle(550, 170, 220, 205), "SCROLLBAR STYLE")

	style := rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.BORDER_WIDTH))
	rl.GuiLabel(rl.NewRectangle(555, 195, 110, 10), "BORDER_WIDTH")
	rl.GuiSpinner(rl.NewRectangle(670, 190, 90, 20), "", &style, 0, 6, false)
	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.BORDER_WIDTH), style)

	style = rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.ARROWS_SIZE))
	rl.GuiLabel(rl.NewRectangle(555, 220, 110, 10), "ARROWS_SIZE")
	rl.GuiSpinner(rl.NewRectangle(670, 215, 90, 20), "", &style, 4, 14, false)
	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.ARROWS_SIZE), style)

	style = rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_PADDING))
	rl.GuiLabel(rl.NewRectangle(555, 245, 110, 10), "SLIDER_PADDING")
	rl.GuiSpinner(rl.NewRectangle(670, 240, 90, 20), "", &style, 0, 14, false)
	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_PADDING), style)

	style = rl.ToInt32(
		rl.GuiCheckBox(
			rl.NewRectangle(565, 280, 20, 20),
			"ARROWS_VISIBLE",
			rl.ToBool(
				rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.ARROWS_VISIBLE)),
			),
		),
	)

	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.ARROWS_VISIBLE), style)

	style = rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_PADDING))
	rl.GuiLabel(rl.NewRectangle(555, 325, 110, 10), "SLIDER_PADDING")
	rl.GuiSpinner(rl.NewRectangle(670, 320, 90, 20), "", &style, 0, 14, false)
	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_PADDING), style)

	style = rl.GuiGetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_WIDTH))
	rl.GuiLabel(rl.NewRectangle(555, 350, 110, 10), "SLIDER_WIDTH")
	rl.GuiSpinner(rl.NewRectangle(670, 345, 90, 20), "", &style, 2, 100, false)
	rl.GuiSetStyle(int32(rl.SCROLLBAR), int32(rl.SLIDER_WIDTH), style)

	var text string
	if rl.GuiGetStyle(int32(rl.LISTVIEW), int32(rl.SCROLLBAR_SIDE)) == int32(rl.SCROLLBAR_LEFT_SIDE) {
		text = "SCROLLBAR: LEFT"
	} else {
		text = "SCROLLBAR: RIGHT"
	}
	style = rl.ToInt32(
		rl.GuiToggle(
			rl.NewRectangle(560, 110, 200, 35),
			text,
			rl.ToBool(
				rl.GuiGetStyle(int32(rl.LISTVIEW), int32(rl.SCROLLBAR_SIDE)),
			),
		),
	)

	rl.GuiSetStyle(int32(rl.LISTVIEW), int32(rl.SCROLLBAR_SIDE), style)

	rl.GuiGroupBox(rl.NewRectangle(550, 20, 220, 135), "SCROLLPANEL STYLE")

	style = rl.GuiGetStyle(int32(rl.LISTVIEW), int32(rl.SCROLLBAR_WIDTH))
	rl.GuiLabel(rl.NewRectangle(555, 35, 110, 10), "SCROLLBAR_WIDTH")
	rl.GuiSpinner(rl.NewRectangle(670, 30, 90, 20), "", &style, 6, 30, false)
	rl.GuiSetStyle(int32(rl.LISTVIEW), int32(rl.SCROLLBAR_WIDTH), style)

	style = rl.GuiGetStyle(int32(rl.DEFAULT), int32(rl.BORDER_WIDTH))
	rl.GuiLabel(rl.NewRectangle(555, 60, 110, 10), "BORDER_WIDTH")
	rl.GuiSpinner(rl.NewRectangle(670, 55, 90, 20), "", &style, 0, 20, false)
	rl.GuiSetStyle(int32(rl.DEFAULT), int32(rl.BORDER_WIDTH), style)
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raygui - GuiScrollPanel()")
	defer rl.CloseWindow()

	panelRec := rl.NewRectangle(20, 40, 200, 150)

	panelContentRec := rl.NewRectangle(0, 0, 340, 340)

	panelScroll := rl.NewVector2(99, -20)

	showContentArea := true

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("[%f, %f]", panelScroll.X, panelScroll.Y), 4, 4, 20, rl.Red)

		view := rl.GuiScrollPanel(panelRec, panelContentRec, &panelScroll)

		rl.BeginScissorMode(int32(view.X), int32(view.Y), int32(view.Width), int32(view.Height))

		rl.GuiGrid(
			rl.NewRectangle(
				panelRec.X+panelScroll.X,
				panelRec.Y+panelScroll.Y,
				panelContentRec.Width,
				panelContentRec.Height,
			), 16, 3)

		rl.EndScissorMode()

		if showContentArea {
			rl.DrawRectangle(
				int32(panelRec.X+panelScroll.X),
				int32(panelRec.Y+panelScroll.Y),
				int32(panelContentRec.Width),
				int32(panelContentRec.Height),
				rl.Fade(rl.Red, 0.1),
			)
		}

		DrawStyleEditControls()

		showContentArea = rl.GuiCheckBox(
			rl.NewRectangle(565, 80, 20, 20),
			"SHOW CONTENT AREA",
			showContentArea,
		)

		panelContentRec.Width = rl.GuiSliderBar(
			rl.NewRectangle(590, 385, 145, 15),
			"WIDTH",
			fmt.Sprintf("%f", panelContentRec.Width),
			1, 600, 1,
		)

		panelContentRec.Height = rl.GuiSliderBar(
			rl.NewRectangle(590, 410, 145, 15),
			"HEIGHT",
			fmt.Sprintf("%f", panelContentRec.Height),
			1, 400, 1,
		)

		rl.EndDrawing()
	}
}
