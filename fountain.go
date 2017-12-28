//
// fountain is a package encoding/decoding fountain formatted screenplays.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package fountain

import (
	"bufio"
	"bytes"
	"fmt" // DEBUG
	"io/ioutil"
	"strings"
)

const (
	Version = `v0.0.1-dev`

	// Types used in ElementSettings and Paragraph elements
	UnknownType = iota
	EmptyType
	TitlePageType
	SceneHeadingType
	ActionType
	CharacterType
	DialogueType
	ParentheticalType
	TransitionType
	ShotType
	LyricType
	NoteType
	BoneyardType

	// Style
	UnderlineStyle
	ItalicStyle
	BoldStyle
	AllCapsStyle
	Strikethrough

	// Alignments
	CenterAlignment
	LeftAlignment
	RightAlignment
)

func typeName(t int) string {
	switch t {
	case UnknownType:
		return "Unknown"
	case EmptyType:
		return "Empty"
	case TitlePageType:
		return "Title Page"
	case SceneHeadingType:
		return "Scene Heading"
	case ActionType:
		return "Action"
	case CharacterType:
		return "Character"
	case DialogueType:
		return "Dialogue"
	case ParentheticalType:
		return "Parenthetical"
	case TransitionType:
		return "Transition"
	case LyricType:
		return "Lyric"
	case NoteType:
		return "Note"
	case BoneyardType:
		return "Boneyard"
	case UnderlineStyle:
		return "Underline"
	case ItalicStyle:
		return "Italic"
	case BoldStyle:
		return "Bold"
	case AllCapsStyle:
		return "AllCaps"
	case Strikethrough:
		return "Strikethrough"
	case CenterAlignment:
		return "Center"
	case LeftAlignment:
		return "Left"
	case RightAlignment:
		return "Right"
	}
	return ""
}

type Fountain struct {
	TitlePage []*Element
	Elements  []*Element
}

type Element struct {
	Type    int    `json:"type"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content"`
}

// String returns an element as a string considering type
func (element *Element) String() string {
	switch element.Type {
	case TitlePageType:
		return element.Name + ":" + element.Content
	default:
		//FIXME: certain types should get a formatting treatment.
		return element.Content
	}
}

// String return a Fountain formatted document as a string
func (doc *Fountain) String() string {
	var s string
	src := []string{}
	if doc.TitlePage != nil {
		for _, elem := range doc.TitlePage {
			s = elem.String()
			src = append(src, s)
		}
		s = "\n"
		src = append(src, s)
	}
	if doc.Elements != nil {
		for _, elem := range doc.Elements {
			s = elem.String()
			src = append(src, s)
		}
		src = append(src, s)
	}
	return strings.Join(src, "\n")
}

// isTitlePage evaluates the current line to see if we're still in the
// title page element.
func isTitlePage(line string, prevType int) bool {
	if prevType == TitlePageType && isSceneHeading(line, prevType) == false {
		return true
	}
	return false
}

// isSceneHeading evaluates a line and return true if it looks like a scene heading or false otherwise
func isSceneHeading(line string, prevType int) bool {
	line = strings.ToUpper(line)
	switch {
	case strings.HasPrefix(line, "."):
		return true
	case strings.HasPrefix(line, "EXT"):
		return true
	case strings.HasPrefix(line, "INT"):
		return true
	case strings.HasPrefix(line, "INT./EXT"):
		return true
	case strings.HasPrefix(line, "INT/EXT"):
		return true
	case strings.HasPrefix(line, "I/E"):
		return true
	default:
		return false
	}
}

// isAction evaluates a line and returns true if it look like an action paragraph or false otherwise
func isAction(line string, prevType int) bool {
	if strings.HasPrefix(line, "!") {
		return true
	}
	if isSceneHeading(line, prevType) == false && isCharacter(line, prevType) == false && isDialogue(line, prevType) == false {
		return true
	}
	return false
}

// isCharacter evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
func isCharacter(line string, prevType int) bool {
	//FIXME: spec requires looking ahead an additional line
	// we only have a prevType known so assuming nextType == EmptyType
	nextType := EmptyType

	if strings.HasPrefix(line, "@") {
		return true
	}
	if line == strings.ToUpper(line) && prevType == EmptyType && nextType != EmptyType {
		return true
	}
	return false
}

// isParenthetical evaluates a prevType and current line
// and returns true if it looks like a Character or false otherwise
func isParenthetical(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "(") == false && strings.HasSuffix(line, ")") == false {
		return false
	}
	switch prevType {
	case CharacterType:
		return true
	case DialogueType:
		return true
	default:
		return false
	}
}

// isDialogue evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
func isDialogue(line string, prevType int) bool {
	switch prevType {
	case CharacterType:
		return true
	case ParentheticalType:
		return true
	default:
		return false
	}
}

// isTransition evaluates a line plus prev/next bool
func isTransition(line string, prevType int) bool {
	//FIXME: We only have one pass so we don't know what
	// the next type is, so hard coding it to assume it is EmptyType
	if strings.HasPrefix(line, ">") == true && strings.HasSuffix(line, "<") == false {
		return true
	}
	if line != strings.ToUpper(line) {
		return false
	}
	if strings.HasSuffix(line, "TO:") && prevType == EmptyType {
		return true
	}
	// FIXME: What about final transitions like "FADE TO BLACK."?
	return false
}

// isLyric evaluates a line to see if it is a lyric.
func isLyric(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "~") == true && strings.HasSuffix(line, "~") == false {
		return true
	}
	return false
}

// isNote evaluates a line if it is a note
func isNote(line string, prevType int) bool {
	//NOTE: a note can span multiple LF but not empty lines
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "[[") && strings.HasSuffix(line, "]]") {
		return true
	}
	if isNoteStart(line, prevType) || isNoteEnd(line, prevType) {
		return true
	}
	return false
}

// isNoteStart evaluates a line if it is a start of a multiline note
func isNoteStart(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if prevType != NoteType && strings.HasPrefix(line, "[[") {
		return true
	}
	return false
}

// isNoteEnd evalutes a line if it is the end of a multiline note
func isNoteEnd(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if prevType == NoteType && strings.HasSuffix(line, "]]") {
		return true
	}
	return false
}

// isBoneyard evaluates if a line is commented out
func isBoneyard(line string, prevType int) bool {
	//NOTE: A comment can span multiple LF like in Go/C/Java
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "/*") && strings.HasSuffix(line, "*/") {
		return true
	}
	if isBoneyardStart(line, prevType) || isBoneyardEnd(line, prevType) {
		return true
	}
	return false
}

// isBoneyardStart evaluates if line is a start of a comment section
func isBoneyardStart(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if prevType != BoneyardType && strings.HasPrefix(line, "/*") {
		return true
	}
	return true
}

// isBoneyardEnd evaluates if line is the end of a comment section
func isBoneyardEnd(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if prevType == BoneyardType && strings.HasSuffix(line, "*/") {
		return true
	}
	return true
}

// getLineType evaluates the current line considering previous line type
// and returns the current line type.
func getLineType(line string, prevType int) int {
	switch {
	case isTitlePage(line, prevType):
		return TitlePageType
	case isSceneHeading(line, prevType):
		return SceneHeadingType
	case isAction(line, prevType):
		return ActionType
	case isCharacter(line, prevType):
		return CharacterType
	case isParenthetical(line, prevType):
		return ParentheticalType
	case isDialogue(line, prevType):
		return DialogueType
	case isTransition(line, prevType):
		return TransitionType
	case isLyric(line, prevType):
		return LyricType
	case isNote(line, prevType):
		return NoteType
	case isBoneyard(line, prevType):
		return BoneyardType
	default:
		return UnknownType
	}
}

// Parse takes []byte and returns a Fountain struct and error
func Parse(src []byte) (*Fountain, error) {
	prevType := TitlePageType
	key, value := "", ""
	document := new(Fountain)
	scanner := bufio.NewScanner(bytes.NewReader(src))
	for scanner.Scan() {
		line := scanner.Text()
		currentType := getLineType(line, prevType)
		fmt.Printf("DEBUG %q, %q, %s\n", typeName(prevType), typeName(currentType), line)
		switch currentType {
		case TitlePageType:
			if strings.Contains(line, ":") {
				parts := strings.SplitN(line, ":", 2)
				key, value = parts[0], parts[1]
				elem := new(Element)
				elem.Type = TitlePageType
				elem.Name = key
				elem.Content = value
				document.TitlePage = append(document.TitlePage, elem)
			} else {
				i := len(document.TitlePage) - 1
				if i < 0 {
					i = 0
					elem := new(Element)
					elem.Type = TitlePageType
					elem.Name = "Unknown"
					elem.Content = line
					document.TitlePage[i] = elem
				} else {
					elem := document.TitlePage[i]
					elem.Content = elem.Content + "\n" + line
					document.TitlePage[i] = elem
				}
			}
		default:
			// If we haven't changed types we don't need to create a new element.
			if prevType == currentType {
				i := len(document.Elements) - 1
				if i < 0 {
					i = 0
					elem := new(Element)
					elem.Type = currentType
					elem.Content = line
					document.Elements[i] = elem
				} else {
					elem := document.Elements[i]
					elem.Content = elem.Content + "\n" + line
					document.Elements[i] = elem
				}
			} else {
				element := new(Element)
				element.Type = currentType
				element.Content = line
				document.Elements = append(document.Elements, element)
			}
		}
		prevType = currentType
	}
	if err := scanner.Err(); err != nil {
		return document, err
	}
	return document, nil
}

// ParseFile takes a filename and returns a Fountain struct and error
func ParseFile(fname string) (*Fountain, error) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return Parse(src)
}
