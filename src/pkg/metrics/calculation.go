package metrics

import (
	"fmt"
	"math"
	"slices"
)

type Metric int

const (
	Mean Metric = iota
	Median
	Mode
	SD
)

const (
	Eps = 1e-9
)

func FindMean(nums []float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}

func FindMedian(nums []float64) float64 {
	var res float64
	half_size := len(nums) >> 1
	if len(nums)&1 == 1 {
		res = nums[half_size]
	} else {
		res = (nums[half_size] + nums[half_size-1]) / float64(2)
	}
	return res
}

func FindMode(nums []float64) float64 {
	m := make(map[float64]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	min_key := Eps
	max_count := 0
	for k, v := range m {
		if max_count < v {
			max_count = v
			min_key = k
		} else if max_count == v && min_key > k {
			min_key = k
		}
	}
	return min_key
}

func FindStandartDeviation(nums []float64) float64 {
	var r float64
	nums_mean := FindMean(nums)
	for _, v := range nums {
		r += math.Pow((v - nums_mean), 2)
	}
	if len(nums) == 1 {
		return .0
	}
	return math.Sqrt(r / float64(len(nums)-1))
}

func CalculateAllMetrics(nums []float64, answer map[Metric]float64) {
	answer[Mean] = FindMean(nums)
	answer[Median] = FindMedian(nums)
	answer[Mode] = FindMode(nums)
	answer[SD] = FindStandartDeviation(nums)
}

func ConvertAnswerToStringSlice(answer map[Metric]float64) []string {
	fanswer := []string{"Mean: %.2f", "Median: %.2f", "Mode: %.2f", "SD: %.2f"}
	for i := 0; i < len(answer); i++ {
		fanswer[i] = fmt.Sprintf(fanswer[i], answer[Metric(i)])
	}
	return fanswer
}

func GetFormatedAnswer(nums []float64) []string {
	slices.Sort(nums)
	metric_vals := make(map[Metric]float64)
	CalculateAllMetrics(nums, metric_vals)
	return ConvertAnswerToStringSlice(metric_vals)
}
