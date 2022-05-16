package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	// Input -> {2: "a", 0: "b", 1: "c"}
	// Output -> ["b", "c", "a"]
	// Input -> {10: "aa", 0: "bb", 500: "cc"}
	// Output -> ["bb", "aa", "cc"]
	keys := make([]int, 0, len(input))
	values := make([]string, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		values = append(values, input[k])
	}
	//Place your code here
	return values
}
