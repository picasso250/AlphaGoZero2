package main

// 死活
import (
// "fmt"
// "strings"
// "errors"
// "log"
)

// 边的结构 E(V_1,V_2)
// 边的类型
//   1. 联通
//   2. 被堵（被对方或者墙壁堵住）
//   3. 开放
// 更新
//   当V被改变（棋下在此点）时，更新与之相关的边
//   更新边有两种情况
//     1. 将开放边变成联通或者被堵
//     2. 新增开放边
// 找出V所在的棋块
// 目标：
//   1. 找出V所在棋块中所有的同色棋子V_c(找出棋块)
//   2. 找出这些棋子所相关的边M_E
//   3. 找出这些边的边缘棋子（中类型为开的）
//   遍历V相关的E
//     1. 如E是联通类型，则
//       a. 将E中的两点加入M_V
//       b. 重复遍历M_V中所有点的相关E，并重复1
//     2. 如E是其他类型，则将之加入M_E，此分支停止
//   将遍历过的E加入新的集合M，这就是棋块所有相关E
//   遍历E中所有的V，根据颜色可以找出棋块中棋子的数量
//   气=qi(M)=count{V|V所属的E是开类型，且V是空点}
// 当然，我们保证：当一个棋子被下在棋盘上时，它相关的4个E都将会存在

// 顶点（代表棋子）
// edge中不会有nil元素，即使它是壁，此时为闭类型
type GoVertex struct {
	i     int
	j     int
	color go_color
	edge  [4]GoEdge // 默认为开
}

func (v *GoVertex) Equal(w *GoVertex) bool {
	return v.i == w.i && v.j == w.j
}
func (v *GoVertex) IsAnyOpenEdge() bool {
	for k := 0; k < 4; k++ {
		if v.edge[k].type_ == OPEN {
			return true
		}
	}
	return false
}
func (v *GoVertex) InitEdge() {
	for k := 0; k < 4; k++ {
		v.edge[k].v1 = v
	}
	i := v.i
	j := v.j
	v.InitEdge_(0, i-1, j)
	v.InitEdge_(1, i, j-1)
	v.InitEdge_(2, i, j+1)
	v.InitEdge_(3, i+1, j)
}
func (v *GoVertex) InitEdge_(index int, i int, j int) {
	if go_pos_in_board(i, j) {
		v.edge[index].v2 = &go_vertex_data[i][j]
	}
}

// 边
// 有可能其中一个点在棋盘外，此时这个点用nil表示
type GoEdge struct {
	v1    *GoVertex
	v2    *GoVertex
	type_ int
}

// 边的类型
const (
	OPEN    = 0
	CONNECT = 1
	BLOCK   = 2
)

type GoEqualAble interface {
	Equal() bool
}

var go_vertex_data [BOARD_SIZE][BOARD_SIZE]GoVertex

func (e *GoEdge) Reverse() GoEdge {
	return GoEdge{v1: e.v2, v2: e.v1}
}
func (e *GoEdge) GetNoneVertex() (v *GoVertex) {
	assert(e.type_ == NONE)
	assert(e.v1 != nil)
	assert(e.v2 != nil)
	v1_is_none := e.v1.color == NONE
	if v1_is_none {
		return e.v1
	} else {
		return e.v2
	}
}
func (e *GoEdge) GetAllPointSlice() (vs []*GoVertex) {
	vs = append(vs, e.v1, e.v2)
	return vs
}
func (e *GoEdge) Equal(ee *GoEdge) bool {
	return e == ee
}
func (e *GoEdge) ContainV(i int, j int) bool {
	return e.IsFirstV(i, j) || e.IsSecondV(i, j)
}
func (e *GoEdge) IsFirstV(i int, j int) bool {
	return e.v1.i == i && e.v1.j == j
}
func (e *GoEdge) IsSecondV(i int, j int) bool {
	return e.v2.i == i && e.v2.j == j
}
func (e *GoEdge) UpdateByV() {
	e._updateEdge()
}
func (e *GoEdge) IsSameColor() bool {
	return e.v1.color == e.v2.color
}
func (e *GoEdge) GetOtherV(i int, j int) *GoVertex {
	if e.v1.i == i && e.v1.j == j {
		return e.v2
	}
	if e.v2.i == i && e.v2.j == j {
		return e.v1
	}
	assert(false)
	return nil
}

func (e *GoEdge) _updateEdge() {
	if e.v2 != nil {
		e.type_ = OPEN
	} else if e.v1.color == e.v2.color {
		e.type_ = CONNECT
	} else {
		e.type_ = BLOCK
	}
}

func go_vetex_data_init() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			go_vertex_data[i][j] = GoVertex{i: i, j: j}

		}
	}
	// init edge
	// 边的初始化一定要在所有的点初始化完毕之后
	// 因为边的初始化过程中使用了邻居点
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			go_vertex_data[i][j].InitEdge()
		}
	}
}

func go_pos_in_board(i int, j int) bool {
	return i >= 0 && i < BOARD_SIZE && j >= 0 && j < BOARD_SIZE
}

// 0 1 2 3
//   0
// 1 x 2
//   3
// 当棋子下在(i,j)上时，更新边
func go_update_edge(i int, j int) {
	v := &go_vertex_data[i][j]
	e0 := go_update_edge_(v, 0, i-1, j)
	e1 := go_update_edge_(v, 1, i, j-1)
	e2 := go_update_edge_(v, 2, i, j+1)
	e3 := go_update_edge_(v, 3, i+1, j)

	go_update_edge_r(3, i-1, j, e0)
	go_update_edge_r(2, i, j-1, e1)
	go_update_edge_r(1, i, j+1, e2)
	go_update_edge_r(0, i+1, j, e3)
}

// 更新v的index边，另一个棋子为(i,j)
func go_update_edge_(v *GoVertex, index int, i int, j int) GoEdge {
	v.edge[index].UpdateByV()
	return v.edge[index]
}
func go_update_edge_r(index int, i int, j int, e GoEdge) {
	if go_pos_in_board(i, j) {
		go_vertex_data[i][j].edge[index] = e.Reverse()
	}
}

func go_get_edge_index_by(v_subject *GoVertex, v_object *GoVertex) int {
	if v_subject.i == v_object.i {
		if v_subject.j-1 == v_object.j {
			return 1
		}
		if v_subject.j+1 == v_object.j {
			return 2
		}
	}
	if v_subject.j == v_object.j {
		if v_subject.i-1 == v_object.i {
			return 0
		}
		if v_subject.i+1 == v_object.i {
			return 3
		}
	}
	return -1
}

// 气=qi(M)=count{V|V所属的E是开类型}
func go_get_qi(i int, j int) int {
	assert(go_data[i][j] != NONE)
	_, es := go_find_color_block(i, j)
	return go_count_open(es)
}
func go_count_open(es *GoEdgeSet) int {
	vs := NewGoVertexSet()
	for v1v2, type_ := range es.m {
		if type_ == OPEN {
			if v1v2[0] != nil {
				vs.Add(v1v2[0])
			} else {
				vs.Add(v1v2[1])
			}
		}
	}
	return len(vs.m)
}

// 获得棋子块相关的所有边和点
// func go_get_all_block_about(i int, j int) ([]*GoVertex, []*GoEdge) {
// 	vs := make([]*GoVertex, 1)
// 	vs[0] = &go_vertex_data[i][j]
// 	vs,es := go_find_color_block(i, j)
// 	vs = go_add_not_color_edge(es)
// 	return vs, es
// }

// 获得棋子块相关的所有边和点（同色）
func go_find_color_block(i int, j int) (*GoVertexSet, *GoEdgeSet) {
	vs := NewGoVertexSet()
	vs.Add(&go_vertex_data[i][j])
	return go_find_color_block_iter(vs, NewGoEdgeSet())
}
func go_find_color_block_iter(
	vs *GoVertexSet, es *GoEdgeSet) (*GoVertexSet, *GoEdgeSet) {
	// var es_ []*GoEdge
	new_coming_v := false
	for v1v2, type_ := range es.m {
		if type_ == CONNECT {
			if vs.AddEdgeByTwoVertex(v1v2) {
				new_coming_v = true
			}
		} else {
			continue
		}
	}
	new_coming_e := false
	for v, _ := range vs.m {
		if es.AddByVertex(v) { // 将v的所有邻居加入
			new_coming_e = true
		}
	}
	if !new_coming_v && !new_coming_e {
		return vs, es
	}
	return go_find_color_block_iter(vs, es)
}

// func go_vertex_comb(a []*GoVertex, b []*GoVertex) (c []*GoVertex, new_coming bool) {
// 	for _, v := range b {
// 		if !go_in_xs(v, a) {
// 			a = append(a, v)
// 			new_coming = true
// 		}
// 	}
// 	return a, new_coming
// }
// go_in_vertex_list
// func go_in_xs(a *GoEqualAble, vs []*GoEqualAble) bool {
// 	for _, x := range xs {
// 		if a.Equal(x) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func go_in_xs(a *GoEqualAble, vs []*GoEqualAble) bool {
// 	for _, x := range xs {
// 		if a.Equal(x) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func go_edge_comb(a []*GoEdge, b []*GoEdge) (c []*GoEdge, new_coming bool) {
// 	for _, v := range b {
// 		if !go_in_xs(v, a) {
// 			a = append(a, v)
// 			new_coming = true
// 		}
// 	}
// 	return a, new_coming
// }
