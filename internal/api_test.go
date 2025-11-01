package internal

import (
	"testing"
	"time"
)

func TestFetchApi(t *testing.T) {
	cases := []struct {
		input    string
		expected User
	}{
		{
			input: "someuser",
			expected: User{
				URL:       "https://github.com/users/someuser",
				Username:  "someuser",
				Repos:     20,
				CreatedAt: time.Now(),
				Social:    "something",
			},
		},
	}

	for _, cas := range cases {
		if len(cas.input) != len(cas.expected.Username) {
			t.Errorf("More than one arguments in the terminal")
		}

		for i := range cas.input {
			if cas.input[i] != cas.expected.Username[i] {
				t.Errorf("%v does not equal %v", cas.input, cas.expected.Username)
			}
		}
	}
}
