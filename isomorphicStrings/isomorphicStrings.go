package isomorphicStrings

import "fmt"

// easy problem

// Isomorphic Strings
// Given two strings s and t, determine if they are isomorphic
// no two characters may map to the same character
// all chars must be replaced with another while preserving order
// a char may map to itself

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Println("mismatched size")
		return false
	}
	// true : egg, add
	// false: foo, bar
	pairsST := make(map[string]string, len(s))
	pairsTS := make(map[string]string, len(s))
	for i, sv := range s {
		sv := string(sv)
		tv := string(t[i])
		if pairedTv, ok := pairsST[sv]; ok && pairedTv != tv {
			return false
		}
		if pairedSV, ok := pairsTS[tv]; ok && pairedSV != sv {
			return false
		}
		pairsST[sv] = tv
		pairsTS[tv] = sv
	}
	return true
}
