package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type card byte
type cards string

func (c card) ToVal() int {
	cards := map[card]int{'J':1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 11, 'K': 12, 'A': 13}
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
	j :=0
	v, ok := same['J'] 
	if ok {
		j = v
	} 
	fmt.Println("Joker: ",j)
	// five_of_a_kind
	if len(same) == 1 {
		return five_of_a_kind
	}

	// four_of_a_kind | full_house (2/3 oder 3/2)
	if len(same) == 2 {
		for _, v := range same {
			if v == 1 || v == 4 {
				if j == 1 || j == 4 {
					return five_of_a_kind
				}
				return four_of_a_kind
			}
			if v == 2 || v == 3 {
				if j == 2 || j == 3 {
					return five_of_a_kind
				}
				return full_house
			}
		}
	}

	// two_pair | three_of_a_kind (2/2/1 oder 3/1/1)
	if len(same) == 3 {
		for _, v := range same {
			if v == 2 {
				if j == 1 { return full_house}
				if j == 2 {return four_of_a_kind}
				return two_pair
			}
			
			if v == 3 {
				if j==3 {return four_of_a_kind}
				if j==1 {return four_of_a_kind}
				return three_of_a_kind
			}
		}
		
	}
	// one_pair (2/1/1/1)
	if len(same) == 4 {
		if j == 1 || j==2 {return three_of_a_kind}
		return one_pair
	}
	// high_card (1/1/1/1/1)
	if len(same) == 5 {
		if j == 1 {return one_pair}
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
	re := regexp.MustCompile(`J+`)
	for i,j := range game {
		sum += int64(int64(i+1)*j.bid)
		if re.MatchString(string(j.hand)) {
			fmt.Println(j)
		}

	}	
	fmt.Println("gold: ", sum)

}
