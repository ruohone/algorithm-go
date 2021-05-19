package sorttest

import "testing"

func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[j] ^= data[i]
				data[i] ^= data[j]
				data[j] ^= data[i]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 4, 432, 67, 89, 4, 3}
	bubbleSort(arr)

	t.Log(arr)
}
