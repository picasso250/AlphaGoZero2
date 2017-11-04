package main

import (
	// "fmt"
	// "errors"
	"log"
)

// 棋盘大小
const BOARD_SIZE = 19

// 棋盘数据
// 0 无棋子
// 1 黑子
// 2 白子
// Golang will init all value to zero
// https://golang.org/ref/spec#Variable_declarations
type go_color byte

var go_data [BOARD_SIZE][BOARD_SIZE]go_color

const (
	NONE  = 0
	BLACK = 1
	WHITE = 2
)

var go_color_repr_map = [3]byte{'.', 'x', 'o'}

func main() {
	var err error
	if err = one_move_(1, 3, WHITE); err != nil {
		log.Fatal(err)
	}
	print_go_board()

	nil_neibours := [...]*QiZi{nil, nil, nil, nil}
	n0 := QiZi{NONE,8,10,nil_neibours}
	n1 := QiZi{NONE,9,9,nil_neibours}
	n2 := QiZi{NONE,9,11,nil_neibours}
	n3 := QiZi{NONE,10,10,nil_neibours}
	qz_neibours := [...]*QiZi{&n0, &n1, &n2, &n3}
	qz := QiZi{2, 9, 10, qz_neibours}
	go_die_qi_print_iter(&qz, 0)
}
