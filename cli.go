package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/Drahoslav7/gobang/logic"
)

func startCLI() {
	exiting := false
	app.OnExit(func() {
		exiting = true
	})

	fmt.Println("CLI started")
	defer fmt.Println("CLI ended")

	clin := bufio.NewReader(os.Stdin)

	app.Entrance(func() {
		fmt.Println("app running")
	})

	for {
		if exiting {
			return
		}
		str, err := clin.ReadString('\n')
		if err != nil {
			return
		}
		str = strings.TrimSpace(str) // trunc \n
		parts := strings.Split(str, " ")

		switch parts[0] {
		case "quit", "q":
			app.Exit()
		case "practice":
			app.Practice(nil)
		case "play":
			size, _ := strconv.Atoi(parts[1])
			app.Play(size, nil)
		}

	}

}
