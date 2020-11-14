package main

func main() {

}

//  so easy
func findKeyInArrayWithIndex(arr []rune, key rune) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	index := 0
	for index < len(arr) {
		if key == arr[index] {
			return index
		} else {
			index++
		}
	}
	return -1
}

func findKeyInArrayWithSentinel(arr []rune, key rune) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	lastOne := len(arr) - 1
	cache := arr[lastOne]
	arr[lastOne] = key

	index := 0
	for arr[index] != key {
		index++
	}

	arr[lastOne] = cache
	if index == (lastOne) {
		return -1
	} else {
		return index
	}
}
