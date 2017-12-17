package main

import (
	// "fmt"
	// "errors"
	// "bufio"
	// "log"
	// "os"
	// "errors"
	"strings"
)

type GoBlock struct {
	data []*GoVertex
}

func NewGoBlock(i int, j int) *GoBlock {
	b := &GoBlock{}
	b.data = append(b.data, &go_vertex_data[i][j])
	return b
}
func (b *GoBlock) Add(v *GoVertex) {
	b.data = append(b.data, v)
}
func (b *GoBlock) Len() int {
	return len(b.data)
}
func (b *GoBlock) String() string {
	ret := make([]string, 0)
	for _, v := range b.data {
		ret = append(ret, v.String())
	}
	return "[" + strings.Join(ret, ",") + "]"
}

func GetGoBlock(i int, j int) *GoBlock {
	ns := neibour_same_color(i, j)
	if len(ns) == 0 {
		return NewGoBlock(i, j)
	} else {
		b := ns[0].block
		b.Add(&go_vertex_data[i][j])
		return b
	}
}

func neibour_same_color(i int, j int) (ret []*GoVertex) {
	v := &go_vertex_data[i][j]
	return get_neibour_by_color(v, v.color)
}
