%fountain2json(1) | version 1.0.0 f02259d
% R. S. Doiel
% 2023-05-19

# NAME

fountain2json

# SYNOPSIS

fountain2json [OPTIONS]

# DESCRIPTION

fountain2json is a command line program that reads an fountain document and returns a JSON representation of it.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i
: read from filename

-o
: write to filename

-newline
: add a trailing newline

-width
: set the width of the text

-pretty
: pretty print the output


# EXAMPLES

Render *screenplay.fountain* as *screenplay.json*.

~~~
fountain2json -i screenplay.fountain -o screenplay.json
~~~

Or alternatively

~~~
    cat screenplay.fountain | fountain2json > screenplay.json
~~~

