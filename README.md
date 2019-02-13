
# synoperms

This is basically a recursive chmod which allows you to set file and directory permissions separately,
and which skips special Synology directories (ones which have names starting with "@").

    Usage of synoperms:
    -dirs string
          mode to set for directories (default "0755")
    -files string
          mode to set for ordinary files (default "0644")
    -v    run in verbose mode

Example usage:

    synoperms /volume1/music

It also avoids unnecessary chmod calls. I wrote it because I got fed up with having to run two `find` operations that would touch every file and directory.

[WTFPL](http://www.wtfpl.net/).
