# Changes

Reconstructed for v1.0.1 through v1.0.3 from the tags and the git history; maintained going forward.

## v1.0.4 — 2026-11-18

### Changed

- **The licence is the GNU Affero General Public License, version 3 or (at your option) any later version, throughout.**
  `LICENSE` and `codemeta.json` have said so since v1.0.2 (2024-05-20), but the headers of `fountain.go` and
  `fountain_test.go` still carried the BSD 2-Clause text of v1.0.1 and earlier, and no release note recorded the
  change. The headers now carry the AGPL notice, and `LicenseText` (shown by `-license`) is the AGPL text. v1.0.1 and
  the releases before it remain available under the BSD 2-Clause licence they were published with.
- **The three programs follow the workspace command line conventions.** Exit codes follow workspace DR-0003: 2 for a bad
  command line (an unknown option, a bad value, a surplus argument), 65 for input that is not a Fountain document, 66 for
  a missing input, 73 for an output that cannot be created, 74 for a read or write failure, 77 for a permission refusal.
  Before, every failure exited 1. An error is one line on standard error prefixed with the program name; `-quiet` drops
  the message and keeps the code. A script that tested for exit status 1 must test for non-zero.
- `-h`, `-l` and `-v` join `-help`, `-license` and `-version`; `-input`, `-output` and `-nl` are long and short spellings
  of `-i`, `-o` and `-newline`. The input and output may be given as the first and second argument; a third argument is
  an error (before, it was ignored). `-help`, `-license` and `-version` answer before any file is opened.
- Help text is Markdown in the man page structure with an EXIT STATUS section, kept in `internal/cli/helptext.go`; the
  command code moved to `internal/cli`. The `*.1.md` pages are generated from it by `make`.
- Go 1.27.3 is the minimum version. The Makefile gained `check-go-version` and `check-committed`; a release is built from a
  clean, pushed commit (workspace DR-0012); `release.bash` takes `--dry-run`; the armv7l target and the website's
  pagefind search are gone.

### Fixed

- **`fountain2html -o FILE` wrote the HTML to standard output and left FILE empty.** It now writes FILE.

The parser and the package API are unchanged from v1.0.3.

## v1.0.3 — 2026-10-17

- Boneyards (`/* ... */`, one line or several) are parsed as `Boneyard` elements; before, a one-line boneyard read as Action
  and the lines of a multi-line one as Character, Dialogue and Action.
- Scene headings carry their number (`INT. KITCHEN - DAY #12A#`) in the new `Element.SceneNumber`; the JSON output has
  `scene_number`. New `BoneyardText` returns the text between the delimiters.

## v1.0.2 — 2024-05-20

- Changed the licence from BSD 2-Clause to AGPL-3.0-or-later (`LICENSE`, `codemeta.json`).

## v1.0.1 — 2024-03-09

- BSD 2-Clause licence.
