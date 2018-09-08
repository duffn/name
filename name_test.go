package name

import (
	"testing"
)

var fullName, _ = New("Bob", "Smith")
var firstName, _ = New("Bob", "")

func TestMissingFirstName(t *testing.T) {
	_, err := New("", "Smith")
	expected := "firstName is required"
	actual := err.Error()
	if actual != expected {
		t.Errorf("Error was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestFull(t *testing.T) {
	expected := "Bob Smith"
	actual := fullName.Full()
	if actual != expected {
		t.Errorf("Full was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "Bob"
	actual = firstName.Full()
	if actual != expected {
		t.Errorf("Full was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestFamiliar(t *testing.T) {
	expected := "Bob S."
	actual := fullName.Familiar()
	if actual != expected {
		t.Errorf("Familiar was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "Bob"
	actual = firstName.Familiar()
	if actual != expected {
		t.Errorf("Familiar was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestAbbreviated(t *testing.T) {
	expected := "B. Smith"
	actual := fullName.Abbreviated()
	if actual != expected {
		t.Errorf("Abbreviated was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "Bob"
	actual = firstName.Abbreviated()
	if actual != expected {
		t.Errorf("Abbreviated was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestSorted(t *testing.T) {
	expected := "Smith, Bob"
	actual := fullName.Sorted()
	if actual != expected {
		t.Errorf("Sorted was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "Bob"
	actual = firstName.Sorted()
	if actual != expected {
		t.Errorf("Sorted was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestPossessive(t *testing.T) {
	expected := "Bob Smith's"
	actual := fullName.Possessive()
	if actual != expected {
		t.Errorf("Possessive was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "Bob's"
	actual = firstName.Possessive()
	if actual != expected {
		t.Errorf("Possessive was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestPossessiveWithS(t *testing.T) {
	n, _ := New("Bob", "Smits")
	expected := "Bob Smits'"
	actual := n.Possessive()
	if actual != expected {
		t.Errorf("Possessive with 's' was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestInitials(t *testing.T) {
	expected := "BS"
	actual := fullName.Initials()
	if actual != expected {
		t.Errorf("Initials was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "B"
	actual = firstName.Initials()
	if actual != expected {
		t.Errorf("Initials was incorrect, got %s, want: %s", actual, expected)
	}
}

func TestMentionable(t *testing.T) {
	expected := "bobs"
	actual := fullName.Mentionable()
	if actual != expected {
		t.Errorf("Mentionable was incorrect, got %s, want: %s", actual, expected)
	}

	expected = "bob"
	actual = firstName.Mentionable()
	if actual != expected {
		t.Errorf("Mentionable was incorrect, got %s, want: %s", actual, expected)
	}
}
