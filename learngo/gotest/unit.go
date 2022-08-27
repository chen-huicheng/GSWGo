package gotest

func MySort1(arr []int) {
	if len(arr) <= 0 {
		return
	}
	tmp := arr[0]
	l, r := 0, len(arr)-1
	for l < r {
		for l < r && arr[r] >= tmp {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] < tmp {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = tmp
	MySort1(arr[:l])
	MySort1(arr[l+1:])
}
