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
	if ID == 0 {
		ID = 1
	}
	switch size {
	case "Très Petite" :
		LP := int(float64(ID) * 2.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d4+"+strconv.Itoa(ID*carcats["Con"])
	case "Petite" :
		LP := int(float64(ID) * 3.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d6+"+strconv.Itoa(ID*carcats["Con"])
	case "Moyenne" :
		LP := int(float64(ID) * 4.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d8+"+strconv.Itoa(ID*carcats["Con"])
	case "Grande" :
		LP := int(float64(ID) * 5.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d10+"+strconv.Itoa(ID*carcats["Con"])
	case "Très Grande" :
		LP := int(float64(ID) * 6.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d12+"+strconv.Itoa(ID*carcats["Con"])
	case "Gigantesque" :
		LP := int(float64(ID) * 10.5 + float64(carcats["Con"]*ID))
		return strconv.Itoa(LP)+" "+strconv.Itoa(ID)+"d20+"+strconv.Itoa(ID*carcats["Con"])
	}
	return ""
}

func ResiVulImun() []string {
	degats := []string{"acide","contondants","froid"," feu","force","foudre","nécrotiques","perforants","poison","psychiques","radiants","tranchants"," tonnerre"}
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
