package nlp

import (
	"io/ioutil"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"
)

/*
var tokenizeCases = []struct {
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"", nil},
}
*/

// Exercise: Read test cases from tokenize_cases.toml
// Use github.com/BurntSushi/toml to read TOML

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	data, err := ioutil.ReadFile("tokenize_cases.toml")
	require.NoError(t, err)

	var testCases struct {
		Cases []tokenizeCase
	}

	err = toml.Unmarshal(data, &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	// for _, tc := range tokenizeCases {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	/* Before testify
	// if tokens != expected {// Can't compare slices with == in Go (only to nil)
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
	require.Equal(t, expected, tokens)
}
