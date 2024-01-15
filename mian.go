package main

import (
	"Clipboard/Timed_Task"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("剪贴板")

	history := Timed_Task.NewClipboardHistory(50)
	go Timed_Task.CaptureClipboard(history)

	// 创建一个定时器，定期更新UI
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {

			// 获取所有历史记录
			records := history.GetRecords()
			w.SetContent(widget.NewLabel(strings.Join(records, "")))
		}
	}()

	w.ShowAndRun()
}

func init() {
	os.Setenv("FYNE_FONT", "./40218451935.ttf")
	os.Setenv("FYNE_FONT", "./47714613200.ttf")

}
