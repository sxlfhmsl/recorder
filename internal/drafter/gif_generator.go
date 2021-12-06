package drafter // gif 生成器
import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"recorder/internal/watcher"
	"sync/atomic"
	"time"

	"github.com/andybons/gogif"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	gifSecCount int = 100 // 秒单位
	gifSec      int = 2   // 秒
)

// gifGenerator gif生成器
type gifGenerator struct {
	fileName    string                     // 保存的文件名称
	imgEventsCh chan *watcher.WatcherEvent // 图片事件通知
	imgs        []*image.Paletted          // 所有图片
	delays      []int                      // 所有的延时
	state       int32                      // 标志，用于停止
	clicked     image.Image                // 点击
	flag        image.Image                // 插旗
	isFlagging  bool                       // 是否在打标记
}

// addRawImg 添加图片
func (g *gifGenerator) addRawImg(e *watcher.WatcherEvent) {
	if atomic.LoadInt32(&g.state) != 0 {
		return
	}

	g.imgEventsCh <- e
}

// complete 处理完成
func (g *gifGenerator) complete() {
	atomic.CompareAndSwapInt32(&g.state, 0, 1)
}

// processScreenImg 处理截图
func (g *gifGenerator) processScreenImg(e *watcher.WatcherEvent) (needBreak bool) {
	fmt.Println("任务:" + g.fileName + "    准备处理图片")
	if g.isFlagging && e.GetExtendInfo().Kind == hook.MouseDown && e.GetExtendInfo().Button == robotgo.MouseMap["center"] {
		fmt.Println("任务:" + g.fileName + "    标记一次")
		drawFlag(g.flag, g.imgs[len(g.imgs)-1], image.Pt(int(e.GetExtendInfo().X), int(e.GetExtendInfo().Y)), 0, 0-g.flag.Bounds().Dy())
		return true
	}

	g.isFlagging = false
	fmt.Println("任务:" + g.fileName + "    添加一张图片")
	pImg := switchToPalettedQuick(g, e)
	if e.GetExtendInfo().Kind == hook.MouseDown {
		if e.GetExtendInfo().Button == robotgo.MouseMap["center"] {
			fmt.Println("任务:" + g.fileName + "    开始标记")
			drawFlag(g.flag, pImg, image.Pt(int(e.GetExtendInfo().X), int(e.GetExtendInfo().Y)), 0, 0-g.flag.Bounds().Dy())
			g.isFlagging = true
		} else {
			drawFlag(g.clicked, pImg, image.Pt(int(e.GetExtendInfo().X), int(e.GetExtendInfo().Y)), 0-g.clicked.Bounds().Dx()/2, 0-g.clicked.Bounds().Dy()/2)
		}
	}
	fmt.Println("任务:" + g.fileName + "    添加一张图片完成")
	g.imgs = append(g.imgs, pImg)
	g.delays = append(g.delays, gifSec*gifSecCount)
	return
}

// processRecord 处理录像
func (g *gifGenerator) processRecord(e *watcher.WatcherEvent) {
	fmt.Println("任务:" + g.fileName + "    准备录入一帧")
	g.imgs = append(g.imgs, switchToPalettedQuick(g, e))
	g.delays = append(g.delays, 4)
	fmt.Println("任务:" + g.fileName + "    录入一帧完成")
}

// dumpToGIF 保存为gif
func (g *gifGenerator) dumpToGIF() {
	fmt.Println("任务:" + g.fileName + "    已完成")
	saveGIF := &gif.GIF{LoopCount: len(g.imgs)}
	saveGIF.Delay = g.delays
	saveGIF.Image = g.imgs
	file, _ := os.Create(g.fileName)
	gif.EncodeAll(file, saveGIF)
	file.Close()
	fmt.Println("任务:" + g.fileName + "    GIF已生成")
}

// newGifGen 新的gif生成器
func newGifGen(fileName string, clicked, flag image.Image) (g *gifGenerator) {
	g = &gifGenerator{}
	g.fileName = fileName
	g.imgEventsCh = make(chan *watcher.WatcherEvent, 1000)
	g.imgs = make([]*image.Paletted, 0)
	g.delays = make([]int, 0)
	g.clicked = clicked
	g.flag = flag
	return
}

// runGifGen 接收并保存图片
func runGifGen(g *gifGenerator) {
	if atomic.LoadInt32(&g.state) != 0 {
		return
	}

	fmt.Println("任务:" + g.fileName + "    已启动")
	for {
		waiter := time.After(time.Second)
		select {
		case e := <-g.imgEventsCh:
			if e == nil { // 存图并退出
				g.dumpToGIF()
				return
			} else { // 存图
				if e.GetType() == watcher.SCREENSHOT {
					g.processScreenImg(e)
				} else if e.GetType() == watcher.RECORD {
					g.processRecord(e)
				}
			}
		case <-waiter:
			if atomic.LoadInt32(&g.state) != 0 && len(g.imgEventsCh) == 0 {
				close(g.imgEventsCh)
			}
		}
	}
}

// switchToPalettedQuick 快转换
func switchToPalettedQuick(g *gifGenerator, e *watcher.WatcherEvent) (palettedImage *image.Paletted) {
	bounds := e.GetImg().Bounds()
	palettedImage = image.NewPaletted(bounds, nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	quantizer.Quantize(palettedImage, bounds, e.GetImg(), bounds.Min)
	return
}

// drawFlag 画标记
func drawFlag(flag image.Image, target *image.Paletted, pos image.Point, offsetX, offsetY int) {
	flagBoundsBegin := image.Pt(pos.X+offsetX, pos.Y+offsetY)
	flagBoundsend := image.Pt(flagBoundsBegin.X+flag.Bounds().Max.X, flagBoundsBegin.Y+flag.Bounds().Max.Y)
	draw.Draw(target, image.Rectangle{flagBoundsBegin, flagBoundsend}, flag, flag.Bounds().Min, draw.Over)
}
