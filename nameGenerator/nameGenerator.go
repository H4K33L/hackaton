package main

import (
	"Monster"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Monsters struct {
	Name           string
	ID             int
	MonsterType    string
	Size           string
	Alignment      string
	Caract         map[string]int
	CaractMod      map[string]int
	Mastery        int
	AC             string
	LP             string
	Resistance     []string
	Vulnerability  []string
	Immunity       []string
	AttacBonnus    int
	DD             int
	Speed          map[string]int
	SaveRoll       map[string]int
	StateImmunity  []string
	Sense          []string
	Languages      []string
}

var race = []string{"Aberration", "Bête", "Artéfact", "Dragon", "Céleste", "Élémentaire", "Fée", "Démon", "Géant", "Humanoïde", "Monstruosité", "Plante", "Mort-vivant"}

func main() {
	http.HandleFunc("/list", generateJSON)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func generateJSON(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	monsterType := r.URL.Query().Get("type")
	fmt.Println("Nombre de monstres à générer:", number)
	fmt.Println("Type de monstre à générer:", monsterType)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du nombre", http.StatusBadRequest)
		return
	}
	generate(number, monsterType)
	http.Redirect(w, r, "/", http.StatusSeeOther)

	handler(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	filePath := "monster.json"
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
		return
	}

	var names []Monsters
	err = json.Unmarshal(body, &names)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
		return
	}

	renderHTML(w, names, "index.html")
}

func renderHTML(w http.ResponseWriter, data interface{}, templateFile string) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du modèle HTML", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur lors du rendu HTML", http.StatusInternalServerError)
		return
	}
}

func generate(number int, monsterType string) {
	var monsters []Monsters

	filepath := "monster.json"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
		return
	}

	var existingData []Monsters
	if len(content) > 0 {
		err = json.Unmarshal(content, &existingData)
		if err != nil {
			fmt.Println("Erreur lors de l'analyse JSON du contenu existant:", err)
			return
		}
	}

	for i := 0; i < number; i++ {
		strength := rand.Intn(30) + 1
		armor := rand.Intn(30) + 1
		name := nameGenerator(strength, monsterType, armor)
		ID := 1
		monster := Monster.GenerateMonster(monsterType, ID)
		monsterData := Monsters{
			Name:           name,
			ID:             ID,
			MonsterType:    monster.MonsterType,
			Size:           monster.Size,
			Alignment:      monster.Alignment,
			Caract:         monster.Caract,
			CaractMod:      monster.CaractMod,
			Mastery:        monster.Mastery,
			AC:             monster.AC,
			LP:             monster.LP,
			Resistance:     monster.Resistance,
			Vulnerability:  monster.Vulnerability,
			Immunity:       monster.Immunity,
			AttacBonnus:    monster.AttacBonnus,
			DD:             monster.DD,
			Speed:          monster.Speed,
			SaveRoll:       monster.SaveRoll,
			StateImmunity:  monster.StateImmunity,
			Sense:          monster.Sense,
			Languages:      monster.Languages,
		}

		monsters = append(monsters, monsterData)
	}

	existingData = append(existingData, monsters...)

	newContent, err := json.Marshal(existingData)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage JSON:", err)
		return
	}

	err = ioutil.WriteFile(filepath, newContent, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON:", err)
		return
	}

	fmt.Println("Monstres ajoutés avec succès au fichier JSON.")
}

func setStrength(strength int) string {
	power := ""

	if strength <= 5 {
		power = "faible"
	} else if strength <= 10 {
		power = "moyen"
	} else if strength <= 15 {
		power = "puissant"
	} else if strength <= 20 {
		power = "surpuissant"
	} else if strength <= 25 {
		power = "légendaire"
	} else if strength <= 29{
		power = "mythique"
	} else if strength == 30{
		power = "divin"
	}
	return power
}

func setArmor(armor int) string {
	protection := ""

	if armor <= 5 {
		protection = "sans-defense"
	} else if armor <= 10 {
		protection = "peux protégé"
	} else if armor <= 15 {
		protection = "protégé"
	} else if armor <= 20 {
		protection = "très protégé"
	} else if armor <= 25 {
		protection = "blindé"
	} else if armor <= 29{
		protection = "impenetrable"
	} else if armor == 30{
		protection = "invincible"
	}

	return protection
}

func nameGenerator(strenght int, race string, armor int) string {
	name := ""
	switch race {
	case "Aberration":
		name = generateAberrationName()
	case "Bête":
		name = generateBeastName()
	case "Construction":
		name = generateConstructionName()
	case "Dragon":
		name = generateDragonName()
	case "Céleste":
		name = generateCelestialName()
	case "Élémentaire":
		name = generateElementalName()
	case "Fée":
		name = generateFairyName()
	case "Démon":
		name = generateDemonName()
	case "Géant":
		name = generateGiantName()
	case "Humanoïde":
		name = generateHumanoidName()
	case "Monstruosité":
		name = generateMonstrosityName()
	case "Plante":
		name = generatePlantName()
	case "Mort-vivant":
		name = generateUndeadName()
	}

	fmt.Println("(" + race + " " + setStrength(strenght) + ") " + name + " " + setArmor(armor))
	return "(" + race + " " + setStrength(strenght) +") " + name + " " + setArmor(armor)
}

func generateAberrationName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    aberrationEndings := []string{"grix", "xal", "vor", "zith", "thor", "quex", "lix", "vyr", "nar", "zur", "myr"}

    return generateName(vowels, consonants, aberrationEndings)
}

func generateBeastName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    beastEndings := []string{"fang", "claw", "fur", "mane", "tooth", "hide", "scale", "talon", "growl", "snarl", "roar"}

    return generateName(vowels, consonants, beastEndings)
}

func generateConstructionName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    constructionEndings := []string{"forge", "stone", "hammer", "builder", "craft", "mason", "wright", "arch", "construct", "craftsman", "artisan"}

    return generateName(vowels, consonants, constructionEndings)
}

func generateDragonName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    dragonEndings := []string{"gon", "drak", "myr", "sorn", "fyre", "thor", "cryx", "lorn", "shyx", "wyr", "garr"}

    return generateName(vowels, consonants, dragonEndings)
}

func generateCelestialName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    celestialEndings := []string{"el", "ius", "on", "a", "iel", "or", "an", "eth", "iel", "iel", "ius"}

    return generateName(vowels, consonants, celestialEndings)
}

func generateElementalName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    elementalEndings := []string{"us", "en", "or", "il", "ar", "th", "on", "ix", "al", "ir", "er"}

    return generateName(vowels, consonants, elementalEndings)
}

func generateFairyName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    fairyEndings := []string{"dara", "wyn", "belle", "luna", "star", "ella", "fey", "briar", "thistle", "ivy", "lily"}

    return generateName(vowels, consonants, fairyEndings)
}

func generateDemonName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    demonEndings := []string{"on", "yx", "mor", "thor", "gath", "nyx", "zel", "vex", "zor", "mar", "rak"}

    return generateName(vowels, consonants, demonEndings)
}

func generateGiantName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    giantEndings := []string{"rock", "stone", "hurl", "smash", "crush", "thud", "grind", "brawl", "bash", "quake", "smash"}

    return generateName(vowels, consonants, giantEndings)
}

func generateHumanoidName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    humanoidEndings := []string{"um", "ar", "on", "ix", "el", "yss", "or", "en", "io", "us", "ath"}

    return generateName(vowels, consonants, humanoidEndings)
}

func generateMonstrosityName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    monstrosityEndings := []string{"ith", "gor", "lox", "shun", "fex", "garr", "thor", "maw", "grox", "lith", "sor"}

    return generateName(vowels, consonants, monstrosityEndings)
}

func generatePlantName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    plantEndings := []string{"thorn", "bloom", "leaf", "root", "petal", "stalk", "fern", "vine", "moss", "seed", "twig"}

    return generateName(vowels, consonants, plantEndings)
}

func generateUndeadName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    undeadEndings := []string{"aith", "hade", "tre", "wight", "eaper", "oul", "bie", "oul", "rit", "shee", "per"}

    return generateName(vowels, consonants, undeadEndings)
}

func generateName(vowels, consonants, endings []string) string {
	name := ""
	nameLength := rand.Intn(2) + 4
	for i := 0; i < nameLength; i++ {
		if i%2 == 0 {
			letter := vowels[rand.Intn(len(vowels))]
			name += letter
			if letter == "o" || letter == "u" || letter == "a" {
				if rand.Intn(5) == 0 {
					name += letter
				}
			}
		} else {
			letter := consonants[rand.Intn(len(consonants))]
			name += letter
			if letter == "m" || letter == "n" || letter == "r" || letter == "l" {
				if rand.Intn(5) == 0 {
					name += letter
				}
			}
		}
	}

	if endings != nil {
		if len(name)%2 == 0 {
			name += endings[rand.Intn(len(endings))]
		} else {
			name += endings[rand.Intn(len(endings)-1)]
		}
	}

	return name
}
