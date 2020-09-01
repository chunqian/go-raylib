# go-raylib

### Bindings

Header     | Supported          |
---------  | ------------------ |
raylib.h   | :heavy_check_mark: |
raymath.h  | :heavy_check_mark: |
raygui.h   | :heavy_check_mark: |
ricons.h   | :heavy_check_mark: |
physac.h   | :heavy_check_mark: |

### Examples

Category   | Example                | Supported          |
---------  | ---------------------- | ------------------ |
audio      | module_playing         | :heavy_check_mark: |
audio      | multichannel_sound     | :heavy_check_mark: |
audio      | music_stream           | :heavy_check_mark: |
audio      | raw_stream             | :heavy_check_mark: |
audio      | sound_loading          | :heavy_check_mark: |
core       | 2d_camera              | :heavy_check_mark: |
core       | 2d_camera_platformer   | :heavy_check_mark: |
core       | 3d_camera_first_person | :heavy_check_mark: |
core       | 3d_camera_free         | :heavy_check_mark: |
core       | 3d_camera_mode         | :heavy_check_mark: |
core       | 3d_picking             | :heavy_check_mark: |
core       | basic_window           | :heavy_check_mark: |
core       | drop_files             | :heavy_check_mark: |
core       | input_gestures         | :heavy_check_mark: |
core       | input_keys             | :heavy_check_mark: |
gui        | controls_test_suite    | :heavy_check_mark: |
gui        | scroll_panel           | :heavy_check_mark: |
models     | animation              | :heavy_check_mark: |
models     | cubicmap               | :heavy_check_mark: |
models     | loading                | :heavy_check_mark: |
physac     | demo                   | :heavy_check_mark: |
shaders    | postprocessing         | :heavy_check_mark: |
text       | font_sdf               | :heavy_check_mark: |
text       | raylib_fonts           | :heavy_check_mark: |
text       | unicode                | :heavy_check_mark: |
textures   | bunnymark              | :heavy_check_mark: |

### Version
Synchronized update with [raylib](https://github.com/raysan5/raylib). Current version is 3.1dev.


### Build

```bash
$ git clone https://github.com/chunqian/go-raylib.git
$ cd go-raylib/lib/raylib/src
$ export MACOSX_DEPLOYMENT_TARGET=10.9
$ make
$ cd ../../..
$ go build -o examples/audio/module_playing/module_playing examples/audio/module_playing/module_playing.go
```

### Run

```bash
$ cd examples/audio/module_playing
$ ./module_playing
```
