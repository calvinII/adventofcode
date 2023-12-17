package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card byte
type cards string

func (c card) ToVal() int {
	cards := map[card]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	return cards[c]
}

type DeckType int

const (
	failure DeckType = iota
	high_card
	one_pair
	two_pair
	three_of_a_kind
	full_house
	four_of_a_kind
	five_of_a_kind
)

type deck struct {
	hand cards
	bid  int64
	dt   DeckType
}

func (s cards) GetDeckType() DeckType {
	same := make(map[rune]int)
	for _, c := range s {
		same[c]++
	}

	// five_of_a_kind
	if len(same) == 1 {
		fmt.Println(s, ": five_of_a_kind")
		return five_of_a_kind
	}

	// four_of_a_kind | full_house
	if len(same) == 2 {
		for _, v := range same {
			if v == 1 || v == 4 {
				fmt.Println(s, ": four_of_a_kind")
				return four_of_a_kind
			}
			if v == 2 || v == 3 {
				fmt.Println(s, ": full_house")
				return full_house
			}
		}
	}

	// two_pair | three_of_a_kind
	if len(same) == 3 {
		for _, v := range same {
			if v == 2 {
				fmt.Println(s, ": two_pair")
				return two_pair
			}
			if v == 3 {
				fmt.Println(s, ": three_of_a_kind")
				return three_of_a_kind
			}
		}
	}
	// one_pair
	if len(same) == 4 {
		fmt.Println(s, ": one_pair")
		return one_pair
	}
	// high_card
	if len(same) == 5 {
		fmt.Println(s, ": high_card")
		return high_card
	}
	fmt.Println(s, ": failure")
	return failure
}

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

func lesserDeck(d1,d2 deck) bool {
	if d1.dt == d2.dt {
		for i := 0 ; i < len(d1.hand); i++ {
			if d1.hand[i] == d2.hand[i] {continue}
			return card(d1.hand[i]).ToVal() < card(d2.hand[i]).ToVal()
		}
	} else {
		return d1.dt < d2.dt
	}
	fmt.Println("ERROR SORTING")
	return false
}
func main() {
	game := []deck{}
	for _, l := range readLines() {
		f := strings.Fields(l)
		bid, _ := strconv.ParseInt(f[1], 10, 64)
		game = append(game, deck{cards(f[0]), bid, cards(f[0]).GetDeckType()})
	}
	sort.SliceStable(game,
		func(i int, j int) bool { return lesserDeck(game[i],game[j]) })
	fmt.Println("by bid %", game)
	sum := int64(0)
	for i,j := range game {
		sum += int64(int64(i+1)*j.bid)
	}	
	fmt.Println("silver: ", sum)

}
