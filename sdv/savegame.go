package sdv

import (
	"encoding/xml"
	"io"
)

type SaveGame struct {
	ID string

	Player struct {
		Money int `xml:"money"`
	} `xml:"player"`

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
