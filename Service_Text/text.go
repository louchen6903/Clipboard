package Service_Text

import (
	codes "Clipboard/Message"
	"fmt"
	"github.com/atotto/clipboard"
)

// GetClipboard 读取剪贴板内容
func GetClipboard() (string, error) {
	contexts, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("[Service_Text][GetClipboard]获取剪贴板内容失败:", err)
		fmt.Println(codes.ErrService, err)
		return "", err
	}
	return contexts, nil
}

// SetClipboard 写入剪贴板内容
func SetClipboard(text string) {
	return
}
