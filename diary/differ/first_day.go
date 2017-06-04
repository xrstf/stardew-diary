package differ

import "github.com/xrstf/stardew-diary/sdv"

type FirstDay struct{}

func (d *FirstDay) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	if current.Stats.DaysPlayed == 1 {
		out("I arrived on my farm in Stardew Valley.")
	}
}
