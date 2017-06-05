package differ

import (
	"fmt"

	"github.com/xrstf/stardew-diary/sdv"
)

type MoneyBottomLine struct{}

func (d *MoneyBottomLine) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	moneyDiff := 0

	if previous != nil {
		moneyDiff = current.Player.TotalMoneyEarned - previous.Player.TotalMoneyEarned
	}

	if moneyDiff > 0 {
		out(fmt.Sprintf("I earned %d gold today.", moneyDiff))
	}
}
