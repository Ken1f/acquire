package Tile

import (
        "fmt"
)

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
