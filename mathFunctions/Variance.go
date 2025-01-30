package mathFunctions

func Variance(numbers []float64) float64 {
	avg := Average(numbers)
	var sum float64
	for _, num := range numbers {
		diff := num - avg
		sum += diff * diff
	}
	return sum / float64(len(numbers))
}
