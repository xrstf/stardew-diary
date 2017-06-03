package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "stardew-diary"
	app.Usage = "Keeps backups of your savegames."
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "watch",
			Usage:  "Watches savegames for changes. Let this run in the background while you are playing.",
			Action: watchCommand,
		},
		cli.Command{
			Name:      "log",
			Usage:     "Prints a list of backups for a given savegame.",
			Action:    logCommand,
			ArgsUsage: "SAVEGAME",
		},
		cli.Command{
			Name:      "revert",
			Usage:     "Reverts a savegame back to a specific date in the past.",
			Action:    revertCommand,
			ArgsUsage: "SAVEGAME REVISION",
		},
	}

	app.Run(os.Args)
}
