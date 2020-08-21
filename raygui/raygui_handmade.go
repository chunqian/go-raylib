package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	// "fmt"
	"unsafe"
)

func Bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Int2bool(i int) bool {
	if i >= 1 {
		return true
	}
	return false
}

func NewBytes(str string, count int) []byte {

	bts := []byte(str + "\x00")
	bts2 := make([]byte, count)
	for i := 0; i < len(bts); i++ {
		bts2[i] = bts[i]
	}
	return bts2
}

// GuiListViewEx function as declared in src/raygui.h:479
func GuiListViewEx(bounds Rectangle, text *Text, count int32, focus *int32, scrollIndex *int32, active int32) int32 {
	cbounds, _ := bounds.PassValue()
	ctext, _ := text.PassRef()
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cfocus, _ := (*C.int)(unsafe.Pointer(focus)), cgoAllocsUnknown
	cscrollIndex, _ := (*C.int)(unsafe.Pointer(scrollIndex)), cgoAllocsUnknown
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	cret := C.GuiListViewEx(cbounds, ctext.text, ccount, cfocus, cscrollIndex, cactive)
	gret := (int32)(cret)
	return gret
}
