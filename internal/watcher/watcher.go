package watcher

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// Watcher 事件监视器
type Watcher struct {
	eventCh        chan *WatcherEvent // 数据通道
	state          int32              // 状态码
	eFilterCh      chan hook.Event    // 事件通道
	screenIsPaused bool               // 是否已经暂停
	recordStopCh   chan interface{}   // 录像停止信号
}

// Start 启动监听
func (w *Watcher) Start() {
	w.eventCh = make(chan *WatcherEvent)

	// 启动或者关闭任务
	robotgo.EventHook(hook.KeyDown, []string{"command", "alt", "s"}, w.switchState)
	// 暂停或者重启
	robotgo.EventHook(hook.KeyDown, []string{"command", "alt", "z"}, w.pauseScreen)
	// 开始录像 25fps
	robotgo.EventHook(hook.KeyDown, []string{"command", "alt", "r"}, w.actionRecord)
	// 截图不带位置提示
	robotgo.EventHook(hook.KeyDown, []string{"command", "alt", "x"}, w.tellScreenshotNow)
	// 截图带有位置提示-----鼠标位置
	robotgo.EventHook(hook.MouseDown, []string{}, w.tellScreenshotNow)

	w.eFilterCh = robotgo.EventStart()
	fmt.Println("监听器即将启动")
}

// Stop 有效停止
func (w *Watcher) Stop() {
	fmt.Println("停止监听器")
	robotgo.EventEnd()
}

// WaitStop 等待停止监听-----阻塞的方法
func (w *Watcher) WaitStop() {
	fmt.Println("等待监听器停止")
	<-robotgo.EventProcess(w.eFilterCh)
	close(w.eventCh)
	fmt.Println("监听器已经停止")
}

// EventChannel 获取事件通道
func (w *Watcher) EventChannel() chan *WatcherEvent {
	return w.eventCh
}

// switchState 切换状态
func (w *Watcher) switchState(e hook.Event) {
	eventObj := &WatcherEvent{}
	eventObj.event = e
	if w.state == 0 {
		eventObj.eventType = ACTION
		w.state = 1
	} else {
		eventObj.eventType = COMPLETE
		w.state = 0

		// 停止所有-----取消暂停状态
		w.screenIsPaused = false
		// 停止所有-----停止录像
		if w.recordStopCh != nil {
			close(w.recordStopCh)
			w.recordStopCh = nil
		}
	}
	w.eventCh <- eventObj
}

// pauseScreen 暂停所有操作
func (w *Watcher) pauseScreen(e hook.Event) {
	if w.state != 1 || w.recordStopCh != nil {
		return
	}

	fmt.Println("切换截图暂停:", w.screenIsPaused)
	w.screenIsPaused = !w.screenIsPaused
}

// actionRecord 立即开始录像
func (w *Watcher) actionRecord(e hook.Event) {
	if w.state != 1 || w.screenIsPaused {
		return
	}

	if w.recordStopCh == nil {
		w.recordStopCh = make(chan interface{})
		go w.record()
	} else {
		close(w.recordStopCh)
		w.recordStopCh = nil
	}
}

// record 具体的录像
func (w *Watcher) record() {
	recordStopCh := w.recordStopCh

	eventObj := &WatcherEvent{}
	eventObj.eventType = RECORD
	eventObj.img = robotgo.CaptureImg()
	w.eventCh <- eventObj

	for {
		screenCh := time.After(time.Millisecond * 40)
		select {
		case <-screenCh: // 录屏
			eventObj := &WatcherEvent{}
			eventObj.eventType = RECORD
			eventObj.img = robotgo.CaptureImg()
			w.eventCh <- eventObj
		case <-recordStopCh: // 停止
			return
		}
	}
}

// tellScreenshotNow 需要立即截屏
func (w *Watcher) tellScreenshotNow(e hook.Event) {
	if w.state != 1 || w.screenIsPaused || w.recordStopCh != nil {
		return
	}

	if e.Kind == hook.MouseDown {
		if e.Button == robotgo.MouseMap["right"] {
			return
		}
	}

	eventObj := &WatcherEvent{}
	eventObj.eventType = SCREENSHOT
	eventObj.event = e
	eventObj.img = robotgo.CaptureImg()
	w.eventCh <- eventObj
}

func New() (w *Watcher) {
	w = &Watcher{}
	return
}
