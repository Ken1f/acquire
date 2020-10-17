package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const XMAX = 12
const YMAX = 9
const MAXBLACK = 30
const MAXBLUE = 36
const MAXGREEN = 30
const MAXRED = 57
const MAXTILES = XMAX * YMAX
const MAXCOLOR = 4
const EMPTY = 0

const (
	BLACK = iota
	BLUE
	GREEN
	RED
	YELLOW
	MAGENTA
	CYAN
	WHITE
)

const (
	COLORSINGLE  = "\033[1;29m%s\033[0m"
	COLORBLACK   = "\033[1;30m%s\033[0m"
	COLORRED     = "\033[1;31m%s\033[0m"
	COLORGREEN   = "\033[1;32m%s\033[0m"
	COLORYELLOW  = "\033[1;33m%s\033[0m"
	COLORBLUE    = "\033[1;34m%s\033[0m"
	COLORMAGENTA = "\033[1;35m%s\033[0m"
	COLORCYAN    = "\033[1;36m%s\033[0m"
	COLORWHITE   = "\033[1;37m%s\033[0m"
)

const (
	MAPSTANDARD = iota
	MAPADVANCE
	MAPTEST
)

const (
	PLAYER1 = iota
	PLAYER2
	PLAYER3
	PLAYER4
)

var TILE map[string]int

func init() {
	TILE = make(map[string]int)

	TILE["EMPTY"] = 0
	TILE["SINGLE"] = 1 // white

	TILE["BLACK"] = 2 //Zeta
	TILE["BLUE"] = 3  //Sackson

	TILE["GREEN"] = 4  //America
	TILE["RED"] = 5    //Hydra
	TILE["YELLOW"] = 6 //Fusion

	TILE["CYAN"] = 7    //Phoenix
	TILE["MAGENTA"] = 8 //Quantum
}

func inBound(x, y int) bool {
	if x >= 0 && x < XMAX && y >= 0 && y < YMAX {
		return true
	} else {
		return false
	}
}

func pr(s string) {
	fmt.Printf(s + "\n")
}

func PrintTile(thisTile int) {
	switch thisTile {
	case TILE["EMPTY"]:
		fmt.Printf(COLORYELLOW, "A")
	case TILE["RED"]:
		fmt.Printf(COLORRED, "T")
	case TILE["GREEN"]:
		fmt.Printf(COLORGREEN, "M")
	case TILE["BLACK"]:
		fmt.Printf(COLORBLACK, "S")
	case TILE["BLUE"]:
		fmt.Printf(COLORBLUE, "F")
	case TILE["YELLOW"]:
		fmt.Printf(COLORYELLOW, "Y")
	case TILE["MAGENTA"]:
		fmt.Printf(COLORMAGENTA, "G")
	case TILE["CYAN"]:
		fmt.Printf(COLORCYAN, "C")
	case TILE["WHITE"]:
		fmt.Printf(COLORWHITE, "W")
	case TILE["SINGLE"]:
		fmt.Printf(COLORSINGLE, "X")
	default:
		fmt.Printf("error")
	}
}

func printGame(p []Player, board Board, bag Bag) { // NOTE: P[] is a slice. Not pointer
	for i, _ := range p {
		p[i].Print()
	}
	board.Print()
	fmt.Print("Remaining Tile ", bag.RemainingTile(), "\n")
}

func printKingdomInfo(k KingdomInfo) {
	color := ""
	fmt.Print("Kingdom info:")
	fmt.Printf(COLORBLACK, " BLACK:")
	fmt.Print(k.tileTotal[BLACK])
	fmt.Printf(COLORBLUE, " BLUE:")
	fmt.Print(k.tileTotal[BLUE])
	fmt.Printf(COLORGREEN, " GREEN:")
	fmt.Print(k.tileTotal[GREEN])
	fmt.Printf(COLORRED, " RED:")
	fmt.Print(k.tileTotal[RED])
	fmt.Printf(COLORYELLOW, " YELLOW:")
	fmt.Print(k.tileTotal[YELLOW])
	fmt.Printf(COLORMAGENTA, " MAGENTA:")
	fmt.Print(k.tileTotal[MAGENTA])
	fmt.Printf(COLORCYAN, " CYAN: ")
	fmt.Print(k.tileTotal[CYAN])
	fmt.Printf(COLORWHITE, " WHITE: ")
	fmt.Print(k.tileTotal[WHITE])

	for i, leader := range k.leader {
		switch i {
		case 0:
			color = COLORBLACK
		case 1:
			color = COLORBLUE
		case 2:
			color = COLORGREEN
		case 3:
			color = COLORRED
		}
		if leader != EMPTY {
			switch leader {
			case PLAYER1 + 1:
				fmt.Printf(color, " P1 ")
			case PLAYER2 + 1:
				fmt.Printf(color, " P2 ")
			case PLAYER3 + 1:
				fmt.Printf(color, " P3 ")
			case PLAYER4 + 1:
				fmt.Printf(color, " P4 ")
			}
		}
	}
	fmt.Printf("\n")
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		fmt.Println("next field", x)
	}
}

func GetSeed() int64 {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	return seed
}

func main() {
	seed := GetSeed()
	fmt.Printf("Hello Acquire: SEED: %d\n", seed)

	var p [2]Player
	var board Board
	var bag Bag
	pr("board init")
	board.Init(MAPSTANDARD)
	pr("bag init")
	bag.Init()
	fmt.Print("Starting Tile ", bag.RemainingTile(), "\n")
	pr("player init")
	p[0].Init(PLAYER1, bag.DrawTiles(6))
	p[1].Init(PLAYER1, bag.DrawTiles(6))
	printGame(p[:], board, bag)

	swap := []int{0, 1, 2} //
	p[0].SwapTiles(bag.DrawTiles(len(swap)), swap)
	swap = []int{0, 1, 2, 3, 4} //
	p[1].SwapTiles(bag.DrawTiles(len(swap)), swap)
	//readInput()
	board.Init(MAPTEST)

	board.PlaceTilePos(16)
	board.PlaceTilePos(3)
	board.PlaceTilePos(12)
	board.PlaceTilePos(20)

	printGame(p[:], board, bag)
	kingdomInfo := board.GetKingdomInfo(1, 1)
	printKingdomInfo(kingdomInfo)
}
