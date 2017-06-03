package sdv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Event struct {
	SaveGameID string
	Kind       string
}

type Game struct {
	rootDir         string
	watchKillSwitch chan struct{}
}

func NewGame() (*Game, error) {
	dir := fmt.Sprintf(`%s\StardewValley\Saves`, os.Getenv("APPDATA"))

	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		return nil, errors.New("Could not find Stardew Valley savegame directory.")
	}

	return &Game{dir, make(chan struct{})}, nil
}

func (g *Game) SaveGameIDs() []string {
	ids := make([]string, 0)

	files, err := ioutil.ReadDir(g.rootDir)
	if err != nil {
		return ids
	}

	for _, file := range files {
		if file.IsDir() {
			ids = append(ids, file.Name())
		}
	}

	return ids
}

func (g *Game) StopWatching() {
	close(g.watchKillSwitch)
}

func (g *Game) WatchForChanges() <-chan Event {
	mtimes := make(map[string]time.Time)
	events := make(chan Event, 10)

	g.watchKillSwitch = make(chan struct{})

	go func() {
		for {
			select {
			case <-g.watchKillSwitch:
				close(events)
				return

			case <-time.After(1 * time.Second):
				savegames := g.SaveGameIDs()

				for _, savegameID := range savegames {
					mainFile := g.saveGameFile(savegameID)

					stats, err := os.Stat(mainFile)
					if err == nil {
						modtime := stats.ModTime()

						oldtime, ok := mtimes[savegameID]
						if ok {
							if modtime.After(oldtime) {
								events <- Event{savegameID, "changed"}
							}
						} else {
							events <- Event{savegameID, "created"}
						}

						mtimes[savegameID] = modtime
					}
				}
			}
		}
	}()

	return events
}

func (g *Game) SaveGameDirectory(id string) string {
	return filepath.Join(g.rootDir, id)
}

func (g *Game) saveGameFile(id string) string {
	return filepath.Join(g.SaveGameDirectory(id), id)
}

func (g *Game) saveGameInfoFile(id string) string {
	return filepath.Join(g.SaveGameDirectory(id), "SaveGameInfo")
}

func (g *Game) SaveGame(id string) (*SaveGame, error) {
	f, err := os.Open(g.saveGameFile(id))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return NewSaveGame(id, f)
}

func (g *Game) SaveGameInfo(id string) (*SaveGameInfo, error) {
	f, err := os.Open(g.saveGameInfoFile(id))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return NewSaveGameInfo(id, f)
}
