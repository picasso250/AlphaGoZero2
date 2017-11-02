package main

import "fmt"

// 棋盘大小
const BOARD_SIZE = 19
// 棋盘数据
// 0 无棋子
// 1 黑子
// 2 白子
// Golang will init all value to zero
// https://golang.org/ref/spec#Variable_declarations
var go_data [BOARD_SIZE][BOARD_SIZE]byte
const (
  NONE  = 0
  BLACK = 1
  WHITE = 2
)
var go_color_repr_map = [3]byte {'.', 'x', 'o'}

// 打印棋盘
func print_go_board() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Printf("%c ", go_color_repr_map[ go_data [i] [j] ])
		}
		fmt.Printf("(%d)\n", BOARD_SIZE-i)
	}
	for i := 0; i < BOARD_SIZE; i++  {
		fmt.Printf("%c ", 'A'+i)
	}
	fmt.Printf("\n")
}

func main() {
    go_data[1][2] = BLACK
    print_go_board()

}
