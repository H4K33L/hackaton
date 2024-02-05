package main

import (
	"fmt"
	"math/rand"
	"encoding/json"
	"io/ioutil"
)

type Name struct {
	Name    string `json:"name"`
	Race    string `json:"race"`	
	Force   int    `json:"force"`
	Armor   int    `json:"armor"`
	Element string `json:"element"`
}

var element = []string{"terre", "feu", "eau", "air", "foudre", "magie"}
var race = []string{"démons", "bête", "mort-vivant", "humanoïde", "Dragon", "Géant"}

func main() {
	data := []Name{}
	for i := 0; i < 10; i++ {
		strenght := rand.Intn(20) + 1
		elementMonster := element[rand.Intn(len(element))]
		raceMonster := race[rand.Intn(len(race))]
		armor := rand.Intn(30) + 1
		name := nameGenerator(strenght, elementMonster, raceMonster, armor)
		data = append(data, Name{name, raceMonster, strenght, armor, elementMonster})
	}
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("monster.json", file, 0644)
	fmt.Println("File created")
}

func setElement(element string) string {
	adj := ""

	fire := []string{"brûlant", "ardent", "enflammé", "incandescent", "brûlé", "rougeoyant", "fumant", "torride"}
	water := []string{"mouillé", "glacial", "gelé", "humide", "ruisselant", "détrempé", "inondé", "débordant"}
	earth := []string{"sauvage", "féroce", "sanglant", "affamé"}
	air := []string{"aérien", "léger", "aérien", "céleste", "nuageux", "venteux", "tempétueux"}
	thunder := []string{"tonnant", "électrique", "foudroyant", "éclatant", "détonnant", "fulgurant", "électrisant"}
	magic := []string{"magique", "ensorcelé", "mystique", "surnaturel", "mystérieux", "enchanté", "sacré"}

	switch element {
	case "terre":
		adj = earth[rand.Intn(len(earth))]
	case "feu":
		adj = fire[rand.Intn(len(fire))]
	case "eau":
		adj = water[rand.Intn(len(water))]
	case "air":
		adj = air[rand.Intn(len(air))]
	case "foudre":
		adj = thunder[rand.Intn(len(thunder))]
	case "magie":
		adj = magic[rand.Intn(len(magic))]
	}

	return adj
}

func setRace(race string, name string, force int) string {
	raceSet := name

	animals := []string{"loup", "ours", "chien", "serpent"}
	undead := []string{"zombie", "squelette", "fantôme", "momie", "vampire", "loup-garou"}

	switch race {
	case "bête":
		raceSet = animals[rand.Intn(len(animals))]
	case "mort-vivant":
		raceSet = undead[rand.Intn(len(undead))]
	}
	return "(" + race + " " + setForce(force) +") " + raceSet 
}

func setForce(force int) string {
	power := ""

	if force <= 5 {
		power = "faible"
	} else if force <= 10 {
		power = "normal"
	} else if force <= 15 {
		power = "puissant"
	} else if force <= 20 {
		power = "surpuissant"
	} else if force <= 25 {
		power = "légendaire"
	} else if force <= 29{
		power = "mythique"
	} else if force >= 30{
		power = "divin"
	}
	return power
}

func setArmor(armor int) string {
	protection := ""

	if armor <= 5 {
		protection = "sans défense"
	} else if armor <= 10 {
		protection = "protégé"
	} else if armor <= 15 {
		protection = "blindé"
	} else if armor >= 20 {
		protection = "impénétrable"
	}

	return protection
}

func nameGenerator(strenght int, element string, race string, armor int) string {
	voyels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
	endWordDemon := []string{"og", "eth", "oth", "uk", "uur", "ach", "uz", "oz", "er", "org"}
	min := 4
	max := 8
	randomNumber := rand.Intn(max - min) + min
	name := ""
	letter := ""

	for i := 0; i < randomNumber; i++ {
		if i%2 == 0 {
			letter = voyels[rand.Intn(len(voyels))]
			name += letter
			if letter == "o" || letter =="u" || letter == "a" {
				if rand.Intn(4) == 0 {
					name += letter
				}
			}
		} else {
			letter = consonants[rand.Intn(len(consonants))]
			name += letter
			if letter ==  "q" && i != randomNumber {
				letter = "u"
				name += letter
			} 
			if letter == "m" || letter == "n" || letter == "r" || letter == "l" {
				if rand.Intn(4) == 0 {
					name += letter
				}
			}
		}
	}

	if race == "démons" {
		if len(name)%2 == 0 {
			name = name + endWordDemon[rand.Intn(len(endWordDemon))]
		} else {
			name = name + endWordDemon[rand.Intn(len(endWordDemon)-1)]
		}
	}

	fmt.Println(setRace(race, name, strenght) + " " + setElement(element) + " et " + setArmor(armor))
	return setRace(race, name, strenght) + " " + setElement(element) + " et " + setArmor(armor)
}

