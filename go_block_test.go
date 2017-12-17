package main

import (
	// "fmt"
	// "errors"
	// "bufio"
	// "log"
	// "os"
	// "errors"
	// "strings"
	"testing"
)

func TestGetGoBlock(t *testing.T) {
	go_vetex_data_init()
	err := one_move_(1, 1, BLACK)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	block := GetGoBlock(1, 1)
	v := block.data[0]
	if v.i != 1 {
		t.Errorf("v.i should be 1 but %d\n", v.i)
	}
	if v.j != 1 {
		t.Errorf("v.j should be 1 but %d\n", v.j)
	}
}
