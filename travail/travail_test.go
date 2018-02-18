package travail_test

import (
	"testing"

	. "github.com/hajimehoshi/jsgen/travail"
)

func TestMap(t *testing.T) {
	m, err := Map("github.com/hajimehoshi/jsplayground/travail", "../travail")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", m)
}
