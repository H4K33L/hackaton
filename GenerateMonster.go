package Monster

import "math/rand"

func GenerateMonster(monsterType string, ID int) monster {
	creature := monster{}
	creature.ID = ID
	creature.MonsterType = monsterType
	for !GoodID(creature) {
		creature = Size(creature)
	}
	AC := GetLP(creature.AC)
	creature.Name = nameGenerator(creature.Caract["For"], monsterType, AC)
	return creature
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

	return "(" + race + " " + setStrength(strenght) +") " + name + " " + setArmor(armor)
}

func Generate(monsterType string, ID int) error {
	monsters := groupMonster{}

	filepath := "Data/monster.json"

	monsters, err := PullMonsters(filepath)
	if err != nil {
		return err
	}

	creature := GenerateMonster(monsterType, ID)
	monsters.Monsters = append(monsters.Monsters, creature)

	err = PushMonsters(filepath, monsters)
	if err != nil {
		return err
	}
	
	return nil
}

/*-----------------------------------------------------------------*/

func GenerateGroup(monsterType string, ID int) groupMonster {
	list := ListID(ID)
	for !GoodListID(list, ID) {
		list = ListID(ID)
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