package main

import (
	"fmt"
	"log"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
	"github.com/xrstf/stardew-diary/diary"
	"github.com/xrstf/stardew-diary/sdv"

	cli "gopkg.in/urfave/cli.v1"
)

func historyCommand(ctx *cli.Context) {
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

	entries, err := diary.History()
	if err != nil {
		log.Fatalln(err)
	}

	for entry := range entries {
		fmt.Printf("  %s\n", entry.IngameDate.String())

		if len(entry.Changes) == 0 {
			fmt.Printf("\n   ...nothing happened...\n\n")
		} else {
			fmt.Println("")

			for _, change := range entry.Changes {
				fmt.Printf(itemize(change, 3))
			}

			fmt.Println("")
		}
	}
}

func itemize(change string, leading int) string {
	wrapped := wordwrap.WrapString(change, 60)
	lines := strings.Split(wrapped, "\n")
	prefix := strings.Repeat(" ", leading)

	output := make([]string, len(lines))

	for idx, line := range lines {
		if idx == 0 {
			output[idx] = prefix + "- " + line
		} else {
			output[idx] = prefix + "  " + line
		}
	}

	return strings.Join(output, "\n") + "\n"
}
