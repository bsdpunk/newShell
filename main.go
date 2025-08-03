package main

import (
	"bufio"
	"fmt"
	"github.com/bsdpunk/macNetTools/shell"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		for scanner.Scan() {
			fmt.Println("Feature to come later")

		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	} else {
		shell.Shell(os.Args)
	}

}
