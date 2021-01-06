package main

import (
	// "fmt"
	"math/rand"
	"time"
)

func main() {
	cards := newDeck()
	// deal, remainigCard := deal(cards, 5)
	// deal.print()
	// fmt.Println("--------------")
	// remainigCard.print()
	// fmt.Println(cards.saveFile("qwe.txt"))
	// text, err := readFile("qwe.txt")
	// if err == nil {
	// 	fmt.Println(text)
	// }
	cards.suffle()
	cards.print()
}
func newCard() string {
	return "five of spades"
}
func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
