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

### Reverting to a Previous Backup

When you look at the output of the ``log`` command

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

you will notice an asterisk, a decreasing number and some random gibberish. You
can basically ignore the gibberish, unless you want to dig into your backups
manually using Fossil.

The asterisk is telling you what day your are currently on. While playing the
game normally, the asterisk will always show the first (i.e. most recent) day.

The decreasing number is just that, a number to identify each day of the
savegame. The fact that the numbers in brackets in the example above are
identical to the days of month in the game is just coincidence.

To revert to a different day, first **close the game**. Then, use the ``revert``
command and supply the savegame (again, abbreviate as much as you like as long
as it's unique) and the number right next to the asterisk. For example, to go
back to the 7th of Spring, you'd type

    stardew-diary.exe revert c 7

(Again: **7** because that's the number in brackets above (``[07]``), not
because it's the day of month!)

After reverting, running ``log`` again shows this:

    stardew-diary.exe log c

      Concerned's Diary
      =================

      [11] [1085e24b9d] 11th of Spring, Year 1 (2143 G)
      [10] [674b39a802] 10th of Spring, Year 1 (272 G)
      [09] [0682111624] 9th of Spring, Year 1 (1816 G)
      [08] [c2cbc14269] 8th of Spring, Year 1 (1068 G)
    * [07] [206951a3e1] 7th of Spring, Year 1 (1023 G)
      [06] [faec669912] 6th of Spring, Year 1 (808 G)
      [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
      [04] [624d3e1724] 4th of Spring, Year 1 (1433 G)
      [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

You can now either start the game and play or revert at will to any other
version of your savegame.

**Important:** When you now play the game and save, you will "lose" all your
newer versions. In the example above where we went back to version 7, saving
would mean a *new* version 8 pops into existence and the old version 8 to 11
would become inaccessible. You've basically forked the timeline for your
savegame and Stardew Diary is -- for simplicity's sake -- only ever dealing with
the current timeline. So after saving, your log output would look like this:

    stardew-diary.exe log c

      Concerned's Diary
      =================

    * [08] [726d1ec98a] 8th of Spring, Year 1 (1507 G)
      [07] [206951a3e1] 7th of Spring, Year 1 (1023 G)
      [06] [faec669912] 6th of Spring, Year 1 (808 G)
      [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
      [04] [624d3e1724] 4th of Spring, Year 1 (1433 G)
      [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

No saved version is ever deleted, though: You can however still reach and restore
**any** version that has ever been recorded, thanks to Fossil. It does require
manual intervention and some doc reading on your end, though.

### Diary Mode

Because each backup contains all versions of a single savegame, SD can diff them
to each other and see the differences for each day. Because it knows where
certain stuff (like levels, buildings, pets, ...) are stored, it can simulate
a diary view for each savegame. Try it out by using the ``history`` command:

    stardew-diary.exe history c

      Concerned's Diary
      =================

      1st of Spring, Year 1

       - I've reached Foraging Level 1!
       - I've spent 300 wood to fix the bridge at the beach.
       - While digging around, I found a Lost Book :o

      2nd of Spring, Year 1

       - I met this nice old man by the beach. We talked about
         fishing for a while, before he gave me a fishing rod :) It's
         just a bamboo pole, but as time comes, I might find better
         rods...

      3rd of Spring, Year 1

       - I've reached Foraging Level 2!

      4th of Spring, Year 1

       ...nothing happened...

      5th of Spring, Year 1

       - I've adopted a dog and named it George!
       - I've reached Farming Level 1!
       - I've stumbled upon some old things and had to visit Gunther
         at the library to learn more about them. He encouraged me to
         search for more artifacts and minerals, so that I could
         donate them to the museum. So I went straight ahead and
         donated my Limestone and Rusty Spur.

      6th of Spring, Year 1

       - I've reached Mining Level 1!

      7th of Spring, Year 1

       - I've reached Combat Level 1!
       - While digging around, I found a Lost Book :o
       - Once again I've found ancient stuff and went to the museum
         to donate the Quartz, Earth Crystal, Nekoite, Topaz, Glass
         Shards and Slate.

      8th of Spring, Year 1

       ...nothing happened...

      9th of Spring, Year 1

       - I've reached Mining Level 2!
       - Once again I've found ancient stuff and went to the museum
         to donate the Amethyst, Calcite and Thunder Egg.

      10th of Spring, Year 1

       - I've reached Farming Level 2!
