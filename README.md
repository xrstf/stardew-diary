# Stardew Diary

This project aims to build a simple backup tool for [Stardew Valley](http://stardewvalley.net/)
savegames (currently limited/only tested on Windows). The backup mechanism makes
use of [Fossil](http://fossil-scm.org/).

## Installation

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

## Usage

Stardew Diary automatically finds savegames and takes care of initializing the
backup repository for you. Backups are kept in the same folder where the
``stardew-diary.exe`` is placed, so that even when the game is deleted, all
backups are still available.

### Creating Backups

When playing the game, you will want to have the ``watch`` mode of SD running
in the background. Use your command prompt to run

    stardew-diary.exe watch

and then just let the window sit there. You can close and restart the process
at any time, just make sure it runs while you're playing.

Everytime you save, the watcher will notice the changed savegame file and
automatically create a backup for you.

### Checking Backups

You can see a list of all backups for a given savegame via

    stardew-diary.exe log [NAME]

The ``NAME`` is your savegame name, which consists of your player name and a
random number. To avoid having to remember the ID all the time, it's sufficient
to give just as many characters of your savegame to make it unique.

For example, suppose you have three savegames: **Testerman_12345**,
**Testerwoman_98765** and **Concerned_42069**. To get the logs for the **Concerned**
savegame, you can just specify the **c**:

    stardew-diary.exe log c

      Concerned's Diary
      =================

    * [11] [1085e24b9d] 11th of Spring, Year 1 (2143 G)
      [10] [674b39a802] 10th of Spring, Year 1 (272 G)
      [09] [0682111624] 9th of Spring, Year 1 (1816 G)
      [08] [c2cbc14269] 8th of Spring, Year 1 (1068 G)
      [07] [206951a3e1] 7th of Spring, Year 1 (1023 G)
      [06] [faec669912] 6th of Spring, Year 1 (808 G)
      [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
      [04] [624d3e1724] 4th of Spring, Year 1 (1433 G)
      [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

To get the **Testerwoman** account, you would need to specify at least **testerw**

    stardew-diary.exe log testerw

SD will tell you when it cannot uniquely identify the savegame:

    stardew-diary.exe log t
    2017/06/25 19:34:33 The savegame name 't' is ambiguous, could mean any of [Testerman_12345, Testerwoman_98765].
