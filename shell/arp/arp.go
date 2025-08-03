package arp

import "os"
import "fmt"
import "strings"
import "os/exec"

func RunProgram(program string) string {
	a := strings.Split(program, " ")
	out, err := exec.Command(a[0], a[1:]...).Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func ArpTable() {
	s := os.Expand("${arp -a}", RunProgram)
	fmt.Print(s)
}
