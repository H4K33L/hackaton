package Monster

import (
    "math/rand"
    "strconv"
)

func Useles()  {
	/*
	A POOP !
	*/
	
}

func Size(creature monster) monster {
	/*
	The size func take a monster struct and set a size to it,
	and it return a full monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Alignement func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	size := []string{"Très Petite","Petite","Moyenne","Grande","Très Grande","Gigantesque"}
	index := rand.Intn(6)
	creature.Size = size[index]
	return Alignment(creature)		// the Alignement func is bellow
}

func Alignment(creature monster) monster {
	/*
	The Alignment func take a monster struct and set a alignment to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Caract func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	line := []string{"Loyal","Neutre","Chaotique"}
	col  := []string{"bon","neutre","mauvais"}
	index := rand.Intn(3)
	index2 := rand.Intn(3)
	creature.Alignment = line[index]+" "+col[index2]
	return Caract(creature)		// the Caract func is bellow
}

func Caract(creature monster) monster{
	/*
	The Caract func take a monster struct and set a caracteristique to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of CaractMod func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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
	return CaractMod(creature)	// the CaractMod func is bellow
}

func CaractMod(creature monster) monster {
	/*
	The CaractMod func take a monster struct and set a caracteristique
	Modificator to it, and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Mastery func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	output := map[string]int{"For":0,"Dex":0,"Con":0,"Int":0,"Sag":0,"Cha":0}
	for key := range creature.Caract {
		output[key] = creature.Caract[key]/2-5
	}
	creature.CaractMod = output
	return Mastery(creature)	// the Mastery func is bellow
}

func Mastery(creature monster) monster {
	/*
	The Mastery func take a monster struct and set the Mastery to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of AC func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	creature.Mastery = int((creature.ID+7)/4)
	return AC(creature)	// the AC func is bellow
}

func AC(creature monster) monster {
	/*
	The AC func take a monster struct and set the AC to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of LP func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	temp := rand.Intn(2)
	if temp == 0 {
		temp = rand.Intn(2)
		if temp == 0 {
			creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+2)+" (bouclier)"
			return LP(creature)	// the LP func is bellow
		} else {
			creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"])
			return LP(creature) // the LP func is bellow
		}
	} else if temp == 1 {
		naturalArmor := rand.Intn(9)+1
		temp := rand.Intn(2)
		if temp == 0 {
			temp = rand.Intn(2)
			if temp == 0 {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle, bouclier) ," +strconv.Itoa(10+creature.CaractMod["Dex"])+" si à terre"
				return LP(creature) // the LP func is bellow
			} else {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle), "+strconv.Itoa(10+creature.CaractMod["Dex"])+" si à terre"
				return LP(creature) // the LP func is bellow
			}
		} else {
			temp = rand.Intn(2)
			if temp == 0 {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor+2)+" (armure naturelle, bouclier)"
				return LP(creature) // the LP func is bellow
			} else {
				creature.AC = strconv.Itoa(10+creature.CaractMod["Dex"]+naturalArmor)+" (armure naturelle)"
				return LP(creature) // the LP func is bellow
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
			return LP(creature) // the LP func is bellow
		} else {
			creature.AC = strconv.Itoa(output)+" ("+name+")"
			return LP(creature) // the LP func is bellow
		}
	}
}

func LP(creature monster) monster {
	/*
	The LP func take a monster struct and set the LP to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of ResiVulImun func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	nb := rand.Intn(40)
	if nb == 0 {
		nb = 1
	}
	switch creature.Size {
	case "Très Petite" :
		LP := int(float64(nb) * 2.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d4+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	case "Petite" :
		LP := int(float64(nb) * 3.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d6+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	case "Moyenne" :
		LP := int(float64(nb) * 4.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d8+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	case "Grande" :
		LP := int(float64(nb) * 5.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d10+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	case "Très Grande" :
		LP := int(float64(nb) * 6.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d12+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	case "Gigantesque" :
		LP := int(float64(nb) * 10.5 + float64(creature.CaractMod["Con"]*nb))
		creature.LP = strconv.Itoa(LP)+" "+strconv.Itoa(nb)+"d20+"+strconv.Itoa(nb*creature.CaractMod["Con"])
		return ResiVulImun(creature) // the ResiVulImun func is bellow
	}
	return ResiVulImun(creature)	// the ResiVulImun func is bellow
}

func ResiVulImun(creature monster) monster {
	/*
	The ResiVulImun func take a monster struct and set the ResiVulImun to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of AttacBonnus func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	degats := []string{"acide","contondants","froid","feu","force","foudre","nécrotiques","perforants","poison","psychiques","radiants","tranchants","tonnerre"}
	creature.Resistance, degats = Dispatch(degats)		// the func Dispatch i bellow in the second part of this file
	creature.Vulnerability, degats = Dispatch(degats)
	creature.Immunity, _ = Dispatch(degats)
	return AttacBonnus(creature)	// the AttacBonnus func is bellow
}

func AttacBonnus(creature monster) monster {
	/*
	The AttacBonnus func take a monster struct and set the AttacBonnus to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of DD func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	if creature.Mastery+creature.CaractMod["Dex"] >= creature.Mastery+creature.CaractMod["For"] {
		creature.AttacBonnus = creature.Mastery+creature.CaractMod["Dex"]
		return DD(creature)		// the DD func is bellow
	} else {
		creature.AttacBonnus = creature.Mastery+creature.CaractMod["For"]
		return DD(creature)		// the DD func is bellow
	}
}

func DD(creature monster) monster {
	/*
	The DD func take a monster struct and set the DD to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Speed func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

	creature.DD = 8+creature.Mastery+creature.CaractMod["Con"]
	return Speed(creature)		// the Speed func is bellow
}

func Speed(creature monster) monster {
	/*
	The Speed func take a monster struct and set the Speed to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of SaveRoll func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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
	return SaveRoll(creature)		// the SaveRoll func is bellow
}

func SaveRoll(creature monster) monster {
	/*
	The SaveRoll func take a monster struct and set the SaveRoll to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of StateImmunity func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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
	return StateImmunity(creature)		// the StateImmunity func is bellow
}

func StateImmunity(creature monster) monster {
	/*
	The StateImmunity func take a monster struct and set the StateImmunity to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Sense func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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
	return Sense(creature)		// the Sense func is bellow
}

func Sense(creature monster) monster {
	/*
	The Sense func take a monster struct and set the Sense to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : the output of Languages func
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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
	return Languages(creature)		// the Languages func is bellow
}

func Languages(creature monster) monster {
	/*
	The Languages func take a monster struct and set the Languages to it,
	and it return a monster struct.
	---------------------------------------------------------
	input : a monster struct
	output : a monster struct
	---------------------------------------------------------
	The func dosen't have error case.
	*/

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

/*-------------------- The second part of the file -----------------------*/

func ListID() []int {
	/*
	The listID func make a slice of random int representing the
	ID for monster in a group.
	------------------------------------------------------------
	input : -
	output : a slice of int
	------------------------------------------------------------
	The func dosen't have error case.
	*/
	
	list := []int{}
	temp := rand.Intn(40)
	for i := 0 ; i <= temp ; i++ {
		list = append(list,rand.Intn(31))
	}
	return list
}

func Dispatch(degats []string) ([]string, []string) {
	/*
	The Dispatch func take a slice and return two slice, one of 
	the slice have a random number on ellement take to the 
	original slice and the other slice is the original slice 
	who loose elements.
	------------------------------------------------------------
	input : a slice of string
	output : two slice of int
	------------------------------------------------------------
	The func dosen't have error case.
	*/

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
			degats = Pop(degats, degats[temp])	// the Pop function is bellow
		}
		return output, degats
	} else {
		return []string{}, degats
	}
}

func Pop(degats []string, elem string) []string {
	/*
	The Pop func take a slice and an a string and return it 
	whithout the string enter in argument.
	------------------------------------------------------------
	input : a slice of string and a string
	output : slice of string
	------------------------------------------------------------
	The func dosen't have error case.
	*/
	
	newDegats := []string{}
			for i := range degats {
				if degats[i] != elem {
					newDegats = append(newDegats, degats[i])
				}
			}
	return newDegats
}

func IsIn(elem string, list []string) bool {
	/*
	The IsIn function take a string and a slice of string 
	and return a bool, false if the string isn't in the 
	slice and true in the other case.
	------------------------------------------------------
	input : a string and a slice of string
	output : a boolean
	------------------------------------------------------
	The func dosen't have error case.
	*/

	for i := range list {
		if elem == list[i] {
			return true
		}
	}
	return false
}