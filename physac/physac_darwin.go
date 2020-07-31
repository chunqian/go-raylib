// +build darwin

package physac

/*
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/src -lraylib
#cgo darwin CFLAGS: -I../lib/raylib/src -DPHYSAC_IMPLEMENTATION -DPHYSAC_STATIC -DPHYSAC_NO_THREADS
*/
import "C"
