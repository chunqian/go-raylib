package raylib

/*
#include "../lib/raygui/src/raygui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ToBool(i interface{}) bool {
	switch i.(type) {
	case int:
		if i.(int) >= 1 {
			return true
		}
		return false
	case int32:
		if i.(int32) >= 1 {
			return true
		}
		return false
	default:
		panic(fmt.Sprintf("unable to convert %#v of type %T to bool", i, i))
	}
}

func ToInt32(i interface{}) int32 {
	switch i.(type) {
	case bool:
		if i.(bool) {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("unable to convert %#v of type %T to int32", i, i))
	}
}

// GuiListViewEx function as declared in src/raygui.h:479
func GuiListViewEx(bounds Rectangle, text *MultiText, count int32, focus *int32, scrollIndex *int32, active int32) int32 {
	cbounds := *(*C.Rectangle)(unsafe.Pointer(&bounds))
	ctext := (*C.MultiText)(unsafe.Pointer(text))
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cfocus, _ := (*C.int)(unsafe.Pointer(focus)), cgoAllocsUnknown
	cscrollIndex, _ := (*C.int)(unsafe.Pointer(scrollIndex)), cgoAllocsUnknown
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	cret := C.GuiListViewEx(cbounds, ctext.text, ccount, cfocus, cscrollIndex, cactive)
	gret := (int32)(cret)
	return gret
}
