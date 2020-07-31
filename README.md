# go-raylib

### Build

```bash
$ export MACOSX_DEPLOYMENT_TARGET=10.9
$ git clone https://github.com/chunqian/go-raylib.git
$ cd go-raylib/lib/raylib/src
$ make
$ cd ../..
$ go build -o examples/audio/module_playing/module_playing examples/audio/module_playing/module_playing.go
```

### Run

```bash
$ cd examples/audio/module_playing
$ ./module_playing
```
