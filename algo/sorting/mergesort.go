package sorting

type Sorter interface {
	Sort([]int) []int
}

type IntMergeSorter struct {
}

func (ims *IntMergeSorter) mergeSort(array []int, left int, right int) {
	if left < right {
		center := (left + right) / 2
		ims.mergeSort(array, left, center)
		ims.mergeSort(array, center+1, right)
		ims.merge(array, left, right, center)
	}

}
func (ims *IntMergeSorter) merge(a []int, left int, center int, right int) {
	i := left
	j := center + 1
	k := 0
	b := make([]int, right-left+1)

	for i <= center && j <= right {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
		k++
	}
	for i <= center {
		b[k] = a[i]
		i++
		k++
	}
	for j <= right {
		b[k] = a[j]
		j++
		k++
	}
	for k = left; k <= right; k++ {
		a[k] = b[k-left]
	}
}

func (ims *IntMergeSorter) Sort(array []int) {
	if len(array) == 1 {
		return
	}
	ims.mergeSort(array, 0, len(array)-1)
}
