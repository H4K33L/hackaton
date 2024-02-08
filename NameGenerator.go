package Monster

import "math/rand"

func generateAberrationName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    aberrationEndings := []string{"grix", "xal", "vor", "zith", "thor", "quex", "lix", "vyr", "nar", "zur", "myr"}

    return generateName(vowels, consonants, aberrationEndings)
}

func generateBeastName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    beastEndings := []string{"fang", "claw", "fur", "mane", "tooth", "hide", "scale", "talon", "growl", "snarl", "roar"}

    return generateName(vowels, consonants, beastEndings)
}

func generateConstructionName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    constructionEndings := []string{"forge", "stone", "hammer", "builder", "craft", "mason", "wright", "arch", "construct", "craftsman", "artisan"}

    return generateName(vowels, consonants, constructionEndings)
}

func generateDragonName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    dragonEndings := []string{"gon", "drak", "myr", "sorn", "fyre", "thor", "cryx", "lorn", "shyx", "wyr", "garr"}

    return generateName(vowels, consonants, dragonEndings)
}

func generateCelestialName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    celestialEndings := []string{"el", "ius", "on", "a", "iel", "or", "an", "eth", "iel", "iel", "ius"}

    return generateName(vowels, consonants, celestialEndings)
}

func generateElementalName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    elementalEndings := []string{"us", "en", "or", "il", "ar", "th", "on", "ix", "al", "ir", "er"}

    return generateName(vowels, consonants, elementalEndings)
}

func generateFairyName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    fairyEndings := []string{"dara", "wyn", "belle", "luna", "star", "ella", "fey", "briar", "thistle", "ivy", "lily"}

    return generateName(vowels, consonants, fairyEndings)
}

func generateDemonName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    demonEndings := []string{"on", "yx", "mor", "thor", "gath", "nyx", "zel", "vex", "zor", "mar", "rak"}

    return generateName(vowels, consonants, demonEndings)
}

func generateGiantName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    giantEndings := []string{"rock", "stone", "hurl", "smash", "crush", "thud", "grind", "brawl", "bash", "quake", "smash"}

    return generateName(vowels, consonants, giantEndings)
}

func generateHumanoidName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    humanoidEndings := []string{"um", "ar", "on", "ix", "el", "yss", "or", "en", "io", "us", "ath"}

    return generateName(vowels, consonants, humanoidEndings)
}

func generateMonstrosityName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    monstrosityEndings := []string{"ith", "gor", "lox", "shun", "fex", "garr", "thor", "maw", "grox", "lith", "sor"}

    return generateName(vowels, consonants, monstrosityEndings)
}

func generatePlantName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    plantEndings := []string{"thorn", "bloom", "leaf", "root", "petal", "stalk", "fern", "vine", "moss", "seed", "twig"}

    return generateName(vowels, consonants, plantEndings)
}

func generateUndeadName() string {
    vowels := []string{"a", "e", "i", "o", "u", "y"}
    consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
    undeadEndings := []string{"aith", "hade", "tre", "wight", "eaper", "oul", "bie", "oul", "rit", "shee", "per"}

    return generateName(vowels, consonants, undeadEndings)
}

func generateName(vowels, consonants, endings []string) string {
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

func setStrength(strength int) string {
	power := ""

	if strength <= 5 {
		power = "faible"
	} else if strength <= 10 {
		power = "moyen"
	} else if strength <= 15 {
		power = "puissant"
	} else if strength <= 20 {
		power = "surpuissant"
	} else if strength <= 25 {
		power = "légendaire"
	} else if strength <= 29{
		power = "mythique"
	} else if strength == 30{
		power = "divin"
	}
	return power
}

func setArmor(armor int) string {
	protection := ""

	if armor <= 5 {
		protection = "sans-defense"
	} else if armor <= 10 {
		protection = "peux protégé"
	} else if armor <= 15 {
		protection = "protégé"
	} else if armor <= 20 {
		protection = "très protégé"
	} else if armor <= 25 {
		protection = "blindé"
	} else if armor <= 29{
		protection = "impenetrable"
	} else if armor == 30{
		protection = "invincible"
	}

	return protection
}