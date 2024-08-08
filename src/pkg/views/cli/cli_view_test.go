package cli_test

import (
	"io"
	"math"
	"os"
	"testing"

	"github.com/zkhrg/go_day00/pkg/views/cli"
)

func TestGetData(t *testing.T) {
	expected_nums := []float64{1, 2, 3, 4, 5}
	input := "1\n2\n3\n4\n5\n!end\n"

	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}

	originalStdin := os.Stdin
	originalStdout := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdin = originalStdin; os.Stdout = originalStdout }()

	os.Stdin = tmpfile

	nums, _ := cli.GetData()

	if len(nums) != len(expected_nums) {
		t.Fatalf("expected %v lines, got %v lines", len(expected_nums), len(nums))
	}

	for i := 0; i < len(expected_nums); i++ {
		if !eqFloat(nums[i], expected_nums[i]) {
			t.Errorf("expected nums_val %f, got %f", expected_nums[i], nums[i])
		}
	}
}

func TestGetDataError(t *testing.T) {
	input := "1\n2\n3\n4\ns\n!end\n"

	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}

	originalStdin := os.Stdin
	originalStdout := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdin = originalStdin; os.Stdout = originalStdout }()

	os.Stdin = tmpfile

	_, err1 := cli.GetData()
	if err1 == nil {
		t.Fatalf("expected error have nil")
	}
}

func TestHandleUserFormatActions(t *testing.T) {
	tests := []struct {
		user_input string
		fanswer    []string
		is_err     bool
	}{
		{"1234c\n", []string{"3", "!", "3", "3"}, true},
		{"\n", []string{"3", "!", "3", "3"}, true},
		{"123\n", []string{"3", "!", "3"}, true},
		{"123\n", []string{"3", "!", "3"}, true},
		{"123\n", []string{"3", "!", "3", "213"}, false},
		{"123123123123123\n", []string{"3", "!", "3", "213"}, false},
		{"1231231231231235\n", []string{"3", "!", "3", "213"}, true},
		{"111111111111111\n", []string{"3", "!", "3", "213"}, false},
	}

	for i := 0; i < len(tests); i++ {
		originalStdin := os.Stdin
		originalStdout := os.Stdout
		tmpfile := passStringToTmpStream(tests[i].user_input, t)
		os.Stdout = nil

		os.Stdin = tmpfile
		err1 := cli.HandleUserFormatActions(tests[i].fanswer)
		if err1 == nil && tests[i].is_err {
			t.Fatalf("expected error have nil %s", tests[i].user_input)
		} else if err1 != nil && !tests[i].is_err {
			t.Fatalf("expected nil have error %s", tests[i].user_input)
		}
		os.Stdin = originalStdin
		os.Stdout = originalStdout
		os.Remove(tmpfile.Name())
	}
}

func passStringToTmpStream(str string, t *testing.T) *os.File {
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.WriteString(str); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	return tmpfile
}

func eqFloat(a, b float64) bool {
	return math.Abs(a-b) < 1e-8
}
