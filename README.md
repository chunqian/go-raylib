# go-raylib

### Bindings

Header     |    Supported       |
---------  | ------------------ |
raylib.h   | :heavy_check_mark: |
raymath.h  | :heavy_check_mark: |
raygui.h   | :heavy_check_mark: |
ricons.h   | :heavy_check_mark: |
physac.h   | :heavy_check_mark: |

### Version
Synchronized update with [raylib](https://github.com/raysan5/raylib).


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
