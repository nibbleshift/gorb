package main

import (
	"fmt"

	"github.com/nibbleshift/internal/exec"
)

func main() {
	fmt.Println("test")

	exec.Run("ls")
}
