package sorttest

import "testing"

func adjustForMaxHeap(arr []user,i,length int)  {
	toMaxNode := func(k int) int {
		if k + 1 < length && arr[k].Id < arr[k + 1].Id {
			k ++
		}
		return k
	}

	curNode := arr[i]
	for k := i*2 + 1; k < length; k = k * 2 + 1 {
		//   切换到 最大值到结点
		k = toMaxNode(k)

		if curNode.Id < arr[k].Id {
			arr[i] = arr[k]
			i = k
		}
	}
	arr[i] = curNode
}

func heapShort(data []user) {
	for i:= len(data)/2 - 1; i >= 0; i -- {
		adjustForMaxHeap(data,i,len(data))
	}

	for i:=len(data) -1; i >0; i -- {
		swap(data,0,i)
		adjustForMaxHeap(data,0,i)
	}
}

func TestHeapSort(t *testing.T) {
	data := make([]user,len(testData))
	copy(data,testData)
	heapShort(data)

	t.Log(data)
}