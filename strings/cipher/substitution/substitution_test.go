package substitution

import (
	"fmt"
	"testing"
)

func TestSubCipher(t *testing.T) {
	type test struct {
		inputMessage string
		inputKey     string
		want         string
	}
	tests := []test{
		{
			inputKey:     "The quick onyx goblin, grabbing his sword, jumps over the lazy dwarf!",
			inputMessage: "It was all a dream.",
			want:         "Od ptw txx t qsutg.",
		},
	}

	A := []byte("A")[0]
	// Z := []byte("Z")[0]
	// a := []byte("a")[0]
	z := []byte("z")[0]
	for i := A; i <= z; i++ {
		fmt.Printf("%s", string(i))
	}
	fmt.Println("")

	for _, tc := range tests {
		if got := sub(tc.inputMessage, tc.inputKey); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
