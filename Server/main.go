package main

import (
	"Monster"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type monster struct {
	Name				string				`json:"Name"`
	ID					int					`json:"ID"`
	MonsterType			string				`json:"MonsterType"`
	Size				string				`json:"Size"`
	Alignment			string				`json:"Alignment"`
	Caract				map[string]int		`json:"Caract"`
	CaractMod			map[string]int		`json:"CaractMod"`
	Mastery				int					`json:"Mastery"`
	AC					string				`json:"AC"`
	LP					string				`json:"LP"`
	Resistance			[]string			`json:"Resistance"`
	Vulnerability		[]string			`json:"Vulnerability"`
	Immunity			[]string			`json:"Immunity"`
	AttacBonnus			int					`json:"AttacBonnus"`
	DD					int					`json:"DD"`
	Speed				map[string]int		`json:"Speed"`
	SaveRoll			map[string]int		`json:"SaveRoll"`
	StateImmunity		[]string			`json:"StateImmunity"`
	Sense				[]string			`json:"Sense"`
	Languages			[]string			`json:"Languages"`
}

type groupMonster struct {
	Monsters			[]monster			`json:"Monsters"`
}


func main() {
	blue := "\033[1;34m"
	green := "\033[1;32m"
	red := "\033[1;31m"

	fmt.Println(blue, "starting server on port 8080...")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/list", generateJSON)
	http.HandleFunc("/delete", deleteMonster)
	http.HandleFunc("/", handler)

	fmt.Println(string(green), "Server started on port 8080")
	fmt.Println(string(red), "Press Ctrl+C to quit")

	http.ListenAndServe(":8080", nil)
}

func generateJSON(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du nombre", http.StatusBadRequest)
		return
	}
	monsterType := r.URL.Query().Get("type")
	ID, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de l'ID", http.StatusBadRequest)
		return
	}
	Monster.Generate(number, monsterType, ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)

	handler(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	filePath := "../Data/monster.json"
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
		return
	}

	var response groupMonster
	err = json.Unmarshal(body, &response)
		if err != nil {
    http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
    return
	}

	renderHTML(w, response.Monsters, "index.html")
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

func deleteMonster(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    filePath := "../Data/monster.json"
    body, err := ioutil.ReadFile(filePath)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
        return
    }
    
    var response groupMonster
    err = json.Unmarshal(body, &response)
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
        return
    }

    newMonsters := make([]monster, 0)
    for _, monster := range response.Monsters {
        if monster.Name != name {
            newMonsters = append(newMonsters, monster)
        }
    }
    response.Monsters = newMonsters

    newData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
        return
    }

    err = ioutil.WriteFile(filePath, newData, 0644)
    if err != nil {
        http.Error(w, "Erreur lors de l'écriture dans le fichier JSON", http.StatusInternalServerError)
        return
    }

	http.Redirect(w, r, "/", http.StatusSeeOther)

	handler(w, r)
}
