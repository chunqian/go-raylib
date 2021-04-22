package raylib

/*
#include "../lib/raylib/src/rglfw.c"

#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo
#cgo darwin CFLAGS: -x objective-c -mmacosx-version-min=10.9 -I../lib/raylib/src/external/glfw/include -DPLATFORM_DESKTOP

#cgo linux LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src/external/glfw/include -Wno-unused-result -DPLATFORM_DESKTOP

#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -I../lib/raylib/src/external/glfw/include -DPLATFORM_DESKTOP
*/
import "C"
