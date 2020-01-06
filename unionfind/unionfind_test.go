package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	t.Run("n 10 connections", func(t *testing.T) {
		n := 10
		items := make([]byte, n)
		for i := 0; i < n; i++ {
			items[i] = byte(i) + 'A'
		}
		u := NewUnionFind(n)
		if got, _ := u.Find(0, 1); got != false {
			t.Errorf("%s and %s should not be connected, but are", string(items[0]), string(items[1]))
		}
		_ = u.Union(0, 7)
		_ = u.Union(7, 1)
		if got, _ := u.Find(7, 1); got != true {
			t.Errorf("%s and %s should be connected, but are not", string(items[7]), string(items[1]))
		}
		if got, _ := u.Find(0, 1); got != true {
			t.Errorf("%s and %s should be connected, but are not", string(items[0]), string(items[1]))
		}
		if got, _ := u.Root(7); got != 0 {
			t.Errorf("Root of %s should be %s, got %s", string(items[7]), string(items[0]), string(items[got]))
		}
		if got, _ := u.Find(8, 1); got != false {
			t.Errorf("%s and %s should not be connected, but are", string(items[8]), string(items[1]))
		}
		u.Union(8, 1)
		if got, _ := u.Root(8); got != 0 {
			t.Errorf("Root of %s should be %s, got %s", string(items[8]), string(items[0]), string(items[got]))
		}
	})

	t.Run("out of bounds method", func(t *testing.T) {
		u := NewUnionFind(10)
		err := u.outOfBoundsError(22)
		if err != nil {
			err := err.(OutOfBoundsError)
			if err.Upper != 10 {
				t.Errorf("got %d, want error Upper %d", err.Upper, 10)
			}
			if err.Lower != 0 {
				t.Errorf("got %d, want error Lower %d", err.Lower, 0)
			}
			if err.Index != 22 {
				t.Errorf("got %d, want error Index %d", err.Index, 50)
			}
		} else {
			t.Error("expected error")
		}
	})

	t.Run("out of bounds Union", func(t *testing.T) {
		type test struct {
			inputA    int
			inputB    int
			want      int
			expectErr bool
		}
		tests := []test{
			{

				inputA:    0,
				inputB:    50,
				want:      50,
				expectErr: true,
			},
			{
				inputA:    0,
				inputB:    -3,
				want:      -3,
				expectErr: true,
			},
			{
				inputA:    -5,
				inputB:    9,
				want:      -5,
				expectErr: true,
			},
			{
				inputA:    100,
				inputB:    9,
				want:      100,
				expectErr: true,
			},
			{
				inputA:    0,
				inputB:    9,
				expectErr: false,
			},
		}

		for _, tc := range tests {
			n := 10
			u := NewUnionFind(n)
			got := u.Union(tc.inputA, tc.inputB)
			if tc.expectErr && got != nil {
				got, ok := got.(OutOfBoundsError)
				if !ok {
					t.Errorf("expected OutOfBoundsError error, got %v", got)
				} else if got.Index != tc.want {
					t.Errorf("got %d, want error Index %d", got.Index, tc.want)
				}
			} else if tc.expectErr {
				t.Error("expected error")
			} else if got != nil {
				t.Errorf("expected no error, got %s", got.Error())
			}
		}
	})

	t.Run("out of bounds Find", func(t *testing.T) {
		type test struct {
			inputA    int
			inputB    int
			want      int
			expectErr bool
		}
		tests := []test{
			{

				inputA:    0,
				inputB:    50,
				want:      50,
				expectErr: true,
			},
			{
				inputA:    0,
				inputB:    -3,
				want:      -3,
				expectErr: true,
			},
			{
				inputA:    -5,
				inputB:    9,
				want:      -5,
				expectErr: true,
			},
			{
				inputA:    100,
				inputB:    9,
				want:      100,
				expectErr: true,
			},
			{
				inputA:    0,
				inputB:    9,
				expectErr: false,
			},
		}

		for _, tc := range tests {
			n := 10
			u := NewUnionFind(n)
			_, got := u.Find(tc.inputA, tc.inputB)
			if tc.expectErr && got != nil {
				got, ok := got.(OutOfBoundsError)
				if !ok {
					t.Errorf("expected OutOfBoundsError error, got %v", got)
				} else if got.Index != tc.want {
					t.Errorf("got %d, want error Index %d", got.Index, tc.want)
				}
			} else if tc.expectErr {
				t.Error("expected error")
			} else if got != nil {
				t.Errorf("expected no error, got %s", got.Error())
			}
		}
	})

	t.Run("out of bounds Root", func(t *testing.T) {

		type test struct {
			input     int
			want      int
			expectErr bool
		}
		tests := []test{
			{
				input:     50,
				want:      50,
				expectErr: true,
			},
			{
				input:     -3,
				want:      -3,
				expectErr: true,
			},
			{
				input:     7,
				expectErr: false,
			},
		}

		for _, tc := range tests {
			n := 10
			u := NewUnionFind(n)
			_, got := u.Root(tc.input)
			if tc.expectErr && got != nil {
				got, ok := got.(OutOfBoundsError)
				if !ok {
					t.Errorf("expected OutOfBoundsError error, got %v", got)
				} else if got.Index != tc.want {
					t.Errorf("got %d, want error Index %d", got.Index, tc.want)
				}
			} else if tc.expectErr {
				t.Error("expected error")
			} else if got != nil {
				t.Errorf("expected no error, got %s", got.Error())
			}
		}
	})
}
