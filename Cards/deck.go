package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}

	cardValues := []string{"A", "1", "2", "3"}
	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}
	return cards

}

func (d deck) printValue() {
	for i, eachValue := range d {
		fmt.Println(i, eachValue)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	if len(d)-1 > handSize {
		return d[:handSize], d[handSize:]
	} else {
		return d[:], []string{}
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFileToDrive(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromDrive(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	stSlice := strings.Split(string(bs), ",")
	return deck(stSlice)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomIndex := r.Intn(len(d))
		// temp := d[i]
		// d[i] = d[randomIndex]
		// d[randomIndex] = temp
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
	return d
}

// func toByte(d deck) []byte {
// 	return []byte(d)
// }
