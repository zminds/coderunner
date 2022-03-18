package main

import (
	"fmt"
	"strconv"
)

func main() {
	// cf上汉字转为unicode码 按照unicode的字符数顺序排序
	// 姜佛喜 \u59dc\u4f5b\u559c
	// 张敏   \u5f20\u654f
	// 李浩   \u674e\u6d69
	// 许吉友 \u8bb8\u5409\u53cb
	names := []string{"姜佛喜", "张敏", "李浩", "许吉友"}
	for _, name := range names {
		textQuoted := strconv.QuoteToASCII(name)
		textUnquoted := textQuoted[1 : len(textQuoted)-1]
		fmt.Println(textUnquoted)
	}
}
