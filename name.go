package name

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// Name is a struct of a first and last name.
type Name struct {
	FirstName string
	LastName  string
}

// New returns a new Name struct or an error if firstName is empty.
func New(firstName string, lastName string) (*Name, error) {
	if firstName == "" {
		return nil, errors.New("firstName is required")
	}

	return &Name{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

// Full is the full name of a person, e.g. "John Smith".
func (name *Name) Full() string {
	if name.hasLast() {
		return fmt.Sprintf("%s %s", name.FirstName, name.LastName)
	}
	return name.FirstName
}

// Familiar is the first name plus last initial, e.g. "John S.".
func (name *Name) Familiar() string {
	if name.hasLast() {
		return fmt.Sprintf("%s %c.", name.FirstName, []rune(name.LastName)[0])
	}
	return name.FirstName
}

// Abbreviated is the initial of the first name plus the last name, e.g. "J. Smith".
func (name *Name) Abbreviated() string {
	if name.hasLast() {
		return fmt.Sprintf("%c. %s", []rune(name.FirstName)[0], name.LastName)
	}
	return name.FirstName
}

// Sorted sorts the name for sorting, e.g. "Smith, John"
func (name *Name) Sorted() string {
	if name.hasLast() {
		return fmt.Sprintf("%s, %s", name.LastName, name.FirstName)
	}
	return name.FirstName
}

// Possessive adds "'s" or "'" to the end of a full name, e.g. "John Smith's".
func (name *Name) Possessive() string {
	n := name.Full()
	if strings.HasSuffix(n, "s") {
		return fmt.Sprintf("%s'", n)
	}
	return fmt.Sprintf("%s's", n)
}

// Initials provides a name's initials, e.g. "JS".
func (name *Name) Initials() string {
	var b bytes.Buffer
	for _, s := range strings.Split(name.Full(), " ") {
		b.WriteRune([]rune(s)[0])
	}
	return b.String()
}

// Mentionable creates a mentionable name, e.g. "johns".
func (name *Name) Mentionable() string {
	return strings.ToLower(strings.Replace(strings.TrimSuffix(name.Familiar(), "."), " ", "", -1))
}

// Check if the last name of a person is empty.
func (name *Name) hasLast() bool {
	if name.LastName == "" {
		return false
	}
	return true
}
