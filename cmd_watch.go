package main

import (
	"log"
	"time"

	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

func watchCommand(ctx *cli.Context) {
	game, err := sdv.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Watching savegames for changes...")
	storage, err := diariesDirectory()
	if err != nil {
		log.Fatal(err)
	}

	for event := range game.WatchForChanges() {
		<-time.After(1 * time.Second)

		savegame := event.SaveGameID
		diary, err := diary.NewDiary(storage, game, savegame)
		if err != nil {
			log.Println(err)
			continue
		}

		switch event.Kind {
		case "created":
			log.Printf("Found new savegame %s.", savegame)
			diary.Record()

		case "changed":
			log.Printf("Detected new save in %s.", savegame)
			diary.Record()
		}
	}
}
