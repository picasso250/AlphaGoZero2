package main

import (
	// "fmt"
	// "errors"
	// "bufio"
	// "log"
	// "os"
	// "errors"
	// "strings"
)
type GoBlock struct {
  data []*GoVertex
}
func NewGoBlock(i int,j int) *GoBlock {
  b:=&GoBlock{}
  b.data = append(b.data, &go_vertex_data[i][j])
  return b
}
func (b *GoBlock) Add(v *GoVertex)  {
  b.data = append(b.data, v)
}

func GetGoBlock(i int,j int) *GoBlock {
  ns := neibour_same_color(i,j)
  if len(ns) == 0 {
    return NewGoBlock(i,j)
  } else {
    b:=ns[0].block
    b.Add(&go_vertex_data[i][j])
    return b
  }
}

func neibour_same_color(i int,j int) (ret []*GoVertex) {
  v:=&go_vertex_data[i][j]
  return get_neibour_by_color(v, v.color)
}
