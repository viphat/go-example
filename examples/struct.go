package examples

import (
	"fmt"
	"math"
)

type power struct {
	attack  int
	defense int
}

type location struct {
	x float64
	y float64
	z float64
}

type nonPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

type attacker struct {
	attackPower int
	dmgBonus    int
}

type sword struct {
	attacker
	twoHanded bool
}

type gun struct {
	attacker
	bulletsRemaining int
}

type weapon interface {
	Wield() bool
}

func (s sword) Wield() bool {
	fmt.Println("You've wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You've wielded a gun!")
	return true
}

func wielder(w weapon) bool {
	fmt.Println("Wielding...")
	return w.Wield()
}

func (loc location) String() string {
	return fmt.Sprintf("(%f, %f, %f)", loc.x, loc.y, loc.z)
}

func (loc location) euclideanDistance(target location) float64 {
	return math.Sqrt(
		(loc.x-target.x)*(loc.x-target.x) +
			(loc.x-target.y)*(loc.x-target.y) +
			(loc.x-target.z)*(loc.x-target.z))
}

func (npc nonPlayerCharacter) distanceTo(target nonPlayerCharacter) float64 {
	return npc.loc.euclideanDistance(target.loc)
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

	sword1 := sword{attacker: attacker{attackPower: 1, dmgBonus: 5}, twoHanded: true}
	gun1 := gun{attacker: attacker{attackPower: 10, dmgBonus: 20}, bulletsRemaining: 11}
	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)
	wielder(sword1)
	wielder(gun1)
	fmt.Println(demon.loc)
	fmt.Printf("Npc %s is %f units away from Npv %s\n", demon.name, demon.distanceTo(anotherDemon), anotherDemon.name)
}
