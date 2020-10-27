package main

import (
	"fmt"
	"test-wrapper/internal/program"
	"test-wrapper/pkg/wrapper"
)

func main() {
	wrapper := wrapper.Wrapper{}

	for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 10} {
		err := wrapper.ExecFunc(func() error {
			return program.Executor(n)
		})

		fmt.Println(n, err)
	}

	fmt.Println("Printing wrapper:", wrapper)
}
