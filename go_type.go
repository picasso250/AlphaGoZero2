package main

// 数据结构 及 类型
import (
	"fmt"
	// "strings"
	// "errors"
	// "log"
)

// 棋子颜色
type GoColor byte

// 0 无棋子
// 1 黑子
// 2 白子
// Golang will init all value to zero
// https://golang.org/ref/spec#Variable_declarations
const (
	NONE  = 0
	BLACK = 1
	WHITE = 2
)

func (c GoColor) Reverse() GoColor {
	return BLACK + WHITE - c
}

// 顶点（代表棋子）
// edge中不会有nil元素，即使它是壁，此时为闭类型
type GoVertex struct {
	i     int
	j     int
	color GoColor
	edge  [4]GoEdge // 相连的4条边 默认类型为开
	block *GoBlock
}

func (v *GoVertex) ColorStr() string {
	var GoColor_repr_map = GoColorMap()
	return GoColor_repr_map[v.color]
}
func GoColorMap() [3]string {
	return [3]string{".", "x", "o"}
}
func GoColorMapRev() map[string]GoColor {
	m := make(map[string]GoColor, 3)
	for k, v := range GoColorMap() {
		m[v] = GoColor(k)
	}
	return m
}

// 对空点，找出其所属的势力（颜色）
func (v *GoVertex) GetOwnerColor() (color GoColor) {
	assert(v.color == NONE)
	m := make(map[GoColor]bool)
	for k := 0; k < 4; k++ {
		neibour := v.edge[k].v2
		if neibour != nil {
			m[neibour.color] = true
			color = neibour.color
		}
	}
	if len(m) == 1 {
		return color
	}
	return NONE
}
func (v *GoVertex) ToStr() string {
	return fmt.Sprintf("(%d,%d)", v.i, v.j)
}
func (v *GoVertex) String() string {
	return fmt.Sprintf("|%s(%d,%d)", v.ColorStr(), v.i, v.j)
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
	OPEN    = 0 // 开，至少一个点为空
	CONNECT = 1 // 连，同色相连
	BLOCK   = 2 // 闭，异色相连 或者 碰壁
)

func (e *GoEdge) String() string {
	m := [...]string{"开", "连", "挡"}
	return fmt.Sprintf("%s-%s\\(%s)", e.v1, e.v2, m[e.type_])
}
func (e *GoEdge) Reverse() GoEdge {
	return GoEdge{e.v2, e.v1, e.type_}
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
	if e.v2 == nil {
		e.type_ = BLOCK
	} else {
		// 当悔棋的时候，色可能为空
		if e.v1.color == NONE {
			e.type_ = OPEN
		} else {
			if e.v2.color == NONE {
				e.type_ = OPEN
			} else {
				if e.v1.color == e.v2.color {
					e.type_ = CONNECT
				} else {
					e.type_ = BLOCK
				}
			}
		}
	}
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
