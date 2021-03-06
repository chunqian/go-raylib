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
import (
	"runtime"
	"unsafe"
)

// RresLoadData function as declared in src/rres.h:179
func RresLoadData(fileName string, rresId int32) RresData {
	fileName = safeString(fileName)
	cfileName, _ := unpackPCharString(fileName)
	crresId, _ := (C.int)(rresId), cgoAllocsUnknown
	__ret := C.rresLoadData(cfileName, crresId)
	runtime.KeepAlive(fileName)
	__v := *newRresDataRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresUnloadData function as declared in src/rres.h:180
func RresUnloadData(data RresData) {
	cdata, _ := *(*C.rresData)(unsafe.Pointer(&data)), cgoAllocsUnknown
	C.rresUnloadData(cdata)
}

// RresLoadCentralDirectory function as declared in src/rres.h:182
func RresLoadCentralDirectory(fileName string) RresCentralDir {
	fileName = safeString(fileName)
	cfileName, _ := unpackPCharString(fileName)
	__ret := C.rresLoadCentralDirectory(cfileName)
	runtime.KeepAlive(fileName)
	__v := *newRresCentralDirRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresUnloadCentralDirectory function as declared in src/rres.h:183
func RresUnloadCentralDirectory(dir RresCentralDir) {
	cdir, _ := *(*C.rresCentralDir)(unsafe.Pointer(&dir)), cgoAllocsUnknown
	C.rresUnloadCentralDirectory(cdir)
}

// RresGetIdFromFileName function as declared in src/rres.h:184
func RresGetIdFromFileName(dir RresCentralDir, fileName string) int32 {
	cdir, _ := *(*C.rresCentralDir)(unsafe.Pointer(&dir)), cgoAllocsUnknown
	fileName = safeString(fileName)
	cfileName, _ := unpackPCharString(fileName)
	__ret := C.rresGetIdFromFileName(cdir, cfileName)
	runtime.KeepAlive(fileName)
	__v := (int32)(__ret)
	return __v
}

// RresComputeCRC32 function as declared in src/rres.h:186
func RresComputeCRC32(buffer *byte, len int32) uint32 {
	cbuffer, _ := (*C.uchar)(unsafe.Pointer(buffer)), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	__ret := C.rresComputeCRC32(cbuffer, clen)
	__v := (uint32)(__ret)
	return __v
}

// RresLoadRaw function as declared in src/rres-raylib.h:48
func RresLoadRaw(rres RresData, size *int32) unsafe.Pointer {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	csize, _ := (*C.int)(unsafe.Pointer(size)), cgoAllocsUnknown
	__ret := C.rresLoadRaw(crres, csize)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// RresLoadText function as declared in src/rres-raylib.h:49
func RresLoadText(rres RresData) *byte {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadText(crres)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// RresLoadImage function as declared in src/rres-raylib.h:50
func RresLoadImage(rres RresData) Image {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadImage(crres)
	__v := *newImageRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresLoadWave function as declared in src/rres-raylib.h:51
func RresLoadWave(rres RresData) Wave {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadWave(crres)
	__v := *newWaveRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresLoadFont function as declared in src/rres-raylib.h:52
func RresLoadFont(rres RresData) Font {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadFont(crres)
	__v := *newFontRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresLoadMesh function as declared in src/rres-raylib.h:53
func RresLoadMesh(rres RresData) Mesh {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadMesh(crres)
	__v := *newMeshRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresLoadMaterial function as declared in src/rres-raylib.h:54
func RresLoadMaterial(rres RresData) Material {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadMaterial(crres)
	__v := *newMaterialRef(unsafe.Pointer(&__ret)).convert()
	return __v
}

// RresLoadModel function as declared in src/rres-raylib.h:55
func RresLoadModel(rres RresData) Model {
	crres, _ := *(*C.rresData)(unsafe.Pointer(&rres)), cgoAllocsUnknown
	__ret := C.rresLoadModel(crres)
	__v := *newModelRef(unsafe.Pointer(&__ret)).convert()
	return __v
}
