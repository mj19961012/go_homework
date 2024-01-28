package main

func DeleteWithIndex[T any](arr []T, index int) []T {
	if len(arr) < index || index < 0 {
		return arr
	}

	for i := index; i < (len(arr) - 1); i++ {
		arr[index] = arr[index+1]
	}

	var newarr []T = arr
	if (len(arr) * 2) < cap(arr) {
		newarr = make([]T, 0, (cap(arr) / 2))
	} else {
		newarr = make([]T, 0, cap(arr))
	}

	for j := 0; j < len(arr)-1; j++ {
		newarr = append(newarr, arr[j])
	}

	return newarr
}
