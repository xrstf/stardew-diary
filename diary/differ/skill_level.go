package differ

import (
	"fmt"

	"github.com/xrstf/stardew-diary/sdv"
)

var skillLevelNames = map[sdv.PlayerSkill]string{
	sdv.FarmingSkill:  "Farming",
	sdv.MiningSkill:   "Mining",
	sdv.CombatSkill:   "Combat",
	sdv.ForagingSkill: "Foraging",
	sdv.FishingSkill:  "Fishing",
}

type SkillLevel struct{}

func (d *SkillLevel) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	for skillID, skillName := range skillLevelNames {
		curValue := current.SkillLevel(skillID)
		prevValue := 0

		if previous != nil {
			prevValue = previous.SkillLevel(skillID)
		}

		if curValue > prevValue {
			out(fmt.Sprintf("I've reached %s Level %d!", skillName, curValue))
		}
	}
}
