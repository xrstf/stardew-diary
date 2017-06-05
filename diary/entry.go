package diary

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

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
	Changes    []string
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

func (e *Entry) Dump() error {
	dir := e.diary.directory
	suffix := fmt.Sprintf("-%04d-%s.xml", e.Number, e.ID)

	for _, file := range []string{e.diary.saveGameID, "SaveGameInfo"} {
		err := e.pipeFile(addExt(file, "xml"), func(input io.Reader) error {
			out, err := os.Create(filepath.Join(dir, file+suffix))
			if err != nil {
				return err
			}
			defer out.Close()

			_, err = io.Copy(out, input)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
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
