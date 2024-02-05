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
var race = []string{"dragon", "gobelin", "mort-vivant", "bête", "démons", "humanoïde"}

func main() {
	data := []Name{}
	for i := 0; i < 10; i++ {
		strenght := rand.Intn(20) + 1
		elementMonster := element[rand.Intn(len(element))]
		raceMonster := race[rand.Intn(len(race))]
		armor := rand.Intn(10) + 1
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
	air := []string{"aérien", "léger", "volant", "aérien", "céleste", "nuageux", "venteux", "tempétueux"}
	thunder := []string{"tonnant", "électrique", "foudroyant", "éclatant", "détonnant", "fulgurant", "électrisant"}
	magic := []string{"magique", "ensorcelé", "mystique", "surnaturel", "mystérieux", "enchanteur", "divin", "sacré"}

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

func setRace(race string, name string) string {
	raceSet := name + " (" + race +")"

	animals := []string{"loup", "ours", "chien", "serpent"}
	undead := []string{"zombie", "squelette", "mort-vivant", "fantôme"}

	switch race {
	case "bête":
		raceSet = animals[rand.Intn(len(animals))]
	case "mort-vivant":
		raceSet = undead[rand.Intn(len(undead))]
	}
	return raceSet
}

func setForce(force int) string {
	power := ""

	if force <= 10 {
		power = "faible"
	} else {
		power = "puissant"
	}

	return power
}

func setArmor(armor int) string {
	protection := ""

	if armor <= 5 {
		protection = "sans défense"
	} else {
		protection = "blindé"
	}

	return protection
}

func nameGenerator(strenght int, element string, race string, armor int) string {
	voyels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
	randomNumber := rand.Intn(4) + 3
	name := ""
	letter := ""

	for i := 0; i < randomNumber; i++ {
		if i%2 == 0 {
			letter = voyels[rand.Intn(len(voyels))]
			name += letter
		} else {
			letter = consonants[rand.Intn(len(consonants))]
			name += letter
			if letter ==  "q" && i != randomNumber {
				letter = "u"
				name += letter
			} 
		}
	}

	if string(name[len(name)-1]) == "u" && string(name[len(name)-2]) == "q" {
		letter = "e"
		name += letter
	}
	fmt.Println(setForce(strenght) + " " + setRace(race, name) + " " + setElement(element) + " " + setArmor(armor))
	return setForce(strenght) + " " + setRace(race, name) + " " + setElement(element) + " " + setArmor(armor)
}
