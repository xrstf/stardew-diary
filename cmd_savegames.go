package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

type savegameItem struct {
	name   string
	id     string
	status string
}

func (i *savegameItem) UniqueID() int {
	parts := strings.Split(i.id, "_")
	parsed, _ := strconv.Atoi(parts[len(parts)-1])

	return parsed
}

type savegameItemByName []savegameItem

func (a savegameItemByName) Len() int {
	return len(a)
}

func (a savegameItemByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a savegameItemByName) Less(i, j int) bool {
	return a[i].name < a[j].name
}

func savegamesCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := diariesDirectory()
	if err != nil {
		log.Fatal(err)
	}

	alive := game.SaveGameIDs()
	saved, err := diary.DiaryIDs(storage)
	if err != nil {
		log.Fatal(err)
	}

	all := make(map[string]savegameItem)

	for _, sg := range alive {
		info, err := game.SaveGameInfo(sg)
		if err != nil {
			log.Fatal(err)
		}

		item := savegameItem{
			id:     sg,
			name:   info.Name,
			status: "new",
		}

		all[sg] = item
	}

	for _, sg := range saved {
		if item, exists := all[sg]; exists {
			item.status = "alive"
			all[sg] = item
		} else {
			diary, err := diary.NewDiary(storage, game, sg)
			if err != nil {
				log.Fatal(err)
			}

			latest, err := diary.LatestEntry()
			if err != nil {
				log.Fatal(err)
			}

			name := ""

			if latest != nil {
				info, err := latest.SaveGameInfo()
				if err != nil {
					log.Fatal(err)
				} else {
					name = info.Name
				}
			}

			item := savegameItem{
				id:     sg,
				name:   name,
				status: "dead",
			}

			all[sg] = item
		}
	}

	idx := 0
	flat := make([]savegameItem, len(all))
	for _, item := range all {
		flat[idx] = item
		idx++
	}

	sort.Sort(savegameItemByName(flat))

	fmt.Println("")
	fmt.Println("  Savegames")
	fmt.Println("  =========")
	fmt.Println("")

	hasNew := false
	hasDead := 0

	for _, item := range flat {
		switch item.status {
		case "alive":
			fmt.Printf("   - %s (%d)\n", item.name, item.UniqueID())
		case "dead":
			fmt.Printf("   - %s (%d) [dead]\n", item.name, item.UniqueID())
			hasDead = item.UniqueID()
		case "new":
			fmt.Printf("   - %s (%d) [new]\n", item.name, item.UniqueID())
			hasNew = true
		}
	}

	if hasNew {
		fmt.Println("")
		fmt.Println("  New savegames will be processed when the `match` command is next used.")
	}

	if hasDead > 0 {
		if !hasNew {
			fmt.Println("")
		}

		fmt.Println("  Dead savegames can be restored by running:")
		fmt.Println("")
		fmt.Printf("    stardew-diary.exe resurrect %d\n", hasDead)
	}
}
