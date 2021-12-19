package wndcore

import (
	"github.com/cosokit/doit/pkg/comm/utils"
	"github.com/cosokit/doit/pkg/wndcore/custom"
	"github.com/webview/webview"
	"net/http"
	"runtime"
	"time"
	"unsafe"
)

type DTWnd struct {
	assets string
	view   webview.WebView
}

var MainWindow *DTWnd

var c = make(chan string, 1)

// New 新建窗口
// :title 窗口标题
// :w 窗口宽度
// :h 窗口高度
// :assets 静态资源所在的文件目录
// :addr 监听的地址和端口
// :debug 是否开启调试模式
func New(title string, w, h int, assets string, debug bool) *DTWnd {
	view := webview.New(debug)

	go func() {
		for {
			js := <-c
			view.Dispatch(func() {
				view.Eval(js)
			})
		}
	}()

	view.SetTitle(title)
	view.SetSize(w, h, webview.HintNone)
	view.Navigate("http://localhost")

	if runtime.GOOS == "windows" {
		title = utils.ConvertTo(title, "gbk")
	} else if runtime.GOOS == "darwin" {
		// MACOSX系统设置窗体边框和阴影
		wnd := custom.NSWindow{Ptr: view.Window()}
		wnd.SetHasShadow(true)
	}

	return &DTWnd{view: view, assets: assets}
}

func (wnd *DTWnd) Destroy() {
	wnd.view.Destroy()
}

func (wnd *DTWnd) Run() {
	go func() {
		s := &http.Server{
			Handler:        http.FileServer(http.Dir(wnd.assets)),
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		_ = s.ListenAndServe()
	}()

	wnd.view.Run()
}

// Eval
// 在webview页面中执行一段JS
func (wnd *DTWnd) Eval(js string) {
	c <- js
}

// WndPtr
// 获取主窗口的窗口句柄
// When using GTK backend the pointer is GtkWindow pointer
// when using Cocoa backend the pointer is NSWindow pointer
// when using Win32 backend the pointer is HWND pointer
func (wnd *DTWnd) WndPtr() unsafe.Pointer {
	return wnd.view.Window()
}
