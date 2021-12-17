package utils

import (
	"github.com/Tang-RoseChild/mahonia"
)

// ConvertTo
// 字符串转换编码
// param string src: 待转换的字符串
// param string coder: 目标编码
func ConvertTo(src, coder string) string {
	enc := mahonia.NewEncoder(coder)
	return enc.ConvertString(src)
}
