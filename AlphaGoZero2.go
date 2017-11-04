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

	qz_neibour := [...]*qizi{nil, nil, nil, nil}
	qz := qizi{2, 9, 10, qz_neibour}
	go_die_qi_print_iter(&qz, 0)
}
