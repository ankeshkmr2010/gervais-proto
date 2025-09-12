package utils

import (
	"fmt"
	"math"
)

func dotProduct(vec1, vec2 []float64) (float64, error) {
	if len(vec1) != len(vec2) || len(vec1) == 0 {
		return 0, fmt.Errorf("vectors must be of same non-zero length")
	}

	sum := 0.0
	for i := range vec1 {
		sum += vec1[i] * vec2[i]
	}
	return sum, nil
}

func magnitude(vec []float64) float64 {
	sumSquares := 0.0
	for _, v := range vec {
		sumSquares += v * v
	}
	return math.Sqrt(sumSquares)
}

func CosineSimilarity(vec1, vec2 []float64) (float64, error) {
	dotProd, err := dotProduct(vec1, vec2)
	if err != nil {
		return 0, err
	}

	mag1 := magnitude(vec1)
	mag2 := magnitude(vec2)

	if mag1 == 0 || mag2 == 0 {
		return 0, fmt.Errorf("one or both vectors have zero magnitude")
	}

	return dotProd / (mag1 * mag2), nil
}
