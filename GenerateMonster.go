package Monster

import (
	"fmt"
	"math/rand"
)

/*----------- This part of functions arent used to generate monster -----------*/

func GenerateMonster(monsterType string, ID int) monster {
	/*
	The GenerateMonster function take a posible type for monster and an ID,
	the function create a monster struct and fill it. At the end the func
	return the struct completly full whith monster information, generate by
	the "whaterfall" function Size.
	------------------------------------------------------------------------
	input : a string representing the monster type, and the ID of the monster.
	output : a monster stuct fill whith information generate whith ID and
			monster type.
	------------------------------------------------------------------------
	The func dosen't have error case.
	*/

	creature := monster{}						// Initialize the empty monster struct, The monster struct is in structs.go
	creature.ID = ID
	creature.MonsterType = monsterType
	for !GoodID(creature) {						// The loop test if the monster generate follow the rulle corre, GoodID func is in verification.go
		
		creature = Size(creature)				/* The Size func is the begining of the generation chain and it return a full monster, Size func 
												   and all other chained func arent in GenerationFunc.go */
	}

	AC := GetLP(creature.AC)					// The GetLP func is used here to get the AC value, GetLP is in verification.go
	
	creature.Name = nameGenerator(creature.Caract["For"], monsterType, AC)	// The nameGeneration is used to generate the name for monster

	return creature
}

func nameGenerator(strenght int, race string, armor int) string {
	/*
	The nemeGeneration func take some information about monster, strenght,
	race (the type) and the armor. The func use the race to generate name
	folowing name rule and add adjective to the name monster to mesure is power.
	----------------------------------------------------------------------------
	input : information about monster, strenght and armor both int and race a string
	output : the monster name in string 
	----------------------------------------------------------------------------
	The func dosen't have error case.
	*/
	
	name := ""

	/* all the folowing generation"type"Name is in NameGenerator.go, and are used to generate  name acording the race */

	switch race {
	case "Aberration":
		name = generateAberrationName()
	case "Bête":
		name = generateBeastName()
	case "Artéfact":
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

	/* the setStrenght and setArmor are used to add adjective to the monster name, the both function aren in NameGenerator.go */
	return "(" + race + " " + setStrength(strenght) +") " + name + " " + setArmor(armor)
}

func Generate(number int, monsterType string, ID int) error {
	/*
	The generate func take the monster type and the ID and generate 
	a monster, after the generationthe monster is add to the json to 
	feed the monster catalog.
	-----------------------------------------------------------------
	input : the monstr type a string and ID an int
	output : the error see during the function run
	-----------------------------------------------------------------
	the function return error.
	*/

	for i := 0; i < number; i++ {

		monsters := groupMonster{}						// initilize an empty groupMonster struct, the groupMonster struct is in struct.go

		filepath := "../Data/monster.json"
		
		monsters, err := PullMonsters(filepath)			// the PullMonster func is used to get the information in the json catalog
		if err != nil {
			return err
		}

		creature := GenerateMonster(monsterType, ID)
		fmt.Println(creature)	// Generate a monster whith GenerateMonster
		monsters.Monsters = append(monsters.Monsters, creature)

		err = PushMonsters(filepath, monsters)			// the PushMonster func is used to add the new monster to the json catalog
		if err != nil {
			return err
		}

	}
	
	return nil
}

/*----------------------------- The folowing func still unused but can be add to a future update ------------------------------------*/

func GenerateGroup(monsterType string, ID int) groupMonster {
	list := ListID()
	for !GoodListID(list, ID) {
		list = ListID()
	}
	output := groupMonster{}
	for i := range list {
		output.Monsters = append(output.Monsters, GenerateMonster(monsterType, list[i]))
	}
	return output
}

func GenerateEncounter(monsterType string, ID int) encounterTable {
	output := encounterTable{}
	temp := rand.Intn(3)
	switch temp {
	case 0:
		for i := 0 ; i <=6 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	case 1:
		for i := 0 ; i <=8 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	case 2:
		for i := 0 ; i <=10 ; i++ {
			temp := rand.Intn(2)
			switch temp {
			case 0 :
				lonly := groupMonster{}
				lonly.Monsters = append(lonly.Monsters,GenerateMonster(monsterType, ID))
				output.Encounter = append(output.Encounter, lonly)
			case 1 :
				output.Encounter = append(output.Encounter, GenerateGroup(monsterType, ID))
			}
		}
	}
	return output
}