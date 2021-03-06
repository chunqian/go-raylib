// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 22 Apr 2021 14:00:10 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raylib

/*
#include "../lib/raygui/src/raygui.h"
#include "../lib/raygui/src/ricons.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// allocGuiStylePropMemory allocates memory for type C.GuiStyleProp in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGuiStylePropMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGuiStylePropValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGuiStylePropValue = unsafe.Sizeof([1]C.GuiStyleProp{})

// newGuiStylePropRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func newGuiStylePropRef(ref unsafe.Pointer) *gGuiStyleProp {
	if ref == nil {
		return nil
	}
	obj := new(gGuiStyleProp)
	obj.ref8760bcee = (*C.GuiStyleProp)(unsafe.Pointer(ref))
	return obj
}

// passRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *gGuiStyleProp) passRef() (*C.GuiStyleProp, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8760bcee != nil {
		if x.allocs8760bcee != nil {
			return x.ref8760bcee, x.allocs8760bcee.(*cgoAllocMap)
		} else {
			return x.ref8760bcee, nil
		}
	}
	mem8760bcee := unsafe.Pointer(new(C.GuiStyleProp))
	ref8760bcee := (*C.GuiStyleProp)(mem8760bcee)
	allocs8760bcee := new(cgoAllocMap)
	// allocs8760bcee.Add(mem8760bcee)

	var ccontrolId_allocs *cgoAllocMap
	ref8760bcee.controlId, ccontrolId_allocs = (C.ushort)(x.gControlId), cgoAllocsUnknown
	allocs8760bcee.Borrow(ccontrolId_allocs)
	x.gControlId = *new(uint16)

	var cpropertyId_allocs *cgoAllocMap
	ref8760bcee.propertyId, cpropertyId_allocs = (C.ushort)(x.gPropertyId), cgoAllocsUnknown
	allocs8760bcee.Borrow(cpropertyId_allocs)
	x.gPropertyId = *new(uint16)

	var cpropertyValue_allocs *cgoAllocMap
	ref8760bcee.propertyValue, cpropertyValue_allocs = (C.int)(x.gPropertyValue), cgoAllocsUnknown
	allocs8760bcee.Borrow(cpropertyValue_allocs)
	x.gPropertyValue = *new(int32)

	x.ref8760bcee = ref8760bcee
	x.allocs8760bcee = allocs8760bcee

	return ref8760bcee, allocs8760bcee
}

// passValue does the same as passRef except that it will try to dereference the returned pointer.
func (x gGuiStyleProp) passValue() (C.GuiStyleProp, *cgoAllocMap) {
	if x.ref8760bcee != nil {
		return *x.ref8760bcee, nil
	}
	ref, allocs := x.passRef()
	return *ref, allocs
}

// convert struct for mapping C struct unanimous.
func (x *gGuiStyleProp) convert() *GuiStyleProp {
	if x.ref8760bcee != nil {
		return (*GuiStyleProp)(unsafe.Pointer(x.ref8760bcee))
	}
	x.passRef()
	return (*GuiStyleProp)(unsafe.Pointer(x.ref8760bcee))
}

// NewGuiStyleProp new Go object and Mapping to C object.
func NewGuiStyleProp(cControlId uint16, cPropertyId uint16, cPropertyValue int32) GuiStyleProp {
	obj := *new(gGuiStyleProp)
	obj.gControlId = cControlId
	obj.gPropertyId = cPropertyId
	obj.gPropertyValue = cPropertyValue

	ret0, alloc0 := obj.passRef()

	if len(alloc0.m) > 0 {
		panic("Cgo memory alloced, please use func AllocGuiStyleProp.")
	}
	return *(*GuiStyleProp)(unsafe.Pointer(ret0))
}

// AllocGuiStyleProp new Go object and Mapping to C object.
func AllocGuiStyleProp(cControlId uint16, cPropertyId uint16, cPropertyValue int32) (*GuiStyleProp, *cgoAllocMap) {
	obj := *new(gGuiStyleProp)
	obj.gControlId = cControlId
	obj.gPropertyId = cPropertyId
	obj.gPropertyValue = cPropertyValue

	ret0, alloc0 := obj.passRef()
	ret1 := (*GuiStyleProp)(unsafe.Pointer(ret0))
	return ret1, alloc0
}

// Index reads Go data structure out from plain C format.
func (x *GuiStyleProp) Index(index int32) *GuiStyleProp {
	ptr1 := (*GuiStyleProp)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(index)*uintptr(sizeOfGuiStylePropValue)))
	return ptr1
}

// GC is register for garbage collection.
func (x *GuiStyleProp) GC(a *cgoAllocMap, args ...*cgoAllocMap) {
	for i := range args {
		a.Borrow(args[i])
	}
	if len(a.m) > 0 {
		for ptr := range a.m {
			fmt.Printf("INFO: MEMORY: [PTR %p] GC register\n", ptr)
		}
		runtime.SetFinalizer(x, func(*GuiStyleProp) {
			a.Free()
		})
	}
}
