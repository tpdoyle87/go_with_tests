package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Thomas"},
			[]string{"Thomas"},
		}, {
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Thomas", "Orlando"},
			[]string{"Thomas", "Orlando"},
		}, {
			"struct with non name field",
			struct {
				Name string
				Age  int
			}{"Thomas", 36},
			[]string{"Thomas"},
		}, {
			"nested fields",
			Person{
				"Thomas",
				Profile{36, "Orlando"},
			},
			[]string{"Thomas", "Orlando"},
		},
		{
			"Pointers passed in",
			&Person{
				"Thomas",
				Profile{36, "Orlando"},
			},
			[]string{"Thomas", "Orlando"},
		},
		{
			"Slice of Person",
			[]Profile{
				{33, "Dallas"},
				{36, "Orlando"},
			},
			[]string{"Dallas", "Orlando"},
		}, {
			"Arrays",
			[2]Profile{
				{33, "Dallas"},
				{36, "Orlando"},
			},
			[]string{"Dallas", "Orlando"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow": "Moo",
			"Dog": "Woof",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Moo")
		assertContains(t, got, "Woof")

	})

	t.Run("Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Dallas"}
			aChannel <- Profile{36, "Orlando"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Dallas", "Orlando"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Dallas"}, Profile{36, "Orlando"}
		}

		var got []string
		want := []string{"Dallas", "Orlando"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
