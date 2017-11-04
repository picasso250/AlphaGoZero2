// 基本类库

package main

import (
  "fmt"
  "errors"
  "log"
  )

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

func assert(cond bool) {
  if !cond {
    log.Fatal(cond)
  }
}

// 在某处位置走一步棋(计算机版本)
func one_move_(i int, j int, color go_color) (err error) {
	assert(i >= 0 && i < BOARD_SIZE)
	assert(j >= 0 && j < BOARD_SIZE)
	if go_data [i] [j] != NONE {
    return errors.New("can not move on an already point")
  }
	go_data[i][j] = color
	return nil
}
