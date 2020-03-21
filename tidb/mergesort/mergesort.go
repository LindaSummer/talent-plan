package main

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

	lokChan := make(chan struct{})
	rokChan := make(chan struct{})
	mid := len(src)/2
	ls := []int64{}
	rs := []int64{}
	go func() {
		ls = mergeSort(src[:mid])
		lokChan <- struct{}{}
		close(lokChan)
	}()

	go func() {
		rs = mergeSort(src[mid:])
		rokChan <- struct {}{}
		close(rokChan)
	}()

	for {
		select {
		case _, ok :=<-lokChan:
			if !ok {
				lokChan = nil
			}
		case _, ok := <-rokChan:
			if !ok {
				rokChan = nil
			}
		}

		if lokChan == nil && rokChan == nil {
			break
		}
	}
	return mergeTwoSorts(ls, rs, func(lhs, rhs int64) bool {
		return lhs < rhs
	})
}

func mergeTwoSorts(left, right []int64, compare func(lhs, rhs int64) bool) []int64 {
	rtn := []int64{}
	li := 0
	ri := 0

	for {
		if li < len(left) && ri < len(right){
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