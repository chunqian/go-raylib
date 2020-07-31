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
