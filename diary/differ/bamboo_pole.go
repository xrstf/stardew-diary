package differ

import "github.com/xrstf/stardew-diary/sdv"

type BambooPole struct{}

func (d *BambooPole) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	if rod := current.Item("FishingRod"); rod != nil {
		if previous == nil || previous.Item("FishingRod") == nil {
			out("I met this nice old man by the beach. We talked about fishing for a while, before he gave me a fishing rod :) It's just a bamboo pole, but as time comes, I might find better rods...")
		}
	}
}
