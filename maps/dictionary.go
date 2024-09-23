package dictionary

import "errors"

// Dictionary stores words and their definitions.
type Dictionary map[string]string

// Seach returns the definition for a given word.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
