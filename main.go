package main

import (
	"bufio"
	"fmt"
	"github.com/bsdpunk/newShell/shell"
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

	//	scanner := bufio.NewScanner(os.Stdin)
	//	stat, _ := os.Stdin.Stat()
	//	if (stat.Mode() & os.ModeCharDevice) == 0 {
	//		//	fmt.Println("data is being piped to stdin")
	//		for scanner.Scan() {
	//			fmt.Println(scanner.Text())
	//		}
	//		if err := scanner.Err(); err != nil {
	//			log.Println(err)
	//		}
	//	} else {
	//		fmt.Println(shell.Shell(os.Args))
	//	}

}
