package fountain_test

import (
	"fmt"

	"github.com/rsdoiel/fountain"
)

// Parse reads a Fountain document. A scene heading carries its scene number
// and a boneyard is its own element.
func ExampleParse() {
	src := []byte("Title: The Long Fire\n\nINT. KITCHEN - DAY #12A#\n\n/* cut: Anna waits */\n\nAnna enters.\n")
	doc, err := fountain.Parse(src)
	if err != nil {
		panic(err)
	}
	for _, e := range doc.Elements {
		switch e.Type {
		case fountain.SceneHeadingType:
			fmt.Printf("scene %s: %s\n", e.SceneNumber, e.Content)
		case fountain.BoneyardType:
			fmt.Printf("boneyard: %s\n", fountain.BoneyardText(e.Content))
		case fountain.ActionType:
			fmt.Printf("action: %s\n", e.Content)
		}
	}
	// Output:
	// scene 12A: INT. KITCHEN - DAY #12A#
	// boneyard: cut: Anna waits
	// action: Anna enters.
}
