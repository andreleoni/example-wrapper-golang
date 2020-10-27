package program

import "fmt"

func Executor(n int) error {
	if n == 2 || n == 5 || n == 8 {
		return fmt.Errorf("some error")
	}

	return nil
}
