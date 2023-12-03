package main

import (
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

type snow struct {
	speed    int
	stack    bool
	particle rune
	colour   string
}

func (s *snow) show() {

	screen.Clear()
	term := getTerminalAttr()

	currentRows := make(map[int]int)
	snowflakes := make(map[int]*row)

	for {

		targetColumn := rand.Intn(term.columns)

		if targetColumn%2 == 0 {
			continue
		}

		if _, onScreen := snowflakes[targetColumn]; onScreen {

			moveFlake(snowflakes,
				currentRows,
				targetColumn,
				s.stack,
				s.colour,
				term,
				s.particle)

		} else {

			flake := getFlake(s.particle)
			snowflakes[targetColumn] = &row{0, flake}
			printSnowflakeCol(snowflakes, targetColumn, s.colour)
		}

		for existingFlake := range snowflakes {
			moveFlake(snowflakes,
				currentRows,
				existingFlake,
				s.stack,
				s.colour,
				term,
				s.particle)
		}

		time.Sleep(time.Second / time.Duration(s.speed))
	}
}
