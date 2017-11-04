// 死活

package main

import (
	"fmt"
	"strings"
	// "errors"
	// "log"
)

// 棋子和他的四个邻居们
type qizi struct {
	color go_color
	i     int
	j     int
	// 0 1 2 3
	//   0
	// 1 o 2
	//   3
	neibour [4]*qizi
}

// 打印这个结构（辅助调试）
func go_die_qi_print_tab(tab int) {
	fmt.Printf("%s", strings.Repeat("\t", tab))
}
func go_die_qi_print_iter(qz *qizi, tab int) {
	if qz != nil {
		fmt.Printf("(%d,%d) %c\n",
			qz.i, qz.j,
			go_color_repr_map[qz.color])
		for i := 0; i < 4; i++ {
			go_die_qi_print_tab(tab + 1)
			fmt.Printf("| ")
			go_die_qi_print_iter(qz.neibour[i], tab+1)
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("nil\n")
	}
}

func go_die_qi_get_(i int, j int, qz *qizi) {
	// int err = 0;
	// go_color_t color = go_data[i][j];
	// // assert(color != NONE); // 如果用户调用，则要保证不为 NONE
	// qizi_t q;
	// go_die_qi_init_default(&q,color,i,j);
	// if (err = go_die_qi_build_struct(i,j,&q,&q)) {
	// 	assert(err == 0);
	// }
	// go_die_qi_print_iter(&q,0);
	// return go_die_qi_get_iter(&q, color);
}
