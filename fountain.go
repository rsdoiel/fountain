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
	"encoding/xml"
	"strings"
)

const (
	Version = `v0.0.0-dev`

	// Types used in ElementSettings and Paragraph elements
	EmptyLineType = iota
	GeneralType
	SceneHeadingType
	ActionType
	CharacterType
	DialogueType
	ParentheticalType
	TransitionType
	CastListType
	ShotType
	SingingType
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

type Fountain struct {
	TitlePage map[string]string
	Elements  []*Element
}

type Elements struct {
	Type    int    `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

// isSceneHeading evaluates a line and return true if it looks like a scene heading or false otherwise
func isSceneHeading(line string) bool {
	switch strings.ToUpper(line) {
	case strings.HasPrefix(line, "."):
		return true
	case strings.HasPrefix(line, "EXT"):
		return true
	case strings.HasPrefix(line, "INT"):
		return true
	case strings.HasPrefix(line, "INT./EXT"):
		return true
	case strings.HasPrefix(lines, "INT/EXT"):
		return true
	case strings.HasPrefix(lines, "I/E"):
	default:
		return false
	}
}

// isAction evaluates a line and returns true if it look like an action paragraph or false otherwise
func isAction(line string) bool {
	if strings.HasPrefix(line, "!") || (isSceneHeading(line) == false && isCharacter(line) == false && isDialogue(line) == false) {
		return true
	}
	return false
}

// isCharacter evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
func isCharacter(line string, prevLineType int, nextLineType int) bool {
	if strings.HasPrefix(current, "@") == true ||
		(prevLineEmpty == true && nextLine == false && (current == strings.ToUpper(current))) {
		true
	}
	return false
}

// isParenthetical evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
func isParenthetical(current string, prevWasCharacter, prevDialogue bool) bool {
	if strings.HasPrefix(current, "(") && strings.HasSuffix(current, ")") {
		return true
	}
	return false
}

// isDialogue evaluates a prev, current and next lines and returns true if it looks like a Character or false otherwise
func isDialogue(prevWasCharacter, prevParenthetical bool) bool {
	return (prevWasCharacter == true || prevParenthetical == true)
}

// isTransition evaluates a line plus prev/next bool
func isTransition(line string, prevLineEmpty, nextLineEmpty bool) bool {
	if strings.HasPrefix(line, ">") || (line == strings.ToUpper(line) && strings.HasSuffix(line, "TO:") && (prevLineEmpty || nextLineEmpty)) {
		return true
	}
	return false
}

// isCentered evaluates a line plus prev/next bool
func isCentered(line) bool {
	if strings.HasPrefix(line, ">") && strings.HasSuffix(line, "<") {
		return true
	}
	return false
}

// Parse takes []byte and returns a FinalDraft struct and error
func Parse(src []byte) (*FinalDraft, error) {
	var s scanner.Scanner

	document := new(Fountain)
	scanner = bufio.NewReader(src)
	for scanner.Scan() {
		// Read a line and decide what it is...
		line := scanner.Text()
		fmt.Println("DEBUG scan line:", line)
	}
	if err := scanner.Err(); err != nil {
		return document, err
	}
	return document, nil
}

// ParseFile takes a filename and returns a FinalDraft struct and error
func ParseFile(fname string) (*FinalDraft, error) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return Parse(src)
}
