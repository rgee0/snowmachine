package main

import (
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

type tree struct {
	lightDelay   int
	colour       string
	lightsColour string
	particle     rune
	snow         bool
	snowColour   string
	snowParticle rune
	snowSpeed    int
}

type part struct {
	x        int
	y        int
	particle rune
	colour   string
}

func setoutTree(p *[]part, particle rune, colour string, height int, term window) {

	j := 1
	for y := height / 2; y <= height; y++ {
		for k := 0; k < j; k++ {
			x := int((term.columns/2)-(j/2)) + k
			*p = append(*p, part{x, y, particle, colour})
		}
		j += 2
	}
}

func repaintTree(t []part, lc string) []part {

	repaintedTree := []part{}
	repaintedTree = append(repaintedTree, t...)

	for i, p := range t {

		finalParticle := p.particle
		finalColour := p.colour

		if chance := rand.Intn(100); chance < 10 {
			finalParticle = rune('o')
			if finalColour = lc; len(finalColour) == 0 || finalColour == "rainbow" {
				finalColour = getRandomColour()
			}
		}

		repaintedTree[i] = part{p.x, p.y, finalParticle, finalColour}
	}
	return repaintedTree
}

func setoutTrunk(p *[]part, treeHeight int, height int, term window) {
	trunkWidth := 9
	for y := treeHeight; y <= treeHeight+height; y++ {
		for k := 0; k < trunkWidth; k++ {
			x := int((term.columns/2)-(trunkWidth/2)) + k
			*p = append(*p, part{x, y, 'w', "yellow"})
		}
	}
}

func (t *tree) show() {

	screen.Clear()
	term := getTerminalAttr()

	treeParts := []part{}
	trunkParts := []part{}

	particle := getFlake(t.particle)
	trunkHeight := 3
	treeHeight := term.rows - trunkHeight

	setoutTree(&treeParts, particle, t.colour, treeHeight, term)
	setoutTrunk(&trunkParts, treeHeight, trunkHeight, term)

	currentRows := make(map[int]int)
	snowflakes := make(map[int]*row)

	newTreeParts := []part{}
	newTreeParts = append(newTreeParts, treeParts...)

	start := time.Now()

	for {

		if t.snow {

			targetColumn := rand.Intn(term.columns)

			if targetColumn%2 == 0 {
				continue
			}

			if _, onScreen := snowflakes[targetColumn]; onScreen {

				moveFlake(snowflakes,
					currentRows,
					targetColumn,
					false,
					t.snowColour,
					term,
					t.snowParticle)

			} else {

				flake := getFlake(t.snowParticle)
				snowflakes[targetColumn] = &row{0, flake}
				printSnowflakeCol(snowflakes, targetColumn, t.snowColour)

			}
			for existingFlake := range snowflakes {

				moveFlake(snowflakes,
					currentRows,
					existingFlake,
					false,
					t.snowColour,
					term,
					t.snowParticle)
			}
		}

		end := time.Now()
		diff := end.Sub(start).Seconds()

		if diff > float64(t.lightDelay) {

			newTreeParts = repaintTree(treeParts, t.lightsColour)
			start = time.Now()

		}

		wholeTree := append(trunkParts, newTreeParts...)
		for _, p := range wholeTree {
			printCharacter(p.x, p.y, p.particle, p.colour)
		}

		time.Sleep(time.Second / time.Duration(t.snowSpeed))
	}
}
