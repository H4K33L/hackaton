package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Monster struct {
	Name				string         `json:"name"`
	KEY					int            `json:"key"`
	ID					int            `json:"id"`
	MonsterType			string         `json:"monsterType"`
	Size				string         `json:"size"`
	Alignment			string         `json:"alignment"`
	Caract				map[string]int `json:"caract"`
	CaractMod			map[string]int `json:"caractMod"`
	Mastery				int            `json:"mastery"`
	AC					string         `json:"ac"`
	LP					string         `json:"lp"`
	Resistance			[]string       `json:"resistance"`
	Vulnerability		[]string       `json:"vulnerability"`
	Immunity			[]string       `json:"immunity"`
	AttacBonnus			int            `json:"attacBonnus"`
	DD					int            `json:"dd"`
	Speed				map[string]int `json:"speed"`
	SaveRoll			map[string]int `json:"saveRoll"`
	
	StateImmunity		[]string       `json:"stateImmunity"`
	
	Sense				[]string       `json:"sense"`
	Languages			[]string       `json:"languages"`

	Attaques			[]string       `json:"attaques"`
	SpecialDeals		[]string       `json:"specialDeals"`
	Action				[]string       `json:"action"`
	Reaction			[]string       `json:"reaction"`
}

var element = []string{"terre", "feu", "eau", "air", "foudre", "magie"}
var race = []string{"démons", "bête", "mort-vivant", "humanoïde", "Dragon", "Géant"}
var monsterType = []string{"normal", "mythique", "élémentaire", "extraplanaire"}

func main() {
	http.HandleFunc("/list", generateJSON)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func generateJSON(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du nombre", http.StatusBadRequest)
		return
	}
	generate(number)
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

	var names []Monster
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


func generate(number int) {
	data := []Monster{}
	for i := 0; i < number; i++ {
		strength := rand.Intn(30) + 1
		elementMonster := element[rand.Intn(len(element))]
		raceMonster := race[rand.Intn(len(race))]
		monsterType := monsterType[rand.Intn(len(monsterType))]
		armor := rand.Intn(30) + 1
		name := nameGenerator(strength, elementMonster, raceMonster, monsterType, armor)
		data = append(data, Monster{name, 0, 0, "", "", "", nil, nil, 0, "", "", nil, nil, nil, 0, 0, nil, nil, nil, nil, nil, nil, nil, nil, nil})
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("monster.json", file, 0644)
	fmt.Println("File created")
}

func setElement(element string) string {
	adj := ""

	fire := []string{"brûlant", "ardent", "enflammé", "incandescent", "brûlé", "rougeoyant", "fumant", "torride"}
	water := []string{"mouillé", "glacial", "gelé", "humide", "ruisselant", "détrempé", "inondé", "débordant"}
	earth := []string{"féroce", "sanglant", "affamé", "indomptable", "écorché", "primordial" }
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

func nameGenerator(strenght int, element string, race string, monsterType string, armor int) string {
	name := ""
	adj := ""
	switch monsterType {
	case "normal":
		name, adj = generateNormalMonsterName(race)
	case "mythique":
		name, adj = generateMythicalMonsterName(race)
	case "élémentaire":
		name, adj = generateElementalMonsterName(race)
	case "extraplanaire":
		name, adj= generateExtraplanarMonsterName(race)
	}

	fmt.Println("(" + race + " " + adj + " " + setStrength(strenght) + ") " + name + " " + setElement(element) + " et " + setArmor(armor))
	return "(" + race + " " + adj + " " + setStrength(strenght) +") " + name + " " + setElement(element) + " et " + setArmor(armor)
}

func generateNormalMonsterName(race string) (string, string) {
	name := ""
	adjective := ""
	switch race {
	case "démons":
		name = generateMachiavelicPersonName()
		adjective = "démoniaque"
	case "bête":
		name = generateLitlleDogoName()
		adjective = "sauvage"
	case "mort-vivant":
		name = generateZomboboName()
		adjective = "macabre"
	case "humanoïde":
		name = generateMonsieurToutLeMondeName()
		adjective = "basique"
	case "Dragon":
		name = generateAutismoDargonoName()
		adjective = "sauvage"
	case "Géant":
		name = generateVeryTallPersonName()
		adjective = "colossal"
	}
	return name, adjective
}

func generateMythicalMonsterName(race string) (string, string) {
	name := ""
	adjective := ""
	switch race {
	case "démons":
		name = generateMachiavelicPersonName()
		adjective = "des abysses"
	case "bête":
		name = generateLitlleDogoName()
		adjective = "légendaire"
	case "mort-vivant":
		name = generateZomboboName()
		adjective = "maudit"
	case "humanoïde":
		name = generateMonsieurToutLeMondeName()
		adjective = "éternel"
	case "Dragon":
		name = generateAutismoDargonoName()
		adjective = "immortel"
	case "Géant":
		name = generateVeryTallPersonName()
		adjective = "titanesque"
	}
	return name, adjective
}

func generateElementalMonsterName(race string) (string, string) {
    name := ""
	adjective := ""
    switch race {
    case "démons":
        name = generateMachiavelicPersonName()
		adjective = "impétueux"
    case "bête":
        name = generateLitlleDogoName()
		adjective = "élémentaire"
    case "mort-vivant":
        name = generateZomboboName()
		adjective = "éthéré"
    case "humanoïde":
        name = generateMonsieurToutLeMondeName()
		adjective = "tourbillonnant"
    case "Dragon":
        name = generateAutismoDargonoName()
		adjective = "embrasé"
    case "Géant":
        name = generateVeryTallPersonName()
		adjective = "fulgurant"
    }
    return name, adjective
}

func generateExtraplanarMonsterName(race string) (string, string) {
	name := ""
	adjective := ""
	switch race {
	case "démons":
		name = generateMachiavelicPersonName()
		adjective = "des enfers"

	case "bête":
		name = generateLitlleDogoName()
		adjective = "de l'au-delà"
	case "mort-vivant":
		name = generateZomboboName()
		adjective = "spectral"
	case "humanoïde":
		name = generateMonsieurToutLeMondeName()
		adjective = "extradimensionnel"
	case "Dragon":
		name = generateAutismoDargonoName()
		adjective = "interplanétaire"
	case "Géant":
		name = generateVeryTallPersonName()
		adjective = "cosmique"
	}
	return name, adjective
}

func generateMachiavelicPersonName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}
	endings := []string{"og", "eth", "oth", "uk", "uur", "ach", "uz", "oz", "er", "org"}

	name := generateName(vowels, consonants, endings)
	return name
}

func generateLitlleDogoName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}

	name := generateName(vowels, consonants, nil)
	return name
}

func generateZomboboName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}
	endings := []string{"th", "x", "us", "os", "is", "on"}

	name := generateName(vowels, consonants, endings)
	return name
}

func generateMonsieurToutLeMondeName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}
	endings := []string{"us", "ius", "or", "ix", "is", "um"}

	name := generateName(vowels, consonants, endings)
	return name
}

func generateAutismoDargonoName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}
	endings := []string{"rax", "rath", "onax", "yth", "zar", "gon", "ra", "dros"}

	name := generateName(vowels, consonants, endings)
	return name
}

func generateVeryTallPersonName() string {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "w", "z"}
	endings := []string{"nor", "gorn", "farn", "grim", "thor", "mung", "gol", "rung"}

	name := generateName(vowels, consonants, endings)
	return name
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

