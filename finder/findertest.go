package finder

import (
	"testing"
)

func TestIsItGopher(t *testing.T) {
	word := "gopher"
	foo, err := IsItGopher(word)
	if err != nil {
		t.Fatal(err)
	}
	if (!foo) {
		t.Fail()
	}
}
