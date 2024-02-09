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

type Monsters struct {
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

type MonsterResponse struct {
	Monsters []Monsters `json:"Monsters"`
}

func main() {
	blue := "\033[1;34m"
	green := "\033[1;32m"
	red := "\033[1;31m"

	fmt.Println(blue, "starting server on port 8080...")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("C:/Users/louka/Documents/YNOV/hackaton/nameGenerator/static/"))))

	http.HandleFunc("/list", generateJSON)
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
	Monster.Generate(monsterType, ID)
	fmt.Println("Nombre de monstres à générer:", number)
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

	var response MonsterResponse
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