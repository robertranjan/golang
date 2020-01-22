package mix

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// RunShell ...
func RunShell() int {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input... robert ranjan")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	return 0
}

// Ls ...
func Ls() int {
	log.Println("\n\033[32mRunning mix.Ls() method")
	cmd := exec.Command("ls", "-altr", "/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return len(out)
}
