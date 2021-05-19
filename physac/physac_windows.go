// +build windows

package physac

/*
// #cgo windows LDFLAGS: -L${SRCDIR}/../lib/raylib/src -lraylib -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -I../lib/raylib/src -DPHYSAC_IMPLEMENTATION -DPHYSAC_STATIC -DPHYSAC_NO_THREADS
*/
import "C"
