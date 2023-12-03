package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"golang.org/x/crypto/ssh/terminal"
)

type window struct {
	columns int
	rows    int
	fd      int
}

type row struct {
	rowNumber int
	char      rune
}

func (r *row) incrementRowNumber() {
	r.rowNumber += 1
}

func getTerminalAttr() window {

	defTerminal := window{80, 20, int(os.Stdin.Fd())}

	width, height, err := terminal.GetSize(defTerminal.fd)
	if err != nil {
		return defTerminal
	}
	defTerminal.columns = width
	defTerminal.rows = height
	return defTerminal
}

func getFlake(c rune) rune {

	if c != 0 {
		return c
	}

	flakes := []int{10048, 10049, 10050, 10051, 10053, 10054, 10056}

	if runtime.GOOS != "windows" {
		flakes = append(flakes, 10052, 10055)
	}

	flake := flakes[rand.Intn(len(flakes))]

	return rune(flake)
}

func printSnowflakeCol(snowflakes map[int]*row, col int, colour string) {
	printCharacter(col, snowflakes[col].rowNumber, snowflakes[col].char, colour)
}

func printCharacter(x int, y int, character rune, colour string) {
	output := fmt.Sprintf("\033[%d;%dH%c", y, x, character)
	c := color.New(selectColour(colour))
	c.Print(output)
	fmt.Print("\033[1;1H")
	fmt.Print("\033[0m")
}

func getRandomColour() string {
	colours := []string{
		"blue",
		"cyan",
		"green",
		"magenta",
		"red",
		"white",
		"yellow",
	}
	return colours[rand.Intn(len(colours))]
}

func selectColour(colour string) color.Attribute {

	c := strings.ToLower(colour)

	if c == "rainbow" {
		c = getRandomColour()
	}

	switch c {
	case "blue":
		return color.FgHiBlue
	case "cyan":
		return color.FgHiCyan
	case "green":
		return color.FgHiGreen
	case "magenta":
		return color.FgHiMagenta
	case "red":
		return color.FgHiRed
	case "white":
		return color.FgHiWhite
	case "yellow":
		return color.FgHiYellow
	default:
		return color.FgHiWhite
	}
}

func moveFlake(snowflakes map[int]*row, currentRows map[int]int, col int, stack bool, color string, term window, particle rune) {

	currentRow := term.rows

	if stack {

		if _, found := currentRows[col]; !found {
			currentRows[col] = currentRow
		}

		currentRow = currentRows[col]

		if currentRow == 1 {
			currentRow = term.rows
			currentRows[col] = currentRow
		}

		if snowflakes[col].rowNumber+1 == currentRow {
			currentRows[col]--
		}
	}

	// If next row is the end, lets start a new snow flake
	if snowflakes[col].rowNumber+1 == currentRow {

		char := getFlake(particle)
		snowflakes[col] = &row{0, char}
		printSnowflakeCol(snowflakes, col, color)

	} else {

		fmt.Printf("\033[%d;%dH  ", snowflakes[col].rowNumber, col)
		snowflakes[col].incrementRowNumber()
		printSnowflakeCol(snowflakes, col, color)

	}
}

func setParticle(p *string) rune {

	if len(*p) > 0 {
		return []rune(*p)[0]
	}
	return rune(0)
}

func main() {

	snowCmd := flag.NewFlagSet("snow", flag.ExitOnError)
	stackFlag := snowCmd.Bool("stack", false, "Set snow to pile up.")
	speedFlag := snowCmd.Int("speed", 14, "Increase to make it snow faster.")
	particleFlag := snowCmd.String("particle", "", "Change the particle used.")
	colourFlag := snowCmd.String("colour", "white", "Change the colour of the particles. [red|green|blue|magenta|cyan|yellow]")

	treeCmd := flag.NewFlagSet("tree", flag.ExitOnError)
	lightDelayFlag := treeCmd.Int("light-delay", 1, "Seconds between light changes")
	treeColourFlag := treeCmd.String("colour", "green", "Change the colour of the snow particles. [red|green|blue|magenta|cyan|yellow]")
	lightColourFlag := treeCmd.String("light-colour", "rainbow", "Change the color of the lights. [red|green|blue|magenta|cyan|yellow]")
	treeParticleFlag := treeCmd.String("particle", "*", "Change the particle used for the tree.")
	snowFlag := treeCmd.Bool("snow", true, "Whether snow should fall.")
	snowColourFlag := treeCmd.String("snow-colour", "white", "Change the colour of the snow particles. [red|green|blue|magenta|cyan|yellow]")
	snowParticleFlag := treeCmd.String("snow-particle", "", "Change the snow particle used.")
	snowSpeedFlag := treeCmd.Int("snow-speed", 14, "Increase to make it snow faster.")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		screen.Clear()
		os.Exit(0)
	}()

	mode := strings.ToLower(os.Args[1])
	if mode != "tree" && mode != "snow" {
		fmt.Println("supported commands are 'snow' and 'tree'")
		os.Exit(1)
	}

	switch mode {
	case "tree":
		treeCmd.Parse(os.Args[2:])
		particle := setParticle(snowParticleFlag)
		treeParticle := setParticle(treeParticleFlag)
		tree := &tree{
			lightDelay:   *lightDelayFlag,
			colour:       *treeColourFlag,
			lightsColour: *lightColourFlag,
			particle:     treeParticle,
			snow:         *snowFlag,
			snowColour:   *snowColourFlag,
			snowSpeed:    *snowSpeedFlag,
			snowParticle: particle}
		tree.show()
	default:
		snowCmd.Parse(os.Args[2:])
		particle := setParticle(particleFlag)
		snow := &snow{*speedFlag, *stackFlag, particle, *colourFlag}
		snow.show()
	}
}
