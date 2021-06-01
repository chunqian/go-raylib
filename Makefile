all: gen

gen:
# 	../c-for-go/c-for-go -out . raylib.yml
# 	../c-for-go/c-for-go -out . physac.yml
# 	../c-for-go/c-for-go -out . raygui.yml
	../c-for-go/c-for-go -out . rres.yml

clean:
# 	# raylib
# 	rm -f raylib/cgo_helpers.go raylib/cgo_helpers.h raylib/cgo_helpers.c
# 	rm -f raylib/doc.go raylib/types.go raylib/const.go
# 	rm -f raylib/raylib.go
# 	# physac
# 	rm -f physac/cgo_helpers.go physac/cgo_helpers.h physac/cgo_helpers.c
# 	rm -f physac/doc.go physac/types.go physac/const.go
# 	rm -f physac/physac.go
# 	# raygui
# 	rm -f raygui/cgo_helpers.go raygui/cgo_helpers.h raygui/cgo_helpers.c
# 	rm -f raygui/doc.go raygui/types.go raygui/const.go
# 	rm -f raygui/raygui.go
	# rres
	rm -f rres/cgo_helpers.go rres/cgo_helpers.h rres/cgo_helpers.c
	rm -f rres/doc.go rres/types.go rres/const.go
	rm -f rres/rres.go

test:
	go build -o examples/bin/audio_module_playing examples/audio/module_playing/module_playing.go
	go build -o examples/bin/audio_multichannel_sound examples/audio/multichannel_sound/multichannel_sound.go
	go build -o examples/bin/audio_music_stream examples/audio/music_stream/music_stream.go
	go build -o examples/bin/audio_raw_stream examples/audio/raw_stream/raw_stream.go
	go build -o examples/bin/audio_sound_loading examples/audio/sound_loading/sound_loading.go

	go build -o examples/bin/core_2d_camera examples/core/2d_camera/2d_camera.go
	go build -o examples/bin/core_2d_camera_platformer examples/core/2d_camera_platformer/2d_camera_platformer.go
	go build -o examples/bin/core_3d_camera_first_person examples/core/3d_camera_first_person/3d_camera_first_person.go
	go build -o examples/bin/core_3d_camera_free examples/core/3d_camera_free/3d_camera_free.go
	go build -o examples/bin/core_3d_camera_mode examples/core/3d_camera_mode/3d_camera_mode.go
	go build -o examples/bin/core_3d_picking examples/core/3d_picking/3d_picking.go
	go build -o examples/bin/core_basic_window examples/core/basic_window/basic_window.go
	go build -o examples/bin/core_drop_files examples/core/drop_files/drop_files.go
	go build -o examples/bin/core_input_gestures examples/core/input_gestures/input_gestures.go
	go build -o examples/bin/core_input_keys examples/core/input_keys/input_keys.go
	go build -o examples/bin/core_input_mouse examples/core/input_mouse/input_mouse.go
	go build -o examples/bin/core_input_mouse_wheel examples/core/input_mouse_wheel/input_mouse_wheel.go
	go build -o examples/bin/core_input_multitouch examples/core/input_multitouch/input_multitouch.go
	go build -o examples/bin/core_random_values examples/core/random_values/random_values.go
	go build -o examples/bin/core_scissor examples/core/scissor/scissor.go
	go build -o examples/bin/core_storage_values examples/core/storage_values/storage_values.go
	go build -o examples/bin/core_vr_simulator examples/core/vr_simulator/vr_simulator.go
	go build -o examples/bin/core_window_letterbox examples/core/window_letterbox/window_letterbox.go
	go build -o examples/bin/core_world_screen examples/core/world_screen/world_screen.go

	go build -o examples/bin/gui_controls_test_suite examples/gui/controls_test_suite/controls_test_suite.go
	go build -o examples/bin/gui_scroll_panel examples/gui/scroll_panel/scroll_panel.go

	go build -o examples/bin/models_animation examples/models/animation/animation.go
	go build -o examples/bin/models_billboard examples/models/billboard/billboard.go
	go build -o examples/bin/models_box_collisions examples/models/box_collisions/box_collisions.go
	go build -o examples/bin/models_cubicmap examples/models/cubicmap/cubicmap.go
	go build -o examples/bin/models_first_person_maze examples/models/first_person_maze/first_person_maze.go
	go build -o examples/bin/models_loading examples/models/loading/loading.go
	go build -o examples/bin/models_material_pbr examples/models/material_pbr/*.go
	go build -o examples/bin/models_mesh_generation examples/models/mesh_generation/mesh_generation.go
	go build -o examples/bin/models_mesh_picking examples/models/mesh_picking/mesh_picking.go
	go build -o examples/bin/models_orthographic_projection examples/models/orthographic_projection/orthographic_projection.go
	go build -o examples/bin/models_skybox examples/models/skybox/skybox.go
	go build -o examples/bin/models_waving_cubes examples/models/waving_cubes/waving_cubes.go
	go build -o examples/bin/models_yaw_pitch_roll examples/models/yaw_pitch_roll/yaw_pitch_roll.go

	go build -o examples/bin/physac_demo examples/physac/demo/demo.go
	go build -o examples/bin/physac_friction examples/physac/friction/friction.go

	go build -o examples/bin/shaders_postprocessing examples/shaders/postprocessing/postprocessing.go
	go build -o examples/bin/shaders_basic_lighting examples/shaders/basic_lighting/*.go
	go build -o examples/bin/shaders_eratosthenes examples/shaders/eratosthenes/eratosthenes.go
	go build -o examples/bin/shaders_fog examples/shaders/fog/*.go
	go build -o examples/bin/shaders_julia_set examples/shaders/julia_set/julia_set.go
	go build -o examples/bin/shaders_model_shader examples/shaders/model_shader/model_shader.go
	go build -o examples/bin/shaders_palette_switch examples/shaders/palette_switch/palette_switch.go

	go build -o examples/bin/text_font_filters examples/text/font_filters/font_filters.go
	go build -o examples/bin/text_font_loading examples/text/font_loading/font_loading.go
	go build -o examples/bin/text_font_sdf examples/text/font_sdf/font_sdf.go
	go build -o examples/bin/text_font_spritefont examples/text/font_spritefont/font_spritefont.go
	go build -o examples/bin/text_format_text examples/text/format_text/format_text.go
	go build -o examples/bin/text_input_box examples/text/input_box/input_box.go
	go build -o examples/bin/text_raylib_fonts examples/text/raylib_fonts/raylib_fonts.go
	go build -o examples/bin/text_rectangle_bounds examples/text/rectangle_bounds/rectangle_bounds.go
	go build -o examples/bin/text_unicode examples/text/unicode/unicode.go
	go build -o examples/bin/text_writing_anim examples/text/writing_anim/writing_anim.go

	go build -o examples/bin/textures_bunnymark examples/textures/bunnymark/bunnymark.go
	go build -o examples/bin/textures_rectangle examples/textures/rectangle/rectangle.go

testwin:
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/audio_module_playing.exe examples/audio/module_playing/module_playing.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/audio_multichannel_sound.exe examples/audio/multichannel_sound/multichannel_sound.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/audio_music_stream.exe examples/audio/music_stream/music_stream.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/audio_raw_stream.exe examples/audio/raw_stream/raw_stream.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/audio_sound_loading.exe examples/audio/sound_loading/sound_loading.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_2d_camera.exe examples/core/2d_camera/2d_camera.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_2d_camera_platformer.exe examples/core/2d_camera_platformer/2d_camera_platformer.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_3d_camera_first_person.exe examples/core/3d_camera_first_person/3d_camera_first_person.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_3d_camera_free.exe examples/core/3d_camera_free/3d_camera_free.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_3d_camera_mode.exe examples/core/3d_camera_mode/3d_camera_mode.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_3d_picking.exe examples/core/3d_picking/3d_picking.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_basic_window.exe examples/core/basic_window/basic_window.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_drop_files.exe examples/core/drop_files/drop_files.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_input_gestures.exe examples/core/input_gestures/input_gestures.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_input_keys.exe examples/core/input_keys/input_keys.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_input_mouse.exe examples/core/input_mouse/input_mouse.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_input_mouse_wheel.exe examples/core/input_mouse_wheel/input_mouse_wheel.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_input_multitouch.exe examples/core/input_multitouch/input_multitouch.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_random_values.exe examples/core/random_values/random_values.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_scissor.exe examples/core/scissor/scissor.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_storage_values.exe examples/core/storage_values/storage_values.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_vr_simulator.exe examples/core/vr_simulator/vr_simulator.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_window_letterbox.exe examples/core/window_letterbox/window_letterbox.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/core_world_screen.exe examples/core/world_screen/world_screen.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/gui_controls_test_suite.exe examples/gui/controls_test_suite/controls_test_suite.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/gui_scroll_panel.exe examples/gui/scroll_panel/scroll_panel.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_animation.exe examples/models/animation/animation.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_billboard.exe examples/models/billboard/billboard.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_box_collisions.exe examples/models/box_collisions/box_collisions.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_cubicmap.exe examples/models/cubicmap/cubicmap.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_first_person_maze.exe examples/models/first_person_maze/first_person_maze.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_loading.exe examples/models/loading/loading.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_material_pbr.exe examples/models/material_pbr/*.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_mesh_generation.exe examples/models/mesh_generation/mesh_generation.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_mesh_picking.exe examples/models/mesh_picking/mesh_picking.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_orthographic_projection.exe examples/models/orthographic_projection/orthographic_projection.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_skybox.exe examples/models/skybox/skybox.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_waving_cubes.exe examples/models/waving_cubes/waving_cubes.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/models_yaw_pitch_roll.exe examples/models/yaw_pitch_roll/yaw_pitch_roll.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/physac_demo.exe examples/physac/demo/demo.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/physac_friction.exe examples/physac/friction/friction.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_postprocessing.exe examples/shaders/postprocessing/postprocessing.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_basic_lighting.exe examples/shaders/basic_lighting/*.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_eratosthenes.exe examples/shaders/eratosthenes/eratosthenes.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_fog.exe examples/shaders/fog/*.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_julia_set.exe examples/shaders/julia_set/julia_set.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_model_shader.exe examples/shaders/model_shader/model_shader.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/shaders_palette_switch.exe examples/shaders/palette_switch/palette_switch.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_font_filters.exe examples/text/font_filters/font_filters.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_font_loading.exe examples/text/font_loading/font_loading.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_font_sdf.exe examples/text/font_sdf/font_sdf.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_font_spritefont.exe examples/text/font_spritefont/font_spritefont.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_format_text.exe examples/text/format_text/format_text.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_input_box.exe examples/text/input_box/input_box.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_raylib_fonts.exe examples/text/raylib_fonts/raylib_fonts.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_rectangle_bounds.exe examples/text/rectangle_bounds/rectangle_bounds.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_unicode.exe examples/text/unicode/unicode.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/text_writing_anim.exe examples/text/writing_anim/writing_anim.go

	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/textures_bunnymark.exe examples/textures/bunnymark/bunnymark.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ go build -o examples/bin/textures_rectangle.exe examples/textures/rectangle/rectangle.go
