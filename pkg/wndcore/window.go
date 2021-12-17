package wndcore

import (
	"github.com/cosokit/doit/pkg/comm/utils"
	"github.com/webview/webview"
	"net/http"
	"runtime"
	"time"
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

	view.SetSize(w, h, webview.HintNone)
	if runtime.GOOS == "windows" {
		title = utils.ConvertTo(title, "gbk")
	}
	view.SetTitle(title)
	view.Navigate("http://localhost")

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
