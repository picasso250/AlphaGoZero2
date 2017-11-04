// 死活

package main

import (
	"fmt"
	"strings"
	// "errors"
	// "log"
)

// 棋子和他的四个邻居们
type QiZi struct {
	color go_color
	i     int
	j     int
	// 0 1 2 3
	//   0
	// 1 o 2
	//   3
	neibour [4]*QiZi
}

// 打印这个结构（辅助调试）
func go_die_qi_print_tab(tab int) {
	fmt.Printf("%s", strings.Repeat("\t", tab))
}
func go_die_qi_print_iter(qz *QiZi, tab int) {
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

// 构建棋子块（4-树）
// func go_die_qi_build_struct
// (i int, j int, root *qizi, qizi *qizi);
// void go_die_qi_init_default
// (qizi_t *qizi, go_color_t color, int i, int j) {
// 	assert(qizi);
// 	printf("init %lu (%d,%d) %d\n", qizi, i,j, color);
// 	qizi->color = color;
// 	qizi->i = i;
// 	qizi->j = j;
// 	qizi->neibour[0] = qizi->neibour[1] =
// 		qizi->neibour[2] = qizi->neibour[3] = NULL;
// 		printf("init complete\n");
// }
// int go_die_qi_build_struct_helper
// (int i, int j, qizi_t *root, qizi_t *qizi, int index) {
// 	qizi->neibour[index] = malloc(sizeof(qizi_t));
// 	if (qizi->neibour[index] == NULL) {
// 		return 1;
// 	}
// 	go_die_qi_init_default(qizi->neibour[index], go_data[i][j], i, j);
//
// 	if (root->color == qizi->neibour[index]->color) {
// 		go_die_qi_build_struct(i-1,j,root,qizi->neibour[index]);
// 	} else {
// 		// do nothing
// 	}
// 	return 0;
// }

func go_die_qi_get_(i int, j int, qz *QiZi) {
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
