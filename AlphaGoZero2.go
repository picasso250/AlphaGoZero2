package main

import (
	"fmt"
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
	go_vetex_data_init()
	if err = one_move_(9, 10, WHITE); err != nil {
		log.Fatal(err)
	}
	if err = one_move_(9, 9, WHITE); err != nil {
		log.Fatal(err)
	}
	print_go_board()

	// fmt.Printf("qi of (9,9) is %d\n",go_get_qi(9,9))
	s := NewGoEdgeSet()
	_, exists := s.Get(&go_vertex_data[9][10], &go_vertex_data[9][9])
	fmt.Printf("(9,10)-(9,9) exists? %v (should be no)\n", exists)
	e:=GoEdge{&go_vertex_data[9][10], &go_vertex_data[9][9], CONNECT}
	is_add := s.Add(&e)
	fmt.Printf("add (9,10)-(9,9)? %v (should be yes)\n", is_add)
	_, exists = s.Get(&go_vertex_data[9][10], &go_vertex_data[9][9])
	fmt.Printf("(9,10)-(9,9) exists? %v (should be yes)\n", exists)
	_, exists = s.Get(&go_vertex_data[9][9], &go_vertex_data[9][10])
	fmt.Printf("(9,9)-(9,10) exists? %v (should be yes)\n", exists)

	fmt.Printf("OK\n")
}
