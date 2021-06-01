// +build windows

package rres

/*
#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -I../lib/raylib/src -DRRES_IMPLEMENTATION -DRRES_RAYLIB_IMPLEMENTATION
*/
import "C"
