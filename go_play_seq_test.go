package main

// 防止无限循环

import (
	// "fmt"
	"testing"
	// "strings"
	// "errors"
	// "log"
)

func TestGoAppendSeq(t *testing.T) {
	go_vetex_data_init()
	if true == GoAppendSeq(1, 1, BLACK) {
		t.Errorf("1,1 全局同形\n")
	}
	go_play_seq = nil
}
