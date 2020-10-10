package main

import (
	"math/rand"
)

type Bag struct {
	tiles [MAXTILES]int
	total int
}

func (b *Bag) Init() {
	for i := range (*b).tiles {
		(*b).tiles[i] = i
	}

	b.total = MAXTILES
}

func (b *Bag) DrawTile() int { // draw 1 tile
	pos := rand.Intn(b.total + 1)
	lastpos := b.total - 1 // b.total is same last position

	drawtile := (*b).tiles[pos]
	(*b).tiles[pos] = (*b).tiles[lastpos]

	b.total--
	return drawtile
}

func (b *Bag) DrawTiles(numTile int) []int { // draw multiple tiles
	var tiles []int

	for i := 0; i < numTile; i++ {
		tiles = append(tiles, (*b).DrawTile())
	}
	return tiles
}

func (b Bag) RemainingTile() int {
	return b.total
}
