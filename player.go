package main

import (
	"fmt"
)

type Player struct {
	playerNum int
	tiles     [6]int
	//	leader    [4]int
	//	points    [4]int
	money  int
	stocks [7]int
}

func (p *Player) Init(playerNum int, tiles []int) { // randomized 5 tiles
	(*p).playerNum = playerNum
	for i, _ := range p.tiles {
		(*p).tiles[i] = tiles[i] // take tiles from bag to add to player tile
	}

	for i := range (*p).stocks {
		(*p).stocks[i] = 0
	}

	(*p).money = 3000
}

func (p *Player) AddTile(thisTile int) {
	for i, val := range p.tiles { // add tile to 1st empty index
		if val == TILE["EMPTY"] {
			(*p).tiles[i] = thisTile
			break
		}
	}
}

func (p *Player) DrawTiles(thisTile int) {
	(*p).AddTile(thisTile) //
}

func (p *Player) SwapTiles(thisTiles, swapList []int) {
	for i := 0; i < len(swapList); i++ {
		//(p).tiles[swapList[i]] = (*b).DrawTile()
		(p).tiles[swapList[i]] = thisTiles[i]
	}
}

/*
func (p *Player) AddPoint(thisTile int) { // TODO check if leader exist to collect points
	(*p).points[(thisTile-1)%MAXCOLOR]++
}
*/

func (p Player) Print() {
	PrintTile(p.stocks[0]) // black   /// was LEADER
	PrintTile(p.stocks[1]) // blue
	PrintTile(p.stocks[2]) // green
	PrintTile(p.stocks[3]) // red
	PrintTile(p.stocks[4]) // red  //
	PrintTile(p.stocks[5]) // red  //
	PrintTile(p.stocks[6]) // red  //

	fmt.Printf(" ")

	fmt.Printf("\033[1;30m%d\033[0m", p.stocks[0]) // black
	fmt.Printf("\033[1;34m%d\033[0m", p.stocks[1]) // blue
	fmt.Printf("\033[1;32m%d\033[0m", p.stocks[2]) // green
	fmt.Printf("\033[1;31m%d\033[0m", p.stocks[3]) // red
	fmt.Printf("\033[1;34m%d\033[0m", p.stocks[4]) // blue  //
	fmt.Printf("\033[1;32m%d\033[0m", p.stocks[5]) // green  //
	fmt.Printf("\033[1;31m%d\033[0m", p.stocks[6]) // red   //

	fmt.Printf(" ")
	for i := 0; i < 5; i++ {
		PrintTile(p.tiles[i])
	}
	fmt.Printf("\n")

	fmt.Printf("Money: %d]\n", p.money)
}
