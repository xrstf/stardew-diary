package sdv

import (
	"encoding/xml"
	"io"
)

type SGNPC struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
}

type SGGameLocation struct {
	Characters struct {
		NPCs []SGNPC `xml:"NPC"`
	} `xml:"characters"`
}

type SaveGame struct {
	ID string

	Player struct {
		Money int `xml:"money"`
	} `xml:"player"`

	Locations struct {
		GameLocations []SGGameLocation `xml:"GameLocation"`
	} `xml:"locations"`

	Stats struct {
		DaysPlayed int `xml:"daysPlayed"`
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

func (s *SaveGame) Date() Date {
	return Date(s.Stats.DaysPlayed)
}

func (s *SaveGame) Pet() *SGNPC {
	for _, loc := range s.Locations.GameLocations {
		for _, char := range loc.Characters.NPCs {
			if char.Type == "Dog" || char.Type == "Cat" {
				return &char
			}
		}
	}

	return nil
}
