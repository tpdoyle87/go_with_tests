package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Giving a dicitonary with words it should return the definitions of a word", func(t *testing.T) {
		dictionary := map[string]string{"grasshopper": "an cute little bug that hops along nicely."}
		got := Search(dictionary, "grasshopper")
		want := "an cute little bug that hops along nicely."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
