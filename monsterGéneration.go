package Monster

import (
    "math/rand"
    "strconv"
)

func MonsterGenartae(monsterType string, ID int) monster {
	output := monster{}
	output.ID = ID
	output.MonsterType = monsterType
	output.Size = Size()
	output.Alignment = Alignment()
	output.Caract = Caract()
	output.CaractMod = CaractMod(output.Caract)
	output.Mastery = Mastery(output.ID)
	output.AC = AC(output.CaractMod)
	output.LP = LP(output.ID, output.Size, output.CaractMod)
	output.Resistance = ResiVulImun()
	output.Vulnerability = ResiVulImun()
	output.Immunity = ResiVulImun()
	output.AttacBonnus = AttacBonnus(output.Mastery, output.CaractMod)
	output.DD = DD(output.Mastery, output.CaractMod)
	output.Speed = Speed()
	output.SaveRoll = SaveRoll(output.CaractMod)
	output.StateImmunity = StateImmunity()
	output.Sense = Sense(output.CaractMod)
	output.Languages = Languages(output.MonsterType)
	output = monster{}
	for !GoodID(output) {
		output.ID = ID
		output.MonsterType = monsterType
		output.Size = Size()
		output.Alignment = Alignment()
		output.Caract = Caract()
		output.CaractMod = CaractMod(output.Caract)
		output.Mastery = Mastery(output.ID)
		output.AC = AC(output.CaractMod)
		output.LP = LP(output.ID, output.Size, output.CaractMod)
		output.Resistance = ResiVulImun()
		output.Vulnerability = ResiVulImun()
		output.Immunity = ResiVulImun()
		output.AttacBonnus = AttacBonnus(output.Mastery, output.CaractMod)
		output.DD = DD(output.Mastery, output.CaractMod)
		output.Speed = Speed()
		output.SaveRoll = SaveRoll(output.CaractMod)
		output.StateImmunity = StateImmunity()
		output.Sense = Sense(output.CaractMod)
		output.Languages = Languages(output.MonsterType)
	}

	return output
}

func Size() string {
	size := []string{"Très Petite","Petite","Moyenne","Grande","Très Grande","Gigantesque"}
	index := rand.Intn(6)
	return size[index]
}

func Alignment() string {
	line := []string{"Loyal","Neutre","Chaotique"}
	col  := []string{"bon","neutre","mauvais"}
	index := rand.Intn(3)
	index2 := rand.Intn(3)
	return line[index]+" "+col[index2]
}

func Caract() map[string]int{
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	temp := 0
	for key := range output {
		for i := 0; i <= 5 ; i++ {
			temp += rand.Intn(5)+1
		}
		output[key] = temp
		temp = 0
	}
	return output
}

func CaractMod(carcats map[string]int) map[string]int {
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	for key := range carcats {
		output[key] = carcats[key]/2-5
	}
	return output
}

func Mastery(ID int) int {
	return int((ID+7)/4)
}

func AC( carcats map[string]int) string {
	temp := rand.Intn(2)
	if temp == 0 {
		temp = rand.Intn(2)
		if temp == 0 {
			return strconv.Itoa(10+carcats["Dex"]+2)+" (bouclier)"
		} else {
			return strconv.Itoa(10+carcats["Dex"])
		}
	} else if temp == 1 {
		naturalArmor := rand.Intn(9)+1
		temp := rand.Intn(2)
		if temp == 0 {
			temp = rand.Intn(2)
			if temp == 0 {
				return strconv.Itoa(10+carcats["Dex"]+naturalArmor)+" (armure naturelle, bouclier) ," +strconv.Itoa(10+carcats["Dex"])+" si à terre"
			} else {
				return strconv.Itoa(10+carcats["Dex"]+naturalArmor)+" (armure naturelle), "+strconv.Itoa(10+carcats["Dex"])+" si à terre"
			}
		} else {
			temp = rand.Intn(2)
			if temp == 0 {
				return strconv.Itoa(10+carcats["Dex"]+naturalArmor+2)+" (armure naturelle, bouclier)"
			} else {
				return strconv.Itoa(10+carcats["Dex"]+naturalArmor)+" (armure naturelle)"
			}
		}
	} else {
		temp := rand.Intn(12)
		output := 0
		name := ""
		switch temp {
		case 0 :
			output = 11+carcats["Dex"]
			name = "Matelassée"
		case 1 :
			output = 11+carcats["Dex"]
			name = "Cuir"
		case 2 :
			output = 12+carcats["Dex"]
			name = "Cuir clouté"
		case 3 :
			if carcats["Dex"] >= 2 {
				output = 14
			} else {
				output = 12+carcats["Dex"]
			}
			name = "Peaux"
		case 4 :
			if carcats["Dex"] >= 2 {
				output = 15
			} else {
				output = 13+carcats["Dex"]
			}
			name = "Chemise de mailles"
		case 5 :
			if carcats["Dex"] >= 2 {
				output = 16
			} else {
				output = 14+carcats["Dex"]
			}
			name = "écailles"
		case 6 :
			if carcats["Dex"] >= 2 {
				output = 16
			} else {
				output = 14+carcats["Dex"]
			}
			name = "Cuirasse"
		case 7 :
			if carcats["Dex"] >= 2 {
				output = 17
			} else {
				output = 15+carcats["Dex"]
			} 
			name = "Demi-plate"
		case 8 :
			output = 14
			name = "Broigne"
		case 9 :
			output = 16
			name = "Cotte de mailles"
		case 10 :
			output = 17
			name = "Clibanion"
		case 11 :
			output = 18
			name = "Harnois"
		}
		temp = rand.Intn(2)
		if temp == 0 {
			return strconv.Itoa(output+2)+" ("+name+", bouclier)"
		} else {
			return strconv.Itoa(output)+" ("+name+")"
		}
	}
}

func LP(ID int, size string, carcats map[string]int) string {
	nb := rand.Intn(40)
	if nb == 0 {
		nb = 1
	}
	switch size {
	case "Très Petite" :
		LP := int(float64(nb) * 2.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d4+"+strconv.Itoa(nb*carcats["Con"])
	case "Petite" :
		LP := int(float64(nb) * 3.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d6+"+strconv.Itoa(nb*carcats["Con"])
	case "Moyenne" :
		LP := int(float64(nb) * 4.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d8+"+strconv.Itoa(nb*carcats["Con"])
	case "Grande" :
		LP := int(float64(nb) * 5.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d10+"+strconv.Itoa(nb*carcats["Con"])
	case "Très Grande" :
		LP := int(float64(nb) * 6.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d12+"+strconv.Itoa(nb*carcats["Con"])
	case "Gigantesque" :
		LP := int(float64(nb) * 10.5 + float64(carcats["Con"]*nb))
		return strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d20+"+strconv.Itoa(nb*carcats["Con"])
	}
	return ""
}

func ResiVulImun() []string {
	degats := []string{"acide","contondants","froid","feu","force","foudre","nécrotiques","perforants","poison","psychiques","radiants","tranchants","tonnerre"}
	temp := rand.Intn(2)
	if temp != 1 {
		output := []string{}
		nb := rand.Intn(7)
		for i := 0; i <= nb ; i++ {
			temp := rand.Intn(12)
			if !IsIn(degats[temp], output) {
				output = append(output, degats[temp])
			}
		}
		return output
	} else {
		return []string{}
	}
}

func IsIn(elem string, list []string) bool {
	for i := range list {
		if elem == list[i] {
			return true
		}
	}
	return false
}

func AttacBonnus(mastery int, carcats map[string]int) int {
	if mastery+carcats["Dex"] >= mastery+carcats["For"] {
		return mastery+carcats["Dex"]
	} else {
		return mastery+carcats["For"]
	}
}

func DD(mastery int, carcats map[string]int) int {
	return 8+mastery+carcats["Con"]
}

func Speed() map[string]int {
	output := map[string]int{"marche":0,"vol":0,"nage":0,"creusement":0,"escalade":0}
	speeds := []string{"marche","vol","nage","creusement","escalade"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(6)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(5)
			output[speeds[temp]] = int(1.5*float64(rand.Intn(13)))
		}
	}
	return output
}

func SaveRoll(carcats map[string]int) map[string]int {
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	carac := []string{"For","Dex","Con","Int","Sag","Cha"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(6)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(6)
			output[carac[temp]] = carcats[carac[temp]]
		}
	}
	return output
}

func StateImmunity() []string {
	output := []string{}
	list := []string{"À terre","Agrippé","Assourdi","Aveuglé","Charmé","Effrayé","Empoisonné","Entravé","Étourdi","Neutralisé","Inconscient","Paralysé","Pétrifié"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(13)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(6)
			if !IsIn(list[temp], output) {
				output = append(output, list[temp])
			}
		}
	}
	return output
}

func Sense(carcats map[string]int) []string {
	output := []string{}
	list := []string{"vision dans le noir ","vision aveugle ","vision véritable "}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(3)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(3)
			if !IsIn(list[temp], output) {
				output = append(output, list[temp])
			}
		}
	}
	for i := range output {
		output[i] = output[i]+strconv.Itoa(int(1.5*float64(rand.Intn(25))))+" m"
	}
	perception := "  perception passive "+strconv.Itoa(10+carcats["Sag"])
	output = append(output, perception)
	return output
}

func Languages(monsterType string) []string {
	output := []string{}
	langues := []string{"Commun","Elfique","Géant","Gnome","Gobelin","Halfelin","Nain","Orc","Abyssal","Céleste","Commun des profondeurs","Draconique","Infernal","Primordial","Profond","Sylvestre"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(16)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(16)
			if !IsIn(langues[temp], output) {
				output = append(output, langues[temp])
			}
		}
	} else {
		output = append(output,"-")
	}
	return output
}

func GoodID(monstre monster) bool {
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