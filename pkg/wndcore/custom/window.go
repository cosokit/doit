//go:build !darwin
// +build !darwin

package custom

import "C"
import "unsafe"

type NSWindow struct {
	Ptr unsafe.Pointer
}

// SetHasShadow 设置窗体阴影
func (wnd *NSWindow) SetHasShadow(hasShadow bool) {

}
