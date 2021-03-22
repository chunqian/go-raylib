// +build windows

package raygui

/*
#cgo windows LDFLAGS: -L${SRCDIR}/../lib/raylib/a/windows -lraylib -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -I../lib/raylib/src -DRAYGUI_IMPLEMENTATION -DRAYGUI_STATIC -DRAYGUI_SUPPORT_ICONS
*/
import "C"
