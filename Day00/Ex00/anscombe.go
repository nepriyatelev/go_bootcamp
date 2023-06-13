package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"sort"
)

func Anscombe(in []float64) string {

	var result string
	if meanVar {
		result += fmt.Sprintf("Mean: %.2f\n", stat.Mean(in, nil))
	}
	if medianVar {
		result += fmt.Sprintf("Median: %.2f\n", median(in))
	}
	if modeVar {
		result += fmt.Sprintf("Mode: %d\n", mode(in))
	}
	if sdVar {
		result += fmt.Sprintf("Sd: %.2f\n", stat.StdDev(in, nil))
	}
	return result
}

func median(in []float64) float64 {

	sort.Float64s(in)
	if len(in)%2 == 1 {
		midIndex := len(in) / 2
		return in[midIndex]
	}
	firstMidIndex := len(in)/2 - 1
	secondMidIndex := len(in) / 2
	return float64(in[firstMidIndex]+in[secondMidIndex]) / 2
}

func mode(in []float64) int {

	m := make(map[float64]int, len(in))
	var result float64
	var count float64

	for i := range in {
		m[in[i]]++
	}

	for k, v := range m {
		fv := float64(v)
		switch {
		case fv > count:
			result, count = k, fv
		case fv == count && k < result:
			result = k
		}
	}

	return int(result)
}
