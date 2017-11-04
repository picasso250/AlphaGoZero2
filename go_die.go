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
// 对于qizi来讲，我们将要考察他的邻居(index)所在位置为i,j的
func go_die_qi_build_struct_helper(
	i int, j int, root *QiZi, qizi *QiZi, index int) (err error) {
	// go_die_qi_init_default(qizi->neibour[index], go_data[i][j], i, j);
	if root.color == go_data[i][j] {
		return go_die_qi_build_struct(i, j, root, qizi.neibour[index])
	} else {
		qizi.neibour[index] = &QiZi{color: go_data[i][j], i: i, j: j}
	}
	return nil
}

// 检测棋子(i,j)是否在树里面
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

// 如果此邻居不是在棋盘内，则不予操作
func go_die_qi_build_struct_helper_neibour(
	i int, j int, root *QiZi, qizi *QiZi, index int) (err error) {
	if i >= 0 && i < BOARD_SIZE && j >= 0 && j < BOARD_SIZE {
		// printf("(%d,%d) in ? %d\n", i-1,j, go_die_qi_in_(i-1,j,root));
		if go_die_qi_in_(i, j, root) {
			qizi.neibour[index] = nil
		} else {
			// printf("0 do (%d,%d)\n",i-1,j);
			err = go_die_qi_build_struct_helper(i-1, j, root, qizi, 0)
			if err != nil {
				return err
			}
		}
	} else {
		qizi.neibour[index] = nil
	}
	return nil
}
func go_die_qi_build_struct(i int, j int, root *QiZi, qizi *QiZi) (err error) {
	// printf("building for (%d,%d) root@(%d,%d)\n", i,j,root->i,root->j);
	// 0
	err = go_die_qi_build_struct_helper_neibour(i-1, j, root, qizi, 0)
	if err != nil {
		return err
	}
	// 1
	err = go_die_qi_build_struct_helper_neibour(i, j-1, root, qizi, 1)
	if err != nil {
		return err
	}
	// 2
	err = go_die_qi_build_struct_helper_neibour(i, j+1, root, qizi, 2)
	if err != nil {
		return err
	}
	// 3
	err = go_die_qi_build_struct_helper_neibour(i+1, j, root, qizi, 3)
	if err != nil {
		return err
	}
	return nil
}

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
