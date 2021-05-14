package main

import (
	"fmt"
	rl "github.com/chunqian/go-raylib/raylib"
	"runtime"
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

	textBoxText := rl.NewBytes("Text box", 64)
	textBoxEditMode := false

	listViewScrollIndex := int32(0)
	listViewActive := int32(-1)

	listViewExScrollIndex := int32(0)
	listViewExActive := int32(2)
	listViewExFocus := int32(-1)
	listViewExList, mem := rl.AllocMultiText([]string{"This", "is", "a", "list view", "with", "disable", "elements", "amazing!"})
	listViewExList.GC(mem)

	multiTextBoxText := rl.NewBytes("Multi text box", 256)
	multiTextBoxEditMode := false
	// colorPickerValue := *(*rl.Color)(unsafe.Pointer(&rl.Red))
	colorPickerValue := rl.Red

	sliderValue := float32(50)
	sliderBarValue := float32(60)
	progressValue := float32(0.4)

	forceSquaredChecked := false

	alphaValue := float32(0.5)

	comboBoxActive := int32(1)

	toggleGroupActive := int32(0)

	viewScroll := rl.NewVector2(0, 0)

	exitWindow := false
	showMessageBox := false

	textInput := rl.NewBytes("", 256)
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
				rl.GuiLoadStyle(droppedFilePath)
			}
			rl.ClearDroppedFiles()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.GetColor(rl.GuiGetStyle(int32(rl.DEFAULT), int32(rl.BACKGROUND_COLOR))))

		if dropDown000EditMode || dropDown001EditMode {
			rl.GuiLock()
		}

		forceSquaredChecked = rl.GuiCheckBox(rl.NewRectangle(25, 108, 15, 15), "FORCE CHECK!", forceSquaredChecked)

		rl.GuiSetStyle(int32(rl.TEXTBOX), int32(rl.TEXT_ALIGNMENT), int32(rl.GUI_TEXT_ALIGN_CENTER))

		if rl.GuiSpinner(rl.NewRectangle(25, 135, 125, 30), "", &spinner001Value, 0, 100, spinnerEditMode) {
			spinnerEditMode = !spinnerEditMode
		}
		if rl.GuiValueBox(rl.NewRectangle(25, 175, 125, 30), "", &valueBox002Value, 0, 100, valueBoxEditMode) {
			valueBoxEditMode = !valueBoxEditMode
		}

		rl.GuiSetStyle(int32(rl.TEXTBOX), int32(rl.TEXT_ALIGNMENT), int32(rl.GUI_TEXT_ALIGN_LEFT))

		if rl.GuiTextBox(rl.NewRectangle(25, 215, 125, 30), &textBoxText[0], 64, textBoxEditMode) {
			textBoxEditMode = !textBoxEditMode
		}

		rl.GuiSetStyle(int32(rl.BUTTON), int32(rl.TEXT_ALIGNMENT), int32(rl.GUI_TEXT_ALIGN_CENTER))

		rl.GuiSetTooltip("Save current file.")

		if rl.GuiButton(rl.NewRectangle(25, 255, 125, 30), rl.GuiIconText(int32(rl.RICON_FILE_SAVE), "Save File")) {
			showTextInputBox = true
		}
		rl.GuiClearTooltip()

		rl.GuiGroupBox(rl.NewRectangle(25, 310, 125, 150), "STATES")
		rl.GuiLock()
		rl.GuiSetState(int32(rl.GUI_STATE_NORMAL))
		if rl.GuiButton(rl.NewRectangle(30, 320, 115, 30), "NORMAL") {

		}
		rl.GuiSetState(int32(rl.GUI_STATE_FOCUSED))
		if rl.GuiButton(rl.NewRectangle(30, 355, 115, 30), "FOCUSED") {

		}
		rl.GuiSetState(int32(rl.GUI_STATE_PRESSED))
		if rl.GuiButton(rl.NewRectangle(30, 390, 115, 30), "#15#PRESSED") {

		}
		rl.GuiSetState(int32(rl.GUI_STATE_DISABLED))
		if rl.GuiButton(rl.NewRectangle(30, 425, 115, 30), "DISABLED") {

		}
		rl.GuiSetState(int32(rl.GUI_STATE_NORMAL))
		rl.GuiUnlock()

		comboBoxActive = rl.GuiComboBox(rl.NewRectangle(25, 470, 125, 30), "ONE;TWO;THREE;FOUR", comboBoxActive)

		rl.GuiSetStyle(int32(rl.DROPDOWNBOX), int32(rl.TEXT_ALIGNMENT), int32(rl.GUI_TEXT_ALIGN_LEFT))
		if rl.GuiDropdownBox(
			rl.NewRectangle(25, 65, 125, 30),
			"#01#ONE;#02#TWO;#03#THREE;#04#FOUR",
			&dropdownBox001Active,
			dropDown001EditMode,
		) {
			dropDown001EditMode = !dropDown001EditMode
		}

		rl.GuiSetStyle(int32(rl.DROPDOWNBOX), int32(rl.TEXT_ALIGNMENT), int32(rl.GUI_TEXT_ALIGN_CENTER))
		if rl.GuiDropdownBox(
			rl.NewRectangle(25, 25, 125, 30),
			"ONE;TWO;THREE",
			&dropdownBox000Active,
			dropDown000EditMode,
		) {
			dropDown000EditMode = !dropDown000EditMode
		}

		listViewActive = rl.GuiListView(
			rl.NewRectangle(165, 25, 140, 140),
			"Charmander;Bulbasaur;#18#Squirtel;Pikachu;Eevee;Pidgey",
			&listViewScrollIndex,
			listViewActive,
		)
		listViewExActive = rl.GuiListViewEx(
			rl.NewRectangle(165, 180, 140, 200),
			listViewExList,
			8,
			&listViewExFocus,
			&listViewExScrollIndex,
			listViewExActive,
		)

		toggleGroupActive = rl.GuiToggleGroup(rl.NewRectangle(165, 400, 140, 25), "#1#ONE\n#3#TWO\n#8#THREE\n#23#", toggleGroupActive)

		if rl.GuiTextBoxMulti(rl.NewRectangle(320, 25, 225, 140), &multiTextBoxText[0], 256, multiTextBoxEditMode) {
			multiTextBoxEditMode = !multiTextBoxEditMode
		}
		colorPickerValue = rl.GuiColorPicker(rl.NewRectangle(320, 185, 196, 192), colorPickerValue)

		sliderValue = rl.GuiSlider(rl.NewRectangle(355, 400, 165, 20), "TEST", fmt.Sprintf("%.2f", sliderValue), sliderValue, -50, 100)
		sliderBarValue = rl.GuiSliderBar(rl.NewRectangle(320, 430, 200, 20), "", fmt.Sprintf("%d", int32(sliderBarValue)), sliderBarValue, 0, 100)
		progressValue = rl.GuiProgressBar(rl.NewRectangle(320, 460, 200, 20), "", "", progressValue, 0, 1)

		rl.GuiScrollPanel(rl.NewRectangle(560, 25, 100, 160), rl.NewRectangle(560, 25, 200, 400), &viewScroll)

		rl.GuiStatusBar(rl.NewRectangle(0, float32(rl.GetScreenHeight()-20), float32(rl.GetScreenWidth()), 20), "This is a status bar")

		alphaValue = rl.GuiColorBarAlpha(rl.NewRectangle(320, 490, 200, 30), alphaValue)

		if showMessageBox {
			rl.DrawRectangle(0, 0, rl.GetScreenWidth(), rl.GetScreenHeight(), rl.Fade(rl.RayWhite, 0.8))

			result := rl.GuiMessageBox(
				rl.NewRectangle(float32(rl.GetScreenWidth()/2-125), float32(rl.GetScreenHeight()/2-50), 250, 100),
				rl.GuiIconText(int32(rl.RICON_EXIT), "Close Window"),
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

			result := rl.GuiTextInputBox(
				rl.NewRectangle(float32(rl.GetScreenWidth()/2-120), float32(rl.GetScreenHeight()/2-60), 240, 140),
				rl.GuiIconText(int32(rl.RICON_FILE_SAVE), "Save file as..."),
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
				textInput = rl.NewBytes("", 256)
			}
		}

		rl.GuiUnlock()

		rl.EndDrawing()
	}
}
