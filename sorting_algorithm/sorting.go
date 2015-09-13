package sorting_algorithm

import (
	"fmt"
	"math"
)

//直接插入排序 稳定 额外空间 1
// 1.设置监视哨 r[i] = tmp
// 2.从r[1]开始比较
// 3.r[i] 和 r[i-1] 比较; r[i-1] >= tmp r[i]=tmp 即r[i-1]后移； r[i-1] < tmp r[i] = tmp
func insertionSort(r []int) {
	length := len(r)
	for i := 1; i < length; i++ {
		tmp := r[i]
		for j := i - 1; j >= -1; j-- {
			if j == -1 {
				r[0] = tmp
				break
			}
			if tmp <= r[j] {
				r[j+1] = r[j]
			} else {
				r[j+1] = tmp
				break
			}
		}
	}
	fmt.Println(r)
}

//没有监视哨
func insertionSort2(r []int) {
	length := len(r)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if r[j] >= r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
	fmt.Println(r)
}

//希尔排序  不稳定
//首先找出 num/2 进行
func shellSort(r []int) {
	length := len(r)
	d := length / 2
	for d > 1 {
		for i := 0; i < d; i++ {
			for j := i; j+d < length; j++ {
				if r[j] >= r[j+d] {
					r[j], r[j+d] = r[j+d], r[j]
				}
			}
		}
		fmt.Println("d=", d, r)
		d = int(math.Ceil(float64(d) / 2))
	}
	fmt.Println(r)
}

//冒泡排序  稳定
//
func bubbleSort(r []int) {
	fmt.Println("bubbleSort=======")
	length := len(r)
	for i := 0; i < length-1; i++ {
		isChange := false
		for j := 0; j < length-1-i; j++ {
			if r[j] < r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
		fmt.Println(r)
	}
	fmt.Println(r)
}

//快速排序  不稳定
//选定r[0]
func quickSort(r []int) {
	length := len(r)
	if length <= 1 {
		return
	}
	mid, i := r[0], 1
	head, tail := 0, length-1
	for head < tail {
		if r[i] > mid {
			r[i], r[tail] = r[tail], r[i]
			tail--
		} else {
			r[i], r[head] = r[head], r[i]
			head++
			i++
		}
	}
	r[head] = mid
	quickSort(r[:head])
	quickSort(r[head+1:])
	fmt.Println(r)
}

//选择排序  不稳定
func selectionSort(r []int) {
	length := len(r)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if r[j] < r[i] {
				r[j], r[i] = r[i], r[j]
			}
		}
	}
	fmt.Println(r)
}

//归并排序  稳定
//
func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//基数排序  稳定
func radixSort() {

}
