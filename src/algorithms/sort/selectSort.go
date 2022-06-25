package sort

func selectSort(n []int) []int {
	i := 1
	for i < len(n)-1 {
		j := i + 1
		minIndex := i

		if j < len(n) {
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		n[i], n[minIndex] = n[minIndex], n[i]
		i++
	}
	return n
}