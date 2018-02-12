package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slic of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), "\u007c")
}

func toDeck(s string) deck {
	return deck(strings.Split(s, "\u007c"))
}

func newDeckFromFile(filename string) (deck, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	mydeck := toDeck(string(content))
	return mydeck, err

}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {

		rndnumber := r.Intn(len(d) - 1)
		d[i], d[rndnumber] = d[rndnumber], d[i]

	}

}
