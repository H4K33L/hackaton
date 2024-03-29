package Monster

import "strconv"

func GoodID(monstre monster) bool {
	/*
	The GoodID take a monster struct and return a boolean if the ID 
	of monster is correct whith the monster Resistance, vulnerability 
	and imunity an amount of LP. The function return false if the 
	monster dosen't have the good ID an true if the amount of LP 
	concording whith the ID.
	All ID have interval of LP, see Donjon&Dragon Master Rulle.
	------------------------------------------------------------------
	input : a monster struct
	output : a boolean
	------------------------------------------------------------------
	the function dosen't have arror case.
	*/

	LP := float64(GetLP(monstre.LP))
	if monstre.ID <= 4 {
		if len(monstre.Resistance) >= 3 {
			LP *= 2.0
		}
		if len(monstre.Immunity) >= 3 {
			LP *= 2.0
		}
		if len(monstre.Vulnerability) >= 3 {
			LP /= 2.0
		} 
	} else if monstre.ID <= 10 {
		if len(monstre.Resistance) >= 3 {
			LP *= 1.5
		}
		if len(monstre.Immunity) >= 3 {
			LP *= 2.0
		}
		if len(monstre.Vulnerability) >= 3 {
			LP /= 2.0
		} 
	} else if monstre.ID <= 16 {
		if len(monstre.Resistance) >= 3 {
			LP *= 1.25
		}
		if len(monstre.Immunity) >= 3 {
			LP *= 1.5
		}
		if len(monstre.Vulnerability) >= 3 {
			LP /= 2.0
		} 
	} else if monstre.ID >= 17 {
		if len(monstre.Resistance) >= 3 {
			LP *= 1.0
		}
		if len(monstre.Immunity) >= 3 {
			LP *= 1.25
		}
		if len(monstre.Vulnerability) >= 3 {
			LP /= 2.0
		} 
	}
	switch monstre.ID {
	case 0 :
		if LP >= 1 && LP <= 70 {
			return true
		} else {
			return false
		}
	case 1 :
		if LP >= 71 && LP <= 85 {
			return true
		} else {
			return false
		}
	case 2 :
		if LP >= 86 && LP <= 100 {
			return true
		} else {
			return false
		}
	case 3 :
		if LP >= 101 && LP <= 115 {
			return true
		} else {
			return false
		}
	case 4 :
		if LP >= 116 && LP <= 130 {
			return true
		} else {
			return false
		}
	case 5 :
		if LP >= 131 && LP <= 145 {
			return true
		} else {
			return false
		}
	case 6 :
		if LP >= 146 && LP <= 160 {
			return true
		} else {
			return false
		}
	case 7 :
		if LP >= 161 && LP <= 175 {
			return true
		} else {
			return false
		}
	case 8 :
		if LP >= 176 && LP <= 190 {
			return true
		} else {
			return false
		}
	case 9 :
		if LP >= 191 && LP <= 205 {
			return true
		} else {
			return false
		}
	case 10 :
		if LP >= 206 && LP <= 220 {
			return true
		} else {
			return false
		}
	case 11 :
		if LP >= 221 && LP <= 235 {
			return true
		} else {
			return false
		}
	case 12 :
		if LP >= 236 && LP <= 250 {
			return true
		} else {
			return false
		}
	case 13 :
		if LP >= 251 && LP <= 265 {
			return true
		} else {
			return false
		}
	case 14 :
		if LP >= 266 && LP <= 280 {
			return true
		} else {
			return false
		}
	case 15 :
		if LP >= 281 && LP <= 295 {
			return true
		} else {
			return false
		}
	case 16 :
		if LP >= 296 && LP <= 310 {
			return true
		} else {
			return false
		}
	case 17 :
		if LP >= 311 && LP <= 325 {
			return true
		} else {
			return false
		}
	case 18 :
		if LP >= 326 && LP <= 340 {
			return true
		} else {
			return false
		}
	case 19 :
		if LP >= 341 && LP <= 355 {
			return true
		} else {
			return false
		}
	case 20 :
		if LP >= 356 && LP <= 400 {
			return true
		} else {
			return false
		}
	case 21 :
		if LP >= 401 && LP <= 445 {
			return true
		} else {
			return false
		}
	case 22 :
		if LP >= 446 && LP <= 490 {
			return true
		} else {
			return false
		}
	case 23 :
		if LP >= 491 && LP <= 535 {
			return true
		} else {
			return false
		}
	case 24 :
		if LP >= 536 && LP <= 580 {
			return true
		} else {
			return false
		}
	case 25 :
		if LP >= 581 && LP <= 625 {
			return true
		} else {
			return false
		}
	case 26 :
		if LP >= 626 && LP <= 670 {
			return true
		} else {
			return false
		}
	case 27 :
		if LP >= 671 && LP <= 715 {
			return true
		} else {
			return false
		}
	case 28 :
		if LP >= 716 && LP <= 760 {
			return true
		} else {
			return false
		}
	case 29 :
		if LP >= 761 && LP <= 805 {
			return true
		} else {
			return false
		}
	case 30 :
		if LP >= 806 && LP <= 850 {
			return true
		} else {
			return false
		}
	}
	return false
}

func GetLP(info string) int {
	/*
	The GetLP function take a string and return the first number 
	in a string, "12 png" return 12 "j54 lm" return a 0.
	-------------------------------------------------------------
	input : a string
	output : an int
	-------------------------------------------------------------
	there is no error case in this function
	*/

	output := ""
	for i := range info {
		if info[i] == ' ' {
			break
		}
		output = output + string(info[i])
	}
	LP, err := strconv.Atoi(output)
	if err != nil {
		return 0
	}
	return LP
}

func GoodListID(list []int, ID int) bool {
	/*
	The GoodListID take an array of int (representing a group of monster) 
	and an int representing the ID of the group and return boolean. 
	False if the ID of the group is hover the ID in parameters, True
	if the ID of the group is the same as ID in parameter.
	To calculate the ID of a group we use the DnD 5 rules, see 
	Donjon&Dragon Master Rulle.
	------------------------------------------------------------------
	input : an array of int an an int
	output : a boolean
	------------------------------------------------------------------
	the function dosen't have arror case.
	*/

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