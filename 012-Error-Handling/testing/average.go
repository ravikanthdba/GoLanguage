package main

func average(x []float64) float64 {
	sum := float64(0)
	count := 0

	for values := range x {
		sum += x[values]
		count++
	}

	return (float64(sum) / float64(count))
}
