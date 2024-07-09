%fountainfmt(1) | version 1.0.2 25e4291
% R. S. Doiel
% 2024-07-09

# NAME

fountainfmt

# SYNOPSIS

fountainfmt [OPTIONS]

# DESCRIPTION

fountainfmt is a command line program that reads an fountain document and pretty prints it.

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
: set text width

-debug
: display type and element content

-section
: include sections in output

-synopsis
: include synopsis in output

-notes
: include notes in output


# EXAMPLES

Pretty print *screenplay.txt* saving it as *screenplay.fountain*.

~~~
fountainfmt -i screenplay.txt -o screenplay.fountain
~~~

Or alternatively

~~~
cat screenplay.txt | fountainfmt > screenplay.fountain
~~~


