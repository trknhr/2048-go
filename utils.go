package main

func ReverseList(l []int) []int{
	llen := len(l)
	r := make([]int, llen)

	for i := 0; i < llen; i++{
		r[len(l) - i - 1] = l[i]
	}

	return r
}
