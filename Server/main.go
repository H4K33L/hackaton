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
	/*
	The main func launch the server on the port 8080 and handle the different request.
	---------------------------------------------------------
	The func dosen't have error case.
	*/
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
	/*
	The generateJSON func add the monster to the JSON file.
	---------------------------------------------------------
	error case:
		- if the number is not a number
		- if the ID is not a number
	*/
	number, err := strconv.Atoi(r.URL.Query().Get("number")) // Get the number of monster to generate from the URL
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du nombre", http.StatusBadRequest)
		return
	}
	monsterType := r.URL.Query().Get("type") // Get the type of monster to generate from the URL
	ID, err := strconv.Atoi(r.URL.Query().Get("ID")) // Get the ID of the monster to generate from the URL
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de l'ID", http.StatusBadRequest)
		return
	}
	Monster.Generate(number, monsterType, ID) // Generate the monsters
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to the main page

	handler(w, r) // Render the HTML file
}

func handler(w http.ResponseWriter, r *http.Request) {
	/*
	The handler func handle the different request.
	---------------------------------------------------------
	error case:
		- if the file JSON is not found
		- if the JSON file is not valid
		- if the HTML file is not found
		- if the HTML file is not valid
	*/
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

	renderHTML(w, response.Monsters, "index.html") // Render the HTML file
}

func renderHTML(w http.ResponseWriter, data interface{}, templateFile string) {
	/*
	The renderHTML func render the HTML file with the data.
	---------------------------------------------------------
	error case:
		- if the HTML file is not found
		- if the HTML file is not valid
	*/
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
	/*
	The deleteMonster func delete the monster from the JSON file.
	---------------------------------------------------------
	error case:
		- if the file JSON is not found
		- if the JSON file is not valid
	*/
    name := r.URL.Query().Get("name") // Get the name of the monster to delete from the URL
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

    newMonsters := make([]monster, 0) // Create a new slice of monster
    for _, monster := range response.Monsters { // Add all the monsters except the one to delete
        if monster.Name != name { // If the monster is not the one to delete
            newMonsters = append(newMonsters, monster) // Add the monster to the new slice
        }
    }
    response.Monsters = newMonsters // Replace the old slice with the new one

    newData, err := json.Marshal(response) // Convert the new slice to JSON
    if err != nil {
        http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
        return
    }

    err = ioutil.WriteFile(filePath, newData, 0644) // Write the new JSON file
    if err != nil {
        http.Error(w, "Erreur lors de l'écriture dans le fichier JSON", http.StatusInternalServerError)
        return
    }

	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to the main page

	handler(w, r) // Render the HTML file
}
