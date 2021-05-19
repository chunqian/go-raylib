// +build linux,!arm,!arm64

package raylib

/*

#cgo linux CFLAGS: -I../lib/raylib/src -DRAYGUI_IMPLEMENTATION -DRAYGUI_STATIC -DRAYGUI_SUPPORT_ICONS
*/
import "C"
