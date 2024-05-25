package shell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type Shell struct {
	reader *bufio.Reader
}

func NewShell() *Shell {
	return &Shell{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (s *Shell) Run() error {
	for {
		// getting the current directory path
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// getting the hostname
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// getting the current user
		currentUser, err := user.Current()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Printf("\033[31m%s\033[0m@\033[32m%s\033[0m:\033[34m%s\033[0m go-shell> ", currentUser.Username, hostname, dir)
		// read the input
		input, err := s.reader.ReadString('\n') // after every newline

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

	switch args[0] {
	case "cd":
		// gotta make cd to home dir with empty path
		if len(args) < 2 {
			return errors.New("path required")
		}

		// change directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...) // prepare the command to execute

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
