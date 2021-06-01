// +build linux,!arm,!arm64

package rres

/*
#cgo linux LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src -DRRES_IMPLEMENTATION -DRRES_RAYLIB_IMPLEMENTATION
*/
import "C"
