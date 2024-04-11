package maps

import (
	"testing"
)

func TestDictionarySearch(t *testing.T) {
	t.Run("Giving a dicitonary with words it should return the definitions of a word", func(t *testing.T) {
		d := Dictionary{"grasshopper": "an cute little bug that hops along nicely."}
		got, _ := d.Search("grasshopper")
		want := "an cute little bug that hops along nicely."
		assertStrings(t, got, want, d)
	})

	t.Run("Giving a word not in the dictionary", func(t *testing.T) {
		d := Dictionary{"work": "Something we have to do."}
		_, err := d.Search("money")
		want := WordUnknownErr
		assertError(t, err)
		assertStrings(t, err.Error(), want.Error(), d)
	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("Add a word to the dictionary", func(t *testing.T) {
		d := Dictionary{}
		_ = d.Add("Flower", "A plant that grows and blooms")
		want := "A plant that grows and blooms"

		assertDefinition(t, d, "Flower", want)
	})

	t.Run("returns an error if the word already exists", func(t *testing.T) {
		d := Dictionary{}
		_ = d.Add("test", "this is just a test")
		err := d.Add("test", "this is just a test 2")
		assertError(t, err)
		assertStrings(t, err.Error(), ErrWordExists.Error(), d)
		assertDefinition(t, d, "test", "this is just a test")
	})

}

func TestDictionaryUpdate(t *testing.T) {
	t.Run("Updates a dictionary value", func(t *testing.T) {
		d := Dictionary{"test": "this is just a test"}
		err := d.Update("test", "updated test value")
		if err != nil {
			t.Fatal("did not expect an erro")
		}
		want := "updated test value"
		assertDefinition(t, d, "test", want)
	})

	t.Run("returns an error if the value doesn't exist", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update("test", "this should error")
		assertError(t, err)
		assertStrings(t, err.Error(), ErrNoWordExists.Error(), d)
	})
}

func TestDictionaryDelete(t *testing.T) {
	t.Run("Removes a word from the dictionary", func(t *testing.T) {

	})
}

// Helpers

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find the word: ", err)
	}
	assertStrings(t, got, definition, d)
}

func assertStrings(t testing.TB, got, want string, d Dictionary) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q dictionary: %#v", got, want, d)
	}
}

func assertError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("expected to get an error.")
	}
}
