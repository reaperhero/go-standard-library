package chank

import (
	"github.com/axgle/mahonia"
)
/*
将文本转换字符编码并返回
srcStr		待转码的原始字符串
srcEncoding	原始字符串的编码（字符集）
dstEncoding	目标编码（字符集）
*/
func ConvertEncoding(srcStr string, srcEncoding string, dstEncoding string) (dstStr string, err error) {

	//创建指定字符集的解码器
	srcDecoder := mahonia.NewDecoder(srcEncoding)
	dstDecoder := mahonia.NewDecoder(dstEncoding)

	//将内容转换为UTF-8字符串
	utfStr := srcDecoder.ConvertString(srcStr)

	//将UTF-8字节转换为目标字符集的字节
	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
	if err != nil {
		return
	}

	//还原为字符串并返回
	dstStr = string(dstBytes)
	return
}
