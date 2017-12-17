package main

import (
	// "fmt"
	// "strings"
	// "errors"
	// "log"
	"testing"
)

func xTestGoOneStep(t *testing.T) {
	go_vetex_data_init()
	err := GoOneMove(1, 1, WHITE)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if len(go_play_seq) == 0 {
		t.Errorf("go_play_seq's length is 0")
	}
	undo_one_step()
	if len(go_play_seq) != 0 {
		t.Errorf("go_play_seq's length is not 0")
	}
}
