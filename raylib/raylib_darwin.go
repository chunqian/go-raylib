// +build darwin

package raylib

/*
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/a/darwin -lraylib
#cgo darwin CFLAGS: -I../lib/raylib/src
*/
import "C"
