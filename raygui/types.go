// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 31 Aug 2020 21:42:43 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include "../lib/raygui/src/ricons.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// GuiStyleProp as declared in src/raygui.h:261
type GuiStyleProp struct {
	gControlId     uint16
	gPropertyId    uint16
	gPropertyValue int32
	ref8760bcee    *C.GuiStyleProp
	allocs8760bcee interface{}
	This           *guiStyleProp
}
type guiStyleProp struct {
	ControlId     uint16
	PropertyId    uint16
	PropertyValue int32
}

// Vector2 as declared in src/raylib.h:176
type Vector2 struct {
	gX             float32
	gY             float32
	ref29ca61a5    *C.Vector2
	allocs29ca61a5 interface{}
	This           *vector2
}
type vector2 struct {
	X float32
	Y float32
}

// Vector3 as declared in src/raylib.h:183
type Vector3 struct {
	gX             float32
	gY             float32
	gZ             float32
	ref5ecd5133    *C.Vector3
	allocs5ecd5133 interface{}
	This           *vector3
}
type vector3 struct {
	X float32
	Y float32
	Z float32
}

// Color as declared in src/raylib.h:210
type Color struct {
	gR             byte
	gG             byte
	gB             byte
	gA             byte
	refa79767ed    *C.Color
	allocsa79767ed interface{}
	This           *color
}
type color struct {
	R byte
	G byte
	B byte
	A byte
}

// Rectangle as declared in src/raylib.h:218
type Rectangle struct {
	gX             float32
	gY             float32
	gWidth         float32
	gHeight        float32
	refcee8783a    *C.Rectangle
	allocscee8783a interface{}
	This           *rectangle
}
type rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Texture2D as declared in src/raylib.h:238
type Texture2D struct {
	gId            uint32
	gWidth         int32
	gHeight        int32
	gMipmaps       int32
	gFormat        int32
	ref3c51a40b    *C.Texture2D
	allocs3c51a40b interface{}
	This           *texture2D
}
type texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// Font as declared in src/raylib.h:283
type Font struct {
	gBaseSize      int32
	gCharsCount    int32
	gTexture       Texture2D
	gRecs          []Rectangle
	ref70a6a7ec    *C.Font
	allocs70a6a7ec interface{}
	This           *font
}
type font struct {
	BaseSize   int32
	CharsCount int32
	Texture    C.Texture2D
	Recs       *C.Rectangle
}

// Text as declared in src/raylib.h:467
type Text struct {
	gText          []string
	ref9bb908f9    *C.Text
	allocs9bb908f9 interface{}
	This           *text
}
type text struct {
	Text []string
}
