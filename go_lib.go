// 基本类库

package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

// 棋盘大小
const BOARD_SIZE = 2

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
		debug.PrintStack()
		log.Fatal(cond)
	}
}

// 是否是禁着点
func GoOneMove(i int, j int, color GoColor) (err error) {
	err = one_move_(i, j, color)
	if err != nil {
		return err
	}
	// 是否死了，谁死了？
	v := &go_vertex_data[i][j]
	tizi := false
	for k := 0; k < 4; k++ {
		v_other := v.edge[k].v2
		if v_other!=nil && v_other.color == v.color.Reverse() {
			q := GoGetQi(v_other.i, v_other.j)
			// fmt.Printf("%s -> %s =%d(气)\n", v, v_other, q)
			if q == 0 {
				GoTiZi(v_other)
				tizi = true
			}
		}
	}
	if !tizi {
		if GoGetQi(v.i, v.j) == 0 {
			return errors.New("禁着点")
		}
	}
	return nil
}

// 提子
func GoTiZi(v *GoVertex) {
	fmt.Printf("ti %s\n", v)
	vs, _ := go_find_color_block(v.i, v.j)
	for v, _ := range vs.m {
		v.color = NONE
	}
	for v, _ := range vs.m {
		go_update_edge(v.i, v.j)
	}
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

	return nil
}

// 悔棋(计算机版本)
func un_move_(i int, j int) {
	assert(i >= 0 && i < BOARD_SIZE)
	assert(j >= 0 && j < BOARD_SIZE)
	assert(go_vertex_data[i][j].color != NONE)
	go_vertex_data[i][j].color = NONE
	go_update_edge(i, j)
}
// 盘面胜负
func GoGetPanMian() (map[GoColor]int) {
	m:=make(map[GoColor]int)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			v:=go_vertex_data[i][j]
			if v.color!=NONE {
				m[v.color]++
			} else {
				c:=v.GetOwnerColor()
				if c!=NONE {
					m[c]++
				}
			}
		}
	}
	return m
}
func GoPrintPanMian() {
	m:=GoGetPanMian()
	fmt.Printf("黑：%d 白: %d\n", m[BLACK],m[WHITE])
}
