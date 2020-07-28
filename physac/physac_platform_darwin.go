// +build darwin

package physac

/*
#cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo darwin CFLAGS: -x objective-c -DC_FOR_GO -DPHYSAC_IMPLEMENTATION -DPHYSAC_NO_THREADS -DPHYSAC_STANDALONE -DPHYSAC_STATIC -Wno-deprecated-declarations -DPLATFORM_DESKTOP -DMAL_NO_COREAUDIO
*/
import "C"
