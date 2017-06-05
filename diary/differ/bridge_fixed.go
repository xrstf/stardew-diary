package differ

import "github.com/xrstf/stardew-diary/sdv"

type BridgeFixed struct{}

func (d *BridgeFixed) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	if current.Beach().BridgeFixed {
		if previous == nil || !previous.Beach().BridgeFixed {
			out("I've spent 300 wood to fix the bridge at the beach.")
		}
	}
}
