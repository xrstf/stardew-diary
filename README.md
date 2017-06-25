Stardew Diary
=============

This project aims to build a simple backup tool for [Stardew Valley](http://stardewvalley.net/)
savegames (currently limited/only tested on Windows). The backup mechanism makes
use of [Fossil](http://fossil-scm.org/).

Installation
------------

Download the latest release from the Github Release pages and extract it anywhere
you like, for example ``C:\Users\you\stardew-diary``.

Stardew Diary has no graphical user interface, so you must use the Windows
command prompt (or any other shell) to use it. Navigate to the directory where
you extracted the release and run the ``stardew-diary.exe``:

    > stardew-diary.exe
    NAME:
       stardew-diary - Keeps backups of your savegames.

    USAGE:
       stardew-diary.exe [global options] command [command options] [arguments...]

    VERSION:
       1.0.0

    COMMANDS:
         watch    Watches savegames for changes. Let this run in the background while you are playing.
         log      Prints a list of backups for a given savegame.
         history  Prints a nice list of achievements and general changes.
         revert   Reverts a savegame back to a specific date in the past.
         dump     Dumps a given diary revision to disk for further debugging.
         help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version

Usage
-----

Stardew Diary automatically finds savegames and takes care of initializing the
backup repository for you. Backups are kept in the same folder where the
``stardew-diary.exe`` is placed, so that even when the game is deleted, all
backups are still available.

Creating Backups
^^^^^^^^^^^^^^^^

When playing the game, you will want to have the ``watch`` mode of SD running
in the background. Use your command prompt to run

    stardew-diary.exe watch

and then just let the Window sit there. You can close and restart the process
at any time, just make sure it runs while you're playing.

Everytime you save, the watcher will notice the changed savegame file and
automatically create a backup for you.
