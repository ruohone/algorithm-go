package sorttest

import "testing"

func merge(left []user,right []user) []user  {
	var leftIndex,rightIndex = 0,0

	result := make([]user,0,len(left) + len(right))
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex].Id < right[rightIndex].Id {
			result = append(result,left[leftIndex])
			leftIndex ++
		} else {
			result = append(result,right[rightIndex])
			rightIndex ++
		}
	}

	for ;leftIndex < len(left);leftIndex++ {
		result = append(result,left[leftIndex])
	}

	for ;rightIndex < len(right); rightIndex ++ {
		result = append(result,right[rightIndex])
	}

	return result
}

func mergeSort(data []user) []user {
	if len(data) <= 1 {
		return data
	}

	middle := len(data) / 2
	left := data[0:middle]
	right := data[middle:]

	return merge(mergeSort(left),mergeSort(right))
}

func TestMergeSort(t *testing.T)  {
	data := make([]user,len(testData))
	copy(data,testData)
	result := mergeSort(data)

	t.Log(result)
}