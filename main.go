package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "stardew-diary"
	app.Usage = "Keeps backups of your savegames."
	app.Version = "1.0-dev"

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "savegames",
			Usage:  "Prints a list of known savegames.",
			Action: savegamesCommand,
		},
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
			Name:      "history",
			Usage:     "Prints a nice list of achievements and general changes.",
			Action:    historyCommand,
			ArgsUsage: "SAVEGAME",
		},
		cli.Command{
			Name:      "revert",
			Usage:     "Reverts a savegame back to a specific date in the past.",
			Action:    revertCommand,
			ArgsUsage: "SAVEGAME REVISION",
		},
		cli.Command{
			Name:      "resurrect",
			Usage:     "Restores a deleted savegame from backup.",
			Action:    resurrectCommand,
			ArgsUsage: "SAVEGAME_ID",
		},
		cli.Command{
			Name:      "dump",
			Usage:     "Dumps a given diary revision to disk for further debugging.",
			Action:    dumpCommand,
			ArgsUsage: "SAVEGAME REVISION[, ...]",
		},
	}

	app.Run(os.Args)
}
