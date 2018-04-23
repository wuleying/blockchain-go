package main

import "math"

// 全局常量
const (
	VERSION = byte(0x00)

	// Last hash key
	LAST_HASH_KEY = "lastHash"
	// Bucket名称
	BLOCK_BUCKET_NAME = "blocks"

	// 创世币数据
	GENESIS_COIN_BASEDATA = "hello luoliang"

	// 挖矿难度
	MINE_TARGET_BITS = 18
	// 挖矿奖励
	MINE_SUBSIDY = 10
	// Number once最大值
	MAX_NONCE = math.MaxInt64

	// clog skip级别
	CLOG_SKIP_DEFAULT      = 0
	CLOG_SKIP_DISPLAY_INFO = 2
)
