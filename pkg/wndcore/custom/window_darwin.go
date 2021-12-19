//go:build darwin
// +build darwin

package custom

// #cgo darwin CFLAGS: -x objective-c
// #cgo darwin LDFLAGS: -framework Cocoa
// #define COCOA_UIKIT
// #include "ns_window.h"
import "C"
import "unsafe"

type NSWindow struct {
	Ptr unsafe.Pointer
}

// SetHasShadow 设置窗体阴影
func (wnd *NSWindow) SetHasShadow(hasShadow bool) {
	C.SetHasShadow(wnd.Ptr, C._Bool(hasShadow))
}
