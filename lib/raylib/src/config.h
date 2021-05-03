/**********************************************************************************************
*
*   raylib configuration flags
*
*   This file defines all the configuration flags for the different raylib modules
*
*   LICENSE: zlib/libpng
*
*   Copyright (c) 2018-2021 Ahmad Fatoum & Ramon Santamaria (@raysan5)
*
*   This software is provided "as-is", without any express or implied warranty. In no event
*   will the authors be held liable for any damages arising from the use of this software.
*
*   Permission is granted to anyone to use this software for any purpose, including commercial
*   applications, and to alter it and redistribute it freely, subject to the following restrictions:
*
*     1. The origin of this software must not be misrepresented; you must not claim that you
*     wrote the original software. If you use this software in a product, an acknowledgment
*     in the product documentation would be appreciated but is not required.
*
*     2. Altered source versions must be plainly marked as such, and must not be misrepresented
*     as being the original software.
*
*     3. This notice may not be removed or altered from any source distribution.
*
**********************************************************************************************/
#ifndef RL_CONFIG_H
#define RL_CONFIG_H

#define RAYLIB_VERSION  "3.5"

// Edit to control what features Makefile'd raylib is compiled with
// cforgo
// #if defined(RAYLIB_CMAKE)
//     // Edit CMakeOptions.txt for CMake instead
//     #include "cmake/config.h"
// #else

//------------------------------------------------------------------------------------
// Module: core - Configuration Flags
//------------------------------------------------------------------------------------
// Camera module is included (camera.h) and multiple predefined cameras are available: free, 1st/3rd person, orbital
#ifndef SUPPORT_CAMERA_SYSTEM
#define SUPPORT_CAMERA_SYSTEM       1
#endif
// Gestures module is included (gestures.h) to support gestures detection: tap, hold, swipe, drag
#ifndef SUPPORT_GESTURES_SYSTEM
#define SUPPORT_GESTURES_SYSTEM     1
#endif
// Mouse gestures are directly mapped like touches and processed by gestures system
#ifndef SUPPORT_MOUSE_GESTURES
#define SUPPORT_MOUSE_GESTURES      1
#endif
// Reconfigure standard input to receive key inputs, works with SSH connection.
#ifndef SUPPORT_SSH_KEYBOARD_RPI
#define SUPPORT_SSH_KEYBOARD_RPI    1
#endif
// Draw a mouse pointer on screen
#ifndef SUPPORT_MOUSE_CURSOR_NATIVE
#define SUPPORT_MOUSE_CURSOR_NATIVE 1
#endif
// Use busy wait loop for timing sync, if not defined, a high-resolution timer is setup and used
//#define SUPPORT_BUSY_WAIT_LOOP      1
// Use a half-busy wait loop, in this case frame sleeps for some time and runs a busy-wait-loop at the end
#define SUPPORT_HALFBUSY_WAIT_LOOP
// Wait for events passively (sleeping while no events) instead of polling them actively every frame
//#define SUPPORT_EVENTS_WAITING      1
// Allow automatic screen capture of current screen pressing F12, defined in KeyCallback()
#ifndef SUPPORT_SCREEN_CAPTURE
#define SUPPORT_SCREEN_CAPTURE      1
#endif
// Allow automatic gif recording of current screen pressing CTRL+F12, defined in KeyCallback()
#ifndef SUPPORT_GIF_RECORDING
#define SUPPORT_GIF_RECORDING       1
#endif
// Support CompressData() and DecompressData() functions
#ifndef SUPPORT_COMPRESSION_API
#define SUPPORT_COMPRESSION_API     1
#endif
// Support saving binary data automatically to a generated storage.data file. This file is managed internally.
#ifndef SUPPORT_DATA_STORAGE
#define SUPPORT_DATA_STORAGE        1
#endif

// core: Configuration values
//------------------------------------------------------------------------------------
#if defined(__linux__)
    #ifndef MAX_FILEPATH_LENGTH
    #define MAX_FILEPATH_LENGTH     4096        // Maximum length for filepaths (Linux PATH_MAX default value)
    #endif
#else
    #ifndef MAX_FILEPATH_LENGTH
    #define MAX_FILEPATH_LENGTH      512        // Maximum length supported for filepaths
    #endif
#endif

#ifndef MAX_GAMEPADS
#define MAX_GAMEPADS                   4        // Max number of gamepads supported
#endif
#ifndef MAX_GAMEPAD_AXIS
#define MAX_GAMEPAD_AXIS               8        // Max number of axis supported (per gamepad)
#endif
#ifndef MAX_GAMEPAD_BUTTONS
#define MAX_GAMEPAD_BUTTONS           32        // Max bumber of buttons supported (per gamepad)
#endif
#ifndef MAX_TOUCH_POINTS
#define MAX_TOUCH_POINTS              10        // Maximum number of touch points supported
#endif
#ifndef MAX_KEY_PRESSED_QUEUE
#define MAX_KEY_PRESSED_QUEUE         16        // Max number of characters in the key input queue
#endif

#ifndef STORAGE_DATA_FILE
#define STORAGE_DATA_FILE  "storage.data"       // Automatic storage filename
#endif


//------------------------------------------------------------------------------------
// Module: rlgl - Configuration Flags
//------------------------------------------------------------------------------------
// Support VR simulation functionality (stereo rendering)
#ifndef SUPPORT_VR_SIMULATOR
#define SUPPORT_VR_SIMULATOR        1
#endif

// rlgl: Configuration values
//------------------------------------------------------------------------------------
#if defined(GRAPHICS_API_OPENGL_11) || defined(GRAPHICS_API_OPENGL_33)
    #ifndef DEFAULT_BATCH_BUFFER_ELEMENTS
    #define DEFAULT_BATCH_BUFFER_ELEMENTS   8192    // Default internal render batch limits
    #endif
#elif defined(GRAPHICS_API_OPENGL_ES2)
    #ifndef DEFAULT_BATCH_BUFFER_ELEMENTS
    #define DEFAULT_BATCH_BUFFER_ELEMENTS   2048    // Default internal render batch limits
    #endif
#endif

#ifndef DEFAULT_BATCH_BUFFERS
#define DEFAULT_BATCH_BUFFERS            1      // Default number of batch buffers (multi-buffering)
#endif
#ifndef DEFAULT_BATCH_DRAWCALLS
#define DEFAULT_BATCH_DRAWCALLS        256      // Default number of batch draw calls (by state changes: mode, texture)
#endif

#ifndef MAX_MATRIX_STACK_SIZE
#define MAX_MATRIX_STACK_SIZE           32      // Maximum size of internal Matrix stack
#endif
#ifndef MAX_SHADER_LOCATIONS
#define MAX_SHADER_LOCATIONS            32      // Maximum number of shader locations supported
#endif
#ifndef MAX_MATERIAL_MAPS
#define MAX_MATERIAL_MAPS               12      // Maximum number of shader maps supported
#endif

#ifndef RL_CULL_DISTANCE_NEAR
#define RL_CULL_DISTANCE_NEAR         0.01      // Default projection matrix near cull distance
#endif
#ifndef RL_CULL_DISTANCE_FAR
#define RL_CULL_DISTANCE_FAR        1000.0      // Default projection matrix far cull distance
#endif

// Default shader vertex attribute names to set location points
#ifndef DEFAULT_SHADER_ATTRIB_NAME_POSITION
#define DEFAULT_SHADER_ATTRIB_NAME_POSITION    "vertexPosition"    // Binded by default to shader location: 0
#endif
#ifndef DEFAULT_SHADER_ATTRIB_NAME_TEXCOORD
#define DEFAULT_SHADER_ATTRIB_NAME_TEXCOORD    "vertexTexCoord"    // Binded by default to shader location: 1
#endif
#ifndef DEFAULT_SHADER_ATTRIB_NAME_NORMAL
#define DEFAULT_SHADER_ATTRIB_NAME_NORMAL      "vertexNormal"      // Binded by default to shader location: 2
#endif
#ifndef DEFAULT_SHADER_ATTRIB_NAME_COLOR
#define DEFAULT_SHADER_ATTRIB_NAME_COLOR       "vertexColor"       // Binded by default to shader location: 3
#endif
#ifndef DEFAULT_SHADER_ATTRIB_NAME_TANGENT
#define DEFAULT_SHADER_ATTRIB_NAME_TANGENT     "vertexTangent"     // Binded by default to shader location: 4
#endif
#ifndef DEFAULT_SHADER_ATTRIB_NAME_TEXCOORD2
#define DEFAULT_SHADER_ATTRIB_NAME_TEXCOORD2   "vertexTexCoord2"   // Binded by default to shader location: 5
#endif


//------------------------------------------------------------------------------------
// Module: shapes - Configuration Flags
//------------------------------------------------------------------------------------
// Draw rectangle shapes using font texture white character instead of default white texture
// Allows drawing rectangles and text with a single draw call, very useful for GUI systems!
#ifndef SUPPORT_FONT_TEXTURE
#define SUPPORT_FONT_TEXTURE        1
#endif
// Use QUADS instead of TRIANGLES for drawing when possible
// Some lines-based shapes could still use lines
#ifndef SUPPORT_QUADS_DRAW_MODE
#define SUPPORT_QUADS_DRAW_MODE     1
#endif


//------------------------------------------------------------------------------------
// Module: textures - Configuration Flags
//------------------------------------------------------------------------------------
// Selecte desired fileformats to be supported for image data loading
#ifndef SUPPORT_FILEFORMAT_PNG
#define SUPPORT_FILEFORMAT_PNG      1
#endif
//#define SUPPORT_FILEFORMAT_BMP      1
//#define SUPPORT_FILEFORMAT_TGA      1
//#define SUPPORT_FILEFORMAT_JPG      1
#ifndef SUPPORT_FILEFORMAT_GIF
#define SUPPORT_FILEFORMAT_GIF      1
#endif
//#define SUPPORT_FILEFORMAT_PSD      1
#ifndef SUPPORT_FILEFORMAT_DDS
#define SUPPORT_FILEFORMAT_DDS      1
#endif
#ifndef SUPPORT_FILEFORMAT_HDR
#define SUPPORT_FILEFORMAT_HDR      1
#endif
//#define SUPPORT_FILEFORMAT_KTX      1
//#define SUPPORT_FILEFORMAT_ASTC     1
//#define SUPPORT_FILEFORMAT_PKM      1
//#define SUPPORT_FILEFORMAT_PVR      1

// Support image export functionality (.png, .bmp, .tga, .jpg)
#ifndef SUPPORT_IMAGE_EXPORT
#define SUPPORT_IMAGE_EXPORT        1
#endif
// Support procedural image generation functionality (gradient, spot, perlin-noise, cellular)
#ifndef SUPPORT_IMAGE_GENERATION
#define SUPPORT_IMAGE_GENERATION    1
#endif
// Support multiple image editing functions to scale, adjust colors, flip, draw on images, crop...
// If not defined, still some functions are supported: ImageFormat(), ImageCrop(), ImageToPOT()
#ifndef SUPPORT_IMAGE_MANIPULATION
#define SUPPORT_IMAGE_MANIPULATION  1
#endif


//------------------------------------------------------------------------------------
// Module: text - Configuration Flags
//------------------------------------------------------------------------------------
// Default font is loaded on window initialization to be available for the user to render simple text
// NOTE: If enabled, uses external module functions to load default raylib font
#ifndef SUPPORT_DEFAULT_FONT
#define SUPPORT_DEFAULT_FONT        1
#endif
// Selected desired font fileformats to be supported for loading
#ifndef SUPPORT_FILEFORMAT_FNT
#define SUPPORT_FILEFORMAT_FNT      1
#endif
#ifndef SUPPORT_FILEFORMAT_TTF
#define SUPPORT_FILEFORMAT_TTF      1
#endif

// Support text management functions
// If not defined, still some functions are supported: TextLength(), TextFormat()
#ifndef SUPPORT_TEXT_MANIPULATION
#define SUPPORT_TEXT_MANIPULATION   1
#endif

// text: Configuration values
//------------------------------------------------------------------------------------
#ifndef MAX_TEXT_BUFFER_LENGTH
#define MAX_TEXT_BUFFER_LENGTH      1024        // Size of internal static buffers used on some functions:
#endif
                                                // TextFormat(), TextSubtext(), TextToUpper(), TextToLower(), TextToPascal(), TextSplit()
#ifndef MAX_TEXT_UNICODE_CHARS
#define MAX_TEXT_UNICODE_CHARS       512        // Maximum number of unicode codepoints: GetCodepoints()
#endif
#ifndef MAX_TEXTSPLIT_COUNT
#define MAX_TEXTSPLIT_COUNT          128        // Maximum number of substrings to split: TextSplit()
#endif


//------------------------------------------------------------------------------------
// Module: models - Configuration Flags
//------------------------------------------------------------------------------------
// Selected desired model fileformats to be supported for loading
#ifndef SUPPORT_FILEFORMAT_OBJ
#define SUPPORT_FILEFORMAT_OBJ      1
#endif
#ifndef SUPPORT_FILEFORMAT_MTL
#define SUPPORT_FILEFORMAT_MTL      1
#endif
#ifndef SUPPORT_FILEFORMAT_IQM
#define SUPPORT_FILEFORMAT_IQM      1
#endif
#ifndef SUPPORT_FILEFORMAT_GLTF
#define SUPPORT_FILEFORMAT_GLTF     1
#endif
// Support procedural mesh generation functions, uses external par_shapes.h library
// NOTE: Some generated meshes DO NOT include generated texture coordinates
#ifndef SUPPORT_MESH_GENERATION
#define SUPPORT_MESH_GENERATION     1
#endif


//------------------------------------------------------------------------------------
// Module: audio - Configuration Flags
//------------------------------------------------------------------------------------
// Desired audio fileformats to be supported for loading
#ifndef SUPPORT_FILEFORMAT_WAV
#define SUPPORT_FILEFORMAT_WAV      1
#endif
#ifndef SUPPORT_FILEFORMAT_OGG
#define SUPPORT_FILEFORMAT_OGG      1
#endif
#ifndef SUPPORT_FILEFORMAT_XM
#define SUPPORT_FILEFORMAT_XM       1
#endif
//#define SUPPORT_FILEFORMAT_MOD      1
#ifndef SUPPORT_FILEFORMAT_MP3
#define SUPPORT_FILEFORMAT_MP3      1
#endif
//#define SUPPORT_FILEFORMAT_FLAC     1

// audio: Configuration values
//------------------------------------------------------------------------------------
#ifndef AUDIO_DEVICE_FORMAT
#define AUDIO_DEVICE_FORMAT    ma_format_f32    // Device output format (miniaudio: float-32bit)
#endif
#ifndef AUDIO_DEVICE_CHANNELS
#define AUDIO_DEVICE_CHANNELS              2    // Device output channels: stereo
#endif
#ifndef AUDIO_DEVICE_SAMPLE_RATE
#define AUDIO_DEVICE_SAMPLE_RATE       44100    // Device output sample rate
#endif

#ifndef DEFAULT_AUDIO_BUFFER_SIZE
#define DEFAULT_AUDIO_BUFFER_SIZE       4096    // Default audio buffer size for streaming
#endif
#ifndef MAX_AUDIO_BUFFER_POOL_CHANNELS
#define MAX_AUDIO_BUFFER_POOL_CHANNELS    16    // Maximum number of audio pool channels
#endif

//------------------------------------------------------------------------------------
// Module: utils - Configuration Flags
//------------------------------------------------------------------------------------
// Show TRACELOG() output messages
// NOTE: By default LOG_DEBUG traces not shown
#ifndef SUPPORT_TRACELOG
#define SUPPORT_TRACELOG            1
#endif
//#define SUPPORT_TRACELOG_DEBUG      1

// utils: Configuration values
//------------------------------------------------------------------------------------
#ifndef MAX_TRACELOG_MSG_LENGTH
#define MAX_TRACELOG_MSG_LENGTH          128    // Max length of one trace-log message
#endif
#ifndef MAX_UWP_MESSAGES
#define MAX_UWP_MESSAGES                 512    // Max UWP messages to process
#endif

// cforgo
// #endif  //defined(RAYLIB_CMAKE)

#if SUPPORT_CAMERA_SYSTEM == 0
#undef SUPPORT_CAMERA_SYSTEM
#endif

#if SUPPORT_GESTURES_SYSTEM == 0
#undef SUPPORT_GESTURES_SYSTEM
#endif

#if SUPPORT_MOUSE_GESTURES == 0
#undef SUPPORT_MOUSE_GESTURES
#endif

#if SUPPORT_SSH_KEYBOARD_RPI == 0
#undef SUPPORT_SSH_KEYBOARD_RPI
#endif

#if SUPPORT_MOUSE_CURSOR_NATIVE == 0
#undef SUPPORT_MOUSE_CURSOR_NATIVE
#endif

#if SUPPORT_SCREEN_CAPTURE == 0
#undef SUPPORT_SCREEN_CAPTURE
#endif

#if SUPPORT_GIF_RECORDING == 0
#undef SUPPORT_GIF_RECORDING
#endif

#if SUPPORT_COMPRESSION_API == 0
#undef SUPPORT_COMPRESSION_API
#endif

#if SUPPORT_DATA_STORAGE == 0
#undef SUPPORT_DATA_STORAGE
#endif

#if SUPPORT_VR_SIMULATOR == 0
#undef SUPPORT_VR_SIMULATOR
#endif

#if SUPPORT_FONT_TEXTURE == 0
#undef SUPPORT_FONT_TEXTURE
#endif

#if SUPPORT_QUADS_DRAW_MODE == 0
#undef SUPPORT_QUADS_DRAW_MODE
#endif

#if SUPPORT_FILEFORMAT_PNG == 0
#undef SUPPORT_FILEFORMAT_PNG
#endif

#if SUPPORT_FILEFORMAT_GIF == 0
#undef SUPPORT_FILEFORMAT_GIF
#endif

#if SUPPORT_FILEFORMAT_DDS == 0
#undef SUPPORT_FILEFORMAT_DDS
#endif

#if SUPPORT_FILEFORMAT_HDR == 0
#undef SUPPORT_FILEFORMAT_HDR
#endif

#if SUPPORT_IMAGE_EXPORT == 0
#undef SUPPORT_IMAGE_EXPORT
#endif

#if SUPPORT_IMAGE_GENERATION == 0
#undef SUPPORT_IMAGE_GENERATION
#endif

#if SUPPORT_IMAGE_MANIPULATION == 0
#undef SUPPORT_IMAGE_MANIPULATION
#endif

#if SUPPORT_DEFAULT_FONT == 0
#undef SUPPORT_DEFAULT_FONT
#endif

#if SUPPORT_FILEFORMAT_FNT == 0
#undef SUPPORT_FILEFORMAT_FNT
#endif

#if SUPPORT_FILEFORMAT_TTF == 0
#undef SUPPORT_FILEFORMAT_TTF
#endif

#if SUPPORT_TEXT_MANIPULATION == 0
#undef SUPPORT_TEXT_MANIPULATION
#endif

#if SUPPORT_FILEFORMAT_OBJ == 0
#undef SUPPORT_FILEFORMAT_OBJ
#endif

#if SUPPORT_FILEFORMAT_MTL == 0
#undef SUPPORT_FILEFORMAT_MTL
#endif

#if SUPPORT_FILEFORMAT_IQM == 0
#undef SUPPORT_FILEFORMAT_IQM
#endif

#if SUPPORT_FILEFORMAT_GLTF == 0
#undef SUPPORT_FILEFORMAT_GLTF
#endif

#if SUPPORT_MESH_GENERATION == 0
#undef SUPPORT_MESH_GENERATION
#endif

#if SUPPORT_FILEFORMAT_WAV == 0
#undef SUPPORT_FILEFORMAT_WAV
#endif

#if SUPPORT_FILEFORMAT_OGG == 0
#undef SUPPORT_FILEFORMAT_OGG
#endif

#if SUPPORT_FILEFORMAT_XM == 0
#undef SUPPORT_FILEFORMAT_XM
#endif

#if SUPPORT_FILEFORMAT_MP3 == 0
#undef SUPPORT_FILEFORMAT_MP3
#endif

#if SUPPORT_TRACELOG == 0
#undef SUPPORT_TRACELOG
#endif

#endif // RL_CONFIG_H
