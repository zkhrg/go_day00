package metrics

import (
	"errors"
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

func FindMean(nums []float64) (float64, error) {
	var sum float64
	var err error

	if len(nums) == 0 {
		err = errors.New("empty slice")
		return 0, err
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum / float64(len(nums)), err
}

func FindMedian(nums []float64) (float64, error) {
	var err error
	var res float64

	if len(nums) == 0 {
		err = errors.New("empty slice")
		return 0, err
	}

	half_size := len(nums) >> 1
	if len(nums)&1 == 1 {
		res = nums[half_size]
	} else {
		res = (nums[half_size] + nums[half_size-1]) / float64(2)
	}
	return res, err
}

func FindMode(nums []float64) (float64, error) {
	var err error
	m := make(map[float64]int)
	min_key := Eps
	max_count := 0

	if len(nums) == 0 {
		err = errors.New("empty slice")
		return 0, err
	}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if max_count < v {
			max_count = v
			min_key = k
		} else if max_count == v && min_key > k {
			min_key = k
		}
	}
	return min_key, err
}

func FindStandartDeviation(nums []float64) (float64, error) {
	var r float64
	var err error

	if len(nums) == 0 {
		err = errors.New("empty slice")
		return 0, err
	}

	nums_mean, _ := FindMean(nums)
	for _, v := range nums {
		r += math.Pow((v - nums_mean), 2)
	}

	return math.Sqrt(r / float64(len(nums))), err
}

func CalculateAllMetrics(nums []float64, answer map[Metric]float64) error {
	if len(nums) == 0 {
		return errors.New("empty slice")
	}
	answer[Mean], _ = FindMean(nums)
	answer[Median], _ = FindMedian(nums)
	answer[Mode], _ = FindMode(nums)
	answer[SD], _ = FindStandartDeviation(nums)

	return nil
}

func ConvertAnswerToStringSlice(answer map[Metric]float64) []string {
	fanswer := []string{"Mean: %.2f", "Median: %.2f", "Mode: %.2f", "SD: %.2f"}
	for i := 0; i < len(answer); i++ {
		fanswer[i] = fmt.Sprintf(fanswer[i], answer[Metric(i)])
	}
	return fanswer
}

func GetFormatedAnswer(nums []float64) ([]string, error) {
	var err error
	slices.Sort(nums)
	metric_vals := make(map[Metric]float64)
	err = CalculateAllMetrics(nums, metric_vals)
	return ConvertAnswerToStringSlice(metric_vals), err
}
