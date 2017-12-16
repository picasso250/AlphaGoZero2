package main

import (
	"fmt"
	// "strings"
	// "errors"
	// "log"
)

// 围棋基本规则

// 下一步棋，同时查看是否是禁着点
// 行棋规则：气尽提取
func GoOneMove(i int, j int, color GoColor) (err error) {
	err = one_move_(i, j, color)
	if err != nil {
		return err
	}
	// 对周围四个子，是否死了，谁死了？
	tizi := do_tizi_4(i, j)
	fmt.Printf("tizi=%v, qi=%d\n",tizi,GoGetQi(i, j))
	// 禁着点
	if !tizi && GoGetQi(i, j) == 0 {
		return ErrForbidPoint
	}
	// 禁止全局同形
	if GoAppendSeq(i, j, color) {
		return ErrViewSame
	}
	return nil
}

// 谁是胜利者？
// 胜负规则：子多为胜
// 正 黑胜 0 平局 负 白胜
func GoWinner() int {
	m := make(map[GoColor]int)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			m[go_vertex_data[i][j].color] += 1
		}
	}
	return m[BLACK] - m[WHITE] + tiemu
}
