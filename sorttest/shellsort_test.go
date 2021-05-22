package sorttest

import "testing"

/*
把 length / y 分成 x段
其中每段 按 0-y 顺序取出即得 x个数组
即：第一个数组 x0 = [y0[1],y1[1],...yy[1]]
最后一个数组 xx=[y0[x],y1[x],...yy[x]]
*/
func shellSort(data []user)  {
	for gap := len(data) / 2; gap > 0; gap /= 2 {
		for i:=gap; i < len(data); i++ {
			for j:=i; j >= gap; j -= gap {
				if data[j].Id < data[j - gap].Id {
					temp := data[j - gap]
					data[j - gap] = data[j]
					data[j] = temp
				} else {
					break
				}
			}
		}
	}
}

func TestShellSort(t *testing.T)  {
	data := make([]user,len(testData))
	copy(data,testData)
	shellSort(data)

	t.Log(data)
}