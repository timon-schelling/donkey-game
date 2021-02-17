package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	donkey bool
	number int
}

type Player struct {
	upStack   []*Card
	downStack []*Card
}

type Game struct {
	endStackNumber int
	playerNumber   int
	highestCard    int
	startStack     []*Card
	endStacks      [][]*Card
	player         []*Player
}

func main() {

	n := 6
	p := 6
	h := 20

	startStack := make([]*Card, (h*n)+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= 20; j++ {
			startStack[((h*i)+j)-1] = &Card{number: j}
		}
	}
	startStack[len(startStack)-1] = &Card{donkey: true}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(startStack), func(i, j int) { startStack[i], startStack[j] = startStack[j], startStack[i] })

	player := make([]*Player, p)
	for i := 0; i < p; i++ {
		player[i] = &Player{}
	}

	game := &Game{
		endStackNumber: n,
		playerNumber:   p,
		highestCard:    h,
		startStack:     startStack,
		endStacks:      make([][]*Card, n),
		player:         player,
	}
	fmt.Println(game)

	for {
		p := game.player[0]
		for {
			if len(p.upStack) > 0 {
				if tryPlace(p.upStack[len(p.upStack)-1], game) {
					p.upStack = p.upStack[:len(p.upStack)-1]
					continue
				} else {
					break

				}
			} else if len(game.startStack) != 0 {
				p.upStack = append(p.upStack, game.startStack[len(game.startStack)-1])
				game.startStack = game.startStack[:len(game.startStack)-1]
				continue
			} else {
				p.upStack = append(p.upStack, game.startStack[len(game.startStack)-1])
				game.startStack = game.startStack[:len(game.startStack)-1]
				continue
			}
		}
		game.player = append(game.player[1:], game.player[0])
	}
}

func tryPlace(card *Card, game *Game) bool {
	for i, s := range game.endStacks {
		if (len(s) == 0 || s[len(s)-1].number+1 == card.number) && !card.donkey {
			game.endStacks[i] = append(game.endStacks[i], card)
			return true
		}
	}
	for _, p := range game.player {
		ptc := p.upStack[len(p.upStack)-1]
		place := false
		if (card.donkey && ptc.number == game.highestCard) || (ptc.donkey && card.number == game.highestCard) {
			place = true
		} else {
			if ptc.number+1 == card.number || ptc.number-1 == card.number {
				place = true
			}
		}
		if place {
			p.upStack = append(p.upStack, card)
			return true
		}
	}
	return false
}
