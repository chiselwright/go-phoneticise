package phoneticise_test

import (
	"testing"

	"github.com/chiselwright/go-phoneticise"
)

func TestLookup(t *testing.T) {
	t.Parallel()

	want := "Alfa"
	got := phoneticise.Lookup('a')
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

	want = "Foxtrot"
	got = phoneticise.Lookup('f')
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPhoneticise(t *testing.T) {
	t.Parallel()

	var want string
	var got string

	want = "Charlie Foxtrot"
	got = phoneticise.Phoneticise("CF")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	want = "Charlie"
	got = phoneticise.Phoneticise("  C      ")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	want = "Charlie Hotel India Sierra Echo Lima"
	got = phoneticise.Phoneticise("Chisel")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	want = "Lima Mike One Eight (space) Charlie Tango Victor"
	got = phoneticise.Phoneticise("LM18 CTV")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	want = "Oscar November Echo (space) Tango Whiskey Oscar"
	got = phoneticise.Phoneticise("one two")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	want = "Hotel Echo Lima Lima Oscar (space) Whiskey Oscar Romeo Lima Delta {!}"
	got = phoneticise.Phoneticise("Hello World!")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	// 여보세요!
	want = "{여} {보} {세} {요}"
	got = phoneticise.Phoneticise("여보세요")
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}
