package moveto

import "testing"

func TestMoveToCity(t *testing.T) {
	type test struct {
		inputValue float64
		input      []CityPop
		want       string
	}

	tests := []test{
		{
			inputValue: 0.22,
			input: []CityPop{
				CityPop{
					Name:       "New York",
					Population: 40,
				},
				CityPop{
					Name:       "New Orleans",
					Population: 10,
				},
				CityPop{
					Name:       "Boston",
					Population: 30,
				},
				CityPop{
					Name:       "San Franscisco",
					Population: 20,
				},
			},
			want: "San Franscisco",
		},
		{
			inputValue: 0.80,
			input: []CityPop{
				CityPop{
					Name:       "New York",
					Population: 40,
				},
				CityPop{
					Name:       "New Orleans",
					Population: 10,
				},
				CityPop{
					Name:       "Boston",
					Population: 30,
				},
				CityPop{
					Name:       "San Franscisco",
					Population: 20,
				},
			},
			want: "New York",
		},
	}

	for _, tc := range tests {
		if got := moveToCity(tc.inputValue, tc.input); got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
