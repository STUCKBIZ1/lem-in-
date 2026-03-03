package main

import (
	"fmt"
	"lem-in/server"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <filename>")
		return
	}
	colony, err := server.ParseInput(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	paths, err := server.FindPaths(colony)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	moves := server.Simulate(paths, colony.NumAnts)

	// All good — print everything
	for _, line := range colony.RawLines {
		fmt.Println(line)
	}
	fmt.Println()
	for _, move := range moves {
		fmt.Println(move)
	}
}
