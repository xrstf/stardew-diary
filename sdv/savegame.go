package sdv

import (
	"encoding/xml"
	"io"
)

type PlayerSkill int

const (
	FarmingSkill PlayerSkill = iota
	MiningSkill
	CombatSkill
	ForagingSkill
	FishingSkill
)

type SGNPC struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
}

type SGItem struct {
	Type         string `xml:"type,attr"`
	Name         string `xml:"name"`
	UpgradeLevel string `xml:"upgradeLevel"`
}

type SGObject struct {
	Key   interface{} `xml:"key"`
	Value interface{} `xml:"value"`
}

type SGGameLocation struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name"`

	Characters struct {
		NPCs []SGNPC `xml:"NPC"`
	} `xml:"characters"`

	Objects struct {
		Items []SGObject `xml:"item"`
	} `xml:"objects"`

	// type=Beach
	BridgeFixed bool `xml:"bridgeFixed"`
}

type SaveGame struct {
	ID string

	Player struct {
		Money            int `xml:"money"`
		ClubCoins        int `xml:"clubCoins"`
		TotalMoneyEarned int `xml:"totalMoneyEarned"`

		FarmingLevel  int `xml:"farmingLevel"`
		MiningLevel   int `xml:"miningLevel"`
		CombatLevel   int `xml:"combatLevel"`
		ForagingLevel int `xml:"foragingLevel"`
		FishingLevel  int `xml:"fishingLevel"`

		Items struct {
			Items []SGItem `xml:"Item"`
		} `xml:"items"`
	} `xml:"player"`

	Locations struct {
		GameLocations []SGGameLocation `xml:"GameLocation"`
	} `xml:"locations"`

	Stats struct {
		DaysPlayed int `xml:"daysPlayed"`
		NotesFound int `xml:"notesFound"`
	} `xml:"stats"`
}

func NewSaveGame(id string, input io.Reader) (*SaveGame, error) {
	output := &SaveGame{
		ID: id,
	}

	decoder := xml.NewDecoder(input)
	err := decoder.Decode(output)

	return output, err
}

func (s *SaveGame) SkillLevel(skill PlayerSkill) int {
	switch skill {
	case FarmingSkill:
		return s.Player.FarmingLevel
	case MiningSkill:
		return s.Player.MiningLevel
	case CombatSkill:
		return s.Player.CombatLevel
	case ForagingSkill:
		return s.Player.ForagingLevel
	case FishingSkill:
		return s.Player.FishingLevel
	default:
		return -1
	}
}

func (s *SaveGame) Date() Date {
	return Date(s.Stats.DaysPlayed)
}

func (s *SaveGame) NPC(name string) *SGNPC {
	for _, loc := range s.Locations.GameLocations {
		for _, char := range loc.Characters.NPCs {
			if char.Type == name {
				return &char
			}
		}
	}

	return nil
}

func (s *SaveGame) Pet() *SGNPC {
	pet := s.NPC("Dog")
	if pet == nil {
		pet = s.NPC("Cat")
	}
	return pet
}

func (s *SaveGame) Location(name string) *SGGameLocation {
	for _, loc := range s.Locations.GameLocations {
		if loc.Type == name {
			return &loc
		}
	}

	return nil
}

func (s *SaveGame) Beach() *SGGameLocation {
	return s.Location("Beach")
}

func (s *SaveGame) Item(name string) *SGItem {
	for _, item := range s.Player.Items.Items {
		if item.Type == name {
			return &item
		}
	}

	return nil
}
