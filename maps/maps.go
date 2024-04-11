package maps

const (
	WordUnknownErr  = DictionaryErr("the word wasn't found in the dictionary")
	ErrWordExists   = DictionaryErr("The word already exists in the dictionary")
	ErrNoWordExists = DictionaryErr("The word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "nil", WordUnknownErr
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case WordUnknownErr:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case WordUnknownErr:
		return ErrNoWordExists
	case nil:
		d[word] = definition
	default:
		return nil
	}
	return nil
}
