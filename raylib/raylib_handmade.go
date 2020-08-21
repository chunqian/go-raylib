package raylib

/*
#include "../lib/raylib/src/raylib.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	// "fmt"
	"unsafe"
)

// NewColor - Returns new Color
func CreateColor(r, g, b, a uint8) *Color {
	// color := Color{
	// 	R: byte(r),
	// 	G: byte(g),
	// 	B: byte(b),
	// 	A: byte(a),
	// }
	color := NewColor(byte(r), byte(g), byte(b), byte(a))
	return &color
}

var (
	// Light Gray
	LightGray = CreateColor(200, 200, 200, 255)
	// Gray
	Gray = CreateColor(130, 130, 130, 255)
	// Dark Gray
	DarkGray = CreateColor(80, 80, 80, 255)
	// Yellow
	Yellow = CreateColor(253, 249, 0, 255)
	// Gold
	Gold = CreateColor(255, 203, 0, 255)
	// Orange
	Orange = CreateColor(255, 161, 0, 255)
	// Pink
	Pink = CreateColor(255, 109, 194, 255)
	// Red
	Red = CreateColor(230, 41, 55, 255)
	// Maroon
	Maroon = CreateColor(190, 33, 55, 255)
	// Green
	Green = CreateColor(0, 228, 48, 255)
	// Lime
	Lime = CreateColor(0, 158, 47, 255)
	// Dark Green
	DarkGreen = CreateColor(0, 117, 44, 255)
	// Sky Blue
	SkyBlue = CreateColor(102, 191, 255, 255)
	// Blue
	Blue = CreateColor(0, 121, 241, 255)
	// Dark Blue
	DarkBlue = CreateColor(0, 82, 172, 255)
	// Purple
	Purple = CreateColor(200, 122, 255, 255)
	// Violet
	Violet = CreateColor(135, 60, 190, 255)
	// Dark Purple
	DarkPurple = CreateColor(112, 31, 126, 255)
	// Beige
	Beige = CreateColor(211, 176, 131, 255)
	// Brown
	Brown = CreateColor(127, 106, 79, 255)
	// Dark Brown
	DarkBrown = CreateColor(76, 63, 47, 255)
	// White
	White = CreateColor(255, 255, 255, 255)
	// Black
	Black = CreateColor(0, 0, 0, 255)
	// Blank (Transparent)
	Blank = CreateColor(0, 0, 0, 0)
	// Magenta
	Magenta = CreateColor(255, 0, 255, 255)
	// Ray White (RayLib Logo White)
	RayWhite = CreateColor(245, 245, 245, 255)
)

func ConvertFilesPath(names **byte, index int32) string {

	ptr0 := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(names)) + uintptr(index)*uintptr(sizeOfPtr)))
	ptrRow := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(*ptr0))))
	v0 := C.GoString(ptrRow)
	return v0
}
