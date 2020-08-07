package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
// "fmt"
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

	bts := []byte(str)
	bts2 := make([]byte, count)
	for i := 0; i < count; i++ {
		if i < len(bts) {
			bts2[i] = bts[i]
		} else {
			btsT := []byte("\x00")
			bts2[i] = btsT[0]
		}
	}
	return bts2
}
