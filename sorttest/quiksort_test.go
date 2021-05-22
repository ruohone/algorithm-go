package sorttest

import "testing"

func quickSort(data []user)  {
	_quickSort(data,0,len(data)-1)
}

func _quickSort(data[]user,left int,right int)  {
	if left >= right {
		return
	}

	pivot := partition(data,left,right)
	_quickSort(data,left,pivot)
	_quickSort(data,pivot+1,right)
}

func partition(data []user,left,right int) int {
	pivotValue := data[left]
	pivotAfterIndex := left + 1

	for i := pivotAfterIndex; i <= right; i ++ {
		if pivotValue.Id > data[i].Id {
			data[pivotAfterIndex],data[i] = data[i],data[pivotAfterIndex]
			pivotAfterIndex++
		}
	}

	// 如果不交换，第一个元素大于所有后面大元素时，会导致死循环。
	data[pivotAfterIndex- 1],data[left] = data[left],data[pivotAfterIndex- 1]
	return pivotAfterIndex - 1
}

func TestQuickSort(t *testing.T) {
	data := make([]user,len(testData))
	copy(data,testData)
	quickSort(data)

	t.Log(data)
}