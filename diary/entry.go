package diary

import (
	"io"
	"os/exec"

	"github.com/xrstf/stardew-diary/sdv"
)

type Entry struct {
	diary      *Diary
	Number     int
	CommitDate string
	Time       string
	ID         string
	IngameDate sdv.Date
	Special    string
	Properties string
}

func (e *Entry) SaveGame() (savegame *sdv.SaveGame, err error) {
	err = e.pipeFile(addExt(e.diary.saveGameID, "xml"), func(input io.Reader) (err error) {
		savegame, err = sdv.NewSaveGame(e.diary.saveGameID, input)
		return
	})

	return
}

func (e *Entry) SaveGameInfo() (info *sdv.SaveGameInfo, err error) {
	err = e.pipeFile("SaveGameInfo.xml", func(input io.Reader) (err error) {
		info, err = sdv.NewSaveGameInfo(e.diary.saveGameID, input)
		return
	})

	return
}

type fossilConsumer func(io.Reader) error

func (e *Entry) pipeFile(filename string, consumer fossilConsumer) error {
	cmd := exec.Command("fossil.exe", "cat", filename, "-r", e.ID)
	cmd.Dir = e.diary.directory

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := consumer(stdout); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
