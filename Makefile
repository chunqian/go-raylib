all: gen
	
gen:
	../c-for-go/c-for-go -out . raylib.yml
	../c-for-go/c-for-go -out . physac.yml
	../c-for-go/c-for-go -out . raygui.yml

clean:
	# raylib
	rm -f raylib/cgo_helpers.go raylib/cgo_helpers.h raylib/cgo_helpers.c
	rm -f raylib/doc.go raylib/types.go raylib/const.go
	rm -f raylib/raylib.go
	# physac
	rm -f physac/cgo_helpers.go physac/cgo_helpers.h physac/cgo_helpers.c
	rm -f physac/doc.go physac/types.go physac/const.go
	rm -f physac/physac.go
	# raygui
	rm -f raygui/cgo_helpers.go raygui/cgo_helpers.h raygui/cgo_helpers.c
	rm -f raygui/doc.go raygui/types.go raygui/const.go
	rm -f raygui/raygui.go

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
	go build -o examples/bin/physac_demo examples/physac/demo/demo.go
	go build -o examples/bin/shaders_postprocessing examples/shaders/postprocessing/postprocessing.go
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
