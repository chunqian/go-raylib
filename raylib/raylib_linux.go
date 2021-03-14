// +build linux,!arm,!arm64

package raylib

/*
#cgo linux LDFLAGS: -L${SRCDIR}/../lib/raylib -lraylib -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src
*/
import "C"
