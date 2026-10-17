package fountain

import (
	"strings"
	"testing"
)

func elementsOfType(doc *Fountain, t int) []*Element {
	var out []*Element
	for _, e := range doc.Elements {
		if e.Type == t {
			out = append(out, e)
		}
	}
	return out
}

func TestSingleLineBoneyardIsBoneyardNotAction(t *testing.T) {
	src := "INT. KITCHEN - DAY\n\n/* story: pov=Anna */\n\nAnna enters.\n"
	doc, err := Parse([]byte(src))
	assertOK(t, err, "Parse")
	bs := elementsOfType(doc, BoneyardType)
	if len(bs) != 1 || bs[0].Content != "/* story: pov=Anna */" {
		t.Fatalf("boneyards = %+v, want one with the whole /* ... */ line", bs)
	}
	for _, a := range elementsOfType(doc, ActionType) {
		if strings.Contains(a.Content, "/*") {
			t.Errorf("action %q still holds boneyard text", a.Content)
		}
	}
	if as := elementsOfType(doc, ActionType); len(as) != 1 || as[0].Content != "Anna enters." {
		t.Errorf("actions = %+v, want only \"Anna enters.\"", as)
	}
}

func TestMultiLineBoneyardIsOneElement(t *testing.T) {
	src := "INT. KITCHEN - DAY\n\n/*\ncut line one\n\nANNA\ncut line two\n*/\n\nAnna enters.\n"
	doc, err := Parse([]byte(src))
	assertOK(t, err, "Parse")
	bs := elementsOfType(doc, BoneyardType)
	if len(bs) != 1 {
		t.Fatalf("boneyards = %d, want 1: %+v", len(bs), bs)
	}
	for _, want := range []string{"/*", "cut line one", "ANNA", "cut line two", "*/"} {
		if !strings.Contains(bs[0].Content, want) {
			t.Errorf("boneyard %q is missing %q", bs[0].Content, want)
		}
	}
	for _, e := range doc.Elements {
		if e.Type == CharacterType || e.Type == DialogueType {
			t.Errorf("%s %q came out of the boneyard", e.Name, e.Content)
		}
	}
	if as := elementsOfType(doc, ActionType); len(as) != 1 || as[0].Content != "Anna enters." {
		t.Errorf("actions = %+v, want only \"Anna enters.\" after the boneyard", as)
	}
}

func TestBoneyardText(t *testing.T) {
	for in, want := range map[string]string{
		"/* story: pov=Anna */": "story: pov=Anna",
		"/*\ncut\nlines\n*/":    "cut\nlines",
		"/**/":                  "",
		"not a boneyard":        "not a boneyard",
	} {
		if got := BoneyardText(in); got != want {
			t.Errorf("BoneyardText(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSceneHeadingSceneNumber(t *testing.T) {
	src := "INT. KITCHEN - DAY #12A#\n\nAction.\n\nEXT. YARD - NIGHT\n\nMore.\n"
	doc, err := Parse([]byte(src))
	assertOK(t, err, "Parse")
	hs := elementsOfType(doc, SceneHeadingType)
	if len(hs) != 2 {
		t.Fatalf("scene headings = %d, want 2", len(hs))
	}
	if hs[0].SceneNumber != "12A" {
		t.Errorf("SceneNumber = %q, want 12A", hs[0].SceneNumber)
	}
	if hs[0].Content != "INT. KITCHEN - DAY #12A#" {
		t.Errorf("Content = %q, want the heading unchanged", hs[0].Content)
	}
	if hs[1].SceneNumber != "" {
		t.Errorf("unnumbered heading SceneNumber = %q, want empty", hs[1].SceneNumber)
	}
}
