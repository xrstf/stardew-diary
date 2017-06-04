package diary

import (
	"fmt"
	"strings"

	"github.com/xrstf/stardew-diary/sdv"
)

type Differ struct {
	counter int
}

func NewDiffer() *Differ {
	return &Differ{0}
}

func (d *Differ) Diff(previous, current, next *sdv.SaveGame) []string {
	changes := make([]string, 0)
	moneyDiff := 0

	if current.Stats.DaysPlayed == 1 {
		changes = append(changes, "I arrived on my farm in Stardew Valley.")
	}

	if pet := current.Pet(); pet != nil {
		changes = append(changes, fmt.Sprintf("I've adopted a %s and named it %s!", strings.ToLower(pet.Type), pet.Name))
	}

	if previous != nil {
		moneyDiff = current.Player.Money - previous.Player.Money
	}

	if moneyDiff > 0 {
		changes = append(changes, fmt.Sprintf("I earned %d gold.", moneyDiff))
	} else if moneyDiff < 0 {
		changes = append(changes, fmt.Sprintf("I lost %d gold.", -moneyDiff))
	}

	d.counter++

	return changes
}
