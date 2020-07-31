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

		rl.DrawText(fmt.Sprintf("[%f, %f]", *panelScroll.GetX(), *panelScroll.GetY()), 4, 4, 20, *rl.Red)

		view := rg.GuiScrollPanel(panelRec, panelContentRec, &panelScroll)

		rl.BeginScissorMode(int32(*view.GetX()), int32(*view.GetY()), int32(*view.GetWidth()), int32(*view.GetHeight()))
		rg.GuiGrid(
			rg.Rectangle{
				X:      *panelRec.GetX() + *panelScroll.GetX(),
				Y:      *panelRec.GetY() + *panelScroll.GetY(),
				Width:  *panelContentRec.GetWidth(),
				Height: *panelContentRec.GetHeight(),
			}, 16, 3)
		rl.EndScissorMode()

		if showContentArea {
			rl.DrawRectangle(
				int32(*panelRec.GetX())+int32(*panelScroll.GetX()),
				int32(*panelRec.GetY())+int32(*panelScroll.GetY()),
				int32(*panelContentRec.GetWidth()),
				int32(*panelContentRec.GetHeight()),
				rl.Fade(*rl.Red, 0.1),
			)
		}

		DrawStyleEditControls()

		showContentArea = rg.GuiCheckBox(
			rg.Rectangle{X: 565, Y: 80, Width: 20, Height: 20},
			"SHOW CONTENT AREA",
			showContentArea,
		)

		panelContentRec.SetWidth(
			rg.GuiSliderBar(
				rg.Rectangle{X: 590, Y: 385, Width: 145, Height: 15},
				"WIDTH",
				fmt.Sprintf("%f", *panelContentRec.GetWidth()),
				1, 600, 1,
			),
		)
		panelContentRec.SetHeight(
			rg.GuiSliderBar(
				rg.Rectangle{X: 590, Y: 410, Width: 145, Height: 15},
				"HEIGHT",
				fmt.Sprintf("%f", *panelContentRec.GetHeight()),
				1, 400, 1,
			),
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
