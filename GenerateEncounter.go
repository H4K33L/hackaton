package Monster

import (
    "math/rand"
)

func GenerateEncounter(monsterType string, ID int) encounterTable {
	output := encounterTable{}
	temp := rand.Intn(3)
	switch temp {
	case 0:
		for i := 0 ; i <=6 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	case 1:
		for i := 0 ; i <=8 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	case 2:
		for i := 0 ; i <=10 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	}
	return output
}