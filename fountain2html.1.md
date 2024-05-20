%fountain2html(1) | version 1.0.2 da5e106
% R. S. Doiel
% 2024-05-20

# NAME

fountain2html

# SYNOPSIS

fountain2html [OPTIONS] 

# DESCRIPTION

fountain2html is a command line program that reads an fountain document and writes out HTML.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i
: read from input file

-o
: read from output file

-newline
: add a trailing newline

-page
: If true output an HTML page otherwise an HTML fragement

-inline-css
: Add inline CSS

-link-css
: Add a link to CSS (default CSS is fountain.css)

-css
: Include a custom CSS file

-width
: set the width for the text


# EXAMPLES

Convert a *screenplay.fountain* to *screenplay.html*.

~~~
    fountain2html -i screenplay.foutnain -o screenplay.html
~~~

Or alternatively

~~~
    cat screenplay.fountain | fountain2html >screenplay.html
~~~


