package differ

import (
	"fmt"
	"strings"

	"github.com/xrstf/stardew-diary/sdv"
	"github.com/xrstf/stardew-diary/sdv/data"
)

type MuseumAchievement struct{}

func (m *MuseumAchievement) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	prevDonated := m.getDonatedItems(previous)
	curDonated := m.getDonatedItems(current)

	if len(curDonated) > len(prevDonated) {
		text := ""

		if len(prevDonated) == 0 {
			text = "I've stumbled upon some old things and had to visit Gunther at the library to learn more about them. "
			text = text + "He encouraged me to search for more artifacts and minerals, so that I could donate them to the museum. "
			text = text + "So I went straight ahead and donated my "
		} else {
			text = "Once again I've found ancient stuff and went to the museum to donate the "
		}

		itemNames := make([]string, 0)

		for _, item := range m.itemDiff(curDonated, prevDonated) {
			itemNames = append(itemNames, item.Singular)
		}

		text += m.humanJoin(itemNames) + "."

		out(text)
	}
}

func (m *MuseumAchievement) getDonatedItems(sg *sdv.SaveGame) []*data.Item {
	items := make([]*data.Item, 0)

	if sg != nil {
		for _, item := range sg.LibraryMuseum().MuseumPieces.Items {
			items = append(items, data.ItemByID(item.Value.ItemID))
		}
	}

	return items
}

func (m *MuseumAchievement) itemDiff(current, previous []*data.Item) []*data.Item {
	diff := make([]*data.Item, 0)

	for _, cur := range current {
		curID := cur.ID
		previously := false

		for _, prev := range previous {
			if prev.ID == curID {
				previously = true
				break
			}
		}

		if !previously {
			diff = append(diff, cur)
		}
	}

	return diff
}

func (m *MuseumAchievement) humanJoin(items []string) string {
	num := len(items)

	switch num {
	case 0:
		return ""
	case 1:
		return items[0]
	default:
		return fmt.Sprintf("%s and %s", strings.Join(items[0:num-1], ", "), items[num-1])
	}
}
