// 基本类库

package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

// 棋盘大小
const BOARD_SIZE = 3

// 贴目
const tiemu = 0

// 棋盘数据
var go_vertex_data [BOARD_SIZE][BOARD_SIZE]GoVertex

var (
	ErrForbidPoint = errors.New("禁着点")
	ErrViewSame    = errors.New("禁止全局同型")
)

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
		return errors.New("can not move on an already point " + strconv.Itoa(i) + "," + strconv.Itoa(j))
	}
	go_vertex_data[i][j].color = color
	go_update_edge(i, j)

	go_vertex_data[i][j].block = GetGoBlock(i, j)
	return nil
}

// 盘面胜负
func GoGetPanMian() map[GoColor]int {
	m := make(map[GoColor]int)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			v := go_vertex_data[i][j]
			if v.color != NONE {
				m[v.color]++
			} else {
				c := v.GetOwnerColor()
				if c != NONE {
					m[c]++
				}
			}
		}
	}
	return m
}
func GoPrintPanMian() {
	m := GoGetPanMian()
	fmt.Printf("黑：%d 白: %d\n", m[BLACK], m[WHITE])
}

func do_tizi_4(i int, j int) bool {
	v := &go_vertex_data[i][j]
	tizi := false
	for k := 0; k < 4; k++ {
		v_other := v.edge[k].v2
		if v_other != nil && v_other.color == v.color.Reverse() {
			qi := GoGetQi(v_other.i, v_other.j)
			if qi == 0 {
				GoTiZi(v_other)
				tizi = true
			}
		}
	}
	return tizi
}

func undo_one_step() error {
	i := len(go_play_seq) - 2
	if i < 0 {
		return errors.New("undo start")
	}
	shot := go_play_seq[i].shot
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			go_vertex_data[i][j].color = shot[i][j]
		}
	}
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			go_vertex_data[i][j].InitEdge()
		}
	}
	return nil
}

// 悔棋(计算机版本)
// 仅仅用于没有提子的情况
func un_move_(i int, j int) {
	assert(i >= 0 && i < BOARD_SIZE)
	assert(j >= 0 && j < BOARD_SIZE)
	assert(go_vertex_data[i][j].color != NONE)
	go_vertex_data[i][j].color = NONE
	go_update_edge(i, j)
}

type GoPos struct {
	i int
	j int
}

func (p GoPos) InBoard() bool {
	return 0 <= p.i && p.i < BOARD_SIZE && 0 <= p.j && p.j < BOARD_SIZE
}
func get_neibour_pos_list(i int, j int) (ret []GoPos) {
	neibour_pos_list := [4]GoPos{
		GoPos{i - 1, j},
		GoPos{i, j - 1},
		GoPos{i, j + 1},
		GoPos{i + 1, j},
	}
	ret = make([]GoPos, 0, 4)
	for ii := 0; ii < 4; ii++ {
		pos := neibour_pos_list[ii]
		if pos.InBoard() {
			ret = append(ret, pos)
		}
	}
	return ret
}

func get_neibour_by_color(v *GoVertex, color GoColor) (ret []*GoVertex) {
	pos_list := get_neibour_pos_list(v.i, v.j)
	for _, pos := range pos_list {
		p := &go_vertex_data[pos.i][pos.j]
		if p.color == color {
			ret = append(ret, p)
		}
	}
	return ret
}
func D(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
