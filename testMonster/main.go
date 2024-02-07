package main

import (
	"fmt"

	"Monster"
)

func main() {
	//monstre1 := Monster.GenerateMonster("bête",3)
	//fmt.Println(monstre1)
	//bandOfMonster := Monster.GenerateGroup("bête",3200)
	/*for i := range bandOfMonster.Monsters {
		fmt.Println(bandOfMonster.Monsters[i].ID)
	}*/
	encounter := Monster.GenerateEncounter("bête",3) 
	for i := range encounter.Encounter {
		fmt.Println(encounter.Encounter[i])
		fmt.Println()
	}
}