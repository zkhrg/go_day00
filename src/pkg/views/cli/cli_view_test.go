package cli_test

import (
	"io"
	"os"
	"testing"
)

func TestGetData(t *testing.T) {

}

// below is sample of test function need to rewrite
func TestHandleUserFormatActions(t *testing.T) {
	input := "line1\nline2\nline3\n"
	expected := []string{"line1", "line2", "line3"}

	// Создаём временный файл и записываем туда тестовые данные
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // удаляем файл после теста

	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}

	// Сохраняем оригинальный os.Stdin и восстанавливаем его после теста
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	// Перенаправляем os.Stdin на наш временный файл
	os.Stdin = tmpfile

	// Вызываем функцию
	// result := ReadInput(os.Stdin)
	result := []string{" ", " ", " "}

	// Проверяем результат
	if len(result) != len(expected) {
		t.Fatalf("expected %v lines, got %v lines", len(expected), len(result))
	}

	for i, line := range expected {
		if result[i] != line {
			t.Errorf("expected line %q, got %q", line, result[i])
		}
	}
}
