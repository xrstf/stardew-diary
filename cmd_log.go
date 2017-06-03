package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

func logCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	args := ctx.Args()
	if len(args) == 0 {
		log.Fatalln("No savegame given.")
	}

	savegame, err := matchProfile(args[0], game.SaveGameIDs())
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := diariesDirectory()
	if err != nil {
		log.Fatal(err)
	}

	info, err := game.SaveGameInfo(savegame)
	if err != nil {
		log.Fatalln(err)
	}

	heading := fmt.Sprintf("%s Diary", possession(info.Name))

	fmt.Println("")
	fmt.Printf("  %s\n", heading)
	fmt.Printf("  %s\n", strings.Repeat("=", len(heading)))
	fmt.Println("")

	diary, err := diary.NewDiary(storage, game, savegame)
	if err != nil {
		log.Fatalln(err)
	}

	entries, err := diary.Entries()
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range entries {
		savegame, err := entry.SaveGameInfo()
		if err != nil {
			log.Fatalln(err)
		}

		prefix := " "
		if entry.Special == "CURRENT" {
			prefix = "*"
		}

		fmt.Printf("%s [%02d] [%s] %s (%d G)\n", prefix, entry.Number, entry.ID, entry.IngameDate.String(), savegame.Money)
	}
}
