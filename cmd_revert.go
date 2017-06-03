package main

import (
	"log"
	"strconv"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"
	cli "gopkg.in/urfave/cli.v1"
)

func revertCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatalln(err)
	}

	args := ctx.Args()
	if len(args) == 0 {
		log.Fatalln("No savegame given.")
	}
	if len(args) == 1 {
		log.Fatalln("No revision given.")
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

	total := len(entries)
	revision, err := strconv.Atoi(args[1])
	if err != nil || revision > total || revision < 1 {
		log.Fatalf("Revision must be a natural number between 1 and %d.\n", total)
	}

	idx := total - revision
	entry := entries[idx]

	if err := diary.Revert(entry.ID); err != nil {
		log.Fatalln(err)
	}
}
