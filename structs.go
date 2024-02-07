package Monster

type monster struct {
	Name				string
	ID					int
	MonsterType			string
	Size				string
	Alignment			string
	Caract				map[string]int
	CaractMod			map[string]int
	Mastery				int
	AC					string
	LP					string
	Resistance			[]string
	Vulnerability		[]string
	Immunity			[]string
	AttacBonnus			int
	DD					int
	Speed				map[string]int
	SaveRoll			map[string]int
	StateImmunity		[]string
	Sense				[]string
	Languages			[]string
}

type groupMonstre struct {

}

type table_rencontre struct {

}