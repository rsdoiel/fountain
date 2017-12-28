//
// fountain is a package encoding/decoding fountain formatted screenplays
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
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func assertOK(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("%s, %s", err, msg)
	}
}

func TestTypes(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join("testdata", "sample-01.fountain"))
	assertOK(t, err, "ReadFile(testdata/sample-01.fountain)")

	doc, err := Parse(src)
	assertOK(t, err, "Parse(src)")

	expected := []int{
		TransitionType,
		SceneHeadingType,
		ActionType,
		CharacterType,
		ParentheticalType,
		DialogueType,
		TransitionType,
	}
	for i := 0; i < len(doc.Elements); i++ {
		if doc.Elements[i].Type != expected[i] {
			t.Errorf("expected %q, got %q for %q", typeName(expected[i]), typeName(doc.Elements[i].Type), doc.Elements[i].Content)
		}
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
