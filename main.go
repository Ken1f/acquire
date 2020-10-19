package main

import (
	"bufio"
	"fmt"
	"github.com/Ken1f/acquire/bag"
	"github.com/Ken1f/acquire/board"
	"github.com/Ken1f/acquire/player"
	"math/rand"
	"os"
	"strings"
	"time"
	//"github.com/Ken1f/acquire/stocks"
	"github.com/Ken1f/acquire/settings"
	"github.com/Ken1f/acquire/tile"
)

func pr(s string) {
	fmt.Printf(s + "\n")
}

func printGame(p []Player.Player, board Board.Board, bag Bag.Bag) { // NOTE: P[] is a slice. Not pointer
	for i, _ := range p {
		p[i].Print()
	}
	board.Print()
	fmt.Print("Remaining Tile ", bag.RemainingTile(), "\n")
}

func printKingdomInfo(k Board.KingdomInfo) {
	color := ""
	fmt.Print("Kingdom info:")
	fmt.Printf(Tile.COLORBLACK, " BLACK:")
	fmt.Print(k.TileTotal[Tile.BLACK])
	fmt.Printf(Tile.COLORBLUE, " BLUE:")
	fmt.Print(k.TileTotal[Tile.BLUE])
	fmt.Printf(Tile.COLORGREEN, " GREEN:")
	fmt.Print(k.TileTotal[Tile.GREEN])
	fmt.Printf(Tile.COLORRED, " RED:")
	fmt.Print(k.TileTotal[Tile.RED])
	fmt.Printf(Tile.COLORYELLOW, " YELLOW:")
	fmt.Print(k.TileTotal[Tile.YELLOW])
	fmt.Printf(Tile.COLORMAGENTA, " MAGENTA:")
	fmt.Print(k.TileTotal[Tile.MAGENTA])
	fmt.Printf(Tile.COLORCYAN, " CYAN: ")
	fmt.Print(k.TileTotal[Tile.CYAN])
	fmt.Printf(Tile.COLORWHITE, " WHITE: ")
	fmt.Print(k.TileTotal[Tile.WHITE])

	for i, leader := range k.Leader {
		switch i {
		case 0:
			color = Tile.COLORBLACK
		case 1:
			color = Tile.COLORBLUE
		case 2:
			color = Tile.COLORGREEN
		case 3:
			color = Tile.COLORRED
		}
		if leader != Settings.EMPTY {
			switch leader {
			case Settings.PLAYER1 + 1:
				fmt.Printf(color, " P1 ")
			case Settings.PLAYER2 + 1:
				fmt.Printf(color, " P2 ")
			case Settings.PLAYER3 + 1:
				fmt.Printf(color, " P3 ")
			case Settings.PLAYER4 + 1:
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

	var p [2]Player.Player
	var board Board.Board
	var bag Bag.Bag
	pr("board init")
	board.Init(Settings.MAPSTANDARD)
	pr("bag init")
	bag.Init()
	fmt.Print("Starting Tile ", bag.RemainingTile(), "\n")
	pr("player init")
	p[0].Init(Settings.PLAYER1, bag.DrawTiles(6))
	p[1].Init(Settings.PLAYER1, bag.DrawTiles(6))
	printGame(p[:], board, bag)

	swap := []int{0, 1, 2} //
	p[0].SwapTiles(bag.DrawTiles(len(swap)), swap)
	swap = []int{0, 1, 2, 3, 4} //
	p[1].SwapTiles(bag.DrawTiles(len(swap)), swap)
	//readInput()
	board.Init(Settings.MAPTEST)

	board.PlaceTilePos(16)
	board.PlaceTilePos(3)
	board.PlaceTilePos(12)
	board.PlaceTilePos(20)

	printGame(p[:], board, bag)
	kingdomInfo := board.GetKingdomInfo(1, 1)
	printKingdomInfo(kingdomInfo)
}
