// fountain is a package encoding/decoding fountain formatted screenplays
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (C) 2017-2026 R. S. Doiel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package fountain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func assertOK(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("%s, %s", err, msg)
	}
}

func TestDialogue(t *testing.T) {
	src := []byte(`
INT. LAB - DAY

CHARLIE
Bring that to me.

(turns to Wilma)
And you fetch the tongs.

(Wilma turns and leaves the room)
`)

	expected := []int{
		SceneHeadingType,  // INT. LAB - DAY
		EmptyType,         //
		CharacterType,     // CHARLIE
		DialogueType,      // Bring that to me.
		EmptyType,         //
		ParentheticalType, // (turns to Wilma)
		DialogueType,      // And you fetch the tongs.
		EmptyType,         //
		ParentheticalType, // (Wilma turns and leaves the room)
	}

	doc, err := Parse(src)
	assertOK(t, err, "Parse(src)")

	if doc == nil || doc.Elements == nil {
		t.Errorf("Couldn't create doc.Elements from Parse()")
		t.FailNow()
	}
	for i := 0; i < len(doc.Elements); i++ {
		if doc.Elements[i].Type != expected[i] {
			if i > 0 {
				previousElement := doc.Elements[i-1]
				t.Errorf("expected %q, got %q for %q, prev %q for %q", typeName(expected[i]), doc.Elements[i].TypeName(), doc.Elements[i].Content, previousElement.TypeName(), previousElement.Content)
			} else {
				t.Errorf("expected %q, got %q for %q", typeName(expected[i]), typeName(doc.Elements[i].Type), doc.Elements[i].Content)
			}
			t.FailNow()
		}
	}

}

func TestTypes(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join("testdata", "sample-01.fountain"))
	assertOK(t, err, "ReadFile(testdata/sample-01.fountain)")

	doc, err := Parse(src)
	assertOK(t, err, "Parse(src)")

	/*
	   !FADE IN:

	   EXT. LIBRARY - DAY

	   A PROGRAMMER typing at an old laptop

	   PROGRAMMER
	   (excited)
	   Eureka!

	   > FADE TO BLACK.
	*/
	expected := []int{
		//NOTE: This "FADE IN:" should be action type because of the "!"
		ActionType,        // "!FADE IN:"
		EmptyType,         // ""
		SceneHeadingType,  // "EXT. LIBRARY - DAY"
		EmptyType,         // ""
		ActionType,        // "A PROGRAMMER typing at an old laptop"
		EmptyType,         // ""
		CharacterType,     // "PROGRAMMER"
		ParentheticalType, // "(excited)"
		DialogueType,      // "Eureka!"
		EmptyType,         // ""
		TransitionType,    // "> FADE TO BLACK."
		EmptyType,         // ""
	}

	if len(doc.Elements) > len(expected) {
		t.Errorf("Got more elements than expected, %+v", doc.Elements)
		for _, elem := range doc.Elements {
			fmt.Fprintf(os.Stderr, "%s %q\n", typeName(elem.Type), elem.Content)
		}
		t.FailNow()
	}
	for i := 0; i < len(doc.Elements); i++ {
		if doc.Elements[i].Type != expected[i] {
			t.Errorf("(%d) expected %q, got %q for %q", i, typeName(expected[i]), typeName(doc.Elements[i].Type), doc.Elements[i].Content)
			t.FailNow()
		}
	}

}

func TestSamples(t *testing.T) {
	files := []string{
		"sample-01.fountain",
		"sample-02.fountain",
		"sample-03.fountain",
		"sample-04.fountain",
		"sample-05.fountain",
		"sample-06.fountain",
		"sample-07.fountain",
	}

	for i, fName := range files {
		screenplay, err := ParseFile(path.Join("testdata", fName))
		if err != nil {
			t.Errorf("(%d) Should be able to read and parse %s, %s", i, fName, err)
			t.FailNow()
		}
		if screenplay.Elements == nil {
			t.Errorf("(%d) expected elements, got nil for %s", i, fName)
		}
		//FIXME: Check to see if the parse sequence is correct
	}
}

func isInList(elements []string, target string) bool {
	for _, element := range elements {
		if strings.Compare(element, target) == 0 {
			return true
		}
	}
	return false
}

func TestIssue2Scripttool(t *testing.T) {
	scriptFName := path.Join("testdata", "sample-07.fountain")
	charactersFName := path.Join("testdata", "sample-07-characters.txt")
	src, err := ioutil.ReadFile(charactersFName)
	if err != nil {
		t.Errorf("failed to read characters file %q, %s", charactersFName, err)
		t.FailNow()
	}
	screenplay, err := ParseFile(scriptFName)
	if err != nil {
		t.Errorf("failed to parse %q, %s", scriptFName, err)
		t.FailNow()
	}
	expectedNames := strings.Split(fmt.Sprintf("%s", src), "\n")
	lastElement := len(screenplay.Elements) - 1
	prevElementType := TitlePageType
	nextElementType := EmptyType
	for i, element := range screenplay.Elements {
		if element.Type == CharacterType {
			if i > 0 && prevElementType != EmptyType {
				t.Errorf("expected previous EmptyType, got %s -> %q", screenplay.Elements[i-1].TypeName(), screenplay.Elements[i-1].Content)
			}
			if i < lastElement {
				nextElementType = screenplay.Elements[i+1].Type
				if !(nextElementType == DialogueType || nextElementType == ParentheticalType) {
					t.Errorf("expected next type to be dialog or parentethical, got %s -> %q", screenplay.Elements[i+1].TypeName(), screenplay.Elements[i+1].Content)
				}
			}
			characterName := CharacterName(element)
			if !isInList(expectedNames, characterName) {
				t.Errorf("expected (%d) a name, got %q in %q type %q", i, characterName, element.Content, element.TypeName())
			}
		}
		prevElementType = element.Type
	}
}
