package Monster

func GenerateMonster(monsterType string, ID int) monster {
	creature := monster{}
	creature.ID = ID
	creature.MonsterType = monsterType
	for !GoodID(creature) {
		creature = Size(creature)
	}
	return creature
}