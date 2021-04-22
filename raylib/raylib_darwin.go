// +build darwin

package raylib

/*
#include "../lib/raylib/src/core.c"
#include "../lib/raylib/src/shapes.c"
#include "../lib/raylib/src/textures.c"
#include "../lib/raylib/src/text.c"
#include "../lib/raylib/src/utils.c"
#include "../lib/raylib/src/models.c"
#include "../lib/raylib/src/raudio.c"

// #cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/src -lraylib
// #cgo darwin CFLAGS: -I../lib/raylib/src
*/
import "C"
