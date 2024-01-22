package dict

import "errors"

// Dictionary
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errAlreadyExist = errors.New("Already Exist")

// Search for a word
func (dict Dictionary) Search(word string) (string, error) {
	value, exist := dict[word]

	if exist {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (dict Dictionary) Add(word string, definition string) error {
	_, exist := dict.Search(word)
	if exist == errNotFound {
		dict[word] = definition
	} else if exist == nil {
		return errAlreadyExist
	}
	return nil
}
