package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

const (
	eps                  = 1e-9
	start_message        = "Enter the numbers separated by a line break (enter `!end` at the end):"
	after_typing_message = "Enter the metric numbers in the order you want to see their values displayed in one line without any separators."
	choose_message       = "(1) Mean\n(2) Median\n(3) Mode\n(4) Standart deviation\n(*) All above"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]float64)
	fmt.Println(start_message)
	nums := make([]float64, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "!end" {
			handleUserFormatActions(scanner)
		}

		if valid, num := validNumber(strconv.ParseFloat(line, 64)); valid {
			nums = append(nums, num)
		} else {
			fmt.Println("invalid number!")
		}

	}
	slices.Sort(nums) // поставить на нужное место
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}

	fmt.Println(nums)
}

func floatsEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func validNumber(num float64, err error) (bool, float64) {
	if err != nil {
		return false, .0
	}
	return num > -1e5 && num < 1e5, num
}

func handleUserFormatActions(scanner *bufio.Scanner) {
	var user_answer string
	fmt.Println(after_typing_message)
	fmt.Println(choose_message)
	if scanner.Scan() {
		user_answer = scanner.Text()
	}
	for i := 0; i < len(user_answer); i++ {

	}
}

func findMean(nums []float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}

func findMedian(nums []float64) float64 {
	var res float64
	half_size := len(nums) >> 1
	if len(nums)&1 == 1 {
		res = nums[half_size]
	} else {
		res = (nums[half_size] + nums[half_size-1]) / float64(2)
	}
	return res
}

func findMode(nums []float64) float64 {
	m := make(map[float64]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

}

func findStandartDeviation(nums []float64) float64 {

}
