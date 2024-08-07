package main

import (
	"fmt"
	"os"

	"github.com/zkhrg/go_day00/pkg/metrics"
	"github.com/zkhrg/go_day00/pkg/views/cli"
)

func main() {
	nums, err := cli.GetData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	answer, _ := metrics.GetFormatedAnswer(nums)
	if err = cli.HandleUserFormatActions(answer); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
