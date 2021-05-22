package sorttest

import "testing"

func selectSort(data []user)  {
	for i := 0; i < len(data) - 1; i++ {
		for j:=i; j < len(data); j++ {
			if data[i].Id > data[j].Id {
				swap(data,i,j)
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