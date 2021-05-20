package sorttest

import "testing"

func insertSort(data []user)  {
	for i:=1; i<len(data); i++{
		for j:=i; j>=1; j -- {
			if data[j - 1].Id > data[j].Id {
				temp := data[j - 1]
				data[j - 1] = data[j]
				data[j] = temp
			} else {
				break
			}
		}
	}
}

func TestInsertSort(t *testing.T)  {
	data := make([]user,len(testData))
	copy(data,testData)
	insertSort(data)

	t.Log(data)
}