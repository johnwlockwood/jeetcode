package group

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	keys := []string{}
	letterMap := make(map[string][]string, 0)
	for _, v := range strs {
		keyValues := strings.Split(v, "")
		sort.Strings(keyValues)
		key := strings.Join(keyValues, "")
		if _, ok := letterMap[key]; ok {
			letterMap[key] = append(letterMap[key], v)
		} else {
			letterMap[key] = []string{v}
			keys = append(keys, key)
		}
	}
	output := [][]string{}
	for _, k := range keys {
		v := letterMap[k]
		output = append(output, v)
	}
	return output
}
