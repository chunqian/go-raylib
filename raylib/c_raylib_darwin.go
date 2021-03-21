// +build darwin

package raylib

/*
#include "../lib/raylib/src/rglfw.c"
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo
#cgo darwin CFLAGS: -x objective-c -I../lib/raylib/src/external/glfw/include -DPLATFORM_DESKTOP
#cgo darwin CFLAGS: -mmacosx-version-min=10.9

// #cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/src -lraylib
// #cgo darwin CFLAGS: -I../lib/raylib/src
*/
import "C"
