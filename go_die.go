// 死活

package main

import (
	// "fmt"
	// "strings"
	// "errors"
	"log"
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
//   1. 找出V所在棋块中所有的同色棋子V_c
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
type GoVertex struct {
	i     int
	j     int
	color go_color
	edge  [4]*GoEdge
}

// 边
// 有可能其中一个点在棋盘外，此时这个点用nil表示
type GoEdge struct {
	v1  *GoVertex
	v2  *GoVertex
	typ int
}

// 边的类型
const (
	CONNECT = 0
	BLOCK   = 1
	OPEN    = 2
)

type GoEqualAble interface {
	Equal() bool
}

var go_edge_data []GoEdge
var go_vertex_data [BOARD_SIZE][BOARD_SIZE]GoVertex

func (v *GoVertex) Equal(w *GoVertex) bool {
	return v.i == w.i && v.j == w.j
}
func (v *GoVertex) IsAnyOpenEdge() bool {
	for k := 0; k < 4; k++ {
		if v.edge[k].typ == OPEN {
			return true
		}
	}
	return false
}
func (v *GoVertex) GetAllEdgeSlice() (es []*GoEdge) {
	for k := 0; k < 4; k++ {
		es = append(es, v.edge[k])
	}
	return es
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
func (e *GoEdge) UpdateByV(i int, j int) {
	e._updateEdge(i, j)
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

func (e *GoEdge) _updateEdge(i int, j int) {
	if e.IsSameColor() {
		e.typ = CONNECT
	} else {
		e.typ = BLOCK
	}
}

func go_vetex_data_init() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			go_vertex_data[i][j] = GoVertex{i: i, j: j}
		}
	}
}

func goNewEdge(v1 *GoVertex, v2 *GoVertex) *GoEdge {
	var t int
	if v2.color == NONE {
		t = OPEN
	} else if v1.color == v2.color {
		t = CONNECT
	} else {
		t = BLOCK
	}
	e := GoEdge{v1, v2, t}
	go_connect_to(v2, v1, &e)
	return &e
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
func go_update_edge_(v *GoVertex, index int, i int, j int) *GoEdge {
	if v.edge[index] == nil {
		v.edge[index] = goNewEdgeVoid(v, i, j)
	} else {
		v.edge[index].UpdateByV(i, j)
	}
	return v.edge[index]
}
func go_update_edge_r(index int, i int, j int, e *GoEdge) {
	if go_pos_in_board(i, j) {
		go_vertex_data[i][j].edge[index] = e
	}
}

// 新建边（边有可能在棋盘外，故处理这种情况）
func goNewEdgeVoid(v *GoVertex, i int, j int) (e *GoEdge) {
	if !go_pos_in_board(i, j) {
		e = &GoEdge{v, nil, BLOCK}
	} else {
		e = goNewEdge(v, &go_vertex_data[i][j])
	}
	go_edge_data = append(go_edge_data, *e)
	return e
}
func go_connect_to(v_lost_edge *GoVertex, v *GoVertex, e *GoEdge) {
	index := go_get_edge_index_by(v_lost_edge, v)
	if index == -1 {
		log.Fatal("no good index")
	}
	v_lost_edge.edge[index] = e
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

// 找到和(i,j)相关的边
func go_find_all_edges_about(i int, j int) (es []*GoEdge) {
	for k := 0; k < 4; k++ {
		es = append(es, go_vertex_data[i][j].edge[k])
	}
	return es
}

// 气=qi(M)=count{V|V所属的E是开类型}
// func go_get_qi(i int, j int) int {
// 	assert(go_data[i][j] != NONE)
// 	vs, es := go_get_all_block_about(i, j)
// 	return go_count_color_and_open(vs)
// }
// func go_count_color_and_open(vs []*GoVertex) int {
// 	q := 0
// 	for _, v := range vs {
// 		if v.IsAnyOpenEdge() && v.color == NONE {
// 			q++
// 		}
// 	}
// 	return q
// }

// 获得棋子块相关的所有边和点
// func go_get_all_block_about(i int, j int) ([]*GoVertex, []*GoEdge) {
// 	vs := make([]*GoVertex, 1)
// 	vs[0] = &go_vertex_data[i][j]
// 	es := go_find_all_edges_about(i, j)
// 	vs, es = go_get_all_block_about_color_iter(vs, es)
// 	vs = go_add_not_color_edge(vs, es)
// 	return vs, es
// }
// func go_add_not_color_edge(vs []*GoVertex, es []*GoEdge) (r_vs []*GoVertex) {
// 	var vs_ []*GoVertex
// 	for _, e := range es {
// 		vs_ = append(vs_, e.GetAllPointSlice()...)
// 	}
// 	r_vs, _ = go_vertex_comb(vs, vs_)
// 	return r_vs
// }

// 获得棋子块相关的所有边和点（同色）
// func go_get_all_block_about_color_iter(vs []*GoVertex, es []*GoEdge) ([]*GoVertex, []*GoEdge) {
// 	var vs_ []*GoVertex
// 	var es_ []*GoEdge
// 	for _, e := range es {
// 		if e.typ == CONNECT {
// 			vs_ = append(vs_, e.GetAllPointSlice()...)
// 		} else {
// 			return vs_, es_
// 		}
// 	}
// 	vs, new_coming_v := go_vertex_comb(vs, vs_)
// 	if !new_coming_v {
// 		return vs, es
// 	}
// 	for _, v := range vs {
// 		es_ = append(es_, v.GetAllEdgeSlice()...)
// 	}
// 	es, _ = go_edge_comb(es, es_)
// 	return go_get_all_block_about_color_iter(vs, es)
// }
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
