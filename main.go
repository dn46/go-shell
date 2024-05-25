package main

import (
	"fmt"

	"github.com/dn46/go-shell/shell"
)

func main() {
	s := shell.NewShell()

	err := s.Run()

	if err != nil {
		fmt.Println(err)
	}
}
