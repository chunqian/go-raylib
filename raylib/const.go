// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 31 Aug 2020 21:42:41 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raylib

/*
#include "../lib/raylib/src/raylib.h"
#include "../lib/raylib/src/raymath.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// PI as defined in src/raylib.h:96
	PI = 3.14159265358979323846
	// DEG2RAD as defined in src/raylib.h:99
	DEG2RAD = (PI / 180.0)
	// RAD2DEG as defined in src/raylib.h:100
	RAD2DEG = (180.0 / PI)
	// LOC_MAP_DIFFUSE as defined in src/raylib.h:714
	LOC_MAP_DIFFUSE = 0
	// LOC_MAP_SPECULAR as defined in src/raylib.h:715
	LOC_MAP_SPECULAR = 0
	// MAP_DIFFUSE as defined in src/raylib.h:745
	MAP_DIFFUSE = 0
	// MAP_SPECULAR as defined in src/raylib.h:746
	MAP_SPECULAR = 0
)

// ConfigFlag as declared in src/raylib.h:484
type ConfigFlag int32

// ConfigFlag enumeration from src/raylib.h:484
const (
	FLAG_RESERVED           ConfigFlag = 1
	FLAG_FULLSCREEN_MODE    ConfigFlag = 2
	FLAG_WINDOW_RESIZABLE   ConfigFlag = 4
	FLAG_WINDOW_UNDECORATED ConfigFlag = 8
	FLAG_WINDOW_TRANSPARENT ConfigFlag = 16
	FLAG_WINDOW_HIDDEN      ConfigFlag = 128
	FLAG_WINDOW_ALWAYS_RUN  ConfigFlag = 256
	FLAG_MSAA_4X_HINT       ConfigFlag = 32
	FLAG_VSYNC_HINT         ConfigFlag = 64
)

// KeyboardKey as declared in src/raylib.h:610
type KeyboardKey int32

// KeyboardKey enumeration from src/raylib.h:610
const (
	KEY_APOSTROPHE    KeyboardKey = 39
	KEY_COMMA         KeyboardKey = 44
	KEY_MINUS         KeyboardKey = 45
	KEY_PERIOD        KeyboardKey = 46
	KEY_SLASH         KeyboardKey = 47
	KEY_ZERO          KeyboardKey = 48
	KEY_ONE           KeyboardKey = 49
	KEY_TWO           KeyboardKey = 50
	KEY_THREE         KeyboardKey = 51
	KEY_FOUR          KeyboardKey = 52
	KEY_FIVE          KeyboardKey = 53
	KEY_SIX           KeyboardKey = 54
	KEY_SEVEN         KeyboardKey = 55
	KEY_EIGHT         KeyboardKey = 56
	KEY_NINE          KeyboardKey = 57
	KEY_SEMICOLON     KeyboardKey = 59
	KEY_EQUAL         KeyboardKey = 61
	KEY_A             KeyboardKey = 65
	KEY_B             KeyboardKey = 66
	KEY_C             KeyboardKey = 67
	KEY_D             KeyboardKey = 68
	KEY_E             KeyboardKey = 69
	KEY_F             KeyboardKey = 70
	KEY_G             KeyboardKey = 71
	KEY_H             KeyboardKey = 72
	KEY_I             KeyboardKey = 73
	KEY_J             KeyboardKey = 74
	KEY_K             KeyboardKey = 75
	KEY_L             KeyboardKey = 76
	KEY_M             KeyboardKey = 77
	KEY_N             KeyboardKey = 78
	KEY_O             KeyboardKey = 79
	KEY_P             KeyboardKey = 80
	KEY_Q             KeyboardKey = 81
	KEY_R             KeyboardKey = 82
	KEY_S             KeyboardKey = 83
	KEY_T             KeyboardKey = 84
	KEY_U             KeyboardKey = 85
	KEY_V             KeyboardKey = 86
	KEY_W             KeyboardKey = 87
	KEY_X             KeyboardKey = 88
	KEY_Y             KeyboardKey = 89
	KEY_Z             KeyboardKey = 90
	KEY_SPACE         KeyboardKey = 32
	KEY_ESCAPE        KeyboardKey = 256
	KEY_ENTER         KeyboardKey = 257
	KEY_TAB           KeyboardKey = 258
	KEY_BACKSPACE     KeyboardKey = 259
	KEY_INSERT        KeyboardKey = 260
	KEY_DELETE        KeyboardKey = 261
	KEY_RIGHT         KeyboardKey = 262
	KEY_LEFT          KeyboardKey = 263
	KEY_DOWN          KeyboardKey = 264
	KEY_UP            KeyboardKey = 265
	KEY_PAGE_UP       KeyboardKey = 266
	KEY_PAGE_DOWN     KeyboardKey = 267
	KEY_HOME          KeyboardKey = 268
	KEY_END           KeyboardKey = 269
	KEY_CAPS_LOCK     KeyboardKey = 280
	KEY_SCROLL_LOCK   KeyboardKey = 281
	KEY_NUM_LOCK      KeyboardKey = 282
	KEY_PRINT_SCREEN  KeyboardKey = 283
	KEY_PAUSE         KeyboardKey = 284
	KEY_F1            KeyboardKey = 290
	KEY_F2            KeyboardKey = 291
	KEY_F3            KeyboardKey = 292
	KEY_F4            KeyboardKey = 293
	KEY_F5            KeyboardKey = 294
	KEY_F6            KeyboardKey = 295
	KEY_F7            KeyboardKey = 296
	KEY_F8            KeyboardKey = 297
	KEY_F9            KeyboardKey = 298
	KEY_F10           KeyboardKey = 299
	KEY_F11           KeyboardKey = 300
	KEY_F12           KeyboardKey = 301
	KEY_LEFT_SHIFT    KeyboardKey = 340
	KEY_LEFT_CONTROL  KeyboardKey = 341
	KEY_LEFT_ALT      KeyboardKey = 342
	KEY_LEFT_SUPER    KeyboardKey = 343
	KEY_RIGHT_SHIFT   KeyboardKey = 344
	KEY_RIGHT_CONTROL KeyboardKey = 345
	KEY_RIGHT_ALT     KeyboardKey = 346
	KEY_RIGHT_SUPER   KeyboardKey = 347
	KEY_KB_MENU       KeyboardKey = 348
	KEY_LEFT_BRACKET  KeyboardKey = 91
	KEY_BACKSLASH     KeyboardKey = 92
	KEY_RIGHT_BRACKET KeyboardKey = 93
	KEY_GRAVE         KeyboardKey = 96
	KEY_KP_0          KeyboardKey = 320
	KEY_KP_1          KeyboardKey = 321
	KEY_KP_2          KeyboardKey = 322
	KEY_KP_3          KeyboardKey = 323
	KEY_KP_4          KeyboardKey = 324
	KEY_KP_5          KeyboardKey = 325
	KEY_KP_6          KeyboardKey = 326
	KEY_KP_7          KeyboardKey = 327
	KEY_KP_8          KeyboardKey = 328
	KEY_KP_9          KeyboardKey = 329
	KEY_KP_DECIMAL    KeyboardKey = 330
	KEY_KP_DIVIDE     KeyboardKey = 331
	KEY_KP_MULTIPLY   KeyboardKey = 332
	KEY_KP_SUBTRACT   KeyboardKey = 333
	KEY_KP_ADD        KeyboardKey = 334
	KEY_KP_ENTER      KeyboardKey = 335
	KEY_KP_EQUAL      KeyboardKey = 336
)

// AndroidButton as declared in src/raylib.h:618
type AndroidButton int32

// AndroidButton enumeration from src/raylib.h:618
const (
	KEY_BACK        AndroidButton = 4
	KEY_MENU        AndroidButton = 82
	KEY_VOLUME_UP   AndroidButton = 24
	KEY_VOLUME_DOWN AndroidButton = 25
)

// MouseButton as declared in src/raylib.h:625
type MouseButton int32

// MouseButton enumeration from src/raylib.h:625
const (
	MOUSE_LEFT_BUTTON   MouseButton = iota
	MOUSE_RIGHT_BUTTON  MouseButton = 1
	MOUSE_MIDDLE_BUTTON MouseButton = 2
)

// GamepadNumber as declared in src/raylib.h:633
type GamepadNumber int32

// GamepadNumber enumeration from src/raylib.h:633
const (
	GAMEPAD_PLAYER1 GamepadNumber = iota
	GAMEPAD_PLAYER2 GamepadNumber = 1
	GAMEPAD_PLAYER3 GamepadNumber = 2
	GAMEPAD_PLAYER4 GamepadNumber = 3
)

// GamepadButton as declared in src/raylib.h:669
type GamepadButton int32

// GamepadButton enumeration from src/raylib.h:669
const (
	GAMEPAD_BUTTON_UNKNOWN          GamepadButton = iota
	GAMEPAD_BUTTON_LEFT_FACE_UP     GamepadButton = 1
	GAMEPAD_BUTTON_LEFT_FACE_RIGHT  GamepadButton = 2
	GAMEPAD_BUTTON_LEFT_FACE_DOWN   GamepadButton = 3
	GAMEPAD_BUTTON_LEFT_FACE_LEFT   GamepadButton = 4
	GAMEPAD_BUTTON_RIGHT_FACE_UP    GamepadButton = 5
	GAMEPAD_BUTTON_RIGHT_FACE_RIGHT GamepadButton = 6
	GAMEPAD_BUTTON_RIGHT_FACE_DOWN  GamepadButton = 7
	GAMEPAD_BUTTON_RIGHT_FACE_LEFT  GamepadButton = 8
	GAMEPAD_BUTTON_LEFT_TRIGGER_1   GamepadButton = 9
	GAMEPAD_BUTTON_LEFT_TRIGGER_2   GamepadButton = 10
	GAMEPAD_BUTTON_RIGHT_TRIGGER_1  GamepadButton = 11
	GAMEPAD_BUTTON_RIGHT_TRIGGER_2  GamepadButton = 12
	GAMEPAD_BUTTON_MIDDLE_LEFT      GamepadButton = 13
	GAMEPAD_BUTTON_MIDDLE           GamepadButton = 14
	GAMEPAD_BUTTON_MIDDLE_RIGHT     GamepadButton = 15
	GAMEPAD_BUTTON_LEFT_THUMB       GamepadButton = 16
	GAMEPAD_BUTTON_RIGHT_THUMB      GamepadButton = 17
)

// GamepadAxis as declared in src/raylib.h:683
type GamepadAxis int32

// GamepadAxis enumeration from src/raylib.h:683
const (
	GAMEPAD_AXIS_LEFT_X        GamepadAxis = iota
	GAMEPAD_AXIS_LEFT_Y        GamepadAxis = 1
	GAMEPAD_AXIS_RIGHT_X       GamepadAxis = 2
	GAMEPAD_AXIS_RIGHT_Y       GamepadAxis = 3
	GAMEPAD_AXIS_LEFT_TRIGGER  GamepadAxis = 4
	GAMEPAD_AXIS_RIGHT_TRIGGER GamepadAxis = 5
)

// ShaderLocationIndex as declared in src/raylib.h:712
type ShaderLocationIndex int32

// ShaderLocationIndex enumeration from src/raylib.h:712
const (
	LOC_VERTEX_POSITION   ShaderLocationIndex = iota
	LOC_VERTEX_TEXCOORD01 ShaderLocationIndex = 1
	LOC_VERTEX_TEXCOORD02 ShaderLocationIndex = 2
	LOC_VERTEX_NORMAL     ShaderLocationIndex = 3
	LOC_VERTEX_TANGENT    ShaderLocationIndex = 4
	LOC_VERTEX_COLOR      ShaderLocationIndex = 5
	LOC_MATRIX_MVP        ShaderLocationIndex = 6
	LOC_MATRIX_MODEL      ShaderLocationIndex = 7
	LOC_MATRIX_VIEW       ShaderLocationIndex = 8
	LOC_MATRIX_PROJECTION ShaderLocationIndex = 9
	LOC_VECTOR_VIEW       ShaderLocationIndex = 10
	LOC_COLOR_DIFFUSE     ShaderLocationIndex = 11
	LOC_COLOR_SPECULAR    ShaderLocationIndex = 12
	LOC_COLOR_AMBIENT     ShaderLocationIndex = 13
	LOC_MAP_ALBEDO        ShaderLocationIndex = 14
	LOC_MAP_METALNESS     ShaderLocationIndex = 15
	LOC_MAP_NORMAL        ShaderLocationIndex = 16
	LOC_MAP_ROUGHNESS     ShaderLocationIndex = 17
	LOC_MAP_OCCLUSION     ShaderLocationIndex = 18
	LOC_MAP_EMISSION      ShaderLocationIndex = 19
	LOC_MAP_HEIGHT        ShaderLocationIndex = 20
	LOC_MAP_CUBEMAP       ShaderLocationIndex = 21
	LOC_MAP_IRRADIANCE    ShaderLocationIndex = 22
	LOC_MAP_PREFILTER     ShaderLocationIndex = 23
	LOC_MAP_BRDF          ShaderLocationIndex = 24
)

// ShaderUniformDataType as declared in src/raylib.h:728
type ShaderUniformDataType int32

// ShaderUniformDataType enumeration from src/raylib.h:728
const (
	UNIFORM_FLOAT     ShaderUniformDataType = iota
	UNIFORM_VEC2      ShaderUniformDataType = 1
	UNIFORM_VEC3      ShaderUniformDataType = 2
	UNIFORM_VEC4      ShaderUniformDataType = 3
	UNIFORM_INT       ShaderUniformDataType = 4
	UNIFORM_IVEC2     ShaderUniformDataType = 5
	UNIFORM_IVEC3     ShaderUniformDataType = 6
	UNIFORM_IVEC4     ShaderUniformDataType = 7
	UNIFORM_SAMPLER2D ShaderUniformDataType = 8
)

// MaterialMapType as declared in src/raylib.h:743
type MaterialMapType int32

// MaterialMapType enumeration from src/raylib.h:743
const (
	MAP_ALBEDO     MaterialMapType = iota
	MAP_METALNESS  MaterialMapType = 1
	MAP_NORMAL     MaterialMapType = 2
	MAP_ROUGHNESS  MaterialMapType = 3
	MAP_OCCLUSION  MaterialMapType = 4
	MAP_EMISSION   MaterialMapType = 5
	MAP_HEIGHT     MaterialMapType = 6
	MAP_CUBEMAP    MaterialMapType = 7
	MAP_IRRADIANCE MaterialMapType = 8
	MAP_PREFILTER  MaterialMapType = 9
	MAP_BRDF       MaterialMapType = 10
)

// PixelFormat as declared in src/raylib.h:772
type PixelFormat int32

// PixelFormat enumeration from src/raylib.h:772
const (
	UNCOMPRESSED_GRAYSCALE    PixelFormat = 1
	UNCOMPRESSED_GRAY_ALPHA   PixelFormat = 2
	UNCOMPRESSED_R5G6B5       PixelFormat = 3
	UNCOMPRESSED_R8G8B8       PixelFormat = 4
	UNCOMPRESSED_R5G5B5A1     PixelFormat = 5
	UNCOMPRESSED_R4G4B4A4     PixelFormat = 6
	UNCOMPRESSED_R8G8B8A8     PixelFormat = 7
	UNCOMPRESSED_R32          PixelFormat = 8
	UNCOMPRESSED_R32G32B32    PixelFormat = 9
	UNCOMPRESSED_R32G32B32A32 PixelFormat = 10
	COMPRESSED_DXT1_RGB       PixelFormat = 11
	COMPRESSED_DXT1_RGBA      PixelFormat = 12
	COMPRESSED_DXT3_RGBA      PixelFormat = 13
	COMPRESSED_DXT5_RGBA      PixelFormat = 14
	COMPRESSED_ETC1_RGB       PixelFormat = 15
	COMPRESSED_ETC2_RGB       PixelFormat = 16
	COMPRESSED_ETC2_EAC_RGBA  PixelFormat = 17
	COMPRESSED_PVRT_RGB       PixelFormat = 18
	COMPRESSED_PVRT_RGBA      PixelFormat = 19
	COMPRESSED_ASTC_4x4_RGBA  PixelFormat = 20
	COMPRESSED_ASTC_8x8_RGBA  PixelFormat = 21
)

// TextureFilterMode as declared in src/raylib.h:784
type TextureFilterMode int32

// TextureFilterMode enumeration from src/raylib.h:784
const (
	FILTER_POINT           TextureFilterMode = iota
	FILTER_BILINEAR        TextureFilterMode = 1
	FILTER_TRILINEAR       TextureFilterMode = 2
	FILTER_ANISOTROPIC_4X  TextureFilterMode = 3
	FILTER_ANISOTROPIC_8X  TextureFilterMode = 4
	FILTER_ANISOTROPIC_16X TextureFilterMode = 5
)

// CubemapLayoutType as declared in src/raylib.h:794
type CubemapLayoutType int32

// CubemapLayoutType enumeration from src/raylib.h:794
const (
	CUBEMAP_AUTO_DETECT         CubemapLayoutType = iota
	CUBEMAP_LINE_VERTICAL       CubemapLayoutType = 1
	CUBEMAP_LINE_HORIZONTAL     CubemapLayoutType = 2
	CUBEMAP_CROSS_THREE_BY_FOUR CubemapLayoutType = 3
	CUBEMAP_CROSS_FOUR_BY_THREE CubemapLayoutType = 4
	CUBEMAP_PANORAMA            CubemapLayoutType = 5
)

// TextureWrapMode as declared in src/raylib.h:802
type TextureWrapMode int32

// TextureWrapMode enumeration from src/raylib.h:802
const (
	WRAP_REPEAT        TextureWrapMode = iota
	WRAP_CLAMP         TextureWrapMode = 1
	WRAP_MIRROR_REPEAT TextureWrapMode = 2
	WRAP_MIRROR_CLAMP  TextureWrapMode = 3
)

// FontType as declared in src/raylib.h:809
type FontType int32

// FontType enumeration from src/raylib.h:809
const (
	FONT_DEFAULT FontType = iota
	FONT_BITMAP  FontType = 1
	FONT_SDF     FontType = 2
)

// BlendMode as declared in src/raylib.h:817
type BlendMode int32

// BlendMode enumeration from src/raylib.h:817
const (
	BLEND_ALPHA      BlendMode = iota
	BLEND_ADDITIVE   BlendMode = 1
	BLEND_MULTIPLIED BlendMode = 2
	BLEND_ADD_COLORS BlendMode = 3
)

// GestureType as declared in src/raylib.h:833
type GestureType int32

// GestureType enumeration from src/raylib.h:833
const (
	GESTURE_NONE        GestureType = iota
	GESTURE_TAP         GestureType = 1
	GESTURE_DOUBLETAP   GestureType = 2
	GESTURE_HOLD        GestureType = 4
	GESTURE_DRAG        GestureType = 8
	GESTURE_SWIPE_RIGHT GestureType = 16
	GESTURE_SWIPE_LEFT  GestureType = 32
	GESTURE_SWIPE_UP    GestureType = 64
	GESTURE_SWIPE_DOWN  GestureType = 128
	GESTURE_PINCH_IN    GestureType = 256
	GESTURE_PINCH_OUT   GestureType = 512
)

// CameraMode as declared in src/raylib.h:842
type CameraMode int32

// CameraMode enumeration from src/raylib.h:842
const (
	CAMERA_CUSTOM       CameraMode = iota
	CAMERA_FREE         CameraMode = 1
	CAMERA_ORBITAL      CameraMode = 2
	CAMERA_FIRST_PERSON CameraMode = 3
	CAMERA_THIRD_PERSON CameraMode = 4
)

// CameraType as declared in src/raylib.h:848
type CameraType int32

// CameraType enumeration from src/raylib.h:848
const (
	CAMERA_PERSPECTIVE  CameraType = iota
	CAMERA_ORTHOGRAPHIC CameraType = 1
)

// NPatchType as declared in src/raylib.h:855
type NPatchType int32

// NPatchType enumeration from src/raylib.h:855
const (
	NPT_9PATCH            NPatchType = iota
	NPT_3PATCH_VERTICAL   NPatchType = 1
	NPT_3PATCH_HORIZONTAL NPatchType = 2
)
