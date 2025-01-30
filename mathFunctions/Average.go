package mathFunctions

func Average(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += float64(num)
	}
	return sum / float64(len(numbers))
}
