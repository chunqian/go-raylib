// +build linux,!arm,!arm64

package raylib

/*
#include "../lib/raylib/src/rglfw.c"
#cgo linux LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src/external/glfw/include -DPLATFORM_DESKTOP

// #cgo linux LDFLAGS: -L${SRCDIR}/../lib/raylib -lraylib -lGL -lm -pthread -ldl -lrt -lX11
// #cgo linux CFLAGS: -I../lib/raylib/src
*/
import "C"
