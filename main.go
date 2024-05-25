package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // setting the standard input device

	for {
		fmt.Print("go-shell> ")
		// read the input
		input, err := reader.ReadString('\n') // after every newline

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle the execution of the input

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n") // removing the newline character at the end of the input

	args := strings.Split(input, " ") // split the input to separate the command and the arguments (ls -l)

	cmd := exec.Command(args[0], args[1:]...) // prepare the command to execute

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
