package watcher

import (
	"image"

	hook "github.com/robotn/gohook"
)

// EventType 事件类型
type EventType int64

const (
	SCREENSHOT EventType = 0 // 立即截图
	RECORD     EventType = 1 // 录像
	ACTION     EventType = 2 // 启动
	COMPLETE   EventType = 3 // 完成
)

// WatcherEvent 事件
type WatcherEvent struct {
	eventType EventType   // 事件类型
	event     hook.Event  // 附件信息
	img       image.Image // 截图
}

// GetType 获取事件类型
func (e *WatcherEvent) GetType() EventType {
	return e.eventType
}

// GetExtendInfo 获取附加信息
func (e *WatcherEvent) GetExtendInfo() hook.Event {
	return e.event
}

// GetImg 获取截图
func (e *WatcherEvent) GetImg() image.Image {
	return e.img
}
