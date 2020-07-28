package raylib

/*
#include "../lib/raylib/src/raylib.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	// "fmt"
)

// NewColor - Returns new Color
func NewColorStructure(r, g, b, a uint8) *Color {
	color := Color{
		R: byte(r),
		G: byte(g),
		B: byte(b),
		A: byte(a),
	}
	return &color
}

var (
	// Light Gray
	LightGray = NewColorStructure(200, 200, 200, 255)
	// Gray
	Gray = NewColorStructure(130, 130, 130, 255)
	// Dark Gray
	DarkGray = NewColorStructure(80, 80, 80, 255)
	// Yellow
	Yellow = NewColorStructure(253, 249, 0, 255)
	// Gold
	Gold = NewColorStructure(255, 203, 0, 255)
	// Orange
	Orange = NewColorStructure(255, 161, 0, 255)
	// Pink
	Pink = NewColorStructure(255, 109, 194, 255)
	// Red
	Red = NewColorStructure(230, 41, 55, 255)
	// Maroon
	Maroon = NewColorStructure(190, 33, 55, 255)
	// Green
	Green = NewColorStructure(0, 228, 48, 255)
	// Lime
	Lime = NewColorStructure(0, 158, 47, 255)
	// Dark Green
	DarkGreen = NewColorStructure(0, 117, 44, 255)
	// Sky Blue
	SkyBlue = NewColorStructure(102, 191, 255, 255)
	// Blue
	Blue = NewColorStructure(0, 121, 241, 255)
	// Dark Blue
	DarkBlue = NewColorStructure(0, 82, 172, 255)
	// Purple
	Purple = NewColorStructure(200, 122, 255, 255)
	// Violet
	Violet = NewColorStructure(135, 60, 190, 255)
	// Dark Purple
	DarkPurple = NewColorStructure(112, 31, 126, 255)
	// Beige
	Beige = NewColorStructure(211, 176, 131, 255)
	// Brown
	Brown = NewColorStructure(127, 106, 79, 255)
	// Dark Brown
	DarkBrown = NewColorStructure(76, 63, 47, 255)
	// White
	White = NewColorStructure(255, 255, 255, 255)
	// Black
	Black = NewColorStructure(0, 0, 0, 255)
	// Blank (Transparent)
	Blank = NewColorStructure(0, 0, 0, 0)
	// Magenta
	Magenta = NewColorStructure(255, 0, 255, 255)
	// Ray White (RayLib Logo White)
	RayWhite = NewColorStructure(245, 245, 245, 255)
)

func loadModelAnimationsLength(fileName string, animsCount *int32) int32 {
	return *animsCount
}

func getDroppedFilesLength(character *C.char, count *int32) int32 {
	return int32(0)
}

func getDirectoryFilesLength(character *C.char, dirPath string, count *int32) int32 {
	return int32(0)
}
