package main

import (
	"fmt"
)

const MAXCORP = 7  // max corporation

type Stocks struct {
	total  [MAXCORP]int
        value  [MAXCORP]int
}

func (s *Stocks) Init() { // randomized 5 tiles
	for i := range (*s).total {
		(*s).total[i] = 25
	}

        for i := range (*s).value {
                (*s).value[i] = 100        // *TO DO. set stocks for each corporation
        }
}

func (s *Stocks) Buy(thisStock int) {
	(*s).total[thisStock]--

        if (*s).total[thisStock] > 0 {  // TO DO. set stocks for each corporation
                (*s).value[thisStock]++
        }
}

func (s *Stocks) Sell(thisStock int) {
	(*s).total[thisStock]++

        if (*s).total[thisStock] > 0 {  // TO DO. set stocks for each corporation
                (*s).value[thisStock]--
        }

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
