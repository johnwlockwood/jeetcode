package substitution

import (
	"fmt"
	"strings"
)

/*

Your friends are now complaining that it's too hard to make sure the lengths of their status updates are not prime numbers.

You decide to create a substitution cipher. The cipher alphabet is based on a key shared amongst those of your friends who don't mind spoilers.

Suppose the key is:
"The quick onyx goblin, grabbing his sword, jumps over the lazy dwarf!".

We use only the unique letters in this key to set the order of the characters in the substitution table.

T H E Q U I C K O N Y X G B L R A S W D J M P V Z F

(spaces added for readability)

We then align it with the regular alphabet:
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
T H E Q U I C K O N Y X G B L R A S W D J M P V Z F

Which gives us the substitution table: A becomes T, B becomes H, C becomes E, etc.

Write a function that takes a key and a string and encrypts the string with the key.

Example:
key = "The quick onyx goblin, grabbing his sword, jumps over the lazy dwarf!"
encrypt("It was all a dream.", key) -> "Od ptw txx t qsutg."

*/

func sub(message, key string) string {
	keySet := make(map[byte]struct{}, 0)
	keyList := make([]byte, 0, 26)
	A := []byte("A")[0]
	// Z := []byte("Z")[0]
	a := []byte("a")[0]
	z := []byte("z")[0]
	key = strings.ToUpper(key)
	for i := 0; i < len(key); i++ {
		v := key[i]
		if _, ok := keySet[v]; !ok && v >= A && v <= z {
			keyList = append(keyList, v)
		}
		keySet[v] = struct{}{}
	}
	fmt.Println(string(keyList))
	encodeMap := make(map[byte]byte, len(key)*2)

	for i := range keyList {
		v := keyList[i]
		var upper, lower byte
		upper = []byte(strings.ToUpper(string(v)))[0]
		lower = []byte(strings.ToLower(string(v)))[0]

		encodeMap[byte(i+int(A))] = upper
		encodeMap[byte(i+int(a))] = lower

	}
	fmt.Println("Map")
	for k, v := range encodeMap {
		fmt.Printf("%s: %s\n", string(k), string(v))
	}
	fmt.Println("")

	ret := make([]byte, 0, len(message))
	for i := 0; i < len(message); i++ {
		if v, ok := encodeMap[message[i]]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, message[i])
		}
	}
	return string(ret)
}

func routeIt(message string, r, c int) string {
	cipher := make([][]byte, r)
	for i := range cipher {
		cipher[i] = make([]byte, c)
	}
	//   fmt.Printf("%v\n", cipher)

	i := 0
	for rIndex := 0; rIndex < r; rIndex++ {
		for j := 0; j < c && i < len(message); j++ {
			cipher[rIndex][j] = message[i]
			i++
		}
	}

	//   fmt.Printf("after %v\n", cipher)

	ret := make([]byte, 0, len(message))

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			ret = append(ret, cipher[j][i])
		}

	}

	return string(ret)
}
