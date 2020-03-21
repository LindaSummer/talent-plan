package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	res := mergeSort(src)
	for i := range src {
		src[i] = res[i]
	}
}

func mergeSort(src []int64) []int64 {
	if len(src) <= 1 {
		return src
	}

	mid := len(src) / 2
	ls := []int64{}
	rs := []int64{}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ls = mergeSort(src[:mid])
		wg.Done()
	}()

	go func() {
		rs = mergeSort(src[mid:])
		wg.Done()
	}()

	wg.Wait()
	return mergeTwoSorts(ls, rs, func(lhs, rhs int64) bool {
		return lhs < rhs
	})
}

func mergeTwoSorts(left, right []int64, compare func(lhs, rhs int64) bool) []int64 {
	rtn := []int64{}
	li := 0
	ri := 0
	for {
		if li < len(left) && ri < len(right) {
			if compare(left[li], right[ri]) {
				rtn = append(rtn, left[li])
				li++
			} else {
				rtn = append(rtn, right[ri])
				ri++
			}
		} else {
			break
		}
	}
	if li < len(left) {
		rtn = append(rtn, left[li:]...)
	}
	if ri < len(right) {
		rtn = append(rtn, right[ri:]...)
	}
	return rtn
}
