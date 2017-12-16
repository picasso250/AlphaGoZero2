package main

// 数据结构
import (
	"fmt"
	"time"
	// "strings"
	// "errors"
	"log"
	"os"
	// "strconv"
	"text/template"
)

type BoardDraw struct {
	Size int
	Id string
	Width int
	Height int
	DrawCode []string
}

func GoDrawBoardData(s *GoBoardShot) BoardDraw {
	log.SetFlags(log.Lshortfile)
	size := (BOARD_SIZE+1)*20
	drawCode := make([]string, 0)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (s[i][j] != NONE) {
				drawCode = append(drawCode, fmt.Sprintf("drawQizi(ctx,%d,%d,WHITE);", i, j))
			}
		}
	}
	data := BoardDraw{
		 (BOARD_SIZE),
		 fmt.Sprintf("c%ld", time.Now().UnixNano()),
		(size),
		(size),
		drawCode,
	}
	return data
}

type GoPlayTree struct {
	boardShot GoBoardShot
	children []*GoPlayTree
}

func GoDrawTree(f *os.File) {
	tmpl := template.Must(template.ParseFiles("tree.html", "board.html"))

	s,_:=GoGetBoardShot()
	data0 :=GoDrawBoardData(&s)

	GoOneMove(0,0,BLACK)
	s,_=GoGetBoardShot()
	data1 :=GoDrawBoardData(&s)
	undo_one_step()

	GoOneMove(0,1,BLACK)
	s,_=GoGetBoardShot()
	data2 :=GoDrawBoardData(&s)
	undo_one_step()

	GoOneMove(0,3,BLACK)
	s,_=GoGetBoardShot()
	data3 :=GoDrawBoardData(&s)
	undo_one_step()

	data:= map[string]BoardDraw{
		"start": data0,
		"a1": data1,
		"a2": data2,
		"a3": data3,
	}
	err := tmpl.ExecuteTemplate(f, "tree", data)
	if err != nil {
		log.Fatal(err)
	}
}
