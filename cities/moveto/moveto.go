package moveto

import "sort"

type CityPop struct {
	Name       string
	Population int
}

type citySize struct {
	name string
	pop  float64
}

func moveToCity(value float64, cities []CityPop) string {
	total := 0
	citySizes := make([]citySize, len(cities))
	for _, c := range cities {
		total += c.Population
	}
	for _, c := range cities {
		citySizes = append(citySizes, citySize{name: c.Name, pop: float64(c.Population) / float64(total)})
	}
	sort.Slice(citySizes, func(i, j int) bool {
		return citySizes[i].pop < citySizes[j].pop
	})
	t := float64(0)
	name := ""
	for i := 0; i < len(citySizes); i++ {
		t += citySizes[i].pop
		if value <= t {
			return citySizes[i].name
		}
	}
	return name
}
