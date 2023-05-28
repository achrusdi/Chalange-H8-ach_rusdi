package helpers

import (
	"os"
	"os/exec"
)

func CLearScreen() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
