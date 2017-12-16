package main

// 数据结构
import (
	// "fmt"
	// "strings"
	// "errors"
	"log"
	"os"
	"testing"
)

func xTestTree(t *testing.T) {
	GoOneMove(1, 1, WHITE)
	f, err := os.OpenFile("test_tree.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	GoDrawTree(f)
}
