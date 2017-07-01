package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

type savegameItem struct {
	id    string
	alive bool
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

	aliveSavegames := game.SaveGameIDs()

	knownSavegames, err := diary.DiaryIDs(storage)
	if err != nil {
		log.Fatal(err)
	}

	all := make([]string, 0)

	for _, sg := range aliveSavegames {
		all = append(all, sg)
	}

	for _, sg := range knownSavegames {
		lower := strings.ToLower(sg)
		include := true

		for _, s := range aliveSavegames {
			if strings.ToLower(s) == lower {
				include = false
				break
			}
		}

		if include {
			all = append(all, sg)
		}
	}

	sort.Strings(all)

	items := make([]savegameItem, 0)

	for _, sg := range all {
		alive := false

		for _, s := range aliveSavegames {
			if s == sg {
				alive = true
				break
			}
		}

		items = append(items, savegameItem{sg, alive})
	}

	fmt.Println("")
	fmt.Println("  Savegames")
	fmt.Println("  =========")
	fmt.Println("")

	for _, item := range items {
		info, err := game.SaveGameInfo(item.id)
		if err != nil {
			log.Fatalln(err)
		}

		if item.alive {
			fmt.Printf("   - %s (%s)\n", info.Name, item.id)
		} else {
			fmt.Printf("   - %s (%s) [dead]\n", info.Name, item.id)
		}
	}
}
