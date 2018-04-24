package main

import "math"

// 全局常量
const (
	VERSION = byte(0x00)

	// 挖矿难度
	MINE_TARGET_BITS = 18
	// Number once最大值
	MAX_NONCE = math.MaxInt64

	// clog skip级别
	CLOG_SKIP_DEFAULT      = 0
	CLOG_SKIP_DISPLAY_INFO = 2
)
