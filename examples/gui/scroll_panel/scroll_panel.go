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

	rg.GuiGroupBox(rg.NewRectangle(550, 170, 220, 205), "SCROLLBAR STYLE")

	style := rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.BORDER_WIDTH))
	rg.GuiLabel(rg.NewRectangle(555, 195, 110, 10), "BORDER_WIDTH")
	rg.GuiSpinner(rg.NewRectangle(670, 190, 90, 20), "", &style, 0, 6, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.BORDER_WIDTH), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_SIZE))
	rg.GuiLabel(rg.NewRectangle(555, 220, 110, 10), "ARROWS_SIZE")
	rg.GuiSpinner(rg.NewRectangle(670, 215, 90, 20), "", &style, 4, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_SIZE), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING))
	rg.GuiLabel(rg.NewRectangle(555, 245, 110, 10), "SLIDER_PADDING")
	rg.GuiSpinner(rg.NewRectangle(670, 240, 90, 20), "", &style, 0, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING), style)

	style = rg.ToInt32(
		rg.GuiCheckBox(
			rg.NewRectangle(565, 280, 20, 20),
			"ARROWS_VISIBLE",
			rg.ToBool(
				rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_VISIBLE)),
			),
		),
	)

	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.ARROWS_VISIBLE), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING))
	rg.GuiLabel(rg.NewRectangle(555, 325, 110, 10), "SLIDER_PADDING")
	rg.GuiSpinner(rg.NewRectangle(670, 320, 90, 20), "", &style, 0, 14, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_PADDING), style)

	style = rg.GuiGetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_WIDTH))
	rg.GuiLabel(rg.NewRectangle(555, 350, 110, 10), "SLIDER_WIDTH")
	rg.GuiSpinner(rg.NewRectangle(670, 345, 90, 20), "", &style, 2, 100, false)
	rg.GuiSetStyle(int32(rg.SCROLLBAR), int32(rg.SLIDER_WIDTH), style)

	var text string
	if rg.GuiGetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_SIDE)) == int32(rg.SCROLLBAR_LEFT_SIDE) {
		text = "SCROLLBAR: LEFT"
	} else {
		text = "SCROLLBAR: RIGHT"
	}
	style = rg.ToInt32(
		rg.GuiToggle(
			rg.NewRectangle(560, 110, 200, 35),
			text,
			rg.ToBool(
				rg.GuiGetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_SIDE)),
			),
		),
	)

	rg.GuiSetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_SIDE), style)

	rg.GuiGroupBox(rg.NewRectangle(550, 20, 220, 135), "SCROLLPANEL STYLE")

	style = rg.GuiGetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_WIDTH))
	rg.GuiLabel(rg.NewRectangle(555, 35, 110, 10), "SCROLLBAR_WIDTH")
	rg.GuiSpinner(rg.NewRectangle(670, 30, 90, 20), "", &style, 6, 30, false)
	rg.GuiSetStyle(int32(rg.LISTVIEW), int32(rg.SCROLLBAR_WIDTH), style)

	style = rg.GuiGetStyle(int32(rg.DEFAULT), int32(rg.BORDER_WIDTH))
	rg.GuiLabel(rg.NewRectangle(555, 60, 110, 10), "BORDER_WIDTH")
	rg.GuiSpinner(rg.NewRectangle(670, 55, 90, 20), "", &style, 0, 20, false)
	rg.GuiSetStyle(int32(rg.DEFAULT), int32(rg.BORDER_WIDTH), style)
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raygui - GuiScrollPanel()")
	defer rl.CloseWindow()

	panelRec := rg.NewRectangle(20, 40, 200, 150)

	panelContentRec := rg.NewRectangle(0, 0, 340, 340)

	panelScroll := rg.NewVector2(99, -20)

	showContentArea := true

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("[%f, %f]", panelScroll.X, panelScroll.Y), 4, 4, 20, rl.Red)

		view := rg.GuiScrollPanel(panelRec, panelContentRec, &panelScroll)

		rl.BeginScissorMode(int32(view.X), int32(view.Y), int32(view.Width), int32(view.Height))

		rg.GuiGrid(
			rg.NewRectangle(
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

		showContentArea = rg.GuiCheckBox(
			rg.NewRectangle(565, 80, 20, 20),
			"SHOW CONTENT AREA",
			showContentArea,
		)

		panelContentRec.Width = rg.GuiSliderBar(
			rg.NewRectangle(590, 385, 145, 15),
			"WIDTH",
			fmt.Sprintf("%f", panelContentRec.Width),
			1, 600, 1,
		)

		panelContentRec.Height = rg.GuiSliderBar(
			rg.NewRectangle(590, 410, 145, 15),
			"HEIGHT",
			fmt.Sprintf("%f", panelContentRec.Height),
			1, 400, 1,
		)

		rl.EndDrawing()
	}
}
