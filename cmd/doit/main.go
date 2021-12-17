package main

import (
	"github.com/cosokit/doit/configs"
	_ "github.com/cosokit/doit/internal/views"
	"github.com/cosokit/doit/pkg/wndcore"
	"path/filepath"
)

func main() {
	dir, filename := filepath.Split(ConfigureFile)
	if len(dir) <= 0 {
		dir = "."
	}

	configs.InitDoit(dir, filename, "yaml")

	wndcore.MainWindow = wndcore.New("doit", 1200, 900, ResourcesPath, true)
	defer wndcore.MainWindow.Destroy()

	wndcore.MainWindow.Run()
}
