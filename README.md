# Name

Presenting names for English-language applications where a basic model of first and last name(s) combined is sufficient.

Based off of [basecamp/name_of_person](https://github.com/basecamp/name_of_person).

## Installation
```bash
go get -u github.com/duffn/name
```

## Examples
```go
// The only possible error here is not providing a first name.
n, err := name.New("Bob", "Smith")

n.Full() // Bob Smith
n.Familiar() // Bob S.
n.Abbreviated() // B. Smith
n.Sorted() // Smith, Bob
n.Possessive() // Bob Smith's
n.Initials() // BS
n.Mentionable() // bobs
```

## Tests
```bash
go test
```

## License
[MIT](https://opensource.org/licenses/MIT)
