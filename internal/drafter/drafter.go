package drafter

import (
	"container/list"
	"image"
	"recorder/internal/watcher"

	uuid "github.com/satori/go.uuid"

	_ "recorder/internal/statik" // TODO: Replace with the absolute import path
)

// Drafter 绘制器
type Drafter struct {
	screenshotList *list.List    // 截图列表
	gitGen         *gifGenerator // 生成器
	clicked        image.Image   // 点击
	flag           image.Image   // 插旗
}

func Run(d *Drafter, eCh chan *watcher.WatcherEvent) {
	for e := range eCh {
		switch e.GetType() {
		case watcher.ACTION:
			if d.gitGen == nil {
				d.gitGen = newGifGen(uuid.NewV4().String()+".gif", d.clicked, d.flag)
				go runGifGen(d.gitGen)
			}
		case watcher.COMPLETE:
			if d.gitGen != nil {
				d.gitGen.complete()
				d.gitGen = nil
			}
		case watcher.SCREENSHOT, watcher.RECORD:
			if d.gitGen != nil {
				d.gitGen.addRawImg(e)
			}

		}
	}
}

func New(clicked, flag image.Image) (d *Drafter) {
	d = &Drafter{}
	d.screenshotList = list.New()
	d.flag = flag
	d.clicked = clicked
	return
}
