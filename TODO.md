
# Action Items

## Bugs

+ [x] Transition follow by transition at end of script treats the second transition as action, see sample-05.fountain, noticed in HTML output
    + A: OK, the line begins with "!" so it should be "Action"
+ [ ] sample-05.fountain has the character "DOG" line identified as "dialogue"
+ [ ] sample-07.fountain (charade) has parenthiticals sprinkled in the dialogue, need to support this.  It meakes some dialog show up as Action in the JSON stream.


+ [ ] Transitions are almost never recognised: `CUT TO:`, `> CUT TO:` and `DISSOLVE TO:` between blank lines come out as Action (found 2026-11-18 while testing osf's `FromFountain`)
    + Cause: in the type switch (around fountain.go:795-810) `isAction` runs before `isTransition`, and `isAction` is true for any line that is not a heading, character, dialogue or parenthetical, so `isTransition` is only reached for lines `isAction` rejects. `FADE OUT.` comes out as General Text.
    + Expected (fountain.io syntax): an uppercase line ending in `TO:` with blank lines around it is a Transition; a line starting with `>` (and not ending in `<`) is a Transition; `> text <` is centered Action, not a Transition (`isTransition` calls both Transition).
    + Reproduce: `INT. A - DAY`, blank, `Hello there.`, blank, `CUT TO:`, blank, `EXT. B - DAY`. Test first, per workspace DR-0007.
+ [ ] `isSceneHeading` takes any line that starts with `INT`, `EXT` or `I/E`, after upper-casing, so ordinary action is read as a Scene Heading: `Interrupted by a knock, she looks up.` and `Extras fill the room.` both come out as Scene Heading (found 2026-11-18)
    + Expected: the prefix must be followed by `.`, ` ` or `/` (`INT.`, `EXT `, `INT./EXT.`, `I/E.`) and the line must follow a blank line.
+ [ ] A line that starts with `..` or `...` is read as a forced Scene Heading, because `isSceneHeading` accepts any leading `.`: `...and so on.` and `..and so on.` come out as Scene Heading (found 2026-11-18)
    + Expected: a forced heading is `.` followed by a letter; `..` and `...` are ordinary text (Action).
+ [ ] Forcing markers (`!`, `@`, `~`, `.`, `>` and `<`) stay in `Element.Content` (by design, like the scene number); consumers have to strip them. **osf** strips them in `FromFountain` since 0.0.10 (`stripForcedMarker`). Decide whether fountain should offer one function for this (like `BoneyardText`) so **fdx**, **osf** and **scripttool** share the rules.
+ [ ] Title page values keep the whitespace after the key and the indentation of multi-line values (`" TITLE"`, `"\n\tContact\n\tinformation\n\n"`); **osf** trims them in `FromFountain` since 0.0.10 (`titleLines`). Decide whether `Parse` should return clean values.

## Next (v0.1.0)

+ [ ] Check whether the boneyard and scene-number changes (added for the knowledge module's story support, 2026-10-17) have parallels in the other screenplay-format modules and in **scripttool**, which wraps them:
    + **fdx**, **osf** and **scripttool**: do their Fountain conversions drop, keep or mangle a `/* ... */` boneyard now that **fountain** types it as `BoneyardType`? Do they carry `Element.SceneNumber` (FDX `Number` attribute on a scene heading, OSF scene numbers, FadeIn) into and out of Fountain's `#12A#`?
    + Each of them pins an older **fountain** (scripttool v1.0.1, fdx v1.0.1); run their tests against the new release before bumping.
    + Checked 2026-10-17 with **scripttool** built against the local **fountain** 1.0.3: the existing tests of **fdx**, **osf** and **scripttool** pass with it. A Fountain → FDX → Fountain round trip keeps a boneyard, one line or several, as plain text (in FDX and OSF it becomes a visible paragraph, as it did when it was read as Action), and keeps a scene number only inside the heading text (`... #12A#`). **fdx** now writes the scene's FDX `Number` attribute from `Element.SceneNumber` and reads it back (2026-10-20, fdx `Number` attribute commit; unreleased). **osf** has no per-scene number (numbering is a document setting), so there is nothing to carry. Decided 2026-10-20: a boneyard stays a visible paragraph in FDX and OSF (lossless, round-trips as text, prints). A hidden note was considered and not done because it cannot be tested here: Final Draft shows a ScriptNote as a margin marker and leaves it out of print, but no sample FDX holds an anchored note to learn its XML from; OSF has a `<para note="...">` attribute (see `osf/testdata/OSF-2.0.xml`) whose display in Fade In is unverified, and a boneyard would need a rule for which paragraph carries it. Revisit if a Final Draft file with a ScriptNote or a Fade In install becomes available.
+ [ ] **fountain2html**: boneyard text is written into the page as `<div class="boneyard">`, so cut material is visible; decide whether it should be hidden by default, as Notes and Synopses are

+ [ ] Add test to check to validate parse structure of each `testdata/*.fountain` file
+ [ ] Add demo of using **fountain2json** for creating a script report
+ [ ] Improve **fountainfmt** pretty print options
+ [ ] handle unlabeled title pages
+ [x] Write and **fountain2html**
+ [ ] Add CSS dump option for **fountain2html** 
+ [ ] Add option to render [scrippets](https://fountain.io/scrippets) compatible HTML and CSS


## Someday, Maybe

+ [ ] Add support for Markdown front matter for additiona metadata processing
+ [ ] handle general text (outside of notes, boneyard)
    + this could be handled like front matter in Markdown
+ [ ] definable heading prefixes
+ [ ] reports on screenplay 
+ [ ] fountain2pdf

## Completed

+ [x] Write Parse(src []byte) returning a Fountain Struct and error
+ [x] word wrap for dialogue


## Modernization (2026-11-16)

To match the workspace standard (knowledge and opml are the models; Go 1.27.3 across repositories). Do one repository at a time; run the tests and `make -n release` after each step.

- [x] Go 1.27.3: change `go 1.26.4` in go.mod, the `Go >=` lines in codemeta.json, README.md and INSTALL.md; run `go mod tidy`, `go vet ./...` and the tests.
- [x] Makefile: add `check-go-version` (go.mod, codemeta.json and the Go in use agree); `check-release` calls it.
- [x] Drop the armv7l target and 32-bit Pi wording (dropped 2026-10-27).
- [x] release.ps1: take knowledge's (existing-draft check, checksums).
- [x] Port workspace DR-0012 (2026-11-18): `check-committed`, `release` without `generate`, `release.bash` with `--dry-run` that commits and pushes nothing, `makefile_test.go`.
- [x] CLI to the workspace practice (2026-11-18): `internal/cli` behind 12-line mains, `helptext.go` with the man page structure and EXIT STATUS, DR-0003 exit codes, `-h/-l/-v`, `cli_exit_test.go`. Fixed `fountain2html -o` writing to standard output. Not done: a `--json` error mode (DR-0011); fountain2json's output is JSON but it has no JSON error mode, as with opml2json.
