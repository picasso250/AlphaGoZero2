package main

import (
	"fmt"
	// "errors"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var err error
	go_vetex_data_init()
	if err = GoOneMove(9, 9, BLACK); err != nil {
		log.Fatal(err)
	}
	if err = GoOneMove(8, 9, WHITE); err != nil {
		log.Fatal(err)
	}
	if err = GoOneMove(10, 9, WHITE); err != nil {
		log.Fatal(err)
	}
	if err = GoOneMove(9, 8, WHITE); err != nil {
		log.Fatal(err)
	}
	if err = GoOneMove(9, 10, WHITE); err != nil {
		log.Fatal(err)
	}
	print_go_board()
	// fmt.Printf("edge of (9,9):\n")
	// for k := 0; k < 4; k++ {
	// 	fmt.Printf("\t%s\n", &go_vertex_data[9][9].edge[k])
	// }

	// fmt.Printf("qi of (9,9) is %d\n",go_get_qi(9,9))
	// s := NewGoEdgeSet()
	// _, exists := s.Get(&go_vertex_data[9][10], &go_vertex_data[9][9])
	// fmt.Printf("(9,10)-(9,9) exists? %v (should be no)\n", exists)
	// e:=GoEdge{&go_vertex_data[9][10], &go_vertex_data[9][9], CONNECT}
	// is_add := s.Add(&e)
	// fmt.Printf("add (9,10)-(9,9)? %v (should be yes)\n", is_add)
	// _, exists = s.Get(&go_vertex_data[9][10], &go_vertex_data[9][9])
	// fmt.Printf("(9,10)-(9,9) exists? %v (should be yes)\n", exists)
	// _, exists = s.Get(&go_vertex_data[9][9], &go_vertex_data[9][10])
	// fmt.Printf("(9,9)-(9,10) exists? %v (should be yes)\n", exists)

	// vs, es := go_find_color_block(9, 9)
	// fmt.Printf("vs: %s\nes: %s\n", vs, es)
	q := GoGetQi(9, 9)
	fmt.Printf("%s qi=%d\n",&go_vertex_data[9][9], q)

	fmt.Printf("OK\n")
}
