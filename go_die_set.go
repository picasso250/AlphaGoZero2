// 死活 之边的集合 与 点的集合

package main

import (
	"fmt"
	"strings"
	// "errors"
	// "log"
)

type GoEdgeSet struct {
	m map[[2]*GoVertex]int
}

func NewGoEdgeSet() (s *GoEdgeSet) {
	m := make(map[[2]*GoVertex]int)
	return &GoEdgeSet{m}
}
func (s *GoEdgeSet) String() string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, fmt.Sprintf("%s-%s", k[0], k[1]))
	}
	return fmt.Sprintf("(%d)[%s]", len(keys), strings.Join(keys, ", "))
}
func (s *GoEdgeSet) Get(v1 *GoVertex, v2 *GoVertex) (int, bool) {
	type_, ok := s.m[[2]*GoVertex{v1, v2}]
	if ok {
		return type_, ok
	}
	type_, ok = s.m[[2]*GoVertex{v2, v1}]
	return type_, ok
}
func (s *GoEdgeSet) Add(e *GoEdge) bool {
	_, ok := s.Get(e.v1, e.v2)
	if !ok {
		s.m[[2]*GoVertex{e.v1, e.v2}] = e.type_
	}
	return !ok // 是否添加了新元素
}
func (s *GoEdgeSet) AddByVertex(v *GoVertex) bool {
	is_add := false
	for k := 0; k < 4; k++ {
		is_ := s.Add(&v.edge[k])
		// fmt.Printf("add edge %s-%s (add?=%v)\n",v.edge[k].v1, v.edge[k].v2,is_)
		is_add = is_add || is_
	}
	// fmt.Printf("AddByVertex is_add=%v\n",is_add)
	return is_add // 是否添加了新元素
}

// 点的集合
type GoVertexSet struct {
	m map[*GoVertex]bool
}

func NewGoVertexSet() (s *GoVertexSet) {
	m := make(map[*GoVertex]bool)
	return &GoVertexSet{m}
}
func (s *GoVertexSet) String() string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k.String())
	}
	return fmt.Sprintf("(%d)[%s]", len(keys), strings.Join(keys, ", "))
}
func (s *GoVertexSet) Get(v *GoVertex) (bool, bool) {
	v_, ok := s.m[v]
	return v_, ok
}
func (s *GoVertexSet) Add(v *GoVertex) bool {
	_, ok := s.m[v]
	if !ok {
		s.m[v] = true
	}
	return !ok // 是否添加了新元素
}
func (s *GoVertexSet) AddEdge(e *GoEdge) bool {
	is_add1 := false
	is_add2 := false
	if e.v1 != nil {
		is_add1 = s.Add(e.v1)
	}
	if e.v2 != nil {
		is_add2 = s.Add(e.v2)
	}
	return is_add1 || is_add2 // 是否添加了新元素
}
func (s *GoVertexSet) AddEdgeByTwoVertex(vs [2]*GoVertex) bool {
	is_add1 := false
	is_add2 := false
	if vs[0] != nil {
		is_add1 = s.Add(vs[0])
	}
	if vs[1] != nil {
		is_add2 = s.Add(vs[1])
	}
	return is_add1 || is_add2 // 是否添加了新元素
}
