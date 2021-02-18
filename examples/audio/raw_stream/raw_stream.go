package main

import (
	rl "go-raylib/raylib"

	"fmt"
	"math"
	"runtime"
	"unsafe"
)

var (
	MAX_SAMPLES            = 512
	MAX_SAMPLES_PER_UPDATE = 4096
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [audio] example - raw audio streaming")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	stream := rl.InitAudioStream(22050, 16, 1)
	defer rl.CloseAudioStream(stream)

	data := make([]int16, MAX_SAMPLES)
	writeBuf := make([]int16, MAX_SAMPLES_PER_UPDATE)

	rl.PlayAudioStream(stream)

	mousePosition := rl.NewVector2(-100, -100)
	frequency := float32(440.0)
	oldFrequency := float32(1.0)
	readCursor := 0
	waveLength := 1

	position := rl.NewVector2(0, 0)

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {

		mousePosition = rl.GetMousePosition()

		if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {
			fp := (&mousePosition).Y
			frequency = 40.0 + fp
		}
		if frequency != oldFrequency {
			oldWavelength := waveLength
			waveLength = int(float32(22050) / frequency)
			if waveLength > MAX_SAMPLES/2 {
				waveLength = MAX_SAMPLES / 2
			}
			if waveLength < 1 {
				waveLength = 1
			}

			for i := 0; i < int(waveLength)*2; i++ {
				data[i] = int16(math.Sin(2.0*rl.PI*float64(i)/float64(waveLength)) * float64(32000))
			}

			readCursor = readCursor * int(float32(waveLength)/float32(oldWavelength))
			oldFrequency = frequency
		}

		if rl.IsAudioStreamProcessed(stream) {

			writeCursor := 0
			for writeCursor < MAX_SAMPLES_PER_UPDATE {
				writeLength := MAX_SAMPLES_PER_UPDATE - writeCursor
				readLength := waveLength - readCursor

				if writeLength > readLength {
					writeLength = readLength
				}
				
				for i := 0; i < writeLength; i++ {
					writeBuf[writeCursor+i] = data[readCursor+i]
				}

				readCursor = (readCursor + writeLength) % waveLength
				writeCursor += writeLength
			}
			rl.UpdateAudioStream(stream, unsafe.Pointer(&writeBuf[0]), int32(MAX_SAMPLES_PER_UPDATE))
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("sine frequency: %d", int(frequency)), rl.GetScreenWidth()-220, 10, 20, rl.Red)
		rl.DrawText("click mouse button to change frequency", 10, 10, 20, rl.DarkGray)

		for i := int32(0); i < screenWidth; i++ {
			px := float32(i)
			py := float32(250.0 + float32(50)*float32(data[i*int32(MAX_SAMPLES)/screenWidth])/float32(32000))
			position.X = px
			position.Y = py

			rl.DrawPixelV(position, rl.Red)
		}
		rl.EndDrawing()
	}
}
