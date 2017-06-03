package sdv

import (
	"encoding/xml"
	"io"
)

type SaveGameInfo struct {
	ID         string
	Name       string `xml:"name"`
	Money      int    `xml:"money"`
	DateString string `xml:"dateStringForSaveGame"`
}

func NewSaveGameInfo(saveGameID string, input io.Reader) (*SaveGameInfo, error) {
	output := &SaveGameInfo{
		ID: saveGameID,
	}

	decoder := xml.NewDecoder(input)
	err := decoder.Decode(output)

	return output, err
}

func (s *SaveGameInfo) Date() (Date, error) {
	return ParsePrettyDate(s.DateString)
}
