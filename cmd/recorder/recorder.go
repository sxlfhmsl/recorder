//go:generate statik -src=../../assets -dest ../../internal -f
//go:generate go fmt ../../internal/statik/statik.go

package main

import (
	"image/png"
	"os"
	"os/signal"
	"recorder/internal/drafter"
	"recorder/internal/watcher"
	"syscall"

	"github.com/rakyll/statik/fs"
)

var quitChannel chan os.Signal

func WaitIntSignal(w *watcher.Watcher) {
	quitChannel = make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE, syscall.SIGABRT,
		syscall.SIGSEGV, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGBUS)
	<-quitChannel
	w.Stop()
}

// Before buildling, run go generate.
func main() {
	// 读取图片
	statikFS, _ := fs.New()
	clickedFile, _ := statikFS.Open("/clicked.png")
	flagFile, _ := statikFS.Open("/flag.png")
	clickedImg, _ := png.Decode(clickedFile)
	flagImg, _ := png.Decode(flagFile)
	clickedFile.Close()
	flagFile.Close()

	d := drafter.New(clickedImg, flagImg)
	w := watcher.New()
	w.Start()
	go WaitIntSignal(w)
	go drafter.Run(d, w.EventChannel())
	w.WaitStop()
}
