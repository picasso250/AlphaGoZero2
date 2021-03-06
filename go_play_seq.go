package main

// 数据结构
import (
// "fmt"
// "strings"
// "errors"
// "log"
)

type GoPlaySeqNode struct {
	is_xuzhao bool
	v         GoVertex
	count     int         // 棋盘上棋子的数量
	shot      GoBoardShot // 下子之后的局面
}
type GoBoardShot [BOARD_SIZE][BOARD_SIZE]GoColor
type GoPlaySeq []GoPlaySeqNode

var go_play_seq GoPlaySeq

func GoSeq2XuZhao() bool {
	n := len(go_play_seq)
	return n >= 2 && go_play_seq[n-1].is_xuzhao && go_play_seq[n-2].is_xuzhao
}

// 返回是否全局同型
func GoAppendSeq(i int, j int, color GoColor) bool {
	v := GoVertex{i: i, j: j, color: color}
	shot, count := GoGetBoardShot()
	n := GoPlaySeqNode{false, v, count, shot}
	for k := len(go_play_seq) - 2; k >= 0; k -= 2 {
		if n.Equal(&go_play_seq[k]) {
			return true
		}
	}
	go_play_seq = append(go_play_seq, n)
	return false
}
func (n *GoPlaySeqNode) Equal(n2 *GoPlaySeqNode) bool {
	if n.count != n2.count {
		return false
	}
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if n.shot[i][j] != n2.shot[i][j] {
				return false
			}
		}
	}
	return true
}
func GoGetBoardShot() (s [BOARD_SIZE][BOARD_SIZE]GoColor, count int) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			color := go_vertex_data[i][j].color
			s[i][j] = color
			if color != NONE {
				count++
			}
		}
	}
	return s, count
}
