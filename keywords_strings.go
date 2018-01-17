package jsonschema

import (
	"encoding/json"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// MaxLength MUST be a non-negative integer.
// A string instance is valid against this keyword if its length is less than, or equal to, the value of this keyword.
// The length of a string instance is defined as the number of its characters as defined by RFC 7159 [RFC7159].
type MaxLength int

func newMaxLength() Validator {
	return new(MaxLength)
}

// Validate implements the Validator interface for MaxLength
func (m MaxLength) Validate(data interface{}) error {
	if str, ok := data.(string); ok {
		if utf8.RuneCountInString(str) > int(m) {
			return fmt.Errorf("max length of %d characters exceeded: %s", m, str)
		}
	}
	return nil
}

// MinLength MUST be a non-negative integer.
// A string instance is valid against this keyword if its length is greater than, or equal to, the value of this keyword.
// The length of a string instance is defined as the number of its characters as defined by RFC 7159 [RFC7159].
// Omitting this keyword has the same behavior as a value of 0.
type MinLength int

func newMinLength() Validator {
	return new(MinLength)
}

// Validate implements the Validator interface for MinLength
func (m MinLength) Validate(data interface{}) error {
	if str, ok := data.(string); ok {
		if utf8.RuneCountInString(str) < int(m) {
			return fmt.Errorf("min length of %d characters required: %s", m, str)
		}
	}
	return nil
}

// Pattern MUST be a string. This string SHOULD be a valid regular expression,
// according to the ECMA 262 regular expression dialect.
// A string instance is considered valid if the regular expression matches the instance successfully.
// Recall: regular expressions are not implicitly anchored.
type Pattern regexp.Regexp

func newPattern() Validator {
	return &Pattern{}
}

// Validate implements the Validator interface for Pattern
func (p Pattern) Validate(data interface{}) error {
	re := regexp.Regexp(p)
	if str, ok := data.(string); ok {
		if !re.Match([]byte(str)) {
			return fmt.Errorf("regext pattrn %s mismatch on string: %s", re.String(), str)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Pattern
func (p *Pattern) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	ptn, err := regexp.Compile(str)
	if err != nil {
		return err
	}

	*p = Pattern(*ptn)
	return nil
}
