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
version: 1.0.3
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.22

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-10-17
---

About this software
===================

## fountain 1.0.3

Boneyards (/* ... */, one line or several) are parsed as Boneyard elements, where a one-line boneyard used to read as Action and the lines of a multi-line one as Character, Dialogue and Action. Scene headings carry their number (INT. KITCHEN - DAY #12A#) in the new Element.SceneNumber, and the JSON output has scene_number. New BoneyardText returns the text between the delimiters. Added for the knowledge module's story support.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A library and command line program for working with Fountain script/screenplay formatted files.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/fountain)
  - [Issue Tracker](https://github.com/rsdoiel/fountain/issues)

## Programming languages

- Go >= 1.22




## Software Requirements

- Go >= 1.22


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


