package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	start_message           = "Enter the numbers separated by a line break (enter `!end` at the end):"
	after_typing_message    = "Enter the metric numbers you want to see their values displayed in one line without any separators."
	choose_message          = "(1) Mean\n(2) Median\n(3) Mode\n(4) Standart deviation"
	no_nums_message         = "no numbers to calculate metrics"
	invalid_num_message     = "invalid number"
	cant_read_stdio_message = "error with reading stdio"
	invalid_opt_message     = "invalid option"
)

func validNumber(num float64, err error) (float64, bool) {
	if err != nil {
		return .0, false
	}
	return num, num > -1e5 && num < 1e5
}

func HandleUserFormatActions(fanswer []string) error {
	var user_answer string
	var err error = nil
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(after_typing_message)
	fmt.Println(choose_message)

	if scanner.Scan() {
		user_answer = scanner.Text()
	}
	for i := 0; i < len(user_answer); i++ {
		if user_answer[i]-'0' > 0 && user_answer[i]-'0' < 5 {
			fmt.Println(fanswer[int(user_answer[i]-'0')-1])
		} else {
			err = errors.New(invalid_opt_message)
			break
		}
	}
	return err
}

func GetData() ([]float64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(start_message)

	nums := make([]float64, 0)
	var err error = nil

	for scanner.Scan() {
		line := scanner.Text()
		if line == "!end" {
			if len(nums) < 1 {
				err = errors.New(no_nums_message)
			}
			break
		}

		if num, ok := validNumber(strconv.ParseFloat(line, 64)); ok {
			nums = append(nums, num)
		} else {
			err = errors.New(invalid_num_message)
			break
		}
	}

	if err == nil {
		err = scanner.Err()
	}

	return nums, err
}
