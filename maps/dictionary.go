package dictionary

// Dictionary stores words and their definitions.
type Dictionary map[string]string

// Seach returns the definition for a given word.
func (d Dictionary) Search(word string) string {
	return d[word]
}
