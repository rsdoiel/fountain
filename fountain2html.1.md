%fountain2html(1) user manual | version 1.0.3 52955d8
% R. S. Doiel
% 2026-10-17

# NAME

fountain2html — read a Fountain document and write HTML

# SYNOPSIS

fountain2html [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

fountain2html reads a Fountain document and writes it as HTML, a fragment suitable for including like a scrippet or, with -page, a whole page.

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

-page
: output an HTML page, not a fragment

-inline-css
: add inline CSS

-link-css
: add a link to the CSS (the default is fountain.css)

-css FILENAME
: include a custom CSS file (default fountain.css)

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

Convert *screenplay.fountain* to *screenplay.html*.

    fountain2html -i screenplay.fountain -o screenplay.html

Or alternatively

    cat screenplay.fountain | fountain2html >screenplay.html

