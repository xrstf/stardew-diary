package main

import (
	"log"
	"strings"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

func dumpCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	args := ctx.Args()
	if len(args) == 0 {
		log.Fatalln("No savegame given.")
	}
	if len(args) == 1 {
		log.Fatalln("No revision(s) given.")
	}

	savegame, err := matchProfile(args[0], game.SaveGameIDs())
	if err != nil {
		log.Fatalln(err)
	}

	storage, err := diariesDirectory()
	if err != nil {
		log.Fatal(err)
	}

	diary, err := diary.NewDiary(storage, game, savegame)
	if err != nil {
		log.Fatalln(err)
	}

	entries, err := diary.Entries()
	if err != nil {
		log.Fatalln(err)
	}

	revisions := args[1:]
	for idx, rev := range revisions {
		if len(rev) > 10 {
			rev = rev[:10]
		}

		revisions[idx] = strings.ToLower(rev)
	}

	for _, entry := range entries {
		commitID := entry.ID

		for _, rev := range revisions {
			if rev == "all" || strings.HasPrefix(commitID, rev) {
				log.Printf("Dumping revision %s...\n", commitID)
				entry.Dump()
			}
		}
	}
}
