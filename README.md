[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

# fountain

_fountain_ is a Go package and three small command line programs for working with [Fountain](https://fountain.io) formatted screenplays, the plain text format many screenwriting programs read and write. The package parses a document into title page and script elements, pretty prints it, and renders it as HTML, JSON or YAML.

## The package

```go
doc, err := fountain.Parse(src)            // or fountain.ParseFile("script.fountain")
for _, e := range doc.Elements {
    switch e.Type {
    case fountain.SceneHeadingType:
        fmt.Println(e.SceneNumber, e.Content) // 12A INT. KITCHEN - DAY #12A#
    case fountain.BoneyardType:
        fmt.Println(fountain.BoneyardText(e.Content))
    }
}
fmt.Println(doc.String())                  // pretty printed Fountain
fmt.Println(doc.ToHTML())                  // HTML
data, _ := doc.ToJSON()                    // or doc.ToYAML()
```

A `Fountain` has a `TitlePage` and `Elements`; each `Element` has a `Type` (`SceneHeadingType`, `ActionType`, `CharacterType`, `DialogueType`, `ParentheticalType`, `TransitionType`, `NoteType`, `BoneyardType`, `SectionType`, `SynopsisType` and the rest), a `Name`, its `Content`, and, on a scene heading, the `SceneNumber`.

- **Scene numbers.** A heading that ends in a number written between hashes, `INT. KITCHEN - DAY #12A#`, has `SceneNumber` "12A". `Content` keeps the heading as written.
- **Boneyards.** Cut material between `/*` and `*/`, on one line or across several, is one `BoneyardType` element. `BoneyardText` returns the text between the delimiters.
- **Output options.** The package variables `MaxWidth`, `AsHTMLPage`, `InlineCSS`, `LinkCSS`, `CSS`, `ShowSection`, `ShowSynopsis` and `ShowNotes` control `String` and `ToHTML`.

The runnable example in `example_test.go` is the one above in miniature, and is checked by `go test`.

## The programs

[fountain2json](fountain2json.1.md)
: a Fountain to JSON converter, useful for building reports. The JSON has the element fields above, `scene_number` included.

[fountainfmt](fountainfmt.1.md)
: a Fountain pretty printer

[fountain2html](fountain2html.1.md)
: a Fountain to HTML converter

## Install

Get a release from [GitHub](https://github.com/rsdoiel/fountain/releases), or see [INSTALL.md](INSTALL.md), which also covers installing from source. To use the package in a Go program:

```shell
go get github.com/rsdoiel/fountain
```

## Build and test

You need Go 1.26.4 or later. Pandoc and GNU Make are only needed for the man pages and the website, and `cmt` (CMTools) only to regenerate `version.go`, `about.md` and `CITATION.cff` from `codemeta.json`.

```shell
make          # build the programs into bin/ and the man pages into man/
make test     # go test ./...
make check    # gofmt -l and go vet; rewrites nothing
```

`get-optional-testdata.bash` downloads the larger sample screenplays from fountain.io into `testdata/`; see [testdata.md](testdata.md). `make generate` runs `cmt`; `make release` builds the distribution zips, and `release.bash` then drafts the GitHub release.

## Release notes

**1.0.3.** Boneyards are parsed as `BoneyardType` elements. A one-line boneyard used to read as Action, and the lines of a multi-line one as Character, Dialogue and Action. Scene headings carry their number in `Element.SceneNumber`, the JSON output has `scene_number`, and `BoneyardText` is new. Go 1.26.4 or later is required. These were added for the [knowledge](https://github.com/rsdoiel/knowledge) module, which reads Fountain scenes and keeps story metadata in one-line boneyards.

## Known limitations

The parser is line based. The open items are in [TODO.md](TODO.md): a character line such as `DOG` in sample-05 is read as dialogue, and parentheticals scattered through dialogue in sample-07 can turn dialogue into Action. A boneyard is written into HTML output as a `boneyard` element rather than hidden.

## Related

- [fdx](https://github.com/rsdoiel/fdx), [osf](https://github.com/rsdoiel/osf) and [scripttool](https://github.com/rsdoiel/scripttool) convert between Fountain, Final Draft, Open Screenplay Format and FadeIn.
- [knowledge](https://github.com/rsdoiel/knowledge) ingests Fountain files as documents, one section per scene.

## Reference materials

- [fountain.io](https://fountain.io), the official place for all things Fountain
    - [Fountain Syntax](https://fountain.io/syntax)

_fountain_ is licensed under the GNU Affero General Public License, version 3; see [LICENSE](LICENSE).
