package src

import (
	"fmt"
	"os/exec"
)

func RunTWInit() {
	cmd := exec.Command("tailwindcss", "init")
	output, err := cmd.CombinedOutput()
    if err != nil {
        ErrOut(err)
    }
    fmt.Println(string(output))
}
