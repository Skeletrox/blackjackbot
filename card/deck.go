package card

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) DealCard() Card {
	rand.Seed(time.Now().UnixNano())
	position := rand.Int()%len(d.Cards) + 1
	returnable := d.Cards[position]
	if position == len(d.Cards)-1 {
		d.Cards = d.Cards[:position]
	} else {
		d.Cards = append(d.Cards[:position], d.Cards[:position+1]...)
	}
	return returnable
}

func (d *Deck) Init() {
	d.Cards = make([]Card, 52)
	counter := 0
	suits := [4]rune{'H', 'S', 'D', 'C'}
	for _, s := range suits {
		for i := 1; i < 14; i++ {
			d.Cards[counter] = Card{Suit: s, Value: int8(i)}
			counter++
		}
	}
}
