package main

import "fmt"

// Create a function that returns the name of the winner in a fight between two fighters.
// Each fighter takes turns attacking the other and whoever kills the other first is victorious.
// Death is defined as having health <= 0.
// Each fighter will be a Fighter object/instance.
// See the Fighter class below in your chosen language.
// Both health and damagePerAttack (damage_per_attack for python) will be integers larger than 0.
// You can mutate the Fighter objects.
// Your function also receives a third argument, a string, with the name of the fighter that attacks first.

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald"))
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	fighters := make([]Fighter, 2)
	fighters[0] = fighter1
	fighters[1] = fighter2

	fristAtk := func(fighters []Fighter, firstAttacker string) {
		for i, f := range fighters {
			if f.Name == firstAttacker {
				defenderIndex := 1 - i
				fighters[defenderIndex].Health -= fighters[i].DamagePerAttack
				break
			}
		}
	}

	for fighters[0].Health > 0 && fighters[1].Health > 0 {
		fristAtk(fighters, firstAttacker)
		if firstAttacker == fighters[0].Name {
			firstAttacker = fighters[1].Name
		} else {
			firstAttacker = fighters[0].Name
		}
	}

	if fighters[0].Health <= 0 {
		return fighters[1].Name
	} else {
		return fighters[0].Name
	}
}
