package sorttest

import "testing"

func bubbleSort(users []user) {
	for i := len(users); i > 1; i-- {
		for j := 1; j < i; j++ {
			if users[j-1].Id > users[j].Id {
				temp := users[j]
				users[j] = users[j-1]
				users[j-1] = temp
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	data := make([]user,len(testData))
	copy(data,testData)
	bubbleSort(data)

	t.Log(data)
}

/*
110100101010  j
100101011011  i

010001110001  j
100101011011  i

110100101010 i
010001110001 j

100101011011 j

*/
