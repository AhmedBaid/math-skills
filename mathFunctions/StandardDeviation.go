package mathFunctions

import "math"

func StandardDeviation(numbers []float64) float64 {
	return math.Sqrt(float64(Variance(numbers)))
}
