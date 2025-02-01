package mathFunctions

func Median(numbers []float64) float64 {
	sortedNumbers := QuickSort(numbers)
	n := len(sortedNumbers)
	if n%2 == 0 {
		return (sortedNumbers[n/2-1] + sortedNumbers[n/2]) / 2
	}
	return sortedNumbers[n/2]
}

func QuickSort(numbers []float64) []float64 {
	if len(numbers) <= 1 {
		return numbers
	}

	last := numbers[len(numbers)-1]
	var left, right []float64

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < last {
			left = append(left, numbers[i])
		} else {
			right = append(right, numbers[i])
		}
	}
	return append(append(QuickSort(left), last), QuickSort(right)...)
}
