package dictionaries

import (
	"testing"
)

func TestSearch(t *testing.T) {
	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got error %q, want %q", got, want)
		}
	}
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}
func TestAdd(t *testing.T) {
	assertDefinition := func(t *testing.T, dict Dictionary, word, definition string) {
		got, err := dict.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		if got != definition {
			t.Errorf("got %q, want %q", got, definition)
		}
	}
	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if got == nil {
			if want == nil {
				return
			}
			t.Fatal("expected an error.")
		}
	}
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is a test"
		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "do not adjust your television set")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
func TestUpdate(t *testing.T) {
	assertDefinition := func(t *testing.T, dict Dictionary, word, definition string) {
		got, err := dict.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		if got != definition {
			t.Errorf("got %q, want %q", got, definition)
		}
	}
	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if got == nil {
			if want == nil {
				return
			}
			t.Fatal("expected an error.")
		}
	}
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update nonexistent word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}
