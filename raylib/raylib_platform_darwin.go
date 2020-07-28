// +build darwin

package raylib

/*
#include "../lib/raylib/src/external/glfw/src/context.c"
#include "../lib/raylib/src/external/glfw/src/init.c"
#include "../lib/raylib/src/external/glfw/src/input.c"
#include "../lib/raylib/src/external/glfw/src/monitor.c"
#include "../lib/raylib/src/external/glfw/src/vulkan.c"
#include "../lib/raylib/src/external/glfw/src/window.c"

#include "../lib/raylib/src/external/glfw/src/cocoa_init.m"
#include "../lib/raylib/src/external/glfw/src/cocoa_joystick.m"
#include "../lib/raylib/src/external/glfw/src/cocoa_monitor.m"
#include "../lib/raylib/src/external/glfw/src/cocoa_time.c"
#include "../lib/raylib/src/external/glfw/src/cocoa_window.m"
#include "../lib/raylib/src/external/glfw/src/posix_thread.c"
#include "../lib/raylib/src/external/glfw/src/nsgl_context.m"
#include "../lib/raylib/src/external/glfw/src/egl_context.c"
#include "../lib/raylib/src/external/glfw/src/osmesa_context.c"

#include "../lib/raylib/src/core.c"
#include "../lib/raylib/src/models.c"
#include "../lib/raylib/src/raudio.c"
#include "../lib/raylib/src/shapes.c"
#include "../lib/raylib/src/text.c"
#include "../lib/raylib/src/textures.c"
#include "../lib/raylib/src/utils.c"

#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo darwin CFLAGS: -x objective-c -I../lib/raylib/src/external/glfw/include -DC_FOR_GO -D_GLFW_COCOA -D_GLFW_USE_CHDIR -D_GLFW_USE_MENUBAR -D_GLFW_USE_RETINA -Wno-deprecated-declarations -DPLATFORM_DESKTOP -DMAL_NO_COREAUDIO

#cgo darwin,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo darwin,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo darwin,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33
*/
import "C"
