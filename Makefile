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
	go build
