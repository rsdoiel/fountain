%fountain2json(1) | version 1.0.2 da5e106
% R. S. Doiel
% 2024-05-20

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


