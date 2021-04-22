// +build windows

package raylib

/*
#include "../lib/raylib/src/core.c"
#include "../lib/raylib/src/shapes.c"
#include "../lib/raylib/src/textures.c"
#include "../lib/raylib/src/text.c"
#include "../lib/raylib/src/utils.c"
#include "../lib/raylib/src/models.c"
#include "../lib/raylib/src/raudio.c"

// #cgo windows LDFLAGS: -L${SRCDIR}/../lib/raylib/src -lraylib -lopengl32 -lgdi32 -lwinmm -lole32
// #cgo windows CFLAGS: -I../lib/raylib/src
*/
import "C"
