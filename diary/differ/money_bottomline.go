package differ

import (
	"fmt"

	"github.com/xrstf/stardew-diary/sdv"
)

type MoneyBottomLine struct{}

func (d *MoneyBottomLine) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	moneyDiff := 0

	if previous != nil {
		moneyDiff = current.Player.Money - previous.Player.Money
	}

	if moneyDiff > 0 {
		out(fmt.Sprintf("I earned %d gold.", moneyDiff))
	} else if moneyDiff < 0 {
		out(fmt.Sprintf("I lost %d gold.", -moneyDiff))
	}
}
