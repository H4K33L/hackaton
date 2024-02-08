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

var race = []string{"Aberration", "Bête", "Construction", "Dragon", "Céleste", "Élémentaire", "Fée", "Démon", "Géant", "Humanoïde", "Monstruosité", "Plante", "Mort-vivant"}

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