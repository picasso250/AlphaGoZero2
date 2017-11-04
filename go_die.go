// 死活

package main

import (
	"fmt"
	"strings"
	// "errors"
	// "log"
)

// 棋子的气
// 棋子和他的四个邻居们
// 是 4-Tree
type QiZi struct {
	color  go_color
	i      int
	j      int
	root *QiZi
	parent *QiZi
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
			qz.i, qz.j, go_color_repr_map[qz.color])
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
// 对于qizi来讲，我们将要考察他的邻居(index)所在位置为i,j的
func go_die_qi_build_struct_iter_(i int, j int) (qz *QiZi) {
	fmt.Printf("do(%d,%d)\n",i,j)
	q := QiZi{color: go_data[i][j], i: i, j: j}
	q.neibour[0] = go_die_qi_build_neibour_(i-1, j)
	q.neibour[1] = go_die_qi_build_neibour_(i, j-1)
	q.neibour[2] = go_die_qi_build_neibour_(i, j+1)
	q.neibour[3] = go_die_qi_build_neibour_(i+1, j)
	for k := 0; k < 4; k++ {
		q.neibour[k].parent = qz
	}
	return &q
}
func go_die_qi_build_neibour_(i int, j int) (qz *QiZi) {
	fmt.Printf("build neibour(%d,%d)\n",i,j)
	// 超出棋盘是nil，代表叶子节点
	if !go_pos_in_board(i, j) {
		return nil
	}
	// 防止无限递归
	if qz.parent == nil || !go_die_qi_in_(i,j,qz.root) {
		fmt.Printf("build neibour(%d,%d) for real\n",i,j)
		return go_die_qi_build_struct_iter_(i, j)
	}
	return nil
}
func go_pos_in_board(i int, j int) bool {
	return i >= 0 && i < BOARD_SIZE && j >= 0 && j < BOARD_SIZE
}
// 检测棋子(i,j)是否在树里面
// 而且是同颜色的
func go_die_qi_in_(i int, j int, root *QiZi) bool {
	if root == nil {
		return false
	}
	// printf("search (%d,%d) == (%d,%d) ? \n", i,j,root->i,root->j);
	if root.i == i && root.j == j {
		return root.color == go_data[i][j]
	}
	for k := 0; k < 4; k++ {
		// printf("index=%d, pointer to %lu\n", k,root.neibour[k]);
		if go_die_qi_in_(i, j, root.neibour[k]) {
			return true
		}
	}
	return false
}
func go_die_qi_get_(i int, j int, qz *QiZi) {
	// int err = 0;
	// go_color_t color = go_data[i][j];
	// // assert(color != NONE); // 如果用户调用，则要保证不为 NONE
	// qizi_t q;
	// go_die_qi_init_default(&q,color,i,j);
	// if (err = go_die_qi_build_struct_(i,j,&q,&q)) {
	// 	assert(err == 0);
	// }
	// go_die_qi_print_iter(&q,0);
	// return go_die_qi_get_iter(&q, color);
}
