// Package fountain supports encoding/decoding fountain formatted screenplays.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// # BSD 2-Clause License
//
// Copyright (c) 2019, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//   - Redistributions of source code must retain the above copyright notice, this
//     list of conditions and the following disclaimer.
//
//   - Redistributions in binary form must reproduce the above copyright notice,
//     this list of conditions and the following disclaimer in the documentation
//     and/or other materials provided with the distribution.
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
package fountain

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	// 3rd Party package
	"gopkg.in/yaml.v3"
)

const (
	//
	// Types used in ElementSettings and Paragraph elements
	//

	// GeneralTextType - not specific formatting, treat as plain text
	GeneralTextType = iota
	// EmptyType An empty line(s), block of whitepace, can occur on title page or script page(s)
	EmptyType
	// TitlePageType - something that only happens on the title page.
	// NOTE: Title page elents have a .Name value which is my best guess about the content.
	TitlePageType
	// SceneHeadingType - exists in the script pages, it is for scene headings
	SceneHeadingType
	// ActionType - designates an action block in the script page(s)
	ActionType
	// CharacterType - designates a CHARACTER heading for dialog
	CharacterType
	// DialogueType - holds a block of dialogue
	DialogueType
	// ParentheticalType - holds any parenthetical statement after CharacterType and before DialogueType
	ParentheticalType
	// TransitionType - scene transition instructions, these are minimal in most scripts now, e.g. FADE IN:, FADE TO BLACK:, THE END.
	TransitionType
	// ShotType - Goes in the screen heading line
	ShotType
	// LyricType - holds lyrics to be sung
	LyricType
	// NoteType - holds a script note
	NoteType
	// BoneyardType - blocks of cut material that haven't been removed
	BoneyardType
	// SectionType - is a markdown like title/heading, not normally display in output but used for navigation in text
	SectionType
	// SynopsisType - used as internal documentation, not normally displayed in output
	SynopsisType

	//
	// In line Styling, important in hinting for CSS generation, some text modifications in pretty printing
	//

	// UnderlineStyle - show underlined, in CSS "text-decoration:underline;"
	UnderlineStyle
	// ItalicStyle - show as italics (e.g. <i>, <em> in HTML)
	ItalicStyle
	// BoldStyle - show as boldface or strong (e.g. <b>, <strong> in HTML)
	BoldStyle
	// AllCapsStyle - will transform the text to uppercase, could also trigger CSS "text-transform: uppercase;"
	AllCapsStyle
	// Strikethrough - generates strikethrough in CSS
	Strikethrough

	// Alignments

	// CenterAlignment - center text or block
	CenterAlignment
	// LeftAlignment - left align text or block
	LeftAlignment
	// RightAlignment - right align text of block
	RightAlignment

	// PageFeed - inject a page feed or <hr> in HTML
	PageFeed
)

var (
	reSceneNo = regexp.MustCompile(`#*#$`)
	// MaxWidth used to set width for Fountain text output in String()
	MaxWidth = 64
	// AsHTMLPage if true generate the HTML header and footer blocks
	AsHTMLPage = false
	// InlineCSS sets behavior of including style elements with CSS in ToHTML()
	InlineCSS = false
	// LinkCSS sets behavior of including link element pointing to CSS file in ToHTML()
	LinkCSS = false
	// CSS holds the filename to use generating CSS links or reading
	// in a customized version of the CSS. Defaults to "fountain.css".
	CSS = "fountain.css"
	// ShowSection - preserve section markers in output (e.g. when pretty printing a working draft)
	ShowSection = false
	// ShowSynopsis - preserve synopsis in output (e.g. when pretty printing a working draft)
	ShowSynopsis = false
	// ShowNotes - preserve notes in output (e.g. when pretty printing a working draft)
	ShowNotes = false

	// Pretty Print - will pretty print for output (e.g. when turning into
	// JSON, use MarshalIndent)
	PrettyPrint = false
)

// Fountain is the document container. It is the type returned by Parse() and ParseFile()
//
//	screenplay, _ := ParseFile("screenplay.fountain")
//	fmt.Println(screenplay.String())
type Fountain struct {
	TitlePage []*Element
	Elements  []*Element
}

// Element holds the parsed token in either the title page of the document or
// scene list parts.
type Element struct {
	Type    int    `json:"type" yaml:"type"`
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Content string `json:"content" yaml:"content"`
}

func typeName(t int) string {
	switch t {
	case PageFeed:
		return "Page Feed"
	case GeneralTextType:
		return "General Text"
	case EmptyType:
		return "Empty"
	case TitlePageType:
		return "Title Page"
	case TransitionType:
		return "Transition"
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
	case SectionType:
		return "Section"
	case SynopsisType:
		return "Synopsis"
	}
	return ""
}

// TypeName returns the string describing the type of Fountain Element.
func (element *Element) TypeName() string {
	return typeName(element.Type)
}

// CharacterName takes an element of type Character and trims spaces,
// removes parenthetical (e.g. `(O.S.)`) and returns a string
// of the character name(s). NOTE: characters in the form of "JANE AND JOE"
// will be returned as "JANE AND JOE".
func CharacterName(element *Element) string {
	characters := []string{}
	if element.Type == CharacterType {
		content := strings.TrimSpace(element.Content)
		if !(strings.HasPrefix(content, `"`) && strings.HasSuffix(content, `"`)) {
			contentParts := strings.Split(element.Content, " ")
			for _, content := range contentParts {
				content = strings.TrimSpace(content)
				// If not a parenthetical or concatentation record as
				// character name.
				if !((content == "") || (strings.HasPrefix(content, "(") && strings.HasSuffix(content, ")"))) {
					// skip content
					if strings.HasSuffix(content, `'s`) {
						content = strings.TrimSuffix(content, `'s`)
					}
					if strings.HasSuffix(content, `'S`) {
						content = strings.TrimSuffix(content, `'S`)
					}
					characters = append(characters, content)
				}
			}
		}
	}
	if len(characters) > 1 && strings.Compare(strings.ToUpper(characters[len(characters)-1]), "VOICE") == 0 {
		// Drop the trailing "VOICE"
		characters = characters[0 : len(characters)-1]
	}
	return strings.Join(characters, " ")
}

// wordWrap will try to break line at a suitable place if they are equal or
// longer than width.
func wordWrap(line string, width int) string {
	if len(line) <= width {
		return line
	}
	buf := []string{}
	words := strings.Split(line, " ")
	l := 0
	for _, word := range words {
		if l+len(word) < width {
			if len(buf) > 0 {
				buf = append(buf, " ", word)
				l += len(word) + 1
			} else {
				buf = append(buf, word)
				l += len(word)
			}
		} else {
			buf = append(buf, "\n", word)
			l = 0
		}
	}
	return strings.Join(buf, "") + "\n"
}

// blockWrap will add left/right padding and wrap the text in the block
func blockWrap(line, padding string, width int) string {
	// NOTE: We need to adjust width to reflect padding on right
	width = width - len(padding)
	if len(padding)+len(line) <= width {
		return padding + line
	}
	buf := []string{}
	words := strings.Split(line, " ")
	l := 0
	for _, word := range words {
		if l+len(word) < width {
			if len(buf) > 0 {
				buf = append(buf, " ", word)
				l += len(word) + 1
			} else {
				buf = append(buf, padding, word)
				l += len(padding) + len(word)
			}
		} else {
			buf = append(buf, "\n", padding, word)
			l = len(padding) + len(word)
		}
	}
	return strings.Join(buf, "") + "\n"
}

// centerAlignText center align text given a line and width
func centerAlignText(line string, width int) string {
	if len(line) >= width {
		return line
	}
	padLength := (width - len(line)) / 2
	return strings.Repeat(" ", padLength) + line
}

func leftAlignText(line string, width int) string {
	src := []string{}
	if strings.Contains(line, "\n") == false {
		return strings.TrimSpace(line)
	}
	lines := strings.Split(line, "\n")
	for _, line := range lines {
		src = append(src, strings.TrimSpace(line))
	}
	return strings.Join(src, "\n")
}

func rightAlignText(line string, width int) string {
	src := []string{}
	if strings.Contains(line, "\n") == false {
		line = strings.TrimSpace(line)
		l := len(line)
		if l >= width {
			return line
		}
		return strings.Repeat(" ", width-l) + line
	}
	lines := strings.Split(line, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		l := len(line)
		if l >= width {
			src = append(src, line)
		} else {
			src = append(src, strings.Repeat(" ", width-l), line)
		}

	}
	return strings.Join(src, "\n")
}

// String() considers elem.Type and formatting output as string
func (element *Element) String() string {
	switch element.Type {
	case TitlePageType:
		return element.Name + ":" + element.Content
	case SceneHeadingType:
		return strings.ToUpper(strings.TrimSpace(element.Content))
	case ActionType:
		return wordWrap(element.Content, MaxWidth)
	case CharacterType:
		return strings.Repeat("    ", 4) + strings.ToUpper(strings.TrimSpace(element.Content))
	case ParentheticalType:
		return strings.Repeat("    ", 3) + strings.TrimSpace(element.Content)
	case DialogueType:
		return blockWrap(element.Content, strings.Repeat("    ", 2), MaxWidth)
	case TransitionType:
		s := strings.TrimSpace(element.Content)
		if strings.HasSuffix(s, ".") || strings.HasSuffix(s, "IN:") {
			return leftAlignText(s, MaxWidth)
		}
		if strings.HasPrefix(s, ">") && strings.HasSuffix(s, "<") {
			return centerAlignText(strings.ToUpper(element.Content), MaxWidth)
		}
		return rightAlignText(strings.ToUpper(element.Content), MaxWidth)
	case CenterAlignment:
		return centerAlignText(element.Content, MaxWidth)
	case LeftAlignment:
		return leftAlignText(element.Content, MaxWidth)
	case RightAlignment:
		return rightAlignText(element.Content, MaxWidth)
	case NoteType:
		if ShowNotes {
			return element.Content
		}
		return ""
	case SectionType:
		if ShowSection {
			return element.Content
		}
		return ""
	case SynopsisType:
		if ShowSynopsis {
			return element.Content
		}
		return ""
	case PageFeed:
		return "\f"
	default:
		return element.Content
	}
}

// createElement assembles an HTML element with provided classs and content
func createElement(elem string, classes []string, content string) string {
	if len(classes) > 0 {
		if elem == "hr" || elem == "p" {
			return fmt.Sprintf("<%s class=%q>\n", elem, strings.Join(classes, " "))

		}
		return fmt.Sprintf("<%s class=%q>%s</%s>\n", elem, strings.Join(classes, " "), content, elem)

	}
	if elem == "hr" || elem == "p" {
		return fmt.Sprintf("<%s>", elem)
	}
	return fmt.Sprintf("<%s>%s</%s>\n", elem, content, elem)
}

// ToHTML considers elem.Type and formatting output
func (element *Element) ToHTML() string {
	switch element.Type {
	case TitlePageType:
		switch strings.ToLower(element.Name) {
		case "title":
			return createElement("div", []string{"title"}, element.Content)
		case "author":
			return createElement("div", []string{"author"}, element.Content)
		case "draft date":
			return createElement("div", []string{"draft-date"}, element.Content)
		case "date":
			return createElement("div", []string{"draft-date"}, element.Content)
		case "copyright":
			return createElement("div", []string{"copyright"}, element.Content)
		case "contact":
			return createElement("div", []string{"contact"}, element.Content)
		default:
			return createElement("div", []string{"general-text"}, element.Content)
		}
	case SceneHeadingType:
		return createElement("div", []string{"scene-heading"}, strings.ToUpper(strings.TrimSpace(element.Content)))
	case ActionType:
		return createElement("div", []string{"action"}, element.Content)
	case CharacterType:
		return createElement("div", []string{"character"}, strings.ToUpper(strings.TrimSpace(element.Content)))
	case ParentheticalType:
		return createElement("div", []string{"parenthetical"}, strings.TrimSpace(element.Content))
	case DialogueType:
		return createElement("div", []string{"dialogue"}, element.Content)
	case TransitionType:
		s := strings.TrimSpace(element.Content)
		if strings.HasPrefix(s, ">") && strings.HasSuffix(s, "<") {
			return createElement("div", []string{"transition", "centered"}, strings.TrimPrefix(strings.TrimSuffix(s, "<"), ">"))
		}
		if strings.HasPrefix(s, ">") {
			return createElement("div", []string{"transition", "right-align"}, strings.TrimPrefix(s, ">"))
		}
		if strings.HasSuffix(s, ".") || strings.HasSuffix(s, "IN:") {
			return createElement("div", []string{"transition", "left-align"}, s)
		}
		return createElement("div", []string{"transition", "right-align"}, strings.ToUpper(element.Content))
	case CenterAlignment:
		return createElement("div", []string{"centered"}, element.Content)
	case LeftAlignment:
		return createElement("div", []string{"left-align"}, element.Content)
	case RightAlignment:
		return createElement("div", []string{"right-align"}, element.Content)
	case PageFeed:
		return createElement("hr", []string{"page-feed"}, "")
	default:
		return createElement("div", []string{strings.ToLower(strings.Replace(typeName(element.Type), " ", "-", -1))}, element.Content)
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
			switch elem.Type {
			case NoteType:
				if ShowNotes {
					src = append(src, elem.Content)
				}
			case SectionType:
				if ShowSection {
					src = append(src, elem.Content)
				}
			case SynopsisType:
				if ShowSynopsis {
					src = append(src, elem.Content)
				}
			default:
				s = elem.String()
				src = append(src, s)
			}
		}
	}
	return strings.Join(src, "\n")
}

// isTitlePage evaluates the current line to see if we're still in the
// title page element.
func isTitlePage(line string, prevType int) bool {
	if prevType == TitlePageType && isSceneHeading(line, prevType) == false && isTransition(line, prevType) == false {
		return true
	}
	return false
}

// isEmpty evaluates the current line to see if we're an "empty" line or still in title page.
func isEmpty(line string, prevType int) bool {
	if prevType == TitlePageType {
		return false
	}
	if len(strings.TrimSpace(line)) == 0 {
		return true
	}
	return false
}

// isCenterAlignment evaluates the current line to see if it is a Center Alignment type
func isCenterAlignment(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, ">") && strings.HasSuffix(line, "<") {
		return true
	}
	return false
}

// isSceneHeading evaluates a line and return true if it looks like a scene heading or false otherwise
func isSceneHeading(line string, prevType int) bool {
	line = strings.ToUpper(strings.TrimSpace(line))
	switch {
	case strings.HasPrefix(line, "!"):
		// This line must be action
		return false
	case reSceneNo.MatchString(line):
		return true
	case strings.HasPrefix(line, "."):
		return true
	case strings.HasPrefix(line, "EXT"):
		// We have line starting with EXT or EXT.
		return true
	case strings.HasPrefix(line, "INT"):
		// We have line starting with including INT., INT./EXT, INT/EXT
		return true
	case strings.HasPrefix(line, "I/E"):
		return true
	case strings.Contains(line, " -"):
		return true
	case strings.Compare(line, "FADE IN:") == 0:
		return true
	case strings.Compare(line, "THE END.") == 0:
		return true
	case strings.Compare(line, "THE END") == 0:
		return true
	case strings.Compare(line, "LA FIN.") == 0:
		return true
	case strings.Compare(line, "LA FIN") == 0:
		return true
	default:
		return false
	}
}

func isEndOfScript(element *Element) bool {
	if element.Type == SceneHeadingType {
		line := strings.ToUpper(strings.TrimSpace(element.Content))
		switch {
		case strings.Compare(line, "THE END.") == 0:
			return true
		case strings.Compare(line, "THE END") == 0:
			return true
		case strings.Compare(line, "LA FIN.") == 0:
			return true
		case strings.Compare(line, "LA FIN") == 0:
			return true
		}
	}
	return false
}

// isAction evaluates a line and returns true if it look like an action paragraph or false otherwise
func isAction(line string, prevType int) bool {
	// FIXME: isAction will have a empty element before and after, the
	// last non-empty element should be a schene heading or dialog
	if strings.HasPrefix(line, "!") {
		return true
	}
	if len(strings.TrimSpace(line)) == 0 {
		return false
	}
	if isSceneHeading(line, prevType) == false && isCharacter(line, prevType) == false && isDialogue(line, prevType) == false && isParenthetical(line, prevType) == false {
		return true
	}
	return false
}

// isCharacter evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
//
// FIXME: to really know that this is a character line we need
// to know the "next" element type, per definition at
// https://fountain.io/syntax#section-character which states next element
// cannot be an empty line.
func isCharacter(line string, prevType int) bool {
	if strings.HasPrefix(line, "@") {
		return true
	}
	if line == strings.ToUpper(line) && prevType == EmptyType && (isParenthetical(line, prevType) == false) {
		// NOTE: Per https://fountain.io/syntax#section-character
		// The next line should not be empty
		content := strings.ToUpper(strings.TrimSpace(line))
		// FIXME: Issue #2 show that I'm picking up non-character
		// elements as character elements. The upper case test is
		// not sufficient. The directives like `(O.S.)` should be
		// trimmed from the name when evaluating name.

		// If quotes are present then they are not a name.
		if strings.HasPrefix(content, `"`) && strings.HasSuffix(content, `"`) {
			return false
		}
		// Since we don't know if the next element is empty, we try to
		// infer from current element content.
		if strings.Contains(content, "--") ||
			strings.HasPrefix(content, "INT.") ||
			strings.HasPrefix(content, "EXT.") ||
			strings.HasSuffix(content, "ANGLE") ||
			strings.HasSuffix(content, "SHOT") ||
			strings.HasSuffix(content, "P.O.V.") ||
			strings.HasSuffix(content, ":") {
			return false
		}
		return true
	}
	return false
}

// isParenthetical evaluates a prevType and current line
// and returns true if it looks like a Character or false otherwise
func isParenthetical(line string, prevType int) bool {
	if strings.HasPrefix(line, "(") && strings.Contains(line, ")") {
		return true
	}
	return false
}

// isDialogue evaluates a prev, current and next lines and returns true
// if it looks like a Character or false otherwise
func isDialogue(line string, prevType int) bool {
	if strings.TrimSpace(line) == "" {
		return false
	}
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
	// NOTE: an explicit transition starts with a '>'
	if strings.HasPrefix(line, ">") == true {
		return true
	}
	if strings.HasSuffix(line, "TO:") || strings.HasSuffix(line, "IN:") {
		return true
	}
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "FADE TO") || strings.HasPrefix(line, ">") {
		return true
	}
	if strings.Contains(line, "THE END.") {
		return true
	}
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
	return false
}

// isBoneyardEnd evaluates if line is the end of a comment section
func isBoneyardEnd(line string, prevType int) bool {
	line = strings.TrimSpace(line)
	if prevType == BoneyardType && strings.HasSuffix(line, "*/") {
		return true
	}
	return false
}

// isPageFeed
func isPageFeed(line string, prevType int) bool {
	if strings.Compare(strings.TrimSpace(line), "===") == 0 {
		return true
	}
	return false
}

// isSection (not normally displayed in output)
func isSection(line string, prevType int) bool {
	if strings.HasPrefix(strings.TrimSpace(line), "#") {
		return true
	}
	return false
}

// isSynopsis (not normally displayed in output)
func isSynopsis(line string, prevType int) bool {
	if strings.HasPrefix(strings.TrimSpace(line), "=") {
		return true
	}
	return false
}

// getLineType evaluates the current line considering previous line type
// and returns the current line type.
func getLineType(line string, prevType int) int {
	switch {
	case isPageFeed(line, prevType):
		return PageFeed
	case isTitlePage(line, prevType):
		return TitlePageType
	case isSection(line, prevType):
		return SectionType
	case isSynopsis(line, prevType):
		return SynopsisType
	case isNote(line, prevType):
		return NoteType
	case isLyric(line, prevType):
		return LyricType
	case isSceneHeading(line, prevType):
		return SceneHeadingType
	case isAction(line, prevType):
		return ActionType
	case isTransition(line, prevType):
		return TransitionType
	case isCharacter(line, prevType):
		return CharacterType
	case isParenthetical(line, prevType):
		return ParentheticalType
	case isDialogue(line, prevType):
		return DialogueType
	case isBoneyard(line, prevType):
		return BoneyardType
	case isEmpty(line, prevType):
		return EmptyType
	case isCenterAlignment(line, prevType):
		return CenterAlignment
	default:
		return GeneralTextType
	}
}

// Parse takes []byte and returns a Fountain struct and error
func Parse(src []byte) (*Fountain, error) {
	prevType := TitlePageType
	key, value := "", ""
	document := new(Fountain)
	scanner := bufio.NewScanner(bytes.NewReader(src))
	foundEndOfScript := false
	for scanner.Scan() {
		line := scanner.Text()
		if !foundEndOfScript {
			currentType := getLineType(line, prevType)
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
						document.TitlePage = append(document.TitlePage, elem)
					} else {
						elem := document.TitlePage[i]
						elem.Content = elem.Content + "\n" + line
						document.TitlePage[i] = elem
					}
				}
			default:
				// If we haven't changed types we don't need to create
				// a new element.
				if prevType == currentType {
					i := len(document.Elements) - 1
					if i < 0 {
						i = 0
						elem := new(Element)
						elem.Type = currentType
						elem.Name = typeName(elem.Type)
						elem.Content = line
						document.Elements[i] = elem
					} else {
						elem := document.Elements[i]
						elem.Name = typeName(elem.Type)
						elem.Content = elem.Content + "\n" + line
						document.Elements[i] = elem
					}
				} else {
					element := new(Element)
					element.Type = currentType
					element.Name = typeName(element.Type)
					element.Content = line
					document.Elements = append(document.Elements, element)
					if element.Type == SceneHeadingType {
						foundEndOfScript = isEndOfScript(element)
					}
				}
			}
			prevType = currentType
		} else {
			element := new(Element)
			element.Type = GeneralTextType
			element.Name = typeName(element.Type)
			element.Content = line
			document.Elements = append(document.Elements, element)
		}
	}
	if err := scanner.Err(); err != nil {
		return document, err
	}
	// NOTE: Character name lines required look ahead.
	// I need to cleanup miss identified Character elements by
	// applying dialaog is next element rule.
	lastElement := len(document.Elements) - 1
	prevElementType := TitlePageType
	for i, element := range document.Elements {
		// Have we identified the character type correctly?
		if element.Type == CharacterType {
			if prevElementType == EmptyType {
				if i < lastElement {
					nextElementType := document.Elements[i+1].Type
					if !(nextElementType == DialogueType || nextElementType == ParentheticalType) {
						// What type are we?
						element.Type = GeneralTextType
					}
				}
				// NOTE: Character must be followed by dialog or
				// parenthetical but the last element has been identified
				// as a character element, what should this element be?
				// We may just have an imcomplete script.
			}
		}
		// If we're at the end of the script then we zero more characters.
		if element.Type == SceneHeadingType && isEndOfScript(element) {
			break
		}
		prevElementType = element.Type
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

// ToHTML converts a Fountain document based on the Options prvided.
// @param opt *Options a populate struct of options this package supports
// @return string of HTML
func (doc *Fountain) ToHTML() string {
	var err error
	out := []string{}
	// Handle Opening .AsHTMLPage
	src := ""
	if AsHTMLPage {
		if LinkCSS {
			src, err = getCSSLink()
			if err != nil {
				fmt.Fprintf(os.Stderr, "WARNING: %s\n", err)
			}
		}
		if InlineCSS {
			src, err = getCSS()
			if err != nil {
				fmt.Fprintf(os.Stderr, "WARNING: %s, using default CSS\n", err)
				// Fallback to default CSS after printing warning.
				src = createElement("style", []string{}, SourceCSS)
			}
		}
		if LinkCSS || InlineCSS {
			out = append(out, fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
%s
	</head>
	<body>
`, src))
		} else {
			out = append(out, `<!DOCTYPE html>
<html>
	<body>
	    <sectiom class="fountain">
`, src)
		}
	} else {
		if LinkCSS {
			src, err = getCSSLink()
			if err != nil {
				fmt.Fprintf(os.Stderr, "WARNING: %s\n", err)
			}
			out = append(out, src)
		}
		if InlineCSS {
			src, err = getCSS()
			if err != nil {
				log.Printf("%s", err)
			} else {
				out = append(out, src)
			}
		}
		out = append(out, fmt.Sprintf("<section class=%q>\n", "fountain"))
	}
	if doc.TitlePage != nil {
		out = append(out, `<section class="title-page">
`)
		for _, elem := range doc.TitlePage {
			out = append(out, elem.ToHTML())
		}
		out = append(out, `</section>
`)
	}
	if doc.Elements != nil {
		out = append(out, `<section class="script">
`)
		for _, elem := range doc.Elements {
			out = append(out, elem.ToHTML())
		}
		out = append(out, `</section>
`)
	}

	// Handle Closing .AsHTMLPage
	if AsHTMLPage {
		out = append(out, `
        </section>
	</body>
</html>
`)
	} else {
		out = append(out, fmt.Sprintf(`</section>`))
	}
	return strings.Join(out, "")
}

// ToJSON renders a Fountain type documents into a JSON
// serialized data structure.
func (doc *Fountain) ToJSON() ([]byte, error) {
	if PrettyPrint {
		return json.MarshalIndent(doc, "", "    ")
	}
	return json.Marshal(doc)
}

// ToYAML renders a Fountain type document into a YAML serialized data structure.
func (doc *Fountain) ToYAML() ([]byte, error) {
	src := []byte{}
	buf := bytes.NewBuffer(src)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err := encoder.Encode(doc)
	return src, err
}

// Run takes a byte split and returns an HTML fragment appropriate
// to use as a Scrippet with John Augusts' CSS
// https://fountain.io/_css/scrippets.css
func Run(input []byte) ([]byte, error) {
	var (
		out []byte
	)
	doc, err := Parse(input)
	if err != nil {
		out = append(out, input...)
	} else {
		out = append(out, []byte(doc.ToHTML())...)
	}
	return out, err
}
