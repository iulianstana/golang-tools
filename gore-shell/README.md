# Gore shell

[![Source code](../src/watch.jpg)](https://github.com/motemen/gore)

Nice tool to have a shell kind in golang.

## Usage

    gore

After a prompt is shown, enter any Go expressions/statements or commands described below.

To quit the session, type `Ctrl-D`

## REPL Commands

Some functionality are provided as colon-commands:

    :import <package path>  Import package
    :print                  Show current source
    :write [<filename>]     Write out current source to file
    :doc <expr or pkg>      Show document (requires godoc)
    :help                   List commands
    :quit                   Quit the session
