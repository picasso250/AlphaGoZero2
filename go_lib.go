// 基本类库

package main

import (
	"errors"
	"fmt"
	"log"
)

// 棋盘大小
const BOARD_SIZE = 19
// 棋盘数据
var go_vertex_data [BOARD_SIZE][BOARD_SIZE]GoVertex

// 打印棋盘
func print_go_board() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Printf("%s ", go_vertex_data[i][j].ColorStr())
		}
		fmt.Printf("(%d)\n", BOARD_SIZE-i)
	}
	for i := 0; i < BOARD_SIZE; i++ {
		fmt.Printf("%c ", 'A'+i)
	}
	fmt.Printf("\n")
}

func assert(cond bool) {
	if !cond {
		log.Fatal(cond)
	}
}

// 是否是禁着点
func IsForbidPoint(v GoVertex) bool {
	assert(go_vertex_data[v.i][v.j].color==NONE)
	// todo
	// for k:=0;k<4;k++ {
	// 	v.edge[k].v2
	// }
	return true
}

// 在某处位置走一步棋(计算机版本)
// 有可能是在禁着点上，在人类版本上考虑此点
func one_move_(i int, j int, color GoColor) (err error) {
	assert(i >= 0 && i < BOARD_SIZE)
	assert(j >= 0 && j < BOARD_SIZE)
	if go_vertex_data[i][j].color != NONE {
		return errors.New("can not move on an already point")
	}
	go_vertex_data[i][j].color = color
	go_update_edge(i, j)
	// 是否死了，谁死了？
	// 如何提子？
	return nil
}
