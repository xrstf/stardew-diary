package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"
	cli "gopkg.in/urfave/cli.v1"
)

func resurrectCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	args := ctx.Args()
	if len(args) == 0 {
		log.Fatalln("No savegame ID given.")
	}

	id := args[0]
	if _, err := strconv.Atoi(id); err != nil {
		log.Fatalln("Give the numeric savegame ID only.")
	}

	storage, err := diariesDirectory()
	if err != nil {
		log.Fatal(err)
	}

	saved, err := diary.DiaryIDs(storage)
	if err != nil {
		log.Fatal(err)
	}

	savegameID := ""

	for _, savedID := range saved {
		if strings.HasSuffix(savedID, id) {
			savegameID = savedID
			break
		}
	}

	if savegameID == "" {
		log.Fatalln("Could not find a savegame backup with this ID.")
	}

	err = diary.Resurrect(storage, game, savegameID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("The savegame %s has been successfully restored.\n", savegameID)
}
