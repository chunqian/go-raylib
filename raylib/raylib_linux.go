// +build linux,!arm,!arm64

package raylib

/*
#include "../lib/raylib/src/core.c"
#include "../lib/raylib/src/shapes.c"
#include "../lib/raylib/src/textures.c"
#include "../lib/raylib/src/text.c"
#include "../lib/raylib/src/utils.c"
#include "../lib/raylib/src/models.c"
#include "../lib/raylib/src/raudio.c"

// #cgo linux LDFLAGS: -L${SRCDIR}/../lib/raylib/src -lraylib -lGL -lm -pthread -ldl -lrt -lX11
// #cgo linux CFLAGS: -I../lib/raylib/src
*/
import "C"
