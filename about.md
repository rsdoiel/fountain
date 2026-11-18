---
title: fountain
abstract: |-
  A library and command line program for working with Fountain script/screenplay formatted files.
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903


maintainer:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903

repository_code: https://github.com/rsdoiel/fountain
version: 1.0.4
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.27.3

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-11-18
---

About this software
===================

## fountain 1.0.4

Licence: the GNU Affero General Public License, version 3 or (at your option) any later version, throughout. LICENSE and codemeta.json have said so since v1.0.2, but the source file headers still carried the BSD 2-Clause text of v1.0.1 and earlier; they now carry the AGPL notice, and v1.0.1 and earlier stay available under BSD 2-Clause. The three programs follow the workspace command line conventions. Exit codes follow DR-0003 (2 bad command line, 65 not a Fountain document, 66 missing input, 73 output cannot be created, 74 I/O failure, 77 permission); errors are one line on standard error prefixed with the program name, and -quiet drops the message but not the code. -h, -l and -v join -help, -license and -version, with -input, -output and -nl as long spellings; the input and output may be given as the first and second argument; help answers before any file is opened. Help text is Markdown in the man page structure with an EXIT STATUS section, kept in internal/cli/helptext.go. Fixed: fountain2html -o created an empty file and wrote the HTML to standard output. The parser is unchanged. Go 1.27.3 is the minimum version, and the build and release tooling follows the workspace standard (a release is built from a clean, pushed commit).

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A library and command line program for working with Fountain script/screenplay formatted files.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/fountain)
  - [Issue Tracker](https://github.com/rsdoiel/fountain/issues)

## Programming languages

- Go >= 1.27.3




## Software Requirements

- Go >= 1.27.3


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


