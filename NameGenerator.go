package Monster

import "math/rand"

/*---------- This part of the code generate name for monster -----------*/

func NameOMaticAberrationEdition() string {
	/*
	The generateAberrationName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    aberrationEndings := []string{"grix", "xal", "vor", "zith", "thor", "quex", "lix", "vyr", "nar", "zur", "myr"}

    return generateName(vowels, consonants, aberrationEndings)	// Return a name generate with rules by generateName
}

func generateBeastName() string {
	/*
	The generateBeastName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    beastEndings := []string{"fang", "claw", "fur", "mane", "tooth", "hide", "scale", "talon", "growl", "snarl", "roar"}

    return generateName(vowels, consonants, beastEndings)	// Return a name generate with rules by generateName
}

func generateConstructionName() string {
	/*
	The generateConstructionName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    constructionEndings := []string{"forge", "stone", "hammer", "builder", "craft", "mason", "wright", "arch", "construct", "craftsman", "artisan"}

    return generateName(vowels, consonants, constructionEndings)	// Return a name generate with rules by generateName
}

func generateDragonName() string {
	/*
	The generateDragonName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    dragonEndings := []string{"gon", "drak", "myr", "sorn", "fyre", "thor", "cryx", "lorn", "shyx", "wyr", "garr"}

    return generateName(vowels, consonants, dragonEndings)	// Return a name generate with rules by generateName
}

func HeavenlyHandlebarMustacheMonikerMaker() string {
	/*
	The generateCelestialName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    celestialEndings := []string{"el", "ius", "on", "a", "iel", "or", "an", "eth", "iel", "iel", "ius"}

    return generateName(vowels, consonants, celestialEndings)	// Return a name generate with rules by generateName
}

func generateElementalName() string {
	/*
	The generateElementalName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    elementalEndings := []string{"us", "en", "or", "il", "ar", "th", "on", "ix", "al", "ir", "er"}

    return generateName(vowels, consonants, elementalEndings)	// Return a name generate with rules by generateName
}

func WhimsicalWandWieldersNameWhirler() string {
	/*
	The generateFairyName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    fairyEndings := []string{"dara", "wyn", "belle", "luna", "star", "ella", "fey", "briar", "thistle", "ivy", "lily"}

    return generateName(vowels, consonants, fairyEndings)	// Return a name generate with rules by generateName
}

func generateDemonName() string {
	/*
	The generateDemonName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    demonEndings := []string{"on", "yx", "mor", "thor", "gath", "nyx", "zel", "vex", "zor", "mar", "rak"}

    return generateName(vowels, consonants, demonEndings)	// Return a name generate with rules by generateName
}

func TitanicTitleTwister() string {
	/*
	The generateGiantName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    giantEndings := []string{"rock", "stone", "hurl", "smash", "crush", "thud", "grind", "brawl", "bash", "quake", "smash"}

    return generateName(vowels, consonants, giantEndings)	// Return a name generate with rules by generateName
}

func generateHumanoidName() string {
	/*
	The generateHumanoidName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    humanoidEndings := []string{"um", "ar", "on", "ix", "el", "yss", "or", "en", "io", "us", "ath"}

    return generateName(vowels, consonants, humanoidEndings)	// Return a name generate with rules by generateName
}

func FrankensteinsFabulousNameForge() string {
	/*
	The generateMonstrosityName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    monstrosityEndings := []string{"ith", "gor", "lox", "shun", "fex", "garr", "thor", "maw", "grox", "lith", "sor"}

    return generateName(vowels, consonants, monstrosityEndings)	// Return a name generate with rules by generateName
}

func generatePlantName() string {
	/*
	The generatePlantName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    plantEndings := []string{"thorn", "bloom", "leaf", "root", "petal", "stalk", "fern", "vine", "moss", "seed", "twig"}

    return generateName(vowels, consonants, plantEndings)	// Return a name generate with rules by generateName
}

func generateUndeadName() string {
	/*
	The generateUndeadName define parameters to 
	send to the generateName function.
	------------------------------------------------
	input : none
	output : the output of generateName function, a string
	------------------------------------------------
	There is no error case in this function.
	*/

    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    undeadEndings := []string{"aith", "hade", "tre", "wight", "eaper", "oul", "bie", "oul", "rit", "shee", "per"}

    return generateName(vowels, consonants, undeadEndings)	// Return a name generate with rules by generateName
}

func generateName(vowels, consonants, endings []string) string {
	/*
	The generateName take array of string who represent the 
	rules to make name acording to the monster type and 
	generate a name folowing these rules.
	The function return a name in string form.
	--------------------------------------------------------
	input : 3 array of string
	outpur : the name in string format
	--------------------------------------------------------
	The function dosen't have error case.
	*/

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

/*---------- in this part this is the function who define adjective for name monster ----------*/

func setStrength(strength int) string {
	/*
	The setStrength function take the strength modificator 
	of a monster, an int and return a adjective to describe 
	the strength of the monster.
	--------------------------------------------------------
	input : an int
	output : an adjective in string format
	--------------------------------------------------------
	the function dosen't have error case.
	*/

	if strength <= 5 {
		return "faible"
	} else if strength <= 10 {
		return "moyen"
	} else if strength <= 15 {
		return "puissant"
	} else if strength <= 20 {
		return "surpuissant"
	} else if strength <= 25 {
		return "légendaire"
	} else if strength <= 29{
		return "mythique"
	} else {
		return "divin"
	}
}

func setArmor(armor int) string {
	/*
	The setArmor function take the armor value of a 
	monster, an int and return a adjective to describe 
	the armor of the monster.
	--------------------------------------------------------
	input : an int
	output : an adjective in string format
	--------------------------------------------------------
	the function dosen't have error case.
	*/

	if armor <= 5 {
		return "sans-defense"
	} else if armor <= 10 {
		return "peux protégé"
	} else if armor <= 15 {
		return "protégé"
	} else if armor <= 20 {
		return "très protégé"
	} else if armor <= 25 {
		return "blindé"
	} else if armor <= 29{
		return "impenetrable"
	} else {
		return "invincible"
	}
}