package Timed_Task

import (
	"github.com/atotto/clipboard"
	"time"
)

// ClipboardHistory 剪贴板历史
type ClipboardHistory struct {
	records []string // 记录内容
	maxSize int      // 限制大小
}

// NewClipboardHistory 新的剪贴板记录
func NewClipboardHistory(maxSize int) *ClipboardHistory {
	return &ClipboardHistory{
		records: make([]string, 0, maxSize),
		maxSize: maxSize,
	}
}

// AddRecord 实现添加记录功能
func (ch *ClipboardHistory) AddRecord(record string) {
	if len(ch.records) >= ch.maxSize {

		ch.records = ch.records[1:] // 移除最旧的记录
	}
	ch.records = append(ch.records, record)
}

// CaptureClipboard 任务-监听剪贴板记录
func CaptureClipboard(history *ClipboardHistory) {
	var lastContent string
	for {
		content, err := clipboard.ReadAll()
		if err != nil {
			// 处理错误，例如打印日志
			time.Sleep(1 * time.Second)
			continue
		}

		if content != lastContent {
			history.AddRecord(content)
			lastContent = content
		}

		time.Sleep(1 * time.Second) // 设置检查频率，每秒检查一次
	}
}

func (ch *ClipboardHistory) GetRecords() []string {
	return ch.records
}
