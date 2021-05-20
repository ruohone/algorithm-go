package sorttest

import "testing"

func insertSort(data []user)  {
	for i:=1; i<len(data); i++{
		temp := data[i]
		replaceIndex := 0
		for j:=i-1; j>=0; j -- {
			if temp.Id >= data[j].Id {
				replaceIndex = j+1
				break
			}
			data[j+1] = data[j]
		}
		data[replaceIndex] = temp
	}
}

func TestInsertSort(t *testing.T)  {
	data := make([]user,len(testData))
	copy(data,testData)
	insertSort(data)

	t.Log(data)
}