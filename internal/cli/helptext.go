package cli

/** Fountain2htmlHelpText is the fountain2html(1) page, shown by `fountain2html -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by fountain.FmtHelp.
 * The Makefile writes the page to fountain2html.1.md.
 *
 * Example:
 *   fmt.Println(fountain.FmtHelp(cli.Fountain2htmlHelpText, "fountain2html", fountain.Version, fountain.ReleaseDate, fountain.ReleaseHash))
 */
const Fountain2htmlHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — read a Fountain document and write HTML

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a Fountain document and writes it as HTML, a fragment suitable for including like a scrippet or, with -page, a whole page.

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

    {app_name} -i screenplay.fountain -o screenplay.html

Or alternatively

    cat screenplay.fountain | {app_name} >screenplay.html
`

/** Fountain2jsonHelpText is the fountain2json(1) page, shown by `fountain2json -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by fountain.FmtHelp.
 * The Makefile writes the page to fountain2json.1.md.
 *
 * Example:
 *   fmt.Println(fountain.FmtHelp(cli.Fountain2jsonHelpText, "fountain2json", fountain.Version, fountain.ReleaseDate, fountain.ReleaseHash))
 */
const Fountain2jsonHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — read a Fountain document and write JSON

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a Fountain document and writes a JSON representation of it.

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

    {app_name} -i screenplay.fountain -o screenplay.json

Or alternatively

    cat screenplay.fountain | {app_name} >screenplay.json
`

/** FountainfmtHelpText is the fountainfmt(1) page, shown by `fountainfmt -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by fountain.FmtHelp.
 * The Makefile writes the page to fountainfmt.1.md.
 *
 * Example:
 *   fmt.Println(fountain.FmtHelp(cli.FountainfmtHelpText, "fountainfmt", fountain.Version, fountain.ReleaseDate, fountain.ReleaseHash))
 */
const FountainfmtHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — pretty print a Fountain document

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a Fountain document and pretty prints it.

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

    {app_name} -i screenplay.txt -o screenplay.fountain

Or alternatively

    cat screenplay.txt | {app_name} >screenplay.fountain
`
