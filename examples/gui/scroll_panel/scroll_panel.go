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

	style = int32(rg.Bool2int(
		rg.GuiCheckBox(
			rg.NewRectangle(565, 280, 20, 20),
			"ARROWS_VISIBLE",
			rg.Int2bool(int(
				rg.GuiGetStyle(
					int32(rg.SCROLLBAR),
					int32(rg.ARROWS_VISIBLE),
				),
			)),
		)),
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
	style = int32(rg.Bool2int(
		rg.GuiToggle(
			rg.NewRectangle(560, 110, 200, 35),
			text,
			rg.Int2bool(int(
				rg.GuiGetStyle(
					int32(rg.LISTVIEW),
					int32(rg.SCROLLBAR_SIDE),
				),
			)),
		)),
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

	panelRec := rg.NewRectangle(20, 40, 200, 150)
	panelRecT := panelRec.Convert()

	panelContentRec := rg.NewRectangle(0, 0, 340, 340)
	panelContentRecT := panelContentRec.Convert()

	panelScroll := rg.NewVector2(99, -20)
	panelScrollT := panelScroll.Convert()

	showContentArea := true

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.DrawText(fmt.Sprintf("[%f, %f]", panelScrollT.X, panelScrollT.Y), 4, 4, 20, *rl.Red)

		view := rg.GuiScrollPanel(panelRec, panelContentRec, &panelScroll)
		viewT := view.Convert()

		rl.BeginScissorMode(int32(viewT.X), int32(viewT.Y), int32(viewT.Width), int32(viewT.Height))

		rg.GuiGrid(
			rg.NewRectangle(
				panelRecT.X+panelScrollT.X,
				panelRecT.Y+panelScrollT.Y,
				panelContentRecT.Width,
				panelContentRecT.Height,
			), 16, 3)

		rl.EndScissorMode()

		if showContentArea {
			rl.DrawRectangle(
				int32(panelRecT.X+panelScrollT.X),
				int32(panelRecT.Y+panelScrollT.Y),
				int32(panelContentRecT.Width),
				int32(panelContentRecT.Height),
				rl.Fade(*rl.Red, 0.1),
			)
		}

		DrawStyleEditControls()

		showContentArea = rg.GuiCheckBox(
			rg.NewRectangle(565, 80, 20, 20),
			"SHOW CONTENT AREA",
			showContentArea,
		)

		panelContentRecT.Width = rg.GuiSliderBar(
			rg.NewRectangle(590, 385, 145, 15),
			"WIDTH",
			fmt.Sprintf("%f", panelContentRecT.Width),
			1, 600, 1,
		)

		panelContentRecT.Height = rg.GuiSliderBar(
			rg.NewRectangle(590, 410, 145, 15),
			"HEIGHT",
			fmt.Sprintf("%f", panelContentRecT.Height),
			1, 400, 1,
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
