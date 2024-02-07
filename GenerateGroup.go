package Monster

import (
    "math/rand"
)

func ListID(ID int) []int {
	list := []int{}
	temp := rand.Intn(40)
	for i := 0 ; i <= temp ; i++ {
		list = append(list,rand.Intn(31))
	}
	return list
}

func GoodListID(list []int, ID int) bool {
	if len(list) == 1 {
		return list[0] == ID
	} else if len(list) == 2 {
		temp := 0
		for i := range list {
			temp += list[i]
		}
		return int(float64(temp)*1.5) == ID
	} else if len(list) >= 3 && len(list) <= 6 {
		temp := 0
		for i := range list {
			temp += list[i]
		}
		return temp*2 == ID
	} else if len(list) >= 7 && len(list) <= 10 {
		temp := 0
		for i := range list {
			temp += list[i]
		}
		return int(float64(temp)*2.5) == ID
	} else if len(list) >= 11 && len(list) <= 14 {
		temp := 0
		for i := range list {
			temp += list[i]
		}
		return temp*3 == ID
	}  else {
		temp := 0
		for i := range list {
			temp += list[i]
		}
		return temp*4 == ID
	}
}

func GenerateGroup(monsterType string, ID int) groupMonster {
	list := ListID(ID)
	for !GoodListID(list, ID) {
		list = ListID(ID)
	}
	output := groupMonster{}
	for i := range list {
		output.Monsters = append(output.Monsters, GenerateMonster(monsterType, list[i]))
	}
	return output
}