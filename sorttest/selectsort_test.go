package sorttest

import "testing"

func selectSort(data []user)  {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].Id > data[j].Id {
				temp := data[j]
				data[j] = data[i]
				data[i] = temp
			}
		}
	}
}

func TestSelectSort(t *testing.T) {
	data := make([]user,len(testData))
	copy(data,testData)
	selectSort(data)

	t.Log(data)
}