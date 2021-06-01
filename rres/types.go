// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 01 Jun 2021 11:52:25 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package rres

/*
#include "../lib/rres/src/rres.h"
#include "../lib/rres/src/rres-raylib.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// RresDataChunk as declared in src/rres.h:108
type gRresDataChunk struct {
	gType          uint32
	gPropsCount    uint32
	gProps         []int32
	gData          unsafe.Pointer
	reff705e719    *C.rresDataChunk
	allocsf705e719 interface{}
}
type RresDataChunk struct {
	Type       uint32
	PropsCount uint32
	Props      *int32
	Data       unsafe.Pointer
}

// RresData as declared in src/rres.h:114
type gRresData struct {
	gCount         uint32
	gChunks        []gRresDataChunk
	ref3a6e0f31    *C.rresData
	allocs3a6e0f31 interface{}
}
type RresData struct {
	Count  uint32
	Chunks *RresDataChunk
}

// RresDirEntry as declared in src/rres.h:124
type gRresDirEntry struct {
	gId            uint32
	gOffset        uint32
	gFileNameLen   uint32
	gFileName      [512]byte
	ref32db2e90    *C.rresDirEntry
	allocs32db2e90 interface{}
}
type RresDirEntry struct {
	Id          uint32
	Offset      uint32
	FileNameLen uint32
	FileName    [512]byte
}

// RresCentralDir as declared in src/rres.h:130
type gRresCentralDir struct {
	gCount         uint32
	gEntries       []gRresDirEntry
	ref87f3899d    *C.rresCentralDir
	allocs87f3899d interface{}
}
type RresCentralDir struct {
	Count   uint32
	Entries *RresDirEntry
}

// RresFontGlyphsInfo as declared in src/rres.h:138
type gRresFontGlyphsInfo struct {
	gX             int32
	gY             int32
	gWidth         int32
	gHeight        int32
	gValue         int32
	gOffsetX       int32
	gOffsetY       int32
	gAdvanceX      int32
	refd91ac199    *C.rresFontGlyphsInfo
	allocsd91ac199 interface{}
}
type RresFontGlyphsInfo struct {
	X        int32
	Y        int32
	Width    int32
	Height   int32
	Value    int32
	OffsetX  int32
	OffsetY  int32
	AdvanceX int32
}