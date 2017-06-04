package differ

import "github.com/xrstf/stardew-diary/sdv"

type Yielder func(string)

type Differ interface {
	Diff(previous, current, next *sdv.SaveGame, output Yielder)
}
