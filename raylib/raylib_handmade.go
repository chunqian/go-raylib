package raylib

/*
#include "../lib/raylib/src/raylib.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	// "fmt"
	"runtime"
	"unsafe"
)

// NewColor - Returns new Color
func CreateColor(r, g, b, a uint8) Color {
	color := NewColor(byte(r), byte(g), byte(b), byte(a))
	return color
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

func NewBytes(str string, count int) []byte {

	bts := []byte(str + "\x00")
	bts2 := make([]byte, count)
	for i := 0; i < len(bts); i++ {
		bts2[i] = bts[i]
	}
	return bts2
}

func StringFromPPByte(names **byte, index int32) (raw string) {

	ptr0 := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(names)) + uintptr(index)*uintptr(sizeOfPtr)))
	ptrRow := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(*ptr0))))
	raw = C.GoString(ptrRow)
	return
}

func StringFromPString(p0 *string, index int32) (raw string) {

	h0 := (*stringHeader)(unsafe.Pointer(p0))
	ptr0 := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(&(h0.Data))) + uintptr(index)*uintptr(sizeOfPtr)))
	p := *ptr0
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

// GenImageFontAtlas function as declared in src/raylib.h:1218
func GenImageFontAtlas(chars *CharInfo, recs **Rectangle, charsCount int32, fontSize int32, padding int32, packMethod int32) Image {
	// cchars, _ := chars.PassRef()
	// crecs, _ := (*recs).PassMemoryRef()
	// cchars := chars
	crecs := (**C.Rectangle)(unsafe.Pointer(recs))
	cchars := (*C.CharInfo)(unsafe.Pointer(chars))
	ccharsCount, _ := (C.int)(charsCount), cgoAllocsUnknown
	cfontSize, _ := (C.int)(fontSize), cgoAllocsUnknown
	cpadding, _ := (C.int)(padding), cgoAllocsUnknown
	cpackMethod, _ := (C.int)(packMethod), cgoAllocsUnknown
	ret0 := C.GenImageFontAtlas(cchars, crecs, ccharsCount, cfontSize, cpadding, cpackMethod)
	v0 := *newImageRef(unsafe.Pointer(&ret0)).convert()
	return v0
}

// TextJoin function as declared in src/raylib.h:1244
func TextJoin(textList *MultiText, count int32, delimiter string) string {
	// ctextList, _ := textList.PassRef()

	ctextList := (*C.MultiText)(unsafe.Pointer(textList))
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	delimiter = safeString(delimiter)
	cdelimiter, _ := unpackPCharString(delimiter)
	ret0 := C.TextJoin(ctextList.text, ccount, cdelimiter)
	runtime.KeepAlive(delimiter)
	v0 := packPCharString(ret0)
	return v0
}

func UnloadColors(color *Color) {

	// colorRef, _ := color.PassRef()
	C.free(unsafe.Pointer(color))
}

func ToCamera3D(camera Camera) Camera3D {

	ret := *(*Camera3D)(unsafe.Pointer(&camera))
	return ret
}
