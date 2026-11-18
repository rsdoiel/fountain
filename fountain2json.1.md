%fountain2json(1) user manual | version 1.0.4 7bcfa88
% R. S. Doiel
% 2026-11-18

# NAME

fountain2json — read a Fountain document and write JSON

# SYNOPSIS

fountain2json [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

fountain2json reads a Fountain document and writes a JSON representation of it.

The result has a TitlePage list and an Elements list. Each element has a "type" number, a "name" (for example
"Scene Heading", "Action", "Dialogue" or "Boneyard"), and its "content". A scene heading also has
"scene_number" when the heading ends in a number written as #12A#; the heading's content keeps the number.
A boneyard, a block of cut material between /* and */ on one line or several, is one "Boneyard" element.

# OPTIONS

-h, -help
: display this help page

-l, -license
: display the license

-v, -version
: display the version

-i, -input FILENAME
: read from FILENAME; the default, and -, is standard input

-o, -output FILENAME
: write to FILENAME; the default, and -, is standard output

-nl, -newline
: add a trailing newline (the default; -newline=false leaves it off)

-quiet
: suppress error messages; the exit status still says what happened

-width N
: set the width of the text (default 65)

-pretty
: indent the JSON

The input and output may also be given as the first and second argument.

# EXIT STATUS

0
: success

2
: bad command line: an unknown option, a bad option value, a surplus argument

65
: the input is not a valid Fountain document

66
: an input file does not exist

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

Render *screenplay.fountain* as *screenplay.json*.

    fountain2json -i screenplay.fountain -o screenplay.json

Or alternatively

    cat screenplay.fountain | fountain2json >screenplay.json

