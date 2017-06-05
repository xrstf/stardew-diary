package differ

import (
	"fmt"
	"strconv"

	"github.com/xrstf/stardew-diary/sdv"
)

type LostBooks struct{}

func (d *LostBooks) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	booksFound := current.Stats.NotesFound
	prevFound := 0

	if previous != nil {
		prevFound = previous.Stats.NotesFound
	}

	if booksFound > prevFound {
		diff := booksFound - prevFound
		out(fmt.Sprintf("While digging around, I found %s Lost Book%s :o", d.word(diff), d.suffix(diff)))
	}
}

func (d *LostBooks) suffix(num int) string {
	if num == 1 {
		return ""
	}

	return "s"
}

func (d *LostBooks) word(num int) string {
	switch num {
	case 1:
		return "a"
	case 2:
		return "TWO"
	case 3:
		return "T H R E E"
	case 4:
		return "F O U R"
	case 5:
		return "F I V E"
	default:
		return strconv.Itoa(num)
	}
}
