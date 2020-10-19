package Board

import (
	"fmt"
        "github.com/Ken1f/acquire/settings"
        "github.com/Ken1f/acquire/stocks"
        "github.com/Ken1f/acquire/tile"
)

type Board struct {
	board [Settings.YMAX][Settings.XMAX]int
	//	kingdom KingdomInfo
	stocks Stocks.Stocks
}

type KingdomInfo struct {
	TileTotal [8]int
	Leader    [8]int // empty (0), black (1), blue (2), green (3), red (4)
}

func (b *Board) Init(mapchoice int) {
	if mapchoice == Settings.MAPSTANDARD {
		(*b).InitMapStandard()
	} else if mapchoice == Settings.MAPADVANCE {
		(*b).InitMapAdvance()
	} else {
		(*b).InitMapTest()
	}

	b.stocks.Init()
}

func (b *Board) InitMapStandard() {
	for j := 0; j < Settings.YMAX; j++ {
		for i := 0; i < Settings.XMAX; i++ {
			(*b).SetEmpty(i, j)
		}
	}
	/*	for i := 4; i < 9; i++ {
			(*b).SetRiver(i, 0)
		}
		(*b).SetRiver(12, 0)

		(*b).SetRiver(4, 1)
		(*b).SetRiver(12, 1)

		(*b).SetRiver(3, 2)
		(*b).SetRiver(4, 2)
		(*b).SetRiver(12, 2)
		(*b).SetRiver(13, 2)

		for i := 0; i < 4; i++ {
			(*b).SetRiver(i, 3)
		}
		for i := 13; i < 16; i++ {
			(*b).SetRiver(i, 3)
		}

		(*b).SetRiver(14, 4)
		(*b).SetRiver(15, 4)

		(*b).SetRiver(14, 5)

		for i := 0; i < 4; i++ {
			(*b).SetRiver(i, 6)
		}
		for i := 12; i < 15; i++ {
			(*b).SetRiver(i, 6)
		}

		for i := 3; i < 7; i++ {
			(*b).SetRiver(i, 7)
		}
		(*b).SetRiver(12, 7)

		for i := 6; i < 13; i++ {
			(*b).SetRiver(i, 8)
		}

		(*b).SetTemple(10, 0)
		(*b).SetTemple(1, 1)
		(*b).SetTemple(15, 1)
		(*b).SetTemple(5, 2)
		(*b).SetTemple(13, 4)
		(*b).SetTemple(8, 6)
		(*b).SetTemple(1, 7)
		(*b).SetTemple(14, 8)
		(*b).SetTemple(5, 9)
		(*b).SetTemple(10, 10) */
}

func (b *Board) InitMapAdvance() { // TODO advance map. unfinished
	for j := 0; j < Settings.YMAX; j++ {
		for i := 0; i < Settings.XMAX; i++ {
			(*b).board[j][i] = Tile.TILE["EMPTY"]
		}
	}
}

func (b *Board) InitMapTest() { // Map for test
	/*	(*b).SetTemple(0, 0)
		(*b).SetTemple(0, 1)
		(*b).SetTemple(0, 2)
		(*b).SetTemple(1, 0)
		(*b).SetTemple(1, 2)
		(*b).SetTemple(2, 0)
		(*b).SetTemple(2, 1)
		(*b).SetTemple(2, 2)
		(*b).SetTemple(3, 0)

		(*b).SetFarm(2, 3)
		(*b).SetTemple(2, 4)
		(*b).SetMarket(1, 4)
		(*b).SetMarket(3, 4)
		(*b).SetMarket(4, 4)
		(*b).SetSettlement(5, 4)
		(*b).SetSettlement(6, 4)
	*/
	(*b).SetTile(3, 1, Tile.TILE["RED"])
	(*b).SetTile(6, 5, Tile.TILE["GREEN"])
}

func (b *Board) SetTile(x, y, thisTile int) {
	(*b).board[y][x] = thisTile
}

func (b Board) GetTile(x, y int) int {
	return b.board[y][x]
}

func (b *Board) SetEmpty(x, y int) {
	(*b).board[y][x] = Tile.TILE["EMPTY"]
}
/*
func (b *Board) SetRiver(x, y int) {
	(*b).board[y][x] = TILE["RIVER"]
}

func (b *Board) SetFarm(x, y int) {
	(*b).board[y][x] = TILE["BLUE"]
}

func (b *Board) SetTemple(x, y int) {
	(*b).board[y][x] = TILE["RED"]
}

func (b *Board) SetMarket(x, y int) {
	(*b).board[y][x] = TILE["GREEN"]
}

func (b *Board) SetSettlement(x, y int) {
	(*b).board[y][x] = TILE["BLACK"]
}
*/
func (b *Board) IsEmpty(x, y int) bool {
	if (*b).board[y][x] == Tile.TILE["EMPTY"] {
		return true
	} else {
		return false
	}
}
/*
func (b *Board) IsRiver(x, y int) bool {
	if (*b).board[y][x] == TILE["RIVER"] {
		return true
	} else {
		return false
	}
}
*/

func (b *Board) IsNeutralTile(x, y int) bool {
	switch (*b).board[y][x] {
	case Tile.TILE["BLACK"], Tile.TILE["BLUE"], Tile.TILE["GREEN"], Tile.TILE["RED"]:
		return true
	default:
		return false
	}
}

func (b *Board) IsLeader(x, y int) bool {
	switch (*b).board[y][x] {
	case Tile.TILE["P1BLACK"], Tile.TILE["P1BLUE"], Tile.TILE["P1RED"], Tile.TILE["P1GREEN"],
		Tile.TILE["P2BLACK"], Tile.TILE["P2BLUE"], Tile.TILE["P2RED"], Tile.TILE["P2GREEN"],
		Tile.TILE["P3BLACK"], Tile.TILE["P3BLUE"], Tile.TILE["P3RED"], Tile.TILE["P3GREEN"],
		Tile.TILE["P4BLACK"], Tile.TILE["P4BLUE"], Tile.TILE["P4RED"], Tile.TILE["P4GREEN"]:
		return true
	default:
		return false
	}
}

func (b *Board) PlaceTile(x, y int) bool {
	if b.IsEmpty(x, y) {
		hasNeighbor, thisTile := b.HasNeighbor(x, y)
		if hasNeighbor {
			if thisTile == Tile.TILE["SINGLE"] {
				(*b).SetTile(x, y, thisTile) // TODO: form new corporation
			} else {
				(*b).SetTile(x, y, thisTile)    // set to corporation
				(*b).stocks.Size[0]++ // corporation increase +1   // how to access size from Stock?!
			}
		} else {
			(*b).SetTile(x, y, Tile.TILE["SINGLE"]) // single tile
		}
		return true
	} else { // else can't place tile
		return false
	}
}

func (b *Board) PlaceTilePos(pos int) bool { // linear position input
	y := pos / Settings.YMAX
	x := pos % Settings.YMAX
	fmt.Printf(" x %d y %d ", x, y) //
	return (*b).PlaceTile(x, y)
}

/*
func (b *Board) RemoveTile(x, y int) { // TODO: remove LEADER then add to PLAYER
	switch (*b).board[y][x] {
	case TILE["P1BLUE"], TILE["P2BLUE"], TILE["P3BLUE"], TILE["P4BLUE"]: // remove farmer tile, replace with river
		(*b).board[y][x] = TILE["RIVER"]
	default: // remove normal tile, replace with empty
		(*b).board[y][x] = TILE["EMPTY"]
	}
}
*/

func (b Board) IsLeaderPlaceable(x, y int) bool {
	for j := -1; j < 2; j++ { // assuming center is empty :-)
		for i := -1; i < 2; i++ {
			if inBound(x+i, y+j) && b.board[y+j][x+i] == Tile.TILE["RED"] {
				return true
			}
		}
	}
	return false
}

func (b Board) GetKingdomInfo(x, y int) KingdomInfo { // Get tile total using Flood Fill function
	var mark [][]bool
	var kingdomInfo KingdomInfo

	mark = make([][]bool, Settings.YMAX) // allocating memory for slice of 2D array
	for j := range mark {
		mark[j] = make([]bool, Settings.XMAX)
	}

	b.FloodFill(x, y, mark, &kingdomInfo)

	return kingdomInfo
}

func (b Board) FloodFill(x, y int, mark [][]bool, k *KingdomInfo) {
	if !inBound(x, y) { // quit function if not in bound
		return
	}

	if (b.IsNeutralTile(x, y) || b.IsLeader(x, y)) && mark[y][x] == false { // check connecting neutral & leader tile
		mark[y][x] = true
/*
		switch b.board[y][x] {
		case Tile.TILE["BLACK"]:
			(*k).tileTotal[BLACK]++
		case Tile.TILE["BLUE"]:
			(*k).tileTotal[BLUE]++
		case Tile.TILE["GREEN"]:
			(*k).tileTotal[GREEN]++
		case Tile.TILE["RED"]:
			(*k).tileTotal[RED]++
		case Tile.TILE["YELLOW"]:
			(*k).tileTotal[YELLOW]++
		case Tile.TILE["MAGENTA"]:
			(*k).tileTotal[MAGENTA]++
		case Tile.TILE["CYAN"]:
			(*k).tileTotal[CYAN]++
		case Tile.TILE["WHITE"]:
			(*k).tileTotal[WHITE]++
		}
*/
		b.FloodFill(x, y+1, mark, k) // up
		b.FloodFill(x+1, y, mark, k) // right
		b.FloodFill(x, y-1, mark, k) // down
		b.FloodFill(x-1, y, mark, k) // left
	}
}

func inBound(x, y int) bool {
	if x >= 0 && x < Settings.XMAX && y >= 0 && y < Settings.YMAX {
		return true
	} else {
		return false
	}
}

func (b Board) HasNeighbor(x, y int) (bool, int) {
	if (y+1 < Settings.YMAX) && !b.IsEmpty(x, y+1) {
		return true, b.GetTile(x, y+1)
	} else if (y > 0) && !b.IsEmpty(x, y-1) {
		return true, b.GetTile(x, y+1)
	} else if (x+1 > Settings.XMAX) && !b.IsEmpty(x+1, y) {
		return true, b.GetTile(x, y+1)
	} else if (x > 0) && !b.IsEmpty(x-1, y) {
		return true, b.GetTile(x, y+1)
	} else {
		return false, -1
	}
}

func (b Board) Print() { // 16 wide x 11 height
	fmt.Printf("  ")
	for i := 0; i < Settings.XMAX; i++ {
		fmt.Printf("%2c", i+65) // print Alphabet character from unicode
	}
	fmt.Printf("\n")
	for j := 0; j < Settings.YMAX; j++ {
		fmt.Printf("%2d ", j+1)
		for i := 0; i < Settings.XMAX; i++ {
			Tile.PrintTile(b.board[j][i])
			fmt.Printf(" ")
		}
		fmt.Print("\n")
	}
}

func (b Board) PrintTile(x, y int) { // Testing purpose
	Tile.PrintTile(b.board[y][x])
}
