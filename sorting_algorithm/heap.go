package sorting_algorithm

import (
	"errors"
)

type Heap struct {
	IsMin  bool
	Data   []int
	length int
}

//默认大根堆
func NewHeap(r []int) (*Heap, error) {
	if len(r) < 2 {
		return nil, errors.New("err r")
	}
	h := &Heap{false, r, len(r)}
	h.make()
	return h, nil
}

//一个数据并入到有序区间
func (this *Heap) fixUp(n int) {
	//   重整 i结点  其父结点为(n - 1) / 2
	//    parentNode := n-1/2
	//    if

}

//向下重整
func (this *Heap) fixDown(i, n int) {
	//   重整 i结点  子节点为
	j := 2*i + 1
	for j < n {
		if j+1 < n && this.Data[j] < this.Data[j+1] {
			j++
		}
		if this.Data[j] > this.Data[i] {
			this.Data[j], this.Data[i] = this.Data[i], this.Data[j]
		}
		i = j
		j = 2*i + 1
	}
}

//建堆
func (this *Heap) make() {
	for i := this.length/2 - 1; i >= 0; i-- {
		this.fixDown(i, this.length)
	}
}

//添加元素
func (this *Heap) Add(e int) {

}

//删除元素 从根结点将一个数据的“下沉”过程
func (this *Heap) Del() (result int) {
	result = this.Top()
	this.Data[0] = this.Data[this.length-1]
	this.Data = this.Data[:this.length-2]
	this.fixDown(0, this.length)
	return
}

func (this *Heap) Top() int {
	return this.Data[0]
}

func (this *Heap) Sort() {
	for i := this.length - 1; i > 0; i-- {
		this.Data[i], this.Data[0] = this.Data[0], this.Data[i]
		this.fixDown(0, i)
	}
}
