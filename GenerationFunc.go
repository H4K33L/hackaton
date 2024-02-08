package Monster

import (
    "math/rand"
    "strconv"
)

func MonsterGenartae(monsterType string, ID int) monster {
	output := monster{}
	output.ID = ID
	output.MonsterType = monsterType
	output = Size(output)
	output = Alignment(output)
	output = Caract(output)
	output = CaractMod(output)
	output = Mastery(output)
	output = AC(output)
	output = LP(output)
	output = ResiVulImun(output)
	output = AttacBonnus(output)
	output = DD(output)
	output = Speed(output)
	output = SaveRoll(output)
	output = StateImmunity(output)
	output = Sense(output)
	output = Languages(output)
	output = monster{}
	for !GoodID(output) {
		output.ID = ID
		output.MonsterType = monsterType
		output = Size(output)
	}

	return output
}

func Size(creature monster) monster {
	size := []string{"Très Petite","Petite","Moyenne","Grande","Très Grande","Gigantesque"}
	index := rand.Intn(6)
	creature.Size = size[index]
	return Alignment(creature)
}

func Alignment(creature monster) monster {
	line := []string{"Loyal","Neutre","Chaotique"}
	col  := []string{"bon","neutre","mauvais"}
	index := rand.Intn(3)
	index2 := rand.Intn(3)
	creature.Alignment = line[index]+" "+col[index2]
	return Caract(creature)
}

func Caract(creature monster) monster{
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	temp := 0
	for key := range output {
		for i := 0; i <= 5 ; i++ {
			temp += rand.Intn(5)+1
		}
		output[key] = temp
		temp = 0
	}
	creature.Caract = output
	return CaractMod(creature)
}

func CaractMod(creature monster) monster {
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	for key := range creature.Caract {
		output[key] = creature.Caract[key]/2-5
	}
	creature.CaractMod = output
	return Mastery(creature)
}

func Mastery(creature monster) monster {
	creature.Mastery = int((creature.ID+7)/4)
	return AC(creature)
}

func AC(creature monster) monster {
	temp := rand.Intn(2)
	if temp == 0 {
		temp = rand.Intn(2)
		if temp == 0 {
			creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+2)+" (bouclier)"
			return LP(creature)
		} else {
			creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"])
			return LP(creature)
		}
	} else if temp == 1 {
		naturalArmor := rand.Intn(9)+1
		temp := rand.Intn(2)
		if temp == 0 {
			temp = rand.Intn(2)
			if temp == 0 {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle, bouclier) ," +strconv.Itoa(10+creature.CaractMod["Dex"])+" si à terre"
				return LP(creature)
			} else {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle), "+strconv.Itoa(10+creature.CaractMod["Dex"])+" si à terre"
				return LP(creature)
			}
		} else {
			temp = rand.Intn(2)
			if temp == 0 {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor+2)+" (armure naturelle, bouclier)"
				return LP(creature)
			} else {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle)"
				return LP(creature)
			}
		}
	} else {
		temp := rand.Intn(12)
		output := 0
		name := ""
		switch temp {
		case 0 :
			output = 11+creature.CaractMod["Dex"]
			name = "Matelassée"
		case 1 :
			output = 11+creature.CaractMod["Dex"]
			name = "Cuir"
		case 2 :
			output = 12+creature.CaractMod["Dex"]
			name = "Cuir clouté"
		case 3 :
			if creature.CaractMod["Dex"] >= 2 {
				output = 14
			} else {
				output = 12+creature.CaractMod["Dex"]
			}
			name = "Peaux"
		case 4 :
			if creature.CaractMod["Dex"] >= 2 {
				output = 15
			} else {
				output = 13+creature.CaractMod["Dex"]
			}
			name = "Chemise de mailles"
		case 5 :
			if creature.CaractMod["Dex"] >= 2 {
				output = 16
			} else {
				output = 14+creature.CaractMod["Dex"]
			}
			name = "écailles"
		case 6 :
			if creature.CaractMod["Dex"] >= 2 {
				output = 16
			} else {
				output = 14+creature.CaractMod["Dex"]
			}
			name = "Cuirasse"
		case 7 :
			if creature.CaractMod["Dex"] >= 2 {
				output = 17
			} else {
				output = 15+creature.CaractMod["Dex"]
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
			creature.AC = strconv.Itoa(output+2)+" ("+name+", bouclier)"
			return LP(creature)
		} else {
			creature.AC = strconv.Itoa(output)+" ("+name+")"
			return LP(creature)
		}
	}
}

func LP(creature monster) monster {
	nb := rand.Intn(40)
	if nb == 0 {
		nb = 1
	}
	switch creature.Size {
	case "Très Petite" :
		LP := int(float64(nb) * 2.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d4+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	case "Petite" :
		LP := int(float64(nb) * 3.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d6+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	case "Moyenne" :
		LP := int(float64(nb) * 4.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d8+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	case "Grande" :
		LP := int(float64(nb) * 5.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d10+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	case "Très Grande" :
		LP := int(float64(nb) * 6.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d12+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	case "Gigantesque" :
		LP := int(float64(nb) * 10.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d20+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature)
	}
	return ResiVulImun(creature)
}

func ResiVulImun(creature monster) monster {
	degats := []string{"acide","contondants","froid","feu","force","foudre","nécrotiques","perforants","poison","psychiques","radiants","tranchants","tonnerre"}
	creature.Resistance, degats = Dispatch(degats)
	creature.Vulnerability, degats = Dispatch(degats)
	creature.Immunity, _ = Dispatch(degats)
	return AttacBonnus(creature)
}

func Dispatch(degats []string) ([]string, []string) {
	temp := rand.Intn(2)
	if temp != 1 {
		output := []string{}
		if len(degats) == 0 {
			return []string{}, []string{}
		}
		nb := rand.Intn(len(degats))
		for i := 0; i <= nb ; i++ {
			temp := rand.Intn(len(degats))
			output = append(output, degats[temp])
			degats = Pop(degats, degats[temp])
		}
		return output, degats
	} else {
		return []string{}, degats
	}
}

func Pop(degats []string, elem string) []string {
	newDegats := []string{}
			for i := range degats {
				if degats[i] != elem {
					newDegats = append(newDegats, degats[i])
				}
			}
	return newDegats
}

func IsIn(elem string, list []string) bool {
	for i := range list {
		if elem == list[i] {
			return true
		}
	}
	return false
}

func AttacBonnus(creature monster) monster {
	if creature.Mastery+creature.CaractMod["Dex"] >= creature.Mastery+creature.CaractMod["For"] {
		creature.AttacBonnus = creature.Mastery+creature.CaractMod["Dex"]
		return DD(creature)
	} else {
		creature.AttacBonnus = creature.Mastery+creature.CaractMod["For"]
		return DD(creature)
	}
}

func DD(creature monster) monster {
	creature.DD = 8+creature.Mastery+creature.CaractMod["Con"]
	return Speed(creature)
}

func Speed(creature monster) monster {
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
	creature.Speed = output
	return SaveRoll(creature)
}

func SaveRoll(creature monster) monster {
	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	carac := []string{"For","Dex","Con","Int","Sag","Cha"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(6)
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(6)
			output[carac[temp]] = creature.CaractMod[carac[temp]]
		}
	}
	creature.SaveRoll = output
	return StateImmunity(creature)
}

func StateImmunity(creature monster) monster {
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
	creature.StateImmunity = output
	return Sense(creature)
}

func Sense(creature monster) monster {
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
	perception := "  perception passive "+strconv.Itoa(10+creature.CaractMod["Sag"])
	creature.Sense = append(output, perception)
	return Languages(creature)
}

func Languages(creature monster) monster {
	output := []string{}
	langues := []string{"Commun","Elfique","Géant","Gnome","Gobelin","Halfelin","Nain","Orc","Abyssal","Céleste","Commun des profondeurs","Draconique","Infernal","Primordial","Profond","Sylvestre"}
	temp := rand.Intn(2)
	if temp == 1 {
		temp := rand.Intn(len(langues))
		for i := 0; i <= temp; i++ {
			temp := rand.Intn(len(langues))
			if !IsIn(langues[temp], output) {
				output = append(output, langues[temp])
			}
		}
	} else {
		output = append(output,"-")
	}
	creature.Languages = output
	return creature
}

/*-------------------------------------------*/

func ListID(ID int) []int {
	list := []int{}
	temp := rand.Intn(40)
	for i := 0 ; i <= temp ; i++ {
		list = append(list,rand.Intn(31))
	}
	return list
}