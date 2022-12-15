package main

import (
	"fmt"
	"sort"
)

type Dice interface {
	Total() int
	Roll(throws int) int
}

type DerministicDice struct {
	rolls int
	last  int
}

func (d *DerministicDice) Total() int {
	return d.rolls
}

func (d *DerministicDice) Roll(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		d.rolls++
		d.last++
		if d.last%100 == 1 {
			d.last = 1
		}
		//fmt.Println("\t\tRoll", d.rolls, ":", d.last)
		sum += d.last
	}
	return sum
}

type Player struct {
	Pos   int
	Score int
}

func main() {
	p1, p2 := Player{Pos: 4}, Player{Pos: 1}
	dice := &DerministicDice{rolls: 0, last: 0}

	r1 := solve1(10, dice, p1, p2)
	fmt.Println("Part 1:", r1)
}

func solve1(n int, dice Dice, players ...Player) int {
gameloop:
	for {
		for i := range players {
			//fmt.Println("Player", i+1, "turn", dice.Total()/(3*len(players)))
			throw := dice.Roll(3)
			//fmt.Println("\tPlayer", i+1, "rolled", throw)
			pos := (players[i].Pos + throw) % n
			if pos == 0 {
				pos = 10
			}

			players[i].Pos = pos
			players[i].Score += pos
			//fmt.Println("\tPlayer", i+1, "is now at", pos, "with score", players[i].Score)
			if players[i].Score >= 1_000 {
				break gameloop
			}
		}
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Score < players[j].Score
	})

	return players[0].Score * dice.Total()
}
