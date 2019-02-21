package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {

	a := make([]int, 8)
	a[0] = 3
	a[1] = 5
	a[2] = 4
	a[3] = 1
	a[4] = 2
	a[5] = 6
	a[6] = 2
	a[7] = 6

	quickSort(a)

	for _, d := range a {
		fmt.Println(d)
	}
}

func quickSort(a []int) {
	quickSortC(a, 0, len(a)-1)
}

func quickSortC(a []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(a, p, r)
	quickSortC(a, p, q-1)
	quickSortC(a, q+1, r)
}

func partition(a []int, p, r int) int {
	pivot := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			tmp := a[i]
			a[i] = a[j]
			a[j] = tmp
			i++
		}
	}

	a[r] = a[i]
	a[i] = pivot
	return i
}

func mergeSort(a []int) {
	mergeSortC(a, 0, len(a)-1)
}

func mergeSortC(a []int, p, r int) {
	if p >= r {
		return
	}

	q := (p + r) / 2

	mergeSortC(a, p, q)
	mergeSortC(a, q+1, r)
	mergeArray(a, p, r, q)
}

func mergeArray(a []int, p, r, q int) {
	tmp := make([]int, r-p+1)
	var i, j, k = p, q + 1, 0

	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	var start, end int
	if i <= q {
		start = i
		end = q
	} else {
		start = j
		end = r
	}

	for start <= end {
		tmp[k] = a[start]
		k++
		start++
	}

	for i = 0; i < len(tmp); i++ {
		a[p+i] = tmp[i]
	}

}

func selectionSort(a []int) {
	if len(a) <= 1 {
		return
	}

	n := len(a)

	for i := 0; i < n; i++ {
		value := a[i]
		index := i
		for j := i + 1; j < n; j++ {
			if value > a[j] {
				value = a[j]
				index = j
			}
		}

		if index != i {
			tmp := a[i]
			a[i] = a[index]
			a[index] = tmp
		}
	}

}

func insertionSort(a []int) {
	if len(a) <= 1 {
		return
	}
	n := len(a)
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value { //如果等于就不是稳定排序。
				a[j+1] = a[j]
			} else {
				break
			}
		}

		a[j+1] = value
	}
}

func bubbleSort(a []int) {
	if len(a) <= 1 {
		return
	}
	n := len(a)
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] { //交换
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true
			}
		}

		if !flag {
			break
		}
	}

}
