%fountain2json(1) | version 1.0.3 52955d8
% R. S. Doiel
% 2026-10-17

# NAME

fountain2json

# SYNOPSIS

fountain2json [OPTIONS]

# DESCRIPTION

fountain2json is a command line program that reads an fountain document and returns a JSON representation of it.

The result has a TitlePage list and an Elements list. Each element has a "type" number, a "name" (for example
"Scene Heading", "Action", "Dialogue" or "Boneyard"), and its "content". A scene heading also has
"scene_number" when the heading ends in a number written as #12A#; the heading's content keeps the number.
A boneyard, a block of cut material between /* and */ on one line or several, is one "Boneyard" element.

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


