// +build windows

package raylib

/*
#cgo windows LDFLAGS: -L${SRCDIR}/../lib/raylib/a/windows -lraylib -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -I../lib/raylib/src
*/
import "C"
