package diary

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xrstf/stardew-diary/sdv"
	"github.com/xrstf/stardew-diary/xml"
)

type Diary struct {
	directory  string
	repo       string
	game       *sdv.Game
	saveGameID string
}

func diaryFilename(saveGameID string) string {
	return addExt(strings.ToLower(saveGameID), "diary")
}

func NewDiary(storageDirectory string, game *sdv.Game, saveGameID string) (*Diary, error) {
	repo := filepath.Join(storageDirectory, diaryFilename(saveGameID))
	diary := &Diary{
		directory:  game.SaveGameDirectory(saveGameID),
		repo:       repo,
		game:       game,
		saveGameID: saveGameID,
	}

	os.MkdirAll(storageDirectory, 0755)

	diary.fossil("init", repo)
	diary.fossil("open", repo, "--keep")

	return diary, nil
}

func (d *Diary) Entries() ([]Entry, error) {
	output, err := d.fossil("timeline", "ancestors", "tip", "--limit", "0", "--width", "0")
	if err != nil {
		return nil, err
	}

	return d.parseEntries(output), nil
}

func (d *Diary) Entry(entryID string) (Entry, error) {
	output, err := d.fossil("timeline", entryID, "--limit", "1", "--width", "0")
	if err != nil {
		return Entry{}, err
	}

	entries := d.parseEntries(output)
	if len(entries) == 0 {
		return Entry{}, fmt.Errorf("Could not find diary entry %s.", entryID)
	}

	return entries[0], nil
}

func (d *Diary) parseEntries(output string) []Entry {
	dateRegexp := regexp.MustCompile(`^=== (\d\d\d\d-\d\d-\d\d) ===$`)
	entryRegexp := regexp.MustCompile(`^(\d\d:\d\d:\d\d) \[([a-f0-9]+)\] (\*([A-Z]+)\* )?(\d\d\d\d-\d-\d\d) \((.+?)\)$`)
	currentDate := ""
	entries := make([]Entry, 0)

	for _, line := range strings.Split(output, "\n") {
		match := dateRegexp.FindStringSubmatch(line)
		if match != nil {
			currentDate = match[1]
		} else {
			match = entryRegexp.FindStringSubmatch(line)
			if match != nil {
				ingameDate, _ := sdv.ParseDate(match[5])
				newEntry := Entry{
					diary:      d,
					CommitDate: currentDate,
					Time:       match[1],
					ID:         match[2],
					Special:    match[4],
					IngameDate: ingameDate,
					Properties: match[6],
				}

				entries = append(entries, newEntry)
			}
		}
	}

	total := len(entries)

	for idx := range entries {
		entries[idx].Number = total - idx
	}

	return entries
}

func (d *Diary) files() []string {
	return []string{
		d.saveGameID,
		"SaveGameInfo",
	}
}

func (d *Diary) Record() error {
	for _, file := range d.files() {
		xml.FormatFile(filepath.Join(d.directory, file))
		d.fossil("add", addExt(file, "xml"))
	}

	changes, _ := d.fossil("changes")
	if len(changes) > 0 {
		// determine date in the game
		sg, err := d.game.SaveGame(d.saveGameID)
		if err != nil {
			return err
		}

		log.Printf("Recording changes for %s...", d.saveGameID)
		d.fossil("commit", "--allow-fork", "--comment", sg.Date().ID())
	}

	return nil
}

func (d *Diary) Revert(entryID string) error {
	// restore the pretty-printed .xml files
	if _, err := d.fossil("checkout", entryID, "--force"); err != nil {
		return err
	}

	// reconstruct the minified XML files that the game is actually using
	os.Chdir(d.directory)

	for _, file := range d.files() {
		pretty := addExt(file, "xml")
		minified := addExt(pretty, "mini")

		xml.MinifyFile(pretty)
		os.Rename(minified, file)
	}

	return nil
}

func (d *Diary) fossil(args ...string) (string, error) {
	c := exec.Command("fossil.exe", args...)
	c.Dir = d.directory

	out, err := c.CombinedOutput()

	return string(out), err
}
