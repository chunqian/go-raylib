# go-raylib


### Bindings
Header     | Supported          |
---------  | ------------------ |
raylib.h   | :heavy_check_mark: |
raymath.h  | :heavy_check_mark: |
raygui.h   | :heavy_check_mark: |
ricons.h   | :heavy_check_mark: |
physac.h   | :heavy_check_mark: |


### Version
Synchronized update with [raylib](https://github.com/raysan5/raylib). Current version is 3.1-dev.


### Performance
Same performance as the Raylib C version. Enjoy it!.


### Memory
Automated management. You don't need to worry about memory when you write code.


### Difference
There are some differences between the processing in Go and C.

In C
```c
camera.zoom = 1.0f;
```
In Go
```go
camera.This.Zoom = 1.0
```

In C
```c
camera->offset.x = width - (max.x - width/2);
```
In Go
```go
camera.Offset().This.X = float32(width - (int32(max.This.X) - width/2))
```

In C
```c
camera->offset = (Vector2){ width/2, height/2 };
```
In Go
```go
camera.This.Offset, _ = rl.NewVector2(float32(width/2), float32(height/2)).PassValue()
```

In C
```c
const char *listViewExList[8] = { "This", "is", "a", "list view", "with", "disable", "elements", "amazing!" };
```
In Go
```go
listViewExList := rg.NewText([]string{"This", "is", "a", "list view", "with", "disable", "elements", "amazing!"})
```

In C
```c
char multiTextBoxText[256] = "Multi text box";
```
In Go
```go
multiTextBoxText := rg.NewBytes("Multi text box", 256)
```

In C
```c
int dropsCount = 0;
char **droppedFiles = GetDroppedFiles(&dropsCount);
const char *droppedFilePath = droppedFiles[0];
```
In Go
```go
dropsCount := int32(0)
droppedFiles := rl.GetDroppedFiles(&dropsCount)
droppedFilePath := rl.StringFromPPByte(droppedFiles, 0)
```

In C
```c
animFrameCounter++;
UpdateModelAnimation(model, anims[0], animFrameCounter);
if (animFrameCounter >= anims[0].frameCount) animFrameCounter = 0;
```
In Go
```go
animFrameCounter++
rl.UpdateModelAnimation(model, *anims.Index(0), animFrameCounter)
if animFrameCounter >= anims.Index(0).This.FrameCount {
    animFrameCounter = 0
}
```

In C
```c
Texture2D texture = LoadTexture("resources/cubicmap_atlas.png");
model.materials[0].maps[MAP_DIFFUSE].texture = texture;
```
In Go
```go
texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
model.Materials(0).Maps(rl.MAP_DIFFUSE).This.Texture, _ = texture.PassValue()
```

In C
```c
Font fontDefault = { 0 };
fontDefault.baseSize = 16;
fontDefault.charsCount = 95;
fontDefault.chars = LoadFontData("resources/anonymous_pro_bold.ttf", 16, 0, 95, FONT_DEFAULT);
Image atlas = GenImageFontAtlas(fontDefault.chars, &fontDefault.recs, 95, 16, 4, 0);
```
In Go
```go
fontDefault := rl.Font{}
fontDefault.PassRef()
fontDefault.This.BaseSize = 16
fontDefault.This.CharsCount = 95
fontDefault.This.Chars, _ = rl.LoadFontData("../text/resources/anonymous_pro_bold.ttf", 16, nil, 95, int32(rl.FONT_DEFAULT)).PassRef()
stlas := rl.GenImageFontAtlas(fontDefault.Chars(0), &fontDefault.This.Recs, 95, 16, 4, 0)
```

In C
```c
selected = hovered;
selectedPos = hoveredPos;
SetClipboardText(messages[emoji[selected].message].text);
```
In Go
```go
selected = hovered
*selectedPos.This = *hoveredPos.This
rl.SetClipboardText(message[emoji[selected].Message].Text)
```


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
core       | input_mouse            | :heavy_check_mark: |
core       | input_mouse_wheel      | :heavy_check_mark: |
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


### License
go-raylib is licensed under an unmodified zlib/libpng license, which is an OSI-certified, BSD-like license that allows static linking with closed source software. Check [LICENSE](LICENSE) for further details.
