# go-raylib

### Build

```bash
$ git clone https://github.com/chunqian/go-raylib.git
$ cd go-raylib/lib/raylib/src
$ export MACOSX_DEPLOYMENT_TARGET=10.9
$ make
$ cd ../..
$ go build -o examples/audio/module_playing/module_playing examples/audio/module_playing/module_playing.go
```

### Run

```bash
$ cd examples/audio/module_playing
$ ./module_playing
```
