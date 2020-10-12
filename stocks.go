package main

import (
	"fmt"
)

const MAXCORP = 7  // max corporation
const MAXSHARES = 25

type Stocks struct {
	total  [MAXCORP]int     // total of available stock
        value  [MAXCORP]int     // value of stock
        size   [MAXCORP]int     // size of company
        active [MAXCORP]bool    // company active
}

func (s *Stocks) Init() { // randomized 5 tiles
	for i := range (*s).total {
		(*s).total[i] = MAXSHARES
	}

        for i := range (*s).value {
                (*s).value[i] = 100        // *TO DO. set stocks for each corporation
        }

        for i := range (*s).size {
                (*s).size[i] = 0
        }


        for i := range (*s).active {
                (*s).active[i] = false
        }
}

func (s *Stocks) Value(thisStock int) int {

        bonus := 0                              // Corporation 0,1
        if thisStock >= 2 && thisStock <= 4 {
                bonus = 100                     // Corporation 2,3,4
        } else if thisStock >= 5 {
                bonus = 200                     // Corporation 5,6
        }

        switch (s.size[thisStock]) {
                case 0, 1:
                        return 100   // TODO: is 0,1 worth $100?
                case 2:
                        return 200 + bonus
                case 3:
                        return 300 + bonus
                case 4:
                        return 400 + bonus
                case 5:
                        return 500 + bonus
                case 6,7,8,9,10:
                        return 600 + bonus
                case 11,12,13,14,15,16,17,18,19,20:
                        return 700 + bonus
                case 21,22,23,24,25,26,27,28,29,30:
                        return 800 + bonus
                case 31,32,33,34,35,36,37,38,39,40:
                        return 900 + bonus
                default:        // 41+ size
                        return 1000 + bonus
        }
}

func (s *Stocks) UpdateValue(thisStock int) {
        (*s).value[thisStock] = s.Value(thisStock)
}

func (s *Stocks) SizeIncrease(thisStock int) {
        (*s).size[thisStock]++
}

func (s *Stocks) SizeDecrease(thisStock int) {
        (*s).size[thisStock]--
}

func (s *Stocks) Size(thisStock, size int) {
        (*s).value[thisStock] = s.value[thisStock] + size
}

func (s *Stocks) Buy(thisStock int) {
	(*s).total[thisStock]--
}

func (s *Stocks) Sell(thisStock int) {
	(*s).total[thisStock]++
}

func (s Stocks) Print() {
	fmt.Printf(" ")

	fmt.Printf("\033[1;30m%d %d\033[0m", s.total[0], s.value[0]) // black
	fmt.Printf("\033[1;34m%d %d\033[0m", s.total[1], s.value[1]) // blue
	fmt.Printf("\033[1;32m%d %d\033[0m", s.total[2], s.value[2]) // green
	fmt.Printf("\033[1;31m%d %d\033[0m", s.total[3], s.value[3]) // red
	fmt.Printf("\033[1;34m%d %d\033[0m", s.total[4], s.value[4]) // blue  //
	fmt.Printf("\033[1;32m%d %d\033[0m", s.total[5], s.value[5]) // green  //
	fmt.Printf("\033[1;31m%d %d\033[0m", s.total[6], s.value[6]) // red   //
	fmt.Printf("\n")
}
