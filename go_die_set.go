// 死活 之边的集合 与 点的集合

package main

import (
	// "fmt"
	// "strings"
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
func (s *GoEdgeSet) Get(v1 *GoVertex, v2 *GoVertex) (int, bool) {
	type_, ok := s.m[[2]*GoVertex{v1, v2}]
	if ok {
		return type_, ok
	}
	type_, ok = s.m[[2]*GoVertex{v2, v1}]
	return type_, ok
}
func (s *GoEdgeSet) Add(v1 *GoVertex, v2 *GoVertex, type_ int) bool {
	_, ok := s.Get(v1,v2)
	if !ok {
		s.m[[2]*GoVertex{v1, v2}] = type_
	}
	return !ok // 是否添加了新元素
}

type GoVertexSet struct {
	m map[[2]int]go_color
}

func NewGoVertexSet() (s *GoVertexSet) {
	m :=make(map[[2]int]go_color)
	return &GoVertexSet{m}
}
func (s *GoVertexSet) Get(i int, j int) (go_color, bool) {
	color, ok := s.m[[2]int{i, j}]
	if ok {
		return color, ok
	}
	color, ok = s.m[[2]int{j, i}]
	return color, ok
}
func (s *GoVertexSet) Add(i int, j int, color go_color) bool {
	_, ok := s.Get(i, j)
	if !ok {
		s.m[[2]int{j, i}] = color
	}
	return !ok // 是否添加了新元素
}
