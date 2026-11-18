%fountainfmt(1) user manual | version 1.0.3 52955d8
% R. S. Doiel
% 2026-10-17

# NAME

fountainfmt — pretty print a Fountain document

# SYNOPSIS

fountainfmt [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

fountainfmt reads a Fountain document and pretty prints it.

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

-debug
: display each element's type and content instead of the screenplay

-section
: include sections in the output

-synopsis
: include synopses in the output

-notes
: include notes in the output

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

Pretty print *screenplay.txt*, saving it as *screenplay.fountain*.

    fountainfmt -i screenplay.txt -o screenplay.fountain

Or alternatively

    cat screenplay.txt | fountainfmt >screenplay.fountain

