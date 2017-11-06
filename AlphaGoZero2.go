package main

import (
	"fmt"
	// "errors"
	// "bufio"
	"log"
	// "os"
	"errors"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var err error
	go_vetex_data_init()
	if err = GoOneMove(0, 0, BLACK); err != nil {
		log.Fatal(err)
	}
	print_go_board()
	GoPrintPanMian()
	if err = GoOneMove(1, 1, WHITE); err != nil {
		log.Fatal(err)
	}
	// if err = GoOneMove(6, 5, WHITE); err != nil {
	// 	log.Fatal(err)
	// }
	// if err = GoOneMove(5, 4, WHITE); err != nil {
	// 	log.Fatal(err)
	// }
	// if err = GoOneMove(5, 6, WHITE); err != nil {
	// 	log.Fatal(err)
	// }
	print_go_board()
	// fmt.Printf("edge of (5,9):\n")
	// for k := 0; k < 4; k++ {
	// 	fmt.Printf("\t%s\n", &go_vertex_data[9][9].edge[k])
	// }

	// vs, es := go_find_color_block(9, 9)
	// fmt.Printf("vs: %s\nes: %s\n", vs, es)
	// q := GoGetQi(9, 9)

	GoPrintPanMian()
	return

	for {
		fmt.Printf("playing: ")
		var cmd string
		n, err := fmt.Scanf("%s", &cmd)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Printf("empty cmd, reEnter\n")
			continue
		}
		fmt.Printf("cmd: %s\n", cmd)

		switch {
		case true:
			err=GoHumanInteractPlay()
			if err!=nil {
				log.Fatal(err)
			}
			print_go_board()
		}
	}
	fmt.Printf("OK\n")
}
func GoHumanInteractPlay() error {
	var i_d int
	var j_s string
	var color_s string
	n, err := fmt.Scanf("%1s%d%1s ", &j_s, &i_d, &color_s)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("empty cmd, reEnter\n")
	}
	fmt.Printf("move at %s %d ,color = %s\n",j_s,i_d,color_s)
	i := BOARD_SIZE - i_d
	j := int(strings.ToUpper(j_s)[0] - 'A')
	if !go_pos_in_board(i, j) {
		return errors.New("not in board")
	}
	color_s = strings.ToLower(color_s)
	if color_s != "x" && color_s != "o" {
		return errors.New("color must be x or o")
	}
	color := GoColorMapRev()[color_s]
	err = GoOneMove(i, j, color)
	if err != nil {
		return err
	}
	return nil
}
