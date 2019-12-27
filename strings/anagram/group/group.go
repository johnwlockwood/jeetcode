package group

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	letterMap := make(map[string][]string, 0)
	for _, v := range strs {
		keyValues := strings.Split(v, "")
		sort.Strings(keyValues)
		key := strings.Join(keyValues, "")
		if _, ok := letterMap[key]; ok {
			letterMap[key] = append(letterMap[key], v)
		} else {
			letterMap[key] = []string{v}
		}
	}
	output := [][]string{}
	for _, v := range letterMap {
		output = append(output, v)
	}
	return output
}
