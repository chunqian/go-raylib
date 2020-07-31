// +build darwin

package raygui

/*
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/src -lraylib
#cgo darwin CFLAGS: -I../lib/raylib/src -DRAYGUI_IMPLEMENTATION -DRAYGUI_STATIC
*/
import "C"
