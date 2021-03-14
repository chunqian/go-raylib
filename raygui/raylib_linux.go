// +build linux,!arm,!arm64

package raygui

/*
#cgo linux LDFLAGS: -L${SRCDIR}/../lib/raylib -lraylib -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src -DRAYGUI_IMPLEMENTATION -DRAYGUI_STATIC -DRAYGUI_SUPPORT_ICONS
*/
import "C"
