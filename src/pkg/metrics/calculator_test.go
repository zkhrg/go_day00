package metrics_test

import (
	"math"
	"slices"
	"testing"

	"github.com/zkhrg/go_day00/pkg/metrics"
)

func TestMean1(t *testing.T) {
	nums := []float64{123, 3, 5, -123, 54, 23, 54}
	expected := 19.857142857143

	slices.Sort(nums)
	res := metrics.FindMean(nums)
	if math.Abs(res-expected) > metrics.Eps {
		t.Errorf("TestMean1: %f != %f", res, expected)
	}
}
