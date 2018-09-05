package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/viphat/go-example/examples"
)

type dieRollFunc func(int) int

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func fakeDieRoll(size int) int {
	return 42
}

func getDieRolls() []dieRollFunc {
	return []dieRollFunc{
		dieRoll,
		fakeDieRoll,
	}
}

func main() {
	examples.PrintHelloWorld()

	var rolls = getDieRolls()
	for index, rollFunc := range rolls {
		fmt.Printf("Die Roll Attempt #%d, result: %d\n", index, rollFunc(10))
	}

	examples.ShowingNonPlayers()
}
