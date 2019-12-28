package common

import (
	"reflect"
	"testing"
)

func TestFindContiguousHistory(t *testing.T) {
	type test struct {
		name   string
		inputA []string
		inputB []string
		want   []string
	}
	user0 := []string{"/start", "/pink", "/register", "/orange", "/red", "a"}
	user1 := []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"}
	user2 := []string{"a", "/one", "/two"}
	user3 := []string{"/red", "/orange", "/yellow", "/green", "/blue", "/purple", "/white", "/amber", "/HotRodPink", "/BritishRacingGreen"}
	user4 := []string{"/red", "/orange", "/amber", "/random", "/green", "/blue", "/purple", "/white", "/lavender", "/HotRodPink", "/BritishRacingGreen"}
	user5 := []string{"a"}

	tests := []test{
		{
			name:   "user0, user1",
			inputA: user0,
			inputB: user1,
			want:   []string{"/pink", "/register", "/orange"},
		},
		{
			name:   "user0, user2",
			inputA: user0,
			inputB: user2,
			want:   []string{"a"},
		},
		{
			name:   "user2, user0",
			inputA: user2,
			inputB: user0,
			want:   []string{"a"},
		},
		{
			name:   "user3, user4",
			inputA: user3,
			inputB: user4,
			want:   []string{"/green", "/blue", "/purple", "/white"},
		},
		{
			name:   "user5, user2",
			inputA: user5,
			inputB: user2,
			want:   []string{"a"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := findContiguousHistory(tc.inputA, tc.inputB); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
