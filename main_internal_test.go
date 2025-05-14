package main

import "testing"

func TestGreet(t *testing.T) {

	type testCast struct {
		lang language
		want string
	}

	var tests = map[string]testCast{
		"English": {
			lang: "en",
			want: "Hello World",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Spanish": {
			lang: "es",
			want: `unsupported language : "es"`,
		},
		"Empty": {
			lang: "",
			want: `unsupported language : ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("excepted %q, got %q", tc.want, got)
			}
		})
	}
}
