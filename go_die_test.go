package main

// 死活
import (
// "fmt"
// "strings"
// "errors"
// "log"
"testing"
)

func TestGoGetQi(t *testing.T) {
	err := one_move_(1, 1, BLACK)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	// 对周围四个子，是否死了，谁死了？
	tizi := do_tizi_4(1,1)
	if tizi {
		t.Errorf("tizi")
	}
	// 禁着点
	print_go_board();
	if GoGetQi(1,1) != 4 {
		t.Errorf("qi is not 4")
	}
	// 禁止全局同形
	if GoAppendSeq(1,1, BLACK) {
		t.Errorf("全局同型")
	}
	un_move_(1,1)
}
