package Monster

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

type encounterTable struct {
	Encounter			[]groupMonster
}