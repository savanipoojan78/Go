package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubes"}
	cardValues := deck{"Ace", "one", "two", "three"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveFile(filename string) error {
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}
func readFile(filename string) (string,error){
	byte,error := ioutil.ReadFile(filename)
	return string(byte),error
}
