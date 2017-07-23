# Stardew Diary

This project aims to build a simple backup tool for [Stardew Valley](http://stardewvalley.net/)
savegames (currently limited/only tested on Windows). The backup mechanism makes
use of [Fossil](http://fossil-scm.org/).

* [Installation](#installation)
* [Usage](#usage)
  * [Finding Savegames](#finding-savegames)
  * [Creating Backups](#creating-backups)
  * [Checking Backups](#checking-backups)
  * [Reverting to a Previous Backup](#reverting-to-a-previous-backup)
    * [Understanding What Goes On](#understanding-what-goes-on)
  * [Diary Mode](#diary-mode)
  * [Restoring a Deleted Savegame](#restoring-a-deleted-savegame)

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

### Finding Savegames

Run the ``savegames`` command to get a list of all known savegames. These
include your current savegames as well as all backups of previous ones.

    > stardew-diary.exe savegames

      Savegames
      =========

       - Concerned on the "Star Trek Farm" (156295541) [new]
       - Schotty on the "HSV Farm" (142210113)
       - test on the "test Farm" (153507091) [dead]

      New savegames will be processed when the `match` command is next used.
      Dead savegames can be restored by running:

        stardew-diary.exe resurrect 153507091

Savegames can be new (have never been backed up), dead (deleted in the game
but exist as backups) or normal (exists in the game and has backups).

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

    > stardew-diary.exe log c

      Concerned's Diary
      =================

    * [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
      [04] [624d3e1724] 4th of Spring, Year 1 (1433 G)
      [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

To get the **Testerwoman** account, you would need to specify at least **testerw**

    stardew-diary.exe log testerw

SD will tell you when it cannot uniquely identify the savegame:

    > stardew-diary.exe log t
    2017/06/25 19:34:33 The savegame name 't' is ambiguous, could mean any of [Testerman_12345, Testerwoman_98765].

### Reverting to a Previous Backup

When you look at the output of the ``log`` command

    > stardew-diary.exe log c

      Concerned's Diary
      =================

    * [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
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
back to the 3rd of Spring, you'd type

    stardew-diary.exe revert c 3

(Again: **3** because that's the number in brackets above (``[03]``), not
because it's the day of month!)

After reverting, running ``log`` again shows this:

    > stardew-diary.exe log c

      Concerned's Diary
      =================

      [05] [0e3c12c127] 5th of Spring, Year 1 (383 G)
      [04] [624d3e1724] 4th of Spring, Year 1 (1433 G)
    * [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

You can move the asterisk to any day you like, but before you start the game
again and play, make sure you understand what happens when you save. Read the
following chapter carefully.

#### Understanding What Goes On

Think of your progression through the game as a graph, a series of savestates:

![graph 1](http://i.imgur.com/sy1w9Iu.png)

The last day has a thicker border because that's the current day of your
savegame. In the ``log`` command, it's marked with an asterisk. When you reverted
your savegame to day 3, you basically did this:

![graph 2](http://i.imgur.com/RwXCp1r.png)

You moved the marker, but all 5 versions of your game are still visible to
Stardew Diary and you can move freely between all of them.

When you now start the game, load your savegame and then *save*, this is what
happens:

![graph 3](http://i.imgur.com/kFp5BUz.png)

You just split the timeline of your savegame into two branches. The old timeline,
now drawn in grey, is no longer visible to Stardew Diary, **but it still exists
in your backups**. For simplicity's sake, SD is only ever concerned (pun intended)
with your current timeline. When you now run the ``log`` command, you will see
that day 5 vanished and day 4 is the new day 4 (notice the different amounts of
gold and the different random gibberish at the beginning of the line):

    > stardew-diary.exe log c

      Concerned's Diary
      =================

    * [04] [81db1ac8d0] 4th of Spring, Year 1 (2948 G)
      [03] [2ad2ffbd0e] 3rd of Spring, Year 1 (1054 G)
      [02] [0636f2a0b5] 2nd of Spring, Year 1 (674 G)
      [01] [208b80743e] 1st of Spring, Year 1 (500 G)

As you continue to play and save, your new timeline grows as expected:

![graph 4](http://i.imgur.com/INpVupB.png)

You can however still reach and restore **any** version that has ever been
recorded, thanks to Fossil. It does require manual intervention and some doc
reading on your end, though.

### Diary Mode

Because each backup contains all versions of a single savegame, SD can diff them
to each other and see the differences for each day. Because it knows where
certain stuff (like levels, buildings, pets, ...) are stored, it can simulate
a diary view for each savegame. Try it out by using the ``history`` command:

    > stardew-diary.exe history c

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

### Restoring a Deleted Savegame

When a savegame is deleted within the game, the backup file remains intact. You
can get a list of all known savegames by running the ``savegames`` command:

    > stardew-diary.exe savegames

      Savegames
      =========

       - Concerned on the "Star Trek Farm" (156295541)
       - Schotty on the "HSV Farm" (142210113)
       - test on the "test Farm" (153507091) [dead]

      Dead savegames can be restored by running:

        stardew-diary.exe resurrect 153507091

As you can see, we have a savegame that was deleted and can be restored. Do as
SD tells you:

    > stardew-diary.exe resurrect 153507091
    The savegame test_153507091 has been successfully restored.
