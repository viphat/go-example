package examples

import "fmt"

type power struct {
	attack  int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

// ShowingNonPlayers - Display some non-players
func ShowingNonPlayers() {
	fmt.Println("Structs... Demo")

	demon := nonPlayerCharacter{
		name:  "Alfred",
		speed: 21,
		hp:    1000,
		power: power{attack: 75, defense: 50},
		loc:   location{1075.123, 521.123, 211.231},
	}

	fmt.Println(demon)

	anotherDemon := nonPlayerCharacter{
		name:  "Beelzebub",
		speed: 30,
		hp:    5000,
		power: power{attack: 10, defense: 10},
		loc:   location{32.03, 72.45, 65.231},
	}

	fmt.Println(anotherDemon)
}
