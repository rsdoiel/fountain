% fountainfmt(1) fountainfmt user manual
% R. S. Doiel
% August 7, 2022

# NAME

fountainfmt

# SYNOPSIS

fountainfmt [OPTIONS]

# DESCRIPTION

fountainfmt is a command line program that reads an fountain document and pretty prints it.


# OPTIONS

Below are a set of options available.

-debug
: display type and element content

-h, -help
: display help

-i, -input
: set the input filename

-l, -license
: display license

-nl, -newline
: add a trailing newline

-o, -output
: set the output filename

-quiet
: suppress error messages

-v, -version
: display version

-w, -width
: set the width for the text


# EXAMPLES

Pretty print *screenplay.txt* saving it as *screenplay.fountain*.

~~~shell
    fountainfmt -i screenplay.txt -o screenplay.fountain
~~~

Or alternatively

~~~shell
    cat screenplay.txt | foutnainfmt > screenplay.fountain
~~~


