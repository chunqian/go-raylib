package main

import (
	"fmt"
	rg "goray/raygui"
	rl "goray/raylib"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(690)
	screenHeight := int32(560)

	rl.InitWindow(screenWidth, screenHeight, "raygui - controls test suite")
	defer rl.CloseWindow()

	rl.SetExitKey(0)

	dropdownBox000Active := int32(0)
	dropDown000EditMode := false

	dropdownBox001Active := int32(0)
	dropDown001EditMode := false

	spinner001Value := int32(0)
	spinnerEditMode := false

	valueBox002Value := int32(0)
	valueBoxEditMode := false

	textBoxText := rg.NewBytes("Text box", 64)
	textBoxEditMode := false

	listViewScrollIndex := int32(0)
	listViewActive := int32(-1)

	listViewExScrollIndex := int32(0)
	listViewExActive := int32(2)
	listViewExFocus := int32(-1)
	listViewExList, mem := rg.AllocMultiText([]string{"This", "is", "a", "list view", "with", "disable", "elements", "amazing!"})
	listViewExList.GC(mem)

	multiTextBoxText := rg.NewBytes("Multi text box", 256)
	multiTextBoxEditMode := false
	colorPickerValue := *(*rg.Color)(unsafe.Pointer(&rl.Red))

	sliderValue := float32(50)
	sliderBarValue := float32(60)
	progressValue := float32(0.4)

	forceSquaredChecked := false

	alphaValue := float32(0.5)

	comboBoxActive := int32(1)

	toggleGroupActive := int32(0)

	viewScroll := rg.NewVector2(0, 0)

	exitWindow := false
	showMessageBox := false

	textInput := rg.NewBytes("", 256)
	showTextInputBox := false

	// textInputFileName := ""

	rl.SetTargetFPS(30)

	for !exitWindow {
		exitWindow = rl.WindowShouldClose()

		if rl.IsKeyPressed(int32(rl.KEY_ESCAPE)) {
			showMessageBox = !showMessageBox
		}
		if rl.IsKeyDown(int32(rl.KEY_LEFT_CONTROL)) && rl.IsKeyPressed(int32(rl.KEY_S)) {
			showTextInputBox = true
		}
		if rl.IsFileDropped() {
			dropsCount := int32(0)
			droppedFiles := rl.GetDroppedFiles(&dropsCount)
			droppedFilePath := rl.ToString(droppedFiles, 0)

			if dropsCount > 0 && rl.IsFileExtension(droppedFilePath, ".rgs") {
				rg.GuiLoadStyle(droppedFilePath)
			}
			rl.ClearDroppedFiles()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.GetColor(rg.GuiGetStyle(int32(rg.DEFAULT), int32(rg.BACKGROUND_COLOR))))

		if dropDown000EditMode || dropDown001EditMode {
			rg.GuiLock()
		}

		forceSquaredChecked = rg.GuiCheckBox(rg.NewRectangle(25, 108, 15, 15), "FORCE CHECK!", forceSquaredChecked)

		rg.GuiSetStyle(int32(rg.TEXTBOX), int32(rg.TEXT_ALIGNMENT), int32(rg.GUI_TEXT_ALIGN_CENTER))

		if rg.GuiSpinner(rg.NewRectangle(25, 135, 125, 30), "", &spinner001Value, 0, 100, spinnerEditMode) {
			spinnerEditMode = !spinnerEditMode
		}
		if rg.GuiValueBox(rg.NewRectangle(25, 175, 125, 30), "", &valueBox002Value, 0, 100, valueBoxEditMode) {
			valueBoxEditMode = !valueBoxEditMode
		}

		rg.GuiSetStyle(int32(rg.TEXTBOX), int32(rg.TEXT_ALIGNMENT), int32(rg.GUI_TEXT_ALIGN_LEFT))

		if rg.GuiTextBox(rg.NewRectangle(25, 215, 125, 30), &textBoxText[0], 64, textBoxEditMode) {
			textBoxEditMode = !textBoxEditMode
		}

		rg.GuiSetStyle(int32(rg.BUTTON), int32(rg.TEXT_ALIGNMENT), int32(rg.GUI_TEXT_ALIGN_CENTER))

		rg.GuiSetTooltip("Save current file.")

		if rg.GuiButton(rg.NewRectangle(25, 255, 125, 30), rg.GuiIconText(int32(rg.RICON_FILE_SAVE), "Save File")) {
			showTextInputBox = true
		}
		rg.GuiClearTooltip()

		rg.GuiGroupBox(rg.NewRectangle(25, 310, 125, 150), "STATES")
		rg.GuiLock()
		rg.GuiSetState(int32(rg.GUI_STATE_NORMAL))
		if rg.GuiButton(rg.NewRectangle(30, 320, 115, 30), "NORMAL") {

		}
		rg.GuiSetState(int32(rg.GUI_STATE_FOCUSED))
		if rg.GuiButton(rg.NewRectangle(30, 355, 115, 30), "FOCUSED") {

		}
		rg.GuiSetState(int32(rg.GUI_STATE_PRESSED))
		if rg.GuiButton(rg.NewRectangle(30, 390, 115, 30), "#15#PRESSED") {

		}
		rg.GuiSetState(int32(rg.GUI_STATE_DISABLED))
		if rg.GuiButton(rg.NewRectangle(30, 425, 115, 30), "DISABLED") {

		}
		rg.GuiSetState(int32(rg.GUI_STATE_NORMAL))
		rg.GuiUnlock()

		comboBoxActive = rg.GuiComboBox(rg.NewRectangle(25, 470, 125, 30), "ONE;TWO;THREE;FOUR", comboBoxActive)

		rg.GuiSetStyle(int32(rg.DROPDOWNBOX), int32(rg.TEXT_ALIGNMENT), int32(rg.GUI_TEXT_ALIGN_LEFT))
		if rg.GuiDropdownBox(
			rg.NewRectangle(25, 65, 125, 30),
			"#01#ONE;#02#TWO;#03#THREE;#04#FOUR",
			&dropdownBox001Active,
			dropDown001EditMode,
		) {
			dropDown001EditMode = !dropDown001EditMode
		}

		rg.GuiSetStyle(int32(rg.DROPDOWNBOX), int32(rg.TEXT_ALIGNMENT), int32(rg.GUI_TEXT_ALIGN_CENTER))
		if rg.GuiDropdownBox(
			rg.NewRectangle(25, 25, 125, 30),
			"ONE;TWO;THREE",
			&dropdownBox000Active,
			dropDown000EditMode,
		) {
			dropDown000EditMode = !dropDown000EditMode
		}

		listViewActive = rg.GuiListView(
			rg.NewRectangle(165, 25, 140, 140),
			"Charmander;Bulbasaur;#18#Squirtel;Pikachu;Eevee;Pidgey",
			&listViewScrollIndex,
			listViewActive,
		)
		listViewExActive = rg.GuiListViewEx(
			rg.NewRectangle(165, 180, 140, 200),
			listViewExList,
			8,
			&listViewExFocus,
			&listViewExScrollIndex,
			listViewExActive,
		)

		toggleGroupActive = rg.GuiToggleGroup(rg.NewRectangle(165, 400, 140, 25), "#1#ONE\n#3#TWO\n#8#THREE\n#23#", toggleGroupActive)

		if rg.GuiTextBoxMulti(rg.NewRectangle(320, 25, 225, 140), &multiTextBoxText[0], 256, multiTextBoxEditMode) {
			multiTextBoxEditMode = !multiTextBoxEditMode
		}
		colorPickerValue = rg.GuiColorPicker(rg.NewRectangle(320, 185, 196, 192), colorPickerValue)

		sliderValue = rg.GuiSlider(rg.NewRectangle(355, 400, 165, 20), "TEST", fmt.Sprintf("%.2f", sliderValue), sliderValue, -50, 100)
		sliderBarValue = rg.GuiSliderBar(rg.NewRectangle(320, 430, 200, 20), "", fmt.Sprintf("%d", int32(sliderBarValue)), sliderBarValue, 0, 100)
		progressValue = rg.GuiProgressBar(rg.NewRectangle(320, 460, 200, 20), "", "", progressValue, 0, 1)

		rg.GuiScrollPanel(rg.NewRectangle(560, 25, 100, 160), rg.NewRectangle(560, 25, 200, 400), &viewScroll)

		rg.GuiStatusBar(rg.NewRectangle(0, float32(rl.GetScreenHeight()-20), float32(rl.GetScreenWidth()), 20), "This is a status bar")

		alphaValue = rg.GuiColorBarAlpha(rg.NewRectangle(320, 490, 200, 30), alphaValue)

		if showMessageBox {
			rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Fade(rl.RayWhite, 0.8))

			result := rg.GuiMessageBox(
				rg.NewRectangle(float32(rl.GetScreenWidth()/2-125), float32(rl.GetScreenHeight()/2-50), 250, 100),
				rg.GuiIconText(int32(rg.RICON_EXIT), "Close Window"),
				"Do you really want to exit?",
				"Yes;No",
			)
			if result == 0 || result == 2 {
				showMessageBox = false
			} else if result == 1 {
				exitWindow = true
			}
		}

		if showTextInputBox {
			rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Fade(rl.RayWhite, 0.8))

			result := rg.GuiTextInputBox(
				rg.NewRectangle(float32(rl.GetScreenWidth()/2-120), float32(rl.GetScreenHeight()/2-60), 240, 140),
				rg.GuiIconText(int32(rg.RICON_FILE_SAVE), "Save file as..."),
				"Introduce a save file name",
				"Ok;Cancel",
				&textInput[0],
			)
			if result == 1 {
				// TODO
				// textInputFileName = textInput
			}
			if result == 0 || result == 1 || result == 2 {
				showTextInputBox = false
				textInput = rg.NewBytes("", 256)
			}
		}

		rg.GuiUnlock()

		rl.EndDrawing()
	}
}
